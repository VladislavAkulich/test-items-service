package postgresql

import (
	"context"
	"example/test-items-service/properties"
	"example/test-items-service/repositories/client"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	MaxAttempts     = 5
	AttemptDuration = 5 * time.Second
)

func GetInstance(context context.Context) (pool client.Client) {
	log.Println("Creating DB connection pool ...")
	var err error
	props := properties.GetProperties()
	config := newPgConfig(props)
	pool, err = getClient(context, MaxAttempts, AttemptDuration, config)
	if pool == nil || err != nil {
		log.Printf("Unable to create DB client: %v\n", err)
	}
	return pool
}

type pgConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func newPgConfig(properties *properties.Properties) *pgConfig {
	return &pgConfig{
		Username: properties.DB.Username,
		Password: properties.DB.Password,
		Host:     properties.DB.Host,
		Port:     properties.DB.Port,
		Database: properties.DB.Database,
	}
}

func getClient(ctx context.Context, maxAttempts int, maxDelay time.Duration, cfg *pgConfig) (pool client.Client, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)

	pool, err = doWithAttempts(getPgPool, ctx, dsn, maxAttempts, maxDelay)

	if err != nil {
		log.Fatal("All attempts are exceeded. Unable to connect to postgres")
	}

	return pool, nil
}

func getPgPool(ctx context.Context, dsn string) (pool *pgxpool.Pool, err error) {
	ctx, cancel := context.WithTimeout(ctx, AttemptDuration)
	defer cancel()

	pgxCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse config: %v\n", err)
	}

	pool, err = pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		log.Println("Failed to connect to postgres... Going to do the next attempt")

		return nil, err
	}

	return pool, nil
}

type connectionFunc = func(ctx context.Context, dsn string) (pool *pgxpool.Pool, err error)

func doWithAttempts(fn connectionFunc, ctx context.Context, dsn string, maxAttempts int, delay time.Duration) (pool *pgxpool.Pool, err error) {
	for maxAttempts > 0 {
		if pool, err = fn(ctx, dsn); err != nil {
			time.Sleep(delay)
			maxAttempts--
			continue
		}
		return pool, nil
	}
	return nil, err
}

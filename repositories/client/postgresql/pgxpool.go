package postgresql

import (
	"context"
	"example/test-items-service/properties"
	"example/test-items-service/repositories/client"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
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

	err = doWithAttempts(func() error {
		ctx, cancel := context.WithTimeout(ctx, AttemptDuration)
		defer cancel()

		pgxCfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			log.Fatalf("Unable to parse config: %v\n", err)
		}

		pool, err = pgxpool.NewWithConfig(ctx, pgxCfg)
		if err != nil {
			log.Println("Failed to connect to postgres... Going to do the next attempt")

			return err
		}

		return nil
	}, maxAttempts, maxDelay)

	if err != nil {
		log.Fatal("All attempts are exceeded. Unable to connect to postgres")
	}

	return pool, nil
}

func doWithAttempts(fn func() error, maxAttempts int, delay time.Duration) error {
	var err error
	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			maxAttempts--
			continue
		}
		return nil
	}
	return err
}

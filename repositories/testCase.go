package repositories

import (
	"context"
	testItems "example/test-items-service/models"
	"example/test-items-service/repositories/client"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
	uuid2 "github.com/gofrs/uuid"
	"github.com/lib/pq"
)

type TestCaseRepository struct {
	client client.Client
	ctx    context.Context
}

func NewTestCaseRepository(ctx context.Context, client client.Client) *TestCaseRepository {
	return &TestCaseRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *TestCaseRepository) FindAll() (testCases []*testItems.TestCase, err error) {
	sql := `SELECT * FROM "test_case"`
	if err = pgxscan.Select(r.ctx, r.client, &testCases, sql); err != nil {
		log.Println(err.Error())
	}
	return testCases, err
}

func (r *TestCaseRepository) FindOneById(uuid uuid2.UUID) (testCase testItems.TestCase, err error) {
	sql := "SELECT * FROM test_case WHERE id = $1"
	log.Println(sql)
	if err = pgxscan.Get(r.ctx, r.client, &testCase, sql, uuid); err != nil {
		log.Println(err.Error())
	}
	return testCase, err
}

func (r *TestCaseRepository) AddOne(tc testItems.TestCase) (err error) {
	sql := "INSERT INTO test_case (name, steps, preconditions, author) VALUES ($1, $2, $3, $4)"
	log.Println(sql)

	if _, err = r.client.Exec(r.ctx, sql, tc.Name, pq.Array(tc.Steps), tc.Preconditions, tc.Author); err != nil {
		log.Println("Inset failed")
	}
	return nil
}

func (r *TestCaseRepository) UpdateOneById(tc testItems.TestCase) (testItems.TestCase, error) {
	sql := "UPDATE test_case SET name=$1, steps=$2, preconditions=$3, author=$4 WHERE id=$5"
	log.Println(sql)
	_, err := r.client.Exec(r.ctx, sql, tc.Name, pq.Array(tc.Steps), tc.Preconditions, tc.Author, tc.ID)
	if err != nil {
		log.Println(err.Error())
	}
	return tc, err
}

func (r *TestCaseRepository) DeleteOneById(uuid uuid2.UUID) (err error) {
	sql := "DELETE FROM test_case WHERE id = $1"
	log.Println(sql)

	if _, err = r.client.Exec(r.ctx, sql, uuid); err != nil {
		return err
	}
	return nil
}

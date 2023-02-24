package repositories

import (
	"context"
	testItems "example/test-items-service/models"
	"example/test-items-service/repositories/client"
	"github.com/georgysavva/scany/v2/pgxscan"
	uuid2 "github.com/gofrs/uuid"
	"log"
)

type TestCaseRepository struct {
	client  client.Client
	context context.Context
}

func NewTestCaseRepository(context context.Context, client client.Client) *TestCaseRepository {
	return &TestCaseRepository{
		client:  client,
		context: context,
	}
}

func (r *TestCaseRepository) FindAll() (testCases []*testItems.TestCase, err error) {
	sql := `SELECT * FROM "test_case"`
	if err = pgxscan.Select(r.context, r.client, &testCases, sql); err != nil {
		log.Println(err.Error())
	}
	return testCases, err
}

func (r *TestCaseRepository) FindOneById(uuid uuid2.UUID) (testCase testItems.TestCase, err error) {
	sql := `SELECT * FROM "test_case" WHERE test_case.id = $1`
	log.Println(sql)
	if err = pgxscan.Get(r.context, r.client, &testCase, sql, uuid); err != nil {
		log.Println(err.Error())
	}
	return testCase, err
}

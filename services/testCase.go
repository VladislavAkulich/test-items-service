package services

import (
	"context"
	testItems "example/test-items-service/models"
	"example/test-items-service/repositories"
	"example/test-items-service/repositories/client/postgresql"
	"github.com/gofrs/uuid"
)

type TestCaseService struct {
	repository repositories.TestCaseRepository
}

// NewTestCaseService TODO need to refactor constructor and do not keep DB initializing here. Factory or smth else.
// Service should be injected in handler from higher level
func NewTestCaseService(ctx context.Context) (service *TestCaseService) {
	service = &TestCaseService{
		repository: *repositories.NewTestCaseRepository(ctx, postgresql.GetInstance(ctx)),
	}
	return service
}

func (r *TestCaseService) FindAll() (testCases []*testItems.TestCase, err error) {
	return r.repository.FindAll()
}

func (r *TestCaseService) FindOneById(id string) (testCase testItems.TestCase, err error) {
	uuidFromString, err := uuid.FromString(id)
	if err != nil {
		return testCase, err
	}
	return r.repository.FindOneById(uuidFromString)
}

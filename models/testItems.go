package testItems

import "github.com/gofrs/uuid"

type TestCase struct {
	ID            uuid.UUID
	Name          string   `json:"name" binding:"required"`
	Steps         []string `json:"steps" binding:"required"`
	Preconditions string   `json:"preconditions" binding:"required"`
	Author        string   `json:"author" binding:"required"`
}

type CheckList struct {
	ID            uuid.UUID
	Name          string
	Steps         []string
	Preconditions string
	Author        string
}

type TestSuite struct {
	ID        uuid.UUID
	Name      string
	TestCases []string
	Author    string
}

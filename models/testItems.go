package testItems

import "github.com/google/uuid"

type TestCase struct {
	ID            uuid.UUID
	Name          string
	Steps         []string
	Preconditions string
	Author        string
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

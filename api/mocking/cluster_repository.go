package mocking

import (
	"github.com/stretchr/testify/mock"
)

type clusterRepository struct {
	mock.Mock
}

func NewClusterRepository() *clusterRepository {
	return &clusterRepository{}
}

func (mock *clusterRepository) GetAllNames() []string {
	args := mock.Called()

	return args.Get(0).([]string)
}

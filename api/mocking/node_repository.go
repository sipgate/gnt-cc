package mocking

import (
	"github.com/stretchr/testify/mock"
	"gnt-cc/model"
)

type nodeRepository struct {
	mock.Mock
}

func NewNodeRepository() *nodeRepository {
	return &nodeRepository{}
}

func (mock *nodeRepository) Get(clusterName string, nodeName string) (model.NodeResult, error) {
	args := mock.Called(clusterName, nodeName)

	return args.Get(0).(model.NodeResult), args.Error(1)
}

func (mock *nodeRepository) GetAll(clusterName string) ([]model.GntNode, error) {
	args := mock.Called(clusterName)

	return args.Get(0).([]model.GntNode), args.Error(1)
}

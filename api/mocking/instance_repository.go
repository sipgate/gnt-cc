package mocking

import (
	"github.com/stretchr/testify/mock"
	"gnt-cc/model"
)

type instanceRepository struct {
	mock.Mock
}

func NewInstanceRepository() *instanceRepository {
	return &instanceRepository{}
}

func (mock *instanceRepository) Get(clusterName string, instanceName string) (model.InstanceResult, error) {
	args := mock.Called(clusterName, instanceName)

	return args.Get(0).(model.InstanceResult), args.Error(1)
}

func (mock *instanceRepository) GetAll(clusterName string) ([]model.GntInstance, error) {
	args := mock.Called(clusterName)

	return args.Get(0).([]model.GntInstance), args.Error(1)
}


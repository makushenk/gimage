package mocks

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type CommonInfrastructure struct{
	mock.Mock
}

func (_m *CommonInfrastructure) NewUUID() uuid.UUID {
	args := _m.Called()
	return args.Get(0).(uuid.UUID)
}

func (_m *CommonInfrastructure) DecodeBase64String(s string) ([]byte, error) {
	args := _m.Called(s)
	return args.Get(0).([]byte), args.Error(1).(error)
}

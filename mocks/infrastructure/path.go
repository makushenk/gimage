package mocks

import "github.com/stretchr/testify/mock"

type PathInfrastructure struct {
	mock.Mock
}

func (_m *PathInfrastructure) Join(elem ...string) string {
	args := _m.Called(elem)
	return args.Get(0).(string)
}

func (_m *PathInfrastructure) Split(path string) (dir, file string) {
	args := _m.Called(path)
	return args.Get(0).(string), args.Get(1).(string)
}

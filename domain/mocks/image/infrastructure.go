package mocks

import "github.com/stretchr/testify/mock"

type ImageInfrastructure struct {
	mock.Mock
}

func (_m *ImageInfrastructure) GenerateThumbnail(source, target string, x, y int, width, height int) error {
	args := _m.Called(source, target, x, y, width, height)
	return args.Error(0).(error)
}
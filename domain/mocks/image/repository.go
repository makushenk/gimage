package mocks

import (
	"context"

	"github.com/makushenk/gimage/domain"

	"github.com/stretchr/testify/mock"
)


type ImageRepository struct {
	mock.Mock
}

func (_m *ImageRepository) GetByID(ctx context.Context, id string) (domain.Image, error) {
	args := _m.Called(ctx, id)
	return args.Get(0).(domain.Image), args.Error(1)
}

func (_m *ImageRepository) Create(ctx context.Context, name string, data []byte) (string, error) {
	args := _m.Called(ctx, name, data)
	return args.Get(0).(string), args.Error(1)
}

func (_m *ImageRepository) Delete(ctx context.Context, ids []string) (int, error) {
	args := _m.Called(ctx, ids)
	return args.Get(0).(int), args.Error(1)
}

package mocks

import (
	"context"

	"github.com/makushenk/gimage/domain"

	"github.com/stretchr/testify/mock"
)

type ImageUsecase struct {
	mock.Mock
}

func (_m *ImageUsecase) Create(ctx context.Context, name string, data []byte) (string, error) {
	args := _m.Called(ctx, name, data)
	return args.Get(0).(string), args.Error(1)
}

func (_m *ImageUsecase) Delete(ctx context.Context, ids []string) error {
	args := _m.Called(ctx, ids)
	return args.Error(0)
}

func (_m *ImageUsecase) GetByID(ctx context.Context, id string) (domain.Image, error) {
	args := _m.Called(ctx, id)
	return args.Get(0).(domain.Image), args.Error(1)
}

func (_m *ImageUsecase) GetThumbnail(ctx context.Context, image *domain.Image) (domain.Image, error) {
	args := _m.Called(ctx, image)
	return args.Get(0).(domain.Image), args.Error(1)
}

func (_m *ImageUsecase) Rotate(ctx context.Context, image *domain.Image, degree int) (domain.Image, error) {
	args := _m.Called(ctx, image, degree)
	return args.Get(0).(domain.Image), args.Error(1)
}

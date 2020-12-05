package usecase

import (
	"context"

	"github.com/makushenk/gimage/boundaries/repository"
)

func (i *imageUsecase) GetByID(ctx context.Context, id string) (boundaries.Image, error) {
	return i.imageRepository.GetByID(ctx, id)
}

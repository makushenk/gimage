package usecase

import (
	"context"

	"github.com/makushenk/gimage/domain"
)

func (i *imageUsecase) GetByID(ctx context.Context, id string) (domain.Image, error) {
	return i.imageRepository.GetByID(ctx, id)
}

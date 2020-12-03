package usecase

import (
	"context"

	"github.com/makushenk/gimage/domain"
)

func (i *imageUsecase) GetThumbnail(ctx context.Context, image *domain.Image) (domain.Image, error) {
	return domain.Image{}, nil
}

package usecase

import (
	"context"

	"github.com/makushenk/gimage/domain"
)

func (i *imageUsecase) Rotate(ctx context.Context, image *domain.Image, degree int) (domain.Image, error) {
	return domain.Image{}, nil
}

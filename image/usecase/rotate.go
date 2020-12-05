package usecase

import (
	"context"

	"github.com/makushenk/gimage/boundaries/repository"
)

func (i *imageUsecase) Rotate(ctx context.Context, image *boundaries.Image, degree int) (boundaries.Image, error) {
	return boundaries.Image{}, nil
}

package boundaries

import (
	"context"

	"github.com/makushenk/gimage/boundaries/repository"
)

type ImageUsecase interface {
	Create(ctx context.Context, name string, data []byte) (boundaries.Image, error)
	Delete(ctx context.Context, ids []string) error
	GetByID(ctx context.Context, id string) (boundaries.Image, error)
	GetThumbnail(ctx context.Context, id string, x, y int, width, height int) (boundaries.Image, error)
	Rotate(ctx context.Context, image *boundaries.Image, degree int) (boundaries.Image, error)
}

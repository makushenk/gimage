package boundaries

import (
	"context"
)

type Image struct {
	ID string
	Path string
	Name string
}

type ImageRepository interface {
	Create(ctx context.Context, name string, data []byte) (Image, error)
	Delete(ctx context.Context, ids []string) (int, error)
	GetByID(ctx context.Context, id string) (Image, error)
}

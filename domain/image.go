package domain

import (
	"context"
)

type Image struct {
	ID			string
	Name		string
	Data 		[]byte
}

type ImageUsecase interface {
	Create(ctx context.Context, name string, data []byte) (Image, error)
	Delete(ctx context.Context, ids []string) error
	GetByID(ctx context.Context, id string) (Image, error)
	GetThumbnail(ctx context.Context, image *Image) (Image, error)
	Rotate(ctx context.Context, image *Image, degree int) (Image, error)
}

type ImageRepository interface {
	Create(ctx context.Context, name string, data []byte) (string, error)
	Delete(ctx context.Context, ids []string) (int, error)
	GetByID(ctx context.Context, id string) (Image, error)
}

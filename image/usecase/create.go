package usecase

import (
	"context"
	"log"

	"github.com/makushenk/gimage/boundaries/repository"
)

func (i *imageUsecase) Create(ctx context.Context, name string,data []byte) (boundaries.Image, error) {
	img, err := i.imageRepository.Create(ctx, name, data)

	if err != nil {
		log.Fatal(err)
		return boundaries.Image{}, err
	}

	return img, nil
}

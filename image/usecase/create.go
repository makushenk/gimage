package usecase

import (
	"context"
	"log"

	"github.com/makushenk/gimage/domain"
)

func (i *imageUsecase) Create(ctx context.Context, name string,data []byte) (domain.Image, error) {
	id, err := i.imageRepository.Create(ctx, name, data)

	if err != nil {
		log.Fatal(err)
		return domain.Image{}, err
	}

	img := domain.Image{
		ID:		id,
		Name:	name,
		Data:	data,
	}

	return img, err
}

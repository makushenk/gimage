package usecase

import (
	"context"
	"github.com/makushenk/gimage/boundaries/repository"
)

func (i *imageUsecase) GetThumbnail(ctx context.Context, id string, x, y int, width, height int) (boundaries.Image, error) {
	sourceImg, err := i.imageRepository.GetByID(ctx, id)

	if err != nil {
		return boundaries.Image{}, err
	}

	previewName := "preview-" + sourceImg.Name

	previewImg, err := i.imageRepository.Create(context.TODO(), previewName, []byte{})

	if err != nil {
		return boundaries.Image{}, err
	}

	i.imageInfrastructure.GenerateThumbnail(sourceImg.Path, previewImg.Path, x, y, width, height)

	return boundaries.Image{}, nil
}

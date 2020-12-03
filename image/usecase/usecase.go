package usecase

import (
	"github.com/makushenk/gimage/domain"
)

type imageUsecase struct {
	imageRepository domain.ImageRepository
}

func NewImageUsecase(i domain.ImageRepository) domain.ImageUsecase {
	return &imageUsecase{
		imageRepository: i,
	}
}

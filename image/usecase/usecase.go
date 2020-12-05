package usecase

import (
	iboundaries "github.com/makushenk/gimage/boundaries/infrastructure"
	rboundaries "github.com/makushenk/gimage/boundaries/repository"
	uboundaries "github.com/makushenk/gimage/boundaries/usecase"
)

type imageUsecase struct {
	imageRepository     rboundaries.ImageRepository
	imageInfrastructure iboundaries.ImageInfrastructure
}

func NewImageUsecase(r rboundaries.ImageRepository, i iboundaries.ImageInfrastructure) uboundaries.ImageUsecase {
	return &imageUsecase{
		imageRepository:     r,
		imageInfrastructure: i,
	}
}
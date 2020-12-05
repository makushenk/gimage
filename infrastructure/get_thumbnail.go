package infrastructure

import "github.com/makushenk/gimage/boundaries/infrastructure"

type imageInfrastructure struct {}

func NewImageInfrastructure() boundaries.ImageInfrastructure {
	return &imageInfrastructure{}
}

func (i *imageInfrastructure) GenerateThumbnail(source, target string, x, y int, width, height int) error {
	return nil
}
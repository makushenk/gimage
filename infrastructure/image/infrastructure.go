package infrastructure

import boundaries "github.com/makushenk/gimage/boundaries/infrastructure"

type imageInfrastructure struct {}

func NewImageInfrastructure() boundaries.ImageInfrastructure {
	return &imageInfrastructure{}
}

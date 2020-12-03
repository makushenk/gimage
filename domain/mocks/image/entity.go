package mocks

import (
	"encoding/base64"
	"log"

	"github.com/makushenk/gimage/domain"
)

type Image struct {}

func (i Image) GetGIFImage() (domain.Image, error) {
	str := "R0lGODlhAQABAIABAP///wAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=="

	data, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		log.Fatal(err)
		return domain.Image{}, err
	}

	image := domain.Image{
		ID:		"gif",
		Name:	"image.gif",
		Data:	data,
	}

	return image, err
}

func (i Image) GetJPGImage() (domain.Image, error) {
	str := "/9j/4AAQSkZJRgABAQEASABIAAD/2wBDAP//////////////////////////////////////////////////////////////////////////////////////wgALCAABAAEBAREA/8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABPxA="

	data, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		log.Fatal(err)
		return domain.Image{}, err
	}

	image := domain.Image{
		ID:		"jpg",
		Name:	"image.jpg",
		Data:	data,
	}

	return image, err
}

func (i Image) GetPNGImage() (domain.Image, error){
	str := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg=="

	data, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		log.Fatal(err)
		return domain.Image{}, err
	}

	image := domain.Image{
		ID:		"png",
		Name:	"image.png",
		Data:	data,
	}

	return image, err
}

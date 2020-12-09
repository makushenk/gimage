package mocks

import "github.com/stretchr/testify/mock"

type ImageInfrastructure struct {
	mock.Mock
}

const (
	GifBase64Data string = "R0lGODlhAQABAIABAP///wAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=="
	JpgBase64Data string = "/9j/4AAQSkZJRgABAQEASABIAAD/2wBDAP//////////////////////////////////////////////////////////////////////////////////////wgALCAABAAEBAREA/8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABPxA="
	PngBase64Data string = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg=="
)

func (_m *ImageInfrastructure) GenerateThumbnail(source, target string, x, y int, width, height int) error {
	args := _m.Called(source, target, x, y, width, height)
	return args.Error(0).(error)
}
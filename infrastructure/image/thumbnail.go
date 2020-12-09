package infrastructure

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func (i *imageInfrastructure) GenerateThumbnail(source, target string, x, y int, width, height int) error {
	imageFile, err := os.Open(source)
	defer imageFile.Close()
	if err != nil {
		return err
	}

	thumbFile, err := os.Create(target)
	defer thumbFile.Close()
	if err != nil {
		return err
	}

	img, ext, err := image.Decode(imageFile)
	if err != nil {
		return err
	}

	thumb := img.(SubImager).SubImage(image.Rect(x, y, x + width, y + height))

	if ext == "gif" {
		err = gif.Encode(thumbFile, thumb, &gif.Options{
			NumColors: 256,
		})
		if err != nil {
			return err
		}
	} else if ext == "jpg" || ext == "jpeg" {
		err = jpeg.Encode(thumbFile, thumb, &jpeg.Options{
			Quality: 90,
		})
		if err != nil {
			return err
		}
	} else if ext == "png" {
		err = png.Encode(thumbFile, thumb)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unsupported file format %s", ext)
	}

	return nil
}
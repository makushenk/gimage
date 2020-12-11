package infrastructure

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDir string = ".test"

const imagesRoot = "../mocks/files/images/"

func TestImageInfrastructure_GenerateThumbnail(t *testing.T) {
	imgInfrastructure := NewImageInfrastructure()
	t.Run("Creating JPG image", func(t *testing.T) {
		jpgImg := imagesRoot + "source.jpg"
		previewJPGImg := imagesRoot + "source-preview.jpg"
		err := imgInfrastructure.GenerateThumbnail(jpgImg, previewJPGImg, 0,0, 200,200)
		assert.NoError(t, err)
		os.Remove(previewJPGImg)
	})

	t.Run("Creating PNG image", func(t *testing.T) {
		pngImg := imagesRoot + "source.png"
		previewPNGImg := imagesRoot + "source-preview.png"
		err := imgInfrastructure.GenerateThumbnail(pngImg, previewPNGImg, 0,0, 200,200)
		assert.NoError(t, err)
		os.Remove(previewPNGImg)
	})

	t.Run("Creating GIF image", func(t *testing.T) {
		gifImg := imagesRoot + "source.gif"
		previewGIFImg := imagesRoot + "source-preview.gif"
		err := imgInfrastructure.GenerateThumbnail(gifImg, previewGIFImg, 0,0, 200,200)
		assert.NoError(t, err)
		os.Remove(previewGIFImg)
	})

}
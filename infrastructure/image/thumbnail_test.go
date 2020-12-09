package infrastructure

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDir string = ".test"

func TestImageInfrastructure_GenerateThumbnail(t *testing.T) {
	sourceGIF := path.Join(testDir, "source.png")
	targetGIF := path.Join(testDir, "target.png")

	//sourceJPG := path.Join(testDir, "source.jpg")
	//targetJPG := path.Join(testDir, "target.jpg")
	//
	//sourceGIF := path.Join(testDir, "source.png")
	//targetGIF := path.Join(testDir, "target.png")

	t.Run("Generating gif thumbnail", func(t *testing.T) {
		inf := NewImageInfrastructure()
		err := inf.GenerateThumbnail(sourceGIF, targetGIF, 0, 0, 400, 400)
		assert.NoError(t, err)
	})
}
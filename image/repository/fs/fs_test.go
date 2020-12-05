package repository

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/makushenk/gimage/domain"
	"github.com/makushenk/gimage/image/repository/fs/utils"

	"github.com/stretchr/testify/assert"
)

const testFsMountPoint string = ".test"

const (
	gifContent string = "R0lGODlhAQABAIABAP///wAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=="
	jpgContent string = "/9j/4AAQSkZJRgABAQEASABIAAD/2wBDAP//////////////////////////////////////////////////////////////////////////////////////wgALCAABAAEBAREA/8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABPxA="
	pngContent string = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg=="
)

func TestFsImageRepository_Create(t *testing.T) {
	const testImagesCount = 3

	testFsImageRepo := NewFsImageRepository(testFsMountPoint)

	t.Run("Creating GIF image", func(t *testing.T) {
		dataGIF, err := base64.StdEncoding.DecodeString(gifContent)
		assert.NoError(t, err)

		img, err := testFsImageRepo.Create(context.TODO(), "image.gif", dataGIF)
		assert.NoError(t, err)

		assert.FileExists(t, img.Path)

		content, err := ioutil.ReadFile(img.Path)
		assert.Equal(t, dataGIF, content)
	})

	t.Run("Creating JPG image", func(t *testing.T) {
		dataJPG, err := base64.StdEncoding.DecodeString(jpgContent)
		assert.NoError(t, err)

		img , err := testFsImageRepo.Create(context.TODO(), "image.jpg", dataJPG)
		assert.NoError(t, err)

		assert.FileExists(t, img.Path)

		content, err := ioutil.ReadFile(img.Path)
		assert.Equal(t, dataJPG, content)
	})

	t.Run("Creating PNG image", func(t *testing.T) {
		dataPNG, err := base64.StdEncoding.DecodeString(pngContent)
		assert.NoError(t, err)

		img , err := testFsImageRepo.Create(context.TODO(), "image.png",dataPNG)
		assert.NoError(t, err)

		assert.FileExists(t, img.Path)

		content, err := ioutil.ReadFile(img.Path)
		assert.Equal(t, dataPNG, content)
	})

	t.Run("Cleaning up", func(t *testing.T) {
		n, err := utils.ClearDir(testFsMountPoint)
		assert.NoError(t, err)
		assert.Equal(t, testImagesCount, n)
	})
}

func TestFsImageRepository_Delete(t *testing.T) {
	testFsImageRepo := NewFsImageRepository(testFsMountPoint)

	t.Run("Deletion without passed ids", func(t *testing.T) {
		var ids []string

		n, err := testFsImageRepo.Delete(context.TODO(), ids)
		assert.NoError(t, err)
		assert.Equal(t, len(ids), n)
	})

	t.Run("Deletion of 1 image", func(t *testing.T) {
		dataGIF, err := base64.StdEncoding.DecodeString(gifContent)
		assert.NoError(t, err)

		img, err := testFsImageRepo.Create(context.TODO(), "image.gif", dataGIF)
		assert.NoError(t, err)

		ids := []string{img.ID}
		n, err := testFsImageRepo.Delete(context.TODO(), ids)
		assert.NoError(t, err)
		assert.Equal(t, len(ids), n)
	})

	t.Run("Deletion of multiple images", func(t *testing.T) {
		dataGIF, err := base64.StdEncoding.DecodeString(gifContent)
		assert.NoError(t, err)
		img1, err := testFsImageRepo.Create(context.TODO(), "image.gif", dataGIF)
		assert.NoError(t, err)

		dataJPG, err := base64.StdEncoding.DecodeString(jpgContent)
		assert.NoError(t,err)
		img2, err := testFsImageRepo.Create(context.TODO(), "image.jpg", dataJPG)
		assert.NoError(t, err)

		dataPNG, err := base64.StdEncoding.DecodeString(pngContent)
		assert.NoError(t,err)
		img3, err := testFsImageRepo.Create(context.TODO(), "image.png", dataPNG)
		assert.NoError(t, err)

		ids := []string{img1.ID, img2.ID, img3.ID}
		n, err := testFsImageRepo.Delete(context.TODO(), ids)
		assert.NoError(t, err)
		assert.Equal(t, len(ids), n)
	})

	t.Run("Cleaning up", func(t *testing.T) {
		n, err := utils.ClearDir(testFsMountPoint)
		assert.NoError(t, err)
		assert.Equal(t, 0, n)
	})
}

func TestFsImageRepository_GetByID(t *testing.T) {
	testFsImageRepo := NewFsImageRepository(testFsMountPoint)

	t.Run("Getting image by existing id", func(t *testing.T) {
		dataGIF, err := base64.StdEncoding.DecodeString(gifContent)
		assert.NoError(t, err)

		img, err := testFsImageRepo.Create(context.TODO(), "image.gif", dataGIF)
		assert.NoError(t, err)

		gotImg, err := testFsImageRepo.GetByID(context.TODO(), img.ID)
		assert.NoError(t, err)
		assert.Equal(t, img.Name, gotImg.Name)
		// TODO: read file and assert data equals
	})

	t.Run("Getting image by not existing id", func(t *testing.T) {
		id := "non-existing-id"
		img, err := testFsImageRepo.GetByID(context.TODO(), id)

		assert.Empty(t, domain.Image{}, img)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		n, err := utils.ClearDir(testFsMountPoint)
		assert.NoError(t, err)
		assert.Equal(t, 1, n)
	})
}
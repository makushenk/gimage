package repository

import (
	"context"
	"io/ioutil"
	"path"
	"testing"

	"github.com/makushenk/gimage/domain"
	"github.com/makushenk/gimage/domain/mocks/image"
	"github.com/makushenk/gimage/image/repository/fs/utils"

	"github.com/stretchr/testify/assert"
)

const testFsMountPoint string = ".test"

func TestFsImageRepository_Create(t *testing.T) {
	const testImagesCount = 3

	testFsImageRepo := NewFsImageRepository(testFsMountPoint)

	t.Run("Creating GIF image", func(t *testing.T) {
		img, err := mocks.Image{}.GetGIFImage()
		assert.NoError(t, err)

		id , err := testFsImageRepo.Create(context.TODO(), img.Name, img.Data)
		assert.NoError(t, err)

		file := path.Join(testFsMountPoint, id, img.Name)
		assert.FileExists(t, file)

		content, err := ioutil.ReadFile(file)
		assert.Equal(t, img.Data, content)
	})

	t.Run("Creating JPG image", func(t *testing.T) {
		img, err := mocks.Image{}.GetJPGImage()
		assert.NoError(t, err)

		id , err := testFsImageRepo.Create(context.TODO(), img.Name, img.Data)
		assert.NoError(t, err)

		file := path.Join(testFsMountPoint, id, img.Name)
		assert.FileExists(t, file)

		content, err := ioutil.ReadFile(file)
		assert.Equal(t, img.Data, content)
	})

	t.Run("Creating PNG image", func(t *testing.T) {
		img, err := mocks.Image{}.GetPNGImage()
		assert.NoError(t, err)

		id , err := testFsImageRepo.Create(context.TODO(), img.Name, img.Data)
		assert.NoError(t, err)

		file := path.Join(testFsMountPoint, id, img.Name)
		assert.FileExists(t, file)

		content, err := ioutil.ReadFile(file)
		assert.Equal(t, img.Data, content)
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
		img, err := mocks.Image{}.GetGIFImage()
		assert.NoError(t, err)

		id, err := testFsImageRepo.Create(context.TODO(), img.Name, img.Data)
		assert.NoError(t, err)

		ids := []string{id}
		n, err := testFsImageRepo.Delete(context.TODO(), ids)
		assert.NoError(t, err)
		assert.Equal(t, len(ids), n)
	})

	t.Run("Deletion of multiple images", func(t *testing.T) {
		img, err := mocks.Image{}.GetGIFImage()
		assert.NoError(t, err)
		id1, err := testFsImageRepo.Create(context.TODO(), img.Name, img.Data)
		assert.NoError(t, err)

		img2, err := mocks.Image{}.GetJPGImage()
		assert.NoError(t,err)
		id2, err := testFsImageRepo.Create(context.TODO(), img2.Name, img2.Data)
		assert.NoError(t, err)

		img3, err := mocks.Image{}.GetPNGImage()
		assert.NoError(t,err)
		id3, err := testFsImageRepo.Create(context.TODO(), img3.Name, img3.Data)
		assert.NoError(t, err)

		ids := []string{id1, id2, id3}
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
		img, err := mocks.Image{}.GetGIFImage()
		assert.NoError(t, err)

		id, err := testFsImageRepo.Create(context.TODO(), img.Name, img.Data)
		assert.NoError(t, err)

		gotImg, err := testFsImageRepo.GetByID(context.TODO(), id)
		assert.NoError(t, err)
		assert.Equal(t, img.Name, gotImg.Name)
		assert.Equal(t, img.Data, gotImg.Data)
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
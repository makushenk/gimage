package repository

import (
	"context"
	"errors"
	"testing"

	boundaries "github.com/makushenk/gimage/boundaries/repository"
	"github.com/makushenk/gimage/domain"
	"github.com/makushenk/gimage/infrastructure"
	mockimg "github.com/makushenk/gimage/mocks/image"
	mockinf "github.com/makushenk/gimage/mocks/infrastructure"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const testFsMountPoint string = ".test"

var (
	commonInfrastructure = infrastructure.NewCommonInfrastructure()
	osInfrastructure = infrastructure.NewOsInfrastructure()
	pathInfrastructure = infrastructure.NewPathInfrastructure()
)

func TestFsImageRepository_Create(t *testing.T) {
	const testImagesCount = 3

	testFsImageRepo := NewFsImageRepository(testFsMountPoint, commonInfrastructure, osInfrastructure, pathInfrastructure)

	t.Run("Creating GIF image", func(t *testing.T) {
		dataGIF, err := commonInfrastructure.DecodeBase64String(mockimg.GifMock)
		assert.NoError(t, err)

		img, err := testFsImageRepo.Create(context.TODO(), "image.gif", dataGIF)
		assert.NoError(t, err)

		assert.FileExists(t, img.Path)

		content, err := osInfrastructure.ReadFile(img.Path)
		assert.Equal(t, dataGIF, content)
	})

	t.Run("Creating JPG image", func(t *testing.T) {
		dataJPG, err := commonInfrastructure.DecodeBase64String(mockimg.JpgMock)
		assert.NoError(t, err)

		img , err := testFsImageRepo.Create(context.TODO(), "image.jpg", dataJPG)
		assert.NoError(t, err)

		assert.FileExists(t, img.Path)

		content, err := osInfrastructure.ReadFile(img.Path)
		assert.Equal(t, dataJPG, content)
	})

	t.Run("Creating PNG image", func(t *testing.T) {
		dataPNG, err := commonInfrastructure.DecodeBase64String(mockimg.PngMock)
		assert.NoError(t, err)

		img , err := testFsImageRepo.Create(context.TODO(), "image.png",dataPNG)
		assert.NoError(t, err)

		assert.FileExists(t, img.Path)

		content, err := osInfrastructure.ReadFile(img.Path)
		assert.Equal(t, dataPNG, content)
	})

	t.Run("Creating image with making dir error", func(t *testing.T) {
		osInfrastructureMock := new(mockinf.OsInfrastructure)
		testFsImageRepo = NewFsImageRepository(testFsMountPoint, commonInfrastructure, osInfrastructureMock, pathInfrastructure)
		osInfrastructureMock.On(
			"MkdirAll",
			mock.AnythingOfType("string"),
			mock.Anything,
		).Return(errors.New("mkdirall error")).Once()

		dataGIF, err := commonInfrastructure.DecodeBase64String(mockimg.GifMock)
		assert.NoError(t, err)

		img, err := testFsImageRepo.Create(context.TODO(), "image.gif", dataGIF)
		assert.Empty(t, boundaries.Image{}, img)
		assert.Error(t, err)
	})

	t.Run("Creating image with writing file error", func(t *testing.T) {
		osInfrastructureMock := new(mockinf.OsInfrastructure)
		testFsImageRepo = NewFsImageRepository(testFsMountPoint, commonInfrastructure, osInfrastructureMock, pathInfrastructure)
		osInfrastructureMock.On(
			"MkdirAll",
			mock.AnythingOfType("string"),
			mock.Anything,
		).Return(nil).Once()
		osInfrastructureMock.On(
			"WriteFile",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("[]uint8"),
			mock.Anything,
		).Return(errors.New("writefile error")).Once()

		dataGIF, err := commonInfrastructure.DecodeBase64String(mockimg.GifMock)
		assert.NoError(t, err)

		img, err := testFsImageRepo.Create(context.TODO(), "image.gif", dataGIF)
		assert.Empty(t, boundaries.Image{}, img)
		assert.Error(t, err)
	})


	t.Run("Cleaning up", func(t *testing.T) {
		n, err := osInfrastructure.ClearDir(testFsMountPoint)
		assert.NoError(t, err)
		assert.Equal(t, testImagesCount, n)
	})
}

func TestFsImageRepository_Delete(t *testing.T) {
	testFsImageRepo := NewFsImageRepository(testFsMountPoint, commonInfrastructure, osInfrastructure, pathInfrastructure)

	t.Run("Deletion without passed ids", func(t *testing.T) {
		var ids []string

		n, err := testFsImageRepo.Delete(context.TODO(), ids)
		assert.NoError(t, err)
		assert.Equal(t, len(ids), n)
	})

	t.Run("Deletion of 1 image", func(t *testing.T) {
		dataGIF, err := commonInfrastructure.DecodeBase64String(mockimg.GifMock)
		assert.NoError(t, err)

		img, err := testFsImageRepo.Create(context.TODO(), "image.gif", dataGIF)
		assert.NoError(t, err)

		ids := []string{img.ID}
		n, err := testFsImageRepo.Delete(context.TODO(), ids)
		assert.NoError(t, err)
		assert.Equal(t, len(ids), n)
	})

	t.Run("Deletion of multiple images", func(t *testing.T) {
		dataGIF, err := commonInfrastructure.DecodeBase64String(mockimg.GifMock)
		assert.NoError(t, err)
		img1, err := testFsImageRepo.Create(context.TODO(), "image.gif", dataGIF)
		assert.NoError(t, err)

		dataJPG, err := commonInfrastructure.DecodeBase64String(mockimg.JpgMock)
		assert.NoError(t,err)
		img2, err := testFsImageRepo.Create(context.TODO(), "image.jpg", dataJPG)
		assert.NoError(t, err)

		dataPNG, err := commonInfrastructure.DecodeBase64String(mockimg.PngMock)
		assert.NoError(t,err)
		img3, err := testFsImageRepo.Create(context.TODO(), "image.png", dataPNG)
		assert.NoError(t, err)

		ids := []string{img1.ID, img2.ID, img3.ID}
		n, err := testFsImageRepo.Delete(context.TODO(), ids)
		assert.NoError(t, err)
		assert.Equal(t, len(ids), n)
	})

	t.Run("Deletion with removing all error", func(t *testing.T) {
		osInfrastructureMock := new(mockinf.OsInfrastructure)
		testFsImageRepo = NewFsImageRepository(testFsMountPoint, commonInfrastructure, osInfrastructureMock, pathInfrastructure)
		osInfrastructureMock.On(
			"MkdirAll",
			mock.AnythingOfType("string"),
			mock.Anything,
		).Return(nil).Once()
		osInfrastructureMock.On(
			"WriteFile",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("[]uint8"),
			mock.Anything,
		).Return(nil).Once()
		osInfrastructureMock.On(
			"RemoveAll",
			mock.AnythingOfType("string"),
		).Return(errors.New("remove all error")).Once()

		dataGIF, err := commonInfrastructure.DecodeBase64String(mockimg.GifMock)
		assert.NoError(t, err)

		img, err := testFsImageRepo.Create(context.TODO(), "image.gif", dataGIF)
		assert.NoError(t, err)

		ids := []string{img.ID}

		n, err := testFsImageRepo.Delete(context.TODO(), ids)
		assert.Error(t, err)
		assert.Equal(t, 0, n)
	})

	t.Run("Cleaning up", func(t *testing.T) {
		n, err := osInfrastructure.ClearDir(testFsMountPoint)
		assert.NoError(t, err)
		assert.Equal(t, 0, n)
	})
}

func TestFsImageRepository_GetByID(t *testing.T) {
	const testImagesCount = 2

	testFsImageRepo := NewFsImageRepository(testFsMountPoint, commonInfrastructure, osInfrastructure, pathInfrastructure)

	t.Run("Getting image by existing id", func(t *testing.T) {
		dataGIF, err := commonInfrastructure.DecodeBase64String(mockimg.GifMock)
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

	t.Run("Get image with get first file error", func(t *testing.T) {
		dataGIF, err := commonInfrastructure.DecodeBase64String(mockimg.GifMock)
		assert.NoError(t, err)

		img, err := testFsImageRepo.Create(context.TODO(), "image.gif", dataGIF)
		assert.NoError(t, err)

		osInfrastructure.Chmod(testFsMountPoint, 0000)

		gotImg, err := testFsImageRepo.GetByID(context.TODO(), img.ID)
		assert.Error(t, err)
		assert.Empty(t, boundaries.Image{}, gotImg)

		osInfrastructure.Chmod(testFsMountPoint, 0777)
	})

	t.Run("Cleaning up", func(t *testing.T) {
		n, err := osInfrastructure.ClearDir(testFsMountPoint)
		assert.NoError(t, err)
		assert.Equal(t, testImagesCount, n)
	})
}
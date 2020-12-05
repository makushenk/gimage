package usecase

import (
	"context"
	boundaries "github.com/makushenk/gimage/boundaries/repository"
	"testing"

	mocks "github.com/makushenk/gimage/domain/mocks/image"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestImageUsecase_GetByID(t *testing.T) {
	mockImageRepo :=  new(mocks.ImageRepository)
	mockImageInfrastructure := new(mocks.ImageInfrastructure)
	mockImage := boundaries.Image{
		ID: 		"mockID",
		Name:		"mockName",
		Path:		"",
	}

	t.Run("Getting image by id", func(t *testing.T) {
		mockImageRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(mockImage, nil).Once()

		i, err := NewImageUsecase(mockImageRepo, mockImageInfrastructure).GetByID(context.TODO(), mockImage.ID)

		assert.NoError(t, err)
		assert.NotNil(t, i)
		assert.Equal(t, mockImage.ID, i.ID)
		mockImageRepo.AssertExpectations(t)
	})

}

func TestImageUsecase_Create(t *testing.T) {
	mockImageRepo := new(mocks.ImageRepository)
	mockImageInfrastructure := new(mocks.ImageInfrastructure)
	mockImage := boundaries.Image{
		ID:			"mockID",
		Name:		"mockName",
		Path:		"",
	}

	t.Run("Creating image", func(t *testing.T) {
		// TODO: fix AnythingOfType("[]uint8") from uint to byte
		mockImageRepo.On(
			"Create",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("[]uint8"),
		).Return(mockImage, nil).Once()

		i, err := NewImageUsecase(mockImageRepo, mockImageInfrastructure).Create(context.TODO(), mockImage.Name, []byte{})

		assert.NoError(t, err)
		assert.NotNil(t, i)
		mockImageRepo.AssertExpectations(t)
	})
}

func TestImageUsecase_Delete(t *testing.T) {
	mockImageRepo := new(mocks.ImageRepository)
	mockImageInfrastructure := new(mocks.ImageInfrastructure)

	t.Run("Deletion of an image", func(t *testing.T) {
		mockID := "mockID"

		mockImageRepo.On(
			"Delete",
			mock.Anything,
			mock.AnythingOfType("[]string"),
		).Return(1, nil).Once()

		err := NewImageUsecase(mockImageRepo, mockImageInfrastructure).Delete(context.TODO(), []string{mockID})

		assert.NoError(t, err)
		mockImageRepo.AssertExpectations(t)
	})

	t.Run("Not completed deletion", func(t *testing.T) {
		mockID1 := "mockID1"
		mockID2 := "mockID2"

		mockImageRepo.On(
			"Delete",
			mock.Anything,
			mock.AnythingOfType("[]string"),
		).Return(1, nil).Once()

		err := NewImageUsecase(mockImageRepo, mockImageInfrastructure).Delete(context.TODO(), []string{mockID1, mockID2})

		assert.Error(t, err)
		mockImageRepo.AssertExpectations(t)
	})
}

func TestImageUsecase_GetThumbnail(t *testing.T) {

}

func TestImageUsecase_Rotate(t *testing.T) {

}

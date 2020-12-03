package usecase

import (
	"context"
	"testing"

	"github.com/makushenk/gimage/domain"
	mocks "github.com/makushenk/gimage/domain/mocks/image"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestImageUsecase_GetByID(t *testing.T) {
	mockImageRepo :=  new(mocks.ImageRepository)
	mockImage := domain.Image{
		ID: 		"mockID",
		Name:		"mockName",
		Data:		[]byte{1,0,0,0},
	}

	t.Run("Getting image by id", func(t *testing.T) {
		mockImageRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(mockImage, nil).Once()

		i, err := NewImageUsecase(mockImageRepo).GetByID(context.TODO(), mockImage.ID)

		assert.NoError(t, err)
		assert.NotNil(t, i)
		mockImageRepo.AssertExpectations(t)
	})

}

func TestImageUsecase_Create(t *testing.T) {
	mockImageRepo := new(mocks.ImageRepository)
	mockImage := domain.Image{
		ID:			"mockID",
		Name:		"mockName",
		Data:		[]byte{},
	}

	t.Run("Creating image", func(t *testing.T) {
		// TODO: fix AnythingOfType("[]uint8") from uint to byte
		mockImageRepo.On(
			"Create",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("[]uint8"),
		).Return(mockImage.ID, nil).Once()

		i, err := NewImageUsecase(mockImageRepo).Create(context.TODO(), mockImage.Name, mockImage.Data)

		assert.NoError(t, err)
		assert.NotNil(t, i)
		mockImageRepo.AssertExpectations(t)
	})
}

func TestImageUsecase_Delete(t *testing.T) {
	mockImageRepo := new(mocks.ImageRepository)

	t.Run("Deletion of an image", func(t *testing.T) {
		mockID := "mockID"

		mockImageRepo.On(
			"Delete",
			mock.Anything,
			mock.AnythingOfType("[]string"),
		).Return(1, nil).Once()

		err := NewImageUsecase(mockImageRepo).Delete(context.TODO(), []string{mockID})

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

		err := NewImageUsecase(mockImageRepo).Delete(context.TODO(), []string{mockID1, mockID2})

		assert.Error(t, err)
		mockImageRepo.AssertExpectations(t)
	})
}

func TestImageUsecase_GetThumbnail(t *testing.T) {

}

func TestImageUsecase_Rotate(t *testing.T) {

}

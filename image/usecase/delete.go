package usecase

import (
	"context"
	"fmt"
)

func (i *imageUsecase) Delete(ctx context.Context, ids []string) error {
	deletedAmount, err := i.imageRepository.Delete(ctx, ids)

	if err != nil {
		return err
	}

	if deletedAmount != len(ids) {
		return fmt.Errorf("%d images were not deleted", len(ids) - deletedAmount)
	}

	return nil
}


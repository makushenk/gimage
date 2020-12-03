package repository

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/makushenk/gimage/domain"
	"github.com/makushenk/gimage/image/repository/fs/utils"

	"github.com/google/uuid"
)

type fsImageRepository struct {
	MountPoint	string
}

func NewFsImageRepository(mountPoint string) domain.ImageRepository {
	return &fsImageRepository{
		MountPoint: mountPoint,
	}
}

func (f *fsImageRepository) Create(ctx context.Context, name string, data []byte) (string, error) {
	id := uuid.New().String()
	dir := path.Join(f.MountPoint, id)
	file := path.Join(dir, name)

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	err = ioutil.WriteFile(file, data, os.ModePerm)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return id, nil
}

func (f *fsImageRepository) Delete(ctx context.Context, ids []string) (int, error) {
	for i, id := range ids {
		dir := path.Join(f.MountPoint, id)

		err := os.RemoveAll(dir)

		if err != nil {
			log.Fatal(err)
			return i, err
		}

	}

	return len(ids), nil
}

func (f *fsImageRepository) GetByID(ctx context.Context, id string) (domain.Image, error) {
	dir := path.Join(f.MountPoint, id)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return domain.Image{}, fmt.Errorf("image with id %s doesn't exists", id)
	}

	file, err := utils.GetFirstFile(dir)

	if err != nil {
		log.Fatal(err)
		return domain.Image{}, err
	}

	content, err := ioutil.ReadFile(file)
	_, file = path.Split(file)

	if err != nil {
		log.Fatal(err)
		return domain.Image{}, err
	}

	image := domain.Image{
		ID: id,
		Name: file,
		Data: content,
	}

	return image, nil
}

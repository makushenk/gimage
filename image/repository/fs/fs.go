package repository

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/makushenk/gimage/boundaries/repository"
	"github.com/makushenk/gimage/image/repository/fs/utils"

	"github.com/google/uuid"
)

type fsImageRepository struct {
	MountPoint	string
}

func NewFsImageRepository(mountPoint string) boundaries.ImageRepository {
	return &fsImageRepository{
		MountPoint: mountPoint,
	}
}

func (f *fsImageRepository) Create(ctx context.Context, name string, data []byte) (boundaries.Image, error) {
	id := uuid.New().String()
	dir := path.Join(f.MountPoint, id)
	file := path.Join(dir, name)

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		log.Fatal(err)
		return boundaries.Image{}, err
	}

	err = ioutil.WriteFile(file, data, os.ModePerm)

	if err != nil {
		log.Fatal(err)
		return boundaries.Image{}, err
	}

	img := boundaries.Image{
		ID:	id,
		Name: name,
		Path: file,
	}
	return img, nil
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

func (f *fsImageRepository) GetByID(ctx context.Context, id string) (boundaries.Image, error) {
	dir := path.Join(f.MountPoint, id)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return boundaries.Image{}, fmt.Errorf("image with id %s doesn't exists", id)
	}

	file, err := utils.GetFirstFile(dir)

	if err != nil {
		log.Fatal(err)
		return boundaries.Image{}, err
	}

	_, name := path.Split(file)

	if err != nil {
		log.Fatal(err)
		return boundaries.Image{}, err
	}

	image := boundaries.Image{
		Path: file,
		Name: name,
	}

	return image, nil
}

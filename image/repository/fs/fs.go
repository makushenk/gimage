package repository

import (
	"context"
	"fmt"
	"os"

	iboundaries "github.com/makushenk/gimage/boundaries/infrastructure"
	rboundaries "github.com/makushenk/gimage/boundaries/repository"
)

type fsImageRepository struct {
	MountPoint				string
	CommonInfrastructure	iboundaries.CommonInfrastructure
	PathInfrastructure		iboundaries.PathInfrastructure
	OsInfrastructure		iboundaries.OsInfrastructure
}

func NewFsImageRepository(
		mountPoint string,
		commonInf iboundaries.CommonInfrastructure,
		osInf iboundaries.OsInfrastructure,
		pathInf iboundaries.PathInfrastructure,
	) rboundaries.ImageRepository {

	return &fsImageRepository{
		MountPoint: mountPoint,
		CommonInfrastructure: commonInf,
		OsInfrastructure: osInf,
		PathInfrastructure: pathInf,
	}
}

func (f *fsImageRepository) Create(ctx context.Context, name string, data []byte) (rboundaries.Image, error) {
	id := f.CommonInfrastructure.NewUUID().String()
	dir := f.PathInfrastructure.Join(f.MountPoint, id)
	file := f.PathInfrastructure.Join(dir, name)

	err := f.OsInfrastructure.MkdirAll(dir, os.ModePerm)

	if err != nil {
		return rboundaries.Image{}, err
	}

	err = f.OsInfrastructure.WriteFile(file, data, os.ModePerm)

	if err != nil {
		return rboundaries.Image{}, err
	}

	img := rboundaries.Image{
		ID:	id,
		Name: name,
		Path: file,
	}
	return img, nil
}

func (f *fsImageRepository) Delete(ctx context.Context, ids []string) (int, error) {
	for i, id := range ids {
		dir := f.PathInfrastructure.Join(f.MountPoint, id)

		err := f.OsInfrastructure.RemoveAll(dir)

		if err != nil {
			return i, err
		}
	}

	return len(ids), nil
}

func (f *fsImageRepository) GetByID(ctx context.Context, id string) (rboundaries.Image, error) {
	dir := f.PathInfrastructure.Join(f.MountPoint, id)

	if _, err := f.OsInfrastructure.Stat(dir); f.OsInfrastructure.IsNotExist(err) {
		return rboundaries.Image{}, fmt.Errorf("image with id %s doesn't exists", id)
	}

	file, err := f.OsInfrastructure.GetFirstFile(dir)

	if err != nil {
		return rboundaries.Image{}, err
	}

	_, name := f.PathInfrastructure.Split(file)

	image := rboundaries.Image{
		Path: file,
		Name: name,
	}

	return image, nil
}

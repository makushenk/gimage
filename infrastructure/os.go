package infrastructure

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	boundaries "github.com/makushenk/gimage/boundaries/infrastructure"
)

type osInfrastructure struct {}

func NewOsInfrastructure() boundaries.OsInfrastructure {
	return &osInfrastructure{}
}

func (o *osInfrastructure) Chmod (name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

func (o *osInfrastructure) ClearDir(dir string) (int, error) {
	d, err := os.Open(dir)

	if err != nil {
		return 0, err
	}

	defer d.Close()

	names, err := d.Readdirnames(-1)

	if err != nil {
		return 0, err
	}

	for i, name := range names {
		err = os.RemoveAll(path.Join(dir, name))

		if err != nil {
			return i, err
		}
	}

	return len(names), nil
}

func (o *osInfrastructure) GetFirstFile(dir string) (string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		return "", err
	}

	if len(files) < 2 {
		return "", errors.New("the directory is empty")
	}

	return files[1], nil
}

func (o *osInfrastructure) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (o *osInfrastructure) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (o *osInfrastructure) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (o *osInfrastructure) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (o *osInfrastructure) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (o *osInfrastructure) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

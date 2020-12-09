package mocks

import (
	"os"

	"github.com/stretchr/testify/mock"
)

type OsInfrastructure struct {
	mock.Mock
}

func (_m *OsInfrastructure) Chmod(name string, mode os.FileMode) error {
	args := _m.Called(name, mode)
	return args.Error(0)
}

func (_m *OsInfrastructure) ClearDir(dir string) (int, error) {
	args := _m.Called(dir)
	return args.Get(0).(int), args.Error(1)
}

func (_m *OsInfrastructure) GetFirstFile(dir string) (string, error) {
	args := _m.Called(dir)
	return args.Get(0).(string), args.Error(1)
}

func (_m *OsInfrastructure) IsNotExist(err error) bool {
	args := _m.Called(err)
	return args.Get(0).(bool)
}

func (_m *OsInfrastructure) MkdirAll(path string, perm os.FileMode) error {
	args := _m.Called(path, perm)
	return args.Error(0)
}

func (_m *OsInfrastructure) ReadFile(filename string) ([]byte, error) {
	args := _m.Called(filename)
	return args.Get(0).([]byte), args.Error(1)
}

func (_m *OsInfrastructure) RemoveAll(path string) error {
	args := _m.Called(path)
	return args.Error(0)
}

func (_m *OsInfrastructure) Stat(name string) (os.FileInfo, error) {
	args := _m.Called(name)
	return args.Get(0).(os.FileInfo), args.Error(1)
}

func (_m *OsInfrastructure) WriteFile(filename string, data []byte, perm os.FileMode) error {
	args := _m.Called(filename, data, perm)
	return args.Error(0)
}

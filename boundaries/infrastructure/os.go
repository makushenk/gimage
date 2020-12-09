package boundaries

import "os"

type OsInfrastructure interface {
	Chmod(name string, mode os.FileMode) error
	ClearDir(dir string) (int, error)
	GetFirstFile(dir string) (string, error)
	IsNotExist(err error) bool
	MkdirAll(path string, perm os.FileMode) error
	ReadFile(filename string) ([]byte, error)
	RemoveAll(path string) error
	Stat(name string) (os.FileInfo, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

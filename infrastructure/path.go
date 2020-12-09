package infrastructure

import (
	boundaries "github.com/makushenk/gimage/boundaries/infrastructure"
	"path"
	"path/filepath"
)

type pathInfrastructure struct {}

func NewPathInfrastructure() boundaries.PathInfrastructure {
	return &pathInfrastructure{}
}

func (p *pathInfrastructure) Join(elem... string) string {
	return path.Join(elem...)
}

func (p *pathInfrastructure) Split(path string) (dir, file string) {
	return filepath.Split(path)
}
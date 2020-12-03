package utils

import (
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
)

func GetFirstFile(dir string) (string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if len(files) < 2 {
		return "", errors.New("the directory is empty")
	}

	return files[1], nil
}

func ClearDir(dir string) (int, error) {
	d, err := os.Open(dir)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	defer d.Close()

	names, err := d.Readdirnames(-1)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	for i, name := range names {
		err = os.RemoveAll(path.Join(dir, name))

		if err != nil {
			log.Fatal(err)
			return i, err
		}
	}

	return len(names), nil
}
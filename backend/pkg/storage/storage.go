package storage

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
)

const baseUrl = "./static"

type FileStorage struct {
	path string
}

type SaveFunc func(file *multipart.FileHeader, dst string) error

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{path: path}
}

func (s *FileStorage) SaveFile(save SaveFunc, file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	var fileExt string
	detectedFileType := http.DetectContentType(b)
	switch detectedFileType {
	case "image/jpeg":
		fileExt = "jpeg"
	case "image/jpg":
		fileExt = "jpg"
	case "image/png":
		fileExt = "png"
	default:
		return "", fmt.Errorf("invalid file MIME type: %s", detectedFileType)
	}

	file.Filename = uuid.New().String() + "." + fileExt

	return file.Filename, save(file, s.path+"/"+file.Filename)
}

func (s *FileStorage) GetFileLink(name string) string {
	return baseUrl + "/" + name
}

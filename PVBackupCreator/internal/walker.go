package internal

import (
	"os"
	"path/filepath"
	"strings"
)

type FileWalker struct {
	RootDir string
}

var photoOrVideoExtension = [...]string{
	"png",
	"PNG",
	"jpg",
	"JPG",
	"jpeg",
	"raw",
	"RAW",
	"mp4",
	"MP4",
	"mov",
	"MOV",
	"avi",
	"AVI",
	"mkv",
	"MKV",
}

func (w FileWalker) GetAllFilesPaths() []string {
	var allFilePaths []string

	err := filepath.Walk(w.RootDir, func(wPath string, info os.FileInfo, err error) error {
		if wPath == w.RootDir {
			return nil
		}

		if wPath != w.RootDir && !info.IsDir() {
			if w.isPhotoOrVideoFile(wPath) {
				allFilePaths = append(allFilePaths, wPath)
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
	return allFilePaths
}

func (w FileWalker) isPhotoOrVideoFile(path string) bool {
	for _, ext := range photoOrVideoExtension {
		if strings.Contains(path, ext) {
			return true
		}
	}
	return false
}

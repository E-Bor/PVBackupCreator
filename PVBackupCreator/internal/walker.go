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

func (w FileWalker) GetAllFilesPaths(ch *chan string) {
	defer close(*ch)

	err := filepath.Walk(w.RootDir, func(wPath string, info os.FileInfo, err error) error {
		if wPath == w.RootDir {
			return nil
		}

		if wPath != w.RootDir && !info.IsDir() {
			if w.isPhotoOrVideoFile(wPath) {
				*ch <- wPath
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func (w FileWalker) isPhotoOrVideoFile(path string) bool {
	for _, ext := range photoOrVideoExtension {
		if strings.Contains(path, ext) {
			return true
		}
	}
	return false
}

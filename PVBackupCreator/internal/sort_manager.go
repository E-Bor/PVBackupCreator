package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type SortManager struct {
	RootPathToGrouping string
}

const pathTemplate = "%s/%s/%s/"

func (s SortManager) CreateGroupingPath(OriginalPath string) (string, string) {
	year, month := s.assignFolder(s.extractCreationDate(OriginalPath))
	fileName := filepath.Base(OriginalPath)
	return fmt.Sprintf(pathTemplate, s.RootPathToGrouping, year, month), fileName
}

func (s SortManager) extractCreationDate(OriginalPath string) time.Time {
	fileInfo, err := os.Stat(OriginalPath)
	if err != nil {
		panic(err)
	}
	modTime := fileInfo.ModTime()

	return modTime
}

func (s SortManager) assignFolder(FileChangedDate time.Time) (string, string) {
	return strconv.Itoa(FileChangedDate.Year()), strconv.Itoa(int(FileChangedDate.Month()))
}

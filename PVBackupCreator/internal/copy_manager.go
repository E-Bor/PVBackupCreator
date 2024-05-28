package internal

import (
	utils "awesomeProject/pkg/utils"
	"fmt"
	"os"
	"sync"
)

type CopyManager struct {
	walker      FileWalker
	hashChecker HashChecker
	sortManager SortManager
}

func NewCopyManager(rootToSort string, rootToCopy string) *CopyManager {
	hashFileName := "hashes.json"
	copyRoot := fmt.Sprintf("%s/%s", rootToCopy, "dump")

	return &CopyManager{
		walker:      FileWalker{RootDir: rootToSort},
		hashChecker: NewHashChecker(copyRoot, hashFileName),
		sortManager: SortManager{RootPathToGrouping: copyRoot},
	}
}

func (c CopyManager) SortAndCopyFiles() {
	fmt.Println("Copying files...")
	pathsChan := make(chan string, 3)

	var wg sync.WaitGroup

	go c.walker.GetAllFilesPaths(&pathsChan)

	for path := range pathsChan {
		go c.copyIfNeededWorker(path, &wg)
		wg.Add(1)
	}
	wg.Wait()

	defer c.hashChecker.SaveHashes()
}

func (c CopyManager) copyIfNeededWorker(
	filePath string,
	wg *sync.WaitGroup,
) {
	if c.hashChecker.CheckOrCreateHash(filePath) {
		dirToCopy, FileName := c.sortManager.CreateGroupingPath(filePath)
		_, err := os.Stat(dirToCopy)
		copyPath := fmt.Sprintf("%s/%s", dirToCopy, FileName)
		if err != nil {
			err = os.MkdirAll(dirToCopy, os.ModePerm)
			if err != nil {
				panic(err)
			}
			utils.CopyFile(filePath, copyPath)
		} else {
			utils.CopyFile(filePath, copyPath)
		}
	}
	defer wg.Done()
}

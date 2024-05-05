package internal

import (
	utils2 "awesomeProject/pkg/utils"
	"fmt"
	"os"
	"sync"
	"time"
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
	totalFileCounter := len(c.walker.GetAllFilesPaths())
	copiedFilesCounter := 0

	isDoneChan := make(chan bool)
	defer close(isDoneChan)
	var counterMutex sync.Mutex

	var wg sync.WaitGroup
	go c.watchDog(isDoneChan, &copiedFilesCounter, totalFileCounter)
	filePathsGroups := utils2.SplitForNGroups(c.walker.GetAllFilesPaths(), 3)
	isDoneChan <- false

	wg.Add(len(filePathsGroups))

	for _, filePathsGroup := range filePathsGroups {
		go c.copyIfNeededWorker(filePathsGroup, &copiedFilesCounter, &wg, &counterMutex)
	}
	wg.Wait()
	isDoneChan <- true
	defer c.hashChecker.SaveHashes()
}

func (c CopyManager) copyIfNeededWorker(
	fileGroup []string,
	copiedFilesCounter *int,
	wg *sync.WaitGroup,
	counterMutex *sync.Mutex,
) {
	for _, filePath := range fileGroup {
		if c.hashChecker.CheckOrCreateHash(filePath) {
			dirToCopy, FileName := c.sortManager.CreateGroupingPath(filePath)
			_, err := os.Stat(dirToCopy)
			copyPath := fmt.Sprintf("%s/%s", dirToCopy, FileName)
			if err != nil {
				err = os.MkdirAll(dirToCopy, os.ModePerm)
				if err != nil {
					panic(err)
				}
				utils2.CopyFile(filePath, copyPath)
				counterMutex.Lock()
				*copiedFilesCounter += 1
				counterMutex.Unlock()
			} else {
				utils2.CopyFile(filePath, copyPath)
				counterMutex.Lock()
				*copiedFilesCounter += 1
				counterMutex.Unlock()
			}
		}
	}
	defer wg.Done()
}

func (c CopyManager) watchDog(
	isDoneChan chan bool,
	copiedFilesCounter *int,
	totalFileCounter int,
) {
	for {
		fmt.Println("Worker is working...")
		time.Sleep(2 * time.Second)
		select {
		case <-isDoneChan:
			fmt.Println("No more work, exiting...")
			break
		default:
			progress := float64(*copiedFilesCounter) / float64(totalFileCounter) * 100
			fmt.Println(progress)
		}
	}
}

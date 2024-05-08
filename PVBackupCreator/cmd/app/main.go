package main

import (
	src "awesomeProject/internal"
	"fmt"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	testDumpFolder := "/home/xan/Documents/raw_data"
	rootToCopy := "/home/xan/Documents/PH_VD_Backup"

	copyManager := src.NewCopyManager(testDumpFolder, rootToCopy)
	copyManager.SortAndCopyFiles()
	elapsed := time.Since(start).Seconds()
	fmt.Println("elapsed:", elapsed)
}

package main

import (
	src "awesomeProject/internal"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	testDumpFolder := "/home/xan/Documents/raw_data"
	rootToCopy := "/home/xan/Documents/PH_VD_Backup"

	copyManager := src.NewCopyManager(testDumpFolder, rootToCopy)
	copyManager.SortAndCopyFiles()
}

package utils

import "os"

func CopyFile(src string, dst string) {
	input, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(dst, input, 0644)
	if err != nil {
		panic(err)
	}
}

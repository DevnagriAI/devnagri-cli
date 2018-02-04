package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	visit2("./en")
	/*
		searchDir := "./en"

		fileList := []string{}

		filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
			fileList = append(fileList, path)
			return nil
		})

		for _, file := range fileList {
			fmt.Println(file)
		}

	*/
}

func visit2(searchDir string) {
	//searchDir := "./en"

	fileList := []string{}

	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		fmt.Println(file)
	}

}

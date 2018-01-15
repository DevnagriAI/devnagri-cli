package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	var allFiles []string

	root := "./en"
	extension := "pdf"

	err := filepath.Walk(root, visit(&allFiles))
	if err != nil {
		panic(err)
	}

	reqFiles := requiredExtensionFiles(allFiles, extension)
	fmt.Println(reqFiles)
	//TODO return the list of files as a json structure
	//fmt.Println(jsonObj.String())
}

func requiredExtensionFiles(files []string, extension string) []string {

	extensionRegexp := ".*." + extension + "$"

	var reqFiles []string
	for _, file := range files {

		matches, _ := regexp.MatchString(extensionRegexp, file)
		if matches == true {
			reqFiles = append(reqFiles, file)

		}
	}
	return reqFiles
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

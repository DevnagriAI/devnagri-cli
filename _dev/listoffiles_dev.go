package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/Jeffail/gabs"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func main() {
	var allFiles []string

	root := "./en"

	err := filepath.Walk(root, visit(&allFiles))
	if err != nil {
		panic(err)
	}

	jsonObj := gabs.New()
	// Print the list of file to the console
	for _, file := range allFiles {
		matches, _ := regexp.MatchString(".*.pdf$", file)
		if matches == true {

			jsonObj.Set(file)
			fmt.Println(file)
		}
	}

	//TODO return the list of files as a json structure

	fmt.Println(jsonObj.String())
}

func requiredExtensionFiles(files []string) []string {

	var reqFiles []string
	for _, file := range files {

		matches, _ := regexp.MatchString(".*.pdf$", file)
		if matches == true {

			fmt.Println(file)

		}
	}

}

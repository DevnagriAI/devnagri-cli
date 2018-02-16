package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/FourtekIT/devnagri-cli/config"
	"github.com/Jeffail/gabs"
)

func main() {

	files := visit2("./en")
	/*
		for _, file := range files {
			fmt.Println(file)
		}

	*/
	impFiles := requiredExtensionFiles(files, "xml")
	/*
		for _, file := range impFiles {
			fmt.Println(file)
		}
	*/

	//DONE
	fmt.Println(jsonListofFiles(impFiles))

	/*

		//TODO: Populate this struct on a per file basis.
		var fileJSON = map[string]string{
			"file[0][hash]":      sha256Hash(filename),
			"file[0][extension]": Extension,
			"file[0][file_type]": Extension,
			"file[0][location]":  filename,
		}
	*/

	var tempMap = make(map[int]string)

	for i := range []int{1, 2, 3, 4, 5} {
		tempMap[i], _ = strconv.FormatInt(i, 10)

		//fmt.Println(i)

	}

	fmt.Println(tempMap)

}

func visit2(searchDir string) []string {
	//searchDir := "./en"

	fileList := []string{}

	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	return fileList
	/*
		for _, file := range fileList {
			fmt.Println(file)
		}

	*/
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

func jsonListofFiles(files []string) string {

	// TODO Change the following to be initialized form the calling location
	var allFiles []string

	RootDir := config.FetchAndValidate("RootFolder") // returns string
	//RootDir := "./en"
	Extension := config.FetchAndValidate("Extension")
	//Extension := "pdf"

	err := filepath.Walk(RootDir, visit(&allFiles))
	if err != nil {
		panic(err)
	}

	reqFiles := requiredExtensionFiles(allFiles, Extension)
	//fmt.Println(reqFiles)

	jsonObj := gabs.New()
	jsonObj.Set(reqFiles)
	//fmt.Println(jsonObj.String())

	// Return the list of files as a json structure
	return jsonObj.String()

}

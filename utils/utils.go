package utils

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/FourtekIT/devnagri-cli/config"
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

	RootDir := config.FetchAndValidate("RootDir") // returns string
	//root := "./en"
	Extension := config.FetchAndValidate("Extension")
	//extension := "pdf"

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

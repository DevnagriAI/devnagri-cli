package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var data = `
ClientID: ""

ClientSecret: ""

ProjectKey: ""

RootDir: ""

LocalizationFolder: ""

SourceLanguage: "en"

TargetLanguages:
    - "pn"
    - "hi"

GlobalPreferenceInCaseOfMergeConflict: ""
`

type T struct {
	ClientID           string `yaml:"ClientID"`
	ClientSecret       string `yaml:"ClientSecret"`
	ProjectKey         string `yaml:"ProjectKey"`
	RootDir            string `yaml:"RootDir"`
	LocalizationFolder string `yaml:"LocalizationFolder"`
	SourceLanguage     string `yaml:"SourceLanguage"`
	//TODO This needs to be changed as the target languages are multiple
	TargetLanguages                       []string `yaml:"TargetLanguages"`
	GlobalPreferenceInCaseOfMergeConflict string   `yaml:"GlobalPreferenceInCaseOfMergeConflict"`
}

func main() {

	/*
		t := T{}

		d, err := yaml.Marshal(&t)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- t dump:\n%s\n\n", string(d))


	*/

	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//		fmt.Printf("--- m:\n%v\n\n", m)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%s\n\n", string(d))

	file, err := os.Create(".devnagri.yaml")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, string(d))
}

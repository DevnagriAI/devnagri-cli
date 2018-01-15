package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func FetchAndValidate(fieldName string) string {

	viper.SetConfigFile("./.devnagri.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet(fieldName) {
		fmt.Printf("\n%s ", fieldName)
		log.Fatal("missing the field")
	}

	fieldValue := viper.GetString(fieldName) // returns string
	return fieldValue

}

/*
var ClientID = fetchAndValidate("ClientID") // returns string

var ClientSecret = fetchAndValidate("ClientSecret") // returns string

var ProjectKey = fetchAndValidate("ProjectKey") // returns string

var RootDir = fetchAndValidate("RootDir") // returns string

var SourceLanguage = fetchAndValidate("SourceLanguage") // returns string

var TargetLanguages = fetchAndValidate("TargetLanguages") // returns string

var GlobalPreferenceInCaseOfMergeConflict = fetchAndValidate("GlobalPreferenceInCaseOfMergeConflict") // returns string
*/

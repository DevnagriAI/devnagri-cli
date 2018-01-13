package main

// TODO
//package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func ViperConfig() {

	viper.SetConfigFile("./.devnagri.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	validateAllFields()

	ClientID := viper.GetString("ClientID") // returns string
	fmt.Printf("\n\nValue: %v, Type: %T\n", ClientID, ClientID)

	ClientSecret := viper.GetString("ClientSecret") // returns string
	fmt.Printf("\n\nValue: %v, Type: %T\n", ClientSecret, ClientSecret)

	ProjectKey := viper.GetString("ProjectKey") // returns string
	fmt.Printf("\n\nValue: %v, Type: %T\n", ProjectKey, ProjectKey)

	RootDir := viper.GetString("RootDir") // returns string
	fmt.Printf("\n\nValue: %v, Type: %T\n", RootDir, RootDir)

	SourceLanguage := viper.GetString("SourceLanguage") // returns string
	fmt.Printf("\n\nValue: %v, Type: %T\n", SourceLanguage, SourceLanguage)

	LanguagesToBeTranslated := viper.GetString("LanguagesToBeTranslated") // returns string
	fmt.Printf("\n\nValue: %v, Type: %T\n", LanguagesToBeTranslated, LanguagesToBeTranslated)

	GlobalPreferenceInCaseOfMergeConflict := viper.GetString("GlobalPreferenceInCaseOfMergeConflict") // returns string
	fmt.Printf("\n\nValue: %v, Type: %T\n", GlobalPreferenceInCaseOfMergeConflict, GlobalPreferenceInCaseOfMergeConflict)

}

func validateAllFields() {

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("ClientID") {
		fmt.Printf("\n")
		log.Fatal("missing the \"ClientID\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("ClientSecret") {
		fmt.Printf("\n")
		log.Fatal("missing the \"ClientSecret\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("ProjectKey") {
		fmt.Printf("\n")
		log.Fatal("missing the \"ProjectKey\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("RootDir") {
		fmt.Printf("\n")
		log.Fatal("missing the \"RootDir\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("LocalizationFolder") {
		fmt.Printf("\n")
		log.Fatal("missing the \"LocalizationFolder\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("SourceLanguage") {
		fmt.Printf("\n")
		log.Fatal("missing the \"SourceLanguage\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("GlobalPreferenceInCaseOfMergeConflict") {
		fmt.Printf("\n")
		log.Fatal("missing the \"GlobalPreferenceInCaseOfMergeConflict\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("LanguagesToBeTranslated") {
		fmt.Printf("\n")
		log.Fatal("missing the \"LanguagesToBeTranslated\" field")
	}

}

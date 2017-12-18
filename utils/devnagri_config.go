package main

// TODO
//package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {

	viper.SetConfigFile("./.devnagri.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	x := viper.GetString("ClientID") // returns string
	//port := viper.GetInt("prod.port") // returns integer
	fmt.Printf("\n\nValue: %v, Type: %T\n", x, x)

}

func validateAllFields() {

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("ClientID") {
		fmt.Printf("\n")
		log.Fatal("missing the \"none\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("ClientSecret") {
		fmt.Printf("\n")
		log.Fatal("missing the \"none\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("ProjectKey") {
		fmt.Printf("\n")
		log.Fatal("missing the \"none\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("LocalizationFolder") {
		fmt.Printf("\n")
		log.Fatal("missing the \"none\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("GlobalPreferenceInCaseOfMergeConflict") {
		fmt.Printf("\n")
		log.Fatal("missing the \"none\" field")
	}

	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("LanguagesToBeTranslated") {
		fmt.Printf("\n")
		log.Fatal("missing the \"none\" field")
	}

}

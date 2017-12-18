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

	x := viper.Get("client_id") // returns string
	//port := viper.GetInt("prod.port") // returns integer
	fmt.Printf("\n\nValue: %v, Type: %T\n", x, x)



	// To validate the presence of a certain field in the YAML file
	if !viper.IsSet("prod.none") {
		fmt.Printf("\n")
		log.Fatal("missing the \"none\" field")
	}
}



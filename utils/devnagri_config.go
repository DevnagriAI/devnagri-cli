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

	x := viper.Get("auth.username") // returns string
	//port := viper.GetInt("prod.port") // returns integer
	fmt.Printf("\n\nValue: %v, Type: %T\n", x, x)

	if !viper.IsSet("prod.none") {
		fmt.Printf("\n")
		log.Fatal("missing the \"none\" field")
	}
}



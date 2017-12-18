package main

// TODO
//package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"log"
)

func main() {

	viper.SetConfigFile("./env.json")
	port := viper.Get("dev.port") // returns string
	//port := viper.GetInt("prod.port") // returns integer
	fmt.Printf("\n\nValue: %v, Type: %T\n", port, port)

	if !viper.IsSet("prod.none") {
		fmt.Printf("\n")
		log.Fatal("missing the \"none\" field")
	}

}

func createConfigFile() {

}

// TODO
//viper.AddConfigPath("$HOME/configs")
// And then register config file name (no extension)
//viper.SetConfigName("env")

package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("yaml")

	viper.SetConfigFile("./.devnagri.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	fieldValue := viper.Get("TargetLanguages") // returns string
	fmt.Println(fieldValue)
}

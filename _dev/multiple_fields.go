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

	fieldValue := viper.GetStringSlice("TargetLanguages") // returns string
	fmt.Println(fieldValue)

	for _, x := range fieldValue {
		println(x)
	}

	//reflect.ArrayOf(2, fieldValue)

	/*

		b := make([]interface{}, len(fieldValue))
		for i := range fieldValue {
			b[i] = fieldValue[i]
			fmt.Println(b[i])
		}

	*/
}

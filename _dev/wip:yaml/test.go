package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {

	ViperConfig()
	fmt.Println(ClientID)

	viper.SetConfigFile("./.devnagri.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	ClientID := viper.GetString("ClientID") // returns string
	fmt.Printf("\n\nValue: %v, Type: %T\n", ClientID, ClientID)

}

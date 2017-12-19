package main


import (
	
	"fmt"
	"io/ioutil"
	"log"
	"github.com/spf13/viper"
	// TODO https://blog.minio.io/debugging-go-routine-leaks-a1220142d32c
	"github.com/fortytw2/leaktest"


func main()  {

	
	viper.SetConfigFile("./.devnagri.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	RootDir := viper.GetString("RootDir") // returns string
	
	// Here we read the files in the RootDirectory
	files, err := ioutil.ReadDir(RootDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
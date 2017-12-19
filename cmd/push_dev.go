package main


import (
	
	"fmt"
	"io/ioutil"
	"log"
	"github.com/spf13/viper"
	"devnagri-cli/utils"
)

func main()  {


	files := listoffiles_dev.main()

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
package main

import (
	"fmt"
	"github.com/FourtekIT/devnagri-cli/config"
)

func main() {

	fmt.Println(config.ClientID)

	fmt.Println(config.ClientSecret)

	fmt.Println(config.ProjectKey)

	fmt.Println(config.SourceLanguage)

	//	fmt.Println(config.LanguagesToBeTranslated)

}

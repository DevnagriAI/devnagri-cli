package main

import (
	"fmt"
	"github.com/Jeffail/gabs"
	//	"reflect"
)

func main() {

	jsonString := []byte(`
{"languages_status":
{"hi":
{"translatedWordsCount":0,"reviewedWordsCount":0,"totalWordsCount":1283,"pendingWordsCount":1283}},
"project_status":"draft"}
`)

	jsonParsed, _ := gabs.ParseJSON(jsonString)

	fmt.Println(jsonParsed.Path("languages_status").String())

	//projectStatus := jsonParsed.Path("project_status").Data().(string)

	//value := jsonParsed.Path("languages_status.hi.totalWordsCount").Data()

	//fmt.Println(value)

	//fmt.Println(reflect.TypeOf(value))

	/*
			lang, _ := jsonParsed.Path("languages_status").Children()
			langName := lang[0]
			fmt.Println(langName)

		children, _ := jsonParsed.Path("languages_status").Children()
		for _, child := range children {
			fmt.Println(child.Data())
		}

	*/
}

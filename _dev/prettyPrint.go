package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
	//	"reflect"
)

func main() {
	/*
		jsonString := []byte(`
			{"languages_status":
			{"hi":
			{"translatedWordsCount":0,"reviewedWordsCount":0,"totalWordsCount":1283,"pendingWordsCount":1283}},
			"project_status":"draft"}
			`)

	*/

	jsonString := []byte(`

	   {
	   	"languages_status": {

	   		"hi": {
	   			"translatedWordsCount": 0,
	   			"reviewedWordsCount": 0,
	   			"totalWordsCount": 1283,
	   			"pendingWordsCount": 1283
	   		},

	   		"pn": {
	   			"translatedWordsCount": 0,
	   			"reviewedWordsCount": 0,
	   			"totalWordsCount": 1283,
	   			"pendingWordsCount": 1283
	   		}
	   	},
	   	"project_status": "draft"
	   }

	   	   `)

	jsonParsed, _ := gabs.ParseJSON(jsonString)

	//	langName := jsonParsed.Path("languages_status").String()
	//	fmt.Println(langName)

	//value := jsonParsed.Path("languages_status.hi.totalWordsCount").Data()

	//fmt.Println(value)

	//fmt.Println(reflect.TypeOf(value))

	//lang, _ := jsonParsed.Path("languages_status").Children()
	//langName := lang[0]
	//fmt.Println(langName)

	/*
		children, _ := jsonParsed.Path("languages_status").Children()
		for _, child := range children {
			fmt.Println(child.Data())
		}
	*/

	projectStatus := jsonParsed.Path("project_status").Data().(string)

	fmt.Println("\nProject Status : ", projectStatus, "\n")

	childrenMap, _ := jsonParsed.Path("languages_status").ChildrenMap()

	for key, val := range childrenMap {

		//fmt.Println("\n", key, val)

		langName := key

		valChildren, _ := val.ChildrenMap()

		pendingWordsCount := valChildren["pendingWordsCount"]
		reviewedWordsCount := valChildren["reviewedWordsCount"]
		totalWordsCount := valChildren["totalWordsCount"]
		translatedWordsCount := valChildren["translatedWordsCount"]

		fmt.Println("\n~~~~~~~~~\n")
		fmt.Println("langName : ", langName)
		fmt.Println("pendingWordsCount : ", pendingWordsCount)
		fmt.Println("reviewedWordsCount : ", reviewedWordsCount)
		fmt.Println("totalWordsCount : ", totalWordsCount)
		fmt.Println("translatedWordsCount : ", translatedWordsCount)

		/*
			for _, y := range valChildren {
				fmt.Println("\n", y)
			}
		*/
	}

}

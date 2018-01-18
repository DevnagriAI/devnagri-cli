package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

func main() {

	jsonParsed, _ := gabs.ParseJSON([]byte(`{"languages_status":{"hi":{"translatedWordsCount":0,"reviewedWordsCount":0,"totalWordsCount":1283,"pendingWordsCount":1283}},"project_status":"draft"}`))
	value := jsonParsed.Search("totalWordsCount").Data()
	fmt.Println(value)

}

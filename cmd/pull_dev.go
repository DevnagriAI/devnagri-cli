package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	//	"os"
	//	"reflect"

	//	"github.com/Jeffail/gabs"
	"gopkg.in/resty.v1"
)

func main() {
}

func decodeBase64(cypher string) string {
	//println("\n\nRecieved at decodeBase64 function\n\n")
	//	fmt.Println(cypher)

	data, _ := base64.StdEncoding.DecodeString(cypher)

	//fmt.Println(string(data))

	println("Text after decoding")

	//fmt.Println(data)
	fmt.Println(string(data))

	return string(data)

}

func saveResponseAndConvert() string {
	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "multipart/form-data").
		SetAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjE0MmI1YjlkOTI3ZGQ2ZjI3ODcyZDJmNDFmNWQzNjhkOGU3MmVhZjk2YzE2YTdmMTZjM2ZkM2I5ZmVmNjJiMDFmMTBiOWQwMTMxNzJjYTIyIn0.eyJhdWQiOiIxIiwianRpIjoiMTQyYjViOWQ5MjdkZDZmMjc4NzJkMmY0MWY1ZDM2OGQ4ZTcyZWFmOTZjMTZhN2YxNmMzZmQzYjlmZWY2MmIwMWYxMGI5ZDAxMzE3MmNhMjIiLCJpYXQiOjE1MTU1OTA4NDksIm5iZiI6MTUxNTU5MDg0OSwiZXhwIjoxNTQ3MTI2ODQ5LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.inJAzB_9Vjx9YXHQgZYCy3HxbTaSV2VSjKYgtB_s4ZrrS9vucEnEF4Ijh8wPMylxDLI7e8iGfRObK0rGTBqnCw94ElcmeYPM7GfWRqWW5Bz8ctJhXH7-VGml22fT6ahFVHssbVcx636xwPlydhN4fBay0TElXgC5QkSmzINFPMF3FRJ54WaAsTAvwpsnAIUfigZdNggSjMnROp-i7dtF5M9Ty3utLM2kmlEZVC9lhhLUvJujvbWJwCGDOStrBv4iapWQ4qPZ8AuaVTmigW_hVinkZ6GGTp8GT07GGF9iwDjPnSyCzEx7Ta4S0U3nSeJVogxmsO7N1hZrmCuZmUtEkSPMSP0wIhrMrranpzqFRqvsYve6hp8BBE9qxbuiEwc3F-2xFm-R6esmUu5mpWcRlM1imenuvb9C-lotKf0znM9neVimWhMe7OqENcGXTx7puTyZoK0wUmrn-2_vIxfwtU5s9aBomlw6qGQt6k14pgKA69bp2C51C46HyKa_QrEvol-_3YvccjvR3hO-jcAj_SLALYk6HJp1eSqy_tGZZ5QW4-4QvUcPaa3E4sXA4Ye_4OB2dGRcKVuhR5jzpaYJC-_qGRY0dhfYdpKjEgc1o3pitzpNW0iWzImN9iidRThu_YFP_RwPk2Qragch55f086JNdM7jGPkjcEUcgGIs3PE").
		SetFormData(map[string]string{
			"client_id":     "3",
			"client_secret": "3WnUqVSP7Vhs8DU7FInIrwHIVMg9twGshcpswlJW",
			"project_key":   "2e12635aca73c7d39ec76a514d7490a6"}).
		Post("http://dev.devnagri.co.in/api/project/pull")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	fmt.Println("\n\n")

	/*
		resJson, _ := gabs.ParseJSON([]byte(resp.String()))
			children, _ := resJson.S("file_content").Children()
			child := children[0]
	*/
	//fmt.Println(child.String())

	//data, _ := base64.StdEncoding.Decode(child.String())
	//encoded := child.String()
	//fmt.Println(encoded)
	//fmt.Println(reflect.TypeOf(encoded))
	//x := decodeBase64(encoded)
	//fmt.Println(x)

	// This works perfectly
	//fmt.Println("Decoding the string manually")
	//x := "PCEtLSBUcmFuc2xhdGVkIEJ5IERldm5hZ3JpIC0tPgo8IS0tIGh0dHA6Ly9kZXZuYWdyaS5jb20gLS0+CjxyZXNvdXJjZXMgdG9vbHM6aWdub3JlPSJFeHRyYVRyYW5zbGF0aW9uIiB4bWxuczp0b29scz0iaHR0cDovL3NjaGVtYXMuYW5kcm9pZC5jb20vdG9vbHMiPgogICAgPHN0cmluZyBuYW1lPSJhcHBfbmFtZSI+PC9zdHJpbmc+CiAgICA8c3RyaW5nIG5hbWU9ImhpbnRfYWN0dWFsIj48L3N0cmluZz4KIDwvcmVzb3VyY2VzPg=="
	//decodeBase64(x)

	dat, _ := ioutil.ReadFile("./content.txt")
	fmt.Println("reading content.txt")
	decodeBase64(string(dat))

}

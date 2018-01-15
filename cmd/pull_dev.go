package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	//"reflect"
	//	"github.com/Jeffail/gabs"
	//	"gopkg.in/resty.v1"
	//      "encoding/base64"
	"github.com/Jeffail/gabs"
	"gopkg.in/resty.v1"
	"io"
	"log"
	"os"
)

func main() {
	resp, err := resty.R().
		SetHeader("Accept", "application/json").SetHeader("Content-Type", "multipart/form-data").
		SetAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjE0MmI1YjlkOTI3ZGQ2ZjI3ODcyZDJmNDFmNWQzNjhkOGU3MmVhZjk2YzE2YTdmMTZjM2ZkM2I5ZmVmNjJiMDFmMTBiOWQwMTMxNzJjYTIyIn0.eyJhdWQiOiIxIiwianRpIjoiMTQyYjViOWQ5MjdkZDZmMjc4NzJkMmY0MWY1ZDM2OGQ4ZTcyZWFmOTZjMTZhN2YxNmMzZmQzYjlmZWY2MmIwMWYxMGI5ZDAxMzE3MmNhMjIiLCJpYXQiOjE1MTU1OTA4NDksIm5iZiI6MTUxNTU5MDg0OSwiZXhwIjoxNTQ3MTI2ODQ5LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.inJAzB_9Vjx9YXHQgZYCy3HxbTaSV2VSjKYgtB_s4ZrrS9vucEnEF4Ijh8wPMylxDLI7e8iGfRObK0rGTBqnCw94ElcmeYPM7GfWRqWW5Bz8ctJhXH7-VGml22fT6ahFVHssbVcx636xwPlydhN4fBay0TElXgC5QkSmzINFPMF3FRJ54WaAsTAvwpsnAIUfigZdNggSjMnROp-i7dtF5M9Ty3utLM2kmlEZVC9lhhLUvJujvbWJwCGDOStrBv4iapWQ4qPZ8AuaVTmigW_hVinkZ6GGTp8GT07GGF9iwDjPnSyCzEx7Ta4S0U3nSeJVogxmsO7N1hZrmCuZmUtEkSPMSP0wIhrMrranpzqFRqvsYve6hp8BBE9qxbuiEwc3F-2xFm-R6esmUu5mpWcRlM1imenuvb9C-lotKf0znM9neVimWhMe7OqENcGXTx7puTyZoK0wUmrn-2_vIxfwtU5s9aBomlw6qGQt6k14pgKA69bp2C51C46HyKa_QrEvol-_3YvccjvR3hO-jcAj_SLALYk6HJp1eSqy_tGZZ5QW4-4QvUcPaa3E4sXA4Ye_4OB2dGRcKVuhR5jzpaYJC-_qGRY0dhfYdpKjEgc1o3pitzpNW0iWzImN9iidRThu_YFP_RwPk2Qragch55f086JNdM7jGPkjcEUcgGIs3PE").
		SetFormData(map[string]string{
			"client_id":     "3",
			"client_secret": "3WnUqVSP7Vhs8DU7FInIrwHIVMg9twGshcpswlJW",
			"project_key":   "2e12635aca73c7d39ec76a514d7490a6"}).
		Post("http://dev.devnagri.co.in/api/project/pull")
		//Post("http://192.168.60.10/api/project/pull")
		//Post("https://requestb.in/vwvh94vw")
	if err != nil {
		panic(err)
	}

	//fmt.Println("Response \n\n")
	//fmt.Println(resp.String())

	//	fmt.Println("Byte Array from String")
	//	fmt.Println([]byte(resp.String()))

	//fmt.Println(resp.Body())

	//uDec, _ := base64.StdEncoding.DecodeString(resp.Result())
	//	fmt.Println("<< Contents of file >>")
	//	fmt.Println(string(uDec))

	//resJson := gabs.ParseJSON(resp.String())

	resJson, _ := gabs.ParseJSON([]byte(resp.String()))

	//fmt.Println(resJson)

	//	fmt.Println("The base64 contents of the returned file : ")

	//fmt.Println(resJson.Path("file_content").Data())

	//cypher := resJson.Path("file_content").String()
	//	fmt.Println(reflect.TypeOf(cypher))
	//	fmt.Println(cypher)

	//fmt.Println(decodeBase64(cypher))

	// 	data, err := base64.StdEncoding.DecodeString(cypher)
	// 	if err != nil {
	// 		log.Fatal("error:", err)
	// 	}

	// 	fmt.Printf("%q\n", data)

	//children, _ := resJson.S("file_content").Children()
	//	children := resJson.S("file_content").Data()
	//	fmt.Println(children.(string))

	children, _ := resJson.S("file_content").Children()

	/*
		for _, child := range children {
			fmt.Println(child.Data().(string))
		}
	*/

	child := children[0]
	//fmt.Println(reflect.TypeOf(child.String()))

	//fmt.Println(child.String())

	decodeBase64(child.String())

	//fmt.Println(decodeBase64(child.String()))
}

func decodeBase64(cypher string) string {
	println("\n\nRecieved at decodeBase64 function\n\n")
	fmt.Println(cypher)

	data, _ := base64.StdEncoding.DecodeString(cypher)

	//fmt.Println(string(data))

	println("Text after decoding")
	fmt.Println(string(data))

	return string(data)

}

func sha256Hash(fileName string) string {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%x", h.Sum(nil))
	//return h.Sum(nil)

	hashString := hex.EncodeToString(h.Sum(nil))
	return hashString
}

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/resty.v1"
)

func main() {

	//fileBytes, _ := ioutil.ReadFile("CallingPapaPro2.xml")
	//hash := BytesToString(sha256Hash("./CallingPapaPro2.xml"))
	//hash := sha256Hash("./CallingPapaPro2.xml")
	//hash := hex.EncodeToString(sha256Hash("./CallingPapaPro2.xml"))
	//hash := sha256Hash("./CallingPapaPro2.xml")

	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("file", "./CallingPapaPro2.xml").
		SetAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjE0MmI1YjlkOTI3ZGQ2ZjI3ODcyZDJmNDFmNWQzNjhkOGU3MmVhZjk2YzE2YTdmMTZjM2ZkM2I5ZmVmNjJiMDFmMTBiOWQwMTMxNzJjYTIyIn0.eyJhdWQiOiIxIiwianRpIjoiMTQyYjViOWQ5MjdkZDZmMjc4NzJkMmY0MWY1ZDM2OGQ4ZTcyZWFmOTZjMTZhN2YxNmMzZmQzYjlmZWY2MmIwMWYxMGI5ZDAxMzE3MmNhMjIiLCJpYXQiOjE1MTU1OTA4NDksIm5iZiI6MTUxNTU5MDg0OSwiZXhwIjoxNTQ3MTI2ODQ5LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.inJAzB_9Vjx9YXHQgZYCy3HxbTaSV2VSjKYgtB_s4ZrrS9vucEnEF4Ijh8wPMylxDLI7e8iGfRObK0rGTBqnCw94ElcmeYPM7GfWRqWW5Bz8ctJhXH7-VGml22fT6ahFVHssbVcx636xwPlydhN4fBay0TElXgC5QkSmzINFPMF3FRJ54WaAsTAvwpsnAIUfigZdNggSjMnROp-i7dtF5M9Ty3utLM2kmlEZVC9lhhLUvJujvbWJwCGDOStrBv4iapWQ4qPZ8AuaVTmigW_hVinkZ6GGTp8GT07GGF9iwDjPnSyCzEx7Ta4S0U3nSeJVogxmsO7N1hZrmCuZmUtEkSPMSP0wIhrMrranpzqFRqvsYve6hp8BBE9qxbuiEwc3F-2xFm-R6esmUu5mpWcRlM1imenuvb9C-lotKf0znM9neVimWhMe7OqENcGXTx7puTyZoK0wUmrn-2_vIxfwtU5s9aBomlw6qGQt6k14pgKA69bp2C51C46HyKa_QrEvol-_3YvccjvR3hO-jcAj_SLALYk6HJp1eSqy_tGZZ5QW4-4QvUcPaa3E4sXA4Ye_4OB2dGRcKVuhR5jzpaYJC-_qGRY0dhfYdpKjEgc1o3pitzpNW0iWzImN9iidRThu_YFP_RwPk2Qragch55f086JNdM7jGPkjcEUcgGIs3PE").
		SetFormData(map[string]string{
			"client_id":     "3",
			"client_secret": "3WnUqVSP7Vhs8DU7FInIrwHIVMg9twGshcpswlJW",
			"project_key":   "2e12635aca73c7d39ec76a514d7490a6",
			//"file[0][hash]": "3221255c6aecb25c8f73472dcb7c99f42ade9112b8e3029a3e67070233fd101d",
			//"file[0][hash]": hash,
			"file[0][hash]": sha256Hash("./CallingPapaPro2.xml"),
			//"file[0][name]":      "CallingPapaPro2.xml",
			"file[0][extension]": "xml",
			"file[0][file_type]": "xml",
			"file[0][location]":  "./CallingPapaPro2.xml",
			//"file[0][file]":

		}).
		Post("http://dev.devnagri.co.in/api/project/sync")
		//Post("http://192.168.60.10/api/project/sync")
		//Post("https://requestb.in/vwvh94vw")
	if err != nil {
		panic(err)
	}

	//fmt.Println(sha256Hash("push_test_file.xml"))

	fmt.Println(resp)
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

// NOTE: http://www.mrwaggel.be/post/golang-hash-sum-and-checksum-to-string-tutorial-and-examples/
func BytesToString(data []byte) string {
	return string(data[:])
}

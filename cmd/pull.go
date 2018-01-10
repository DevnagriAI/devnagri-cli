// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"reflect"
	//	"github.com/Jeffail/gabs"
	//	"gopkg.in/resty.v1"
	//      "encoding/base64"

	//	"encoding/base64"
	"devnagri-cli/utils"
	"fmt"
	"github.com/Jeffail/gabs"
	"gopkg.in/resty.v1"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "A brief description of pull command",
	Long:  `A long description of pull command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")

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
		//	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
		//	fmt.Printf("\nResponse Status: %v", resp.Status())
		//	fmt.Printf("\nResponse Time: %v", resp.Time())
		//	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
		//	fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())

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
		//fmt.Println(resJson.Path("file_content"))
		cypher := resJson.Path("file_content").String()
		fmt.Println(reflect.TypeOf(cypher))
		fmt.Println(cypher)

		fmt.Println(decodeBase64(cypher))

		// 	data, err := base64.StdEncoding.DecodeString(cypher)
		// 	if err != nil {
		// 		log.Fatal("error:", err)
		// 	}

		// 	fmt.Printf("%q\n", data)

		//children, _ := resJson.S("file_content").Children()

		//children := resJson.S("file_content").Data()
		//fmt.Println(children)

	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

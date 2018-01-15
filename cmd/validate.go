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
	"fmt"
	"github.com/FourtekIT/devnagri-cli/config"
	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"
	"gopkg.in/resty.v1"
	"os"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "This command validates the credentials in the .devnagri file within the local folder.",
	Long:  `A long description of validate command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("validate called")
		validate()
},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



func validate() {
		resp, err := resty.R().
			SetHeader("Content-Type", "multipart/form-data").
			SetFormData(map[string]string{
				"client_id":     config.ClientID,
				"client_secret": config.ClientSecret,
				"project_key":   config.ProjectKey}).
			Post("http://dev.devnagri.co.in/api/key/validations")
		//	Post("http://192.168.60.10/api/key/validations")

		if err != nil {
			panic(err)
		}

		fmt.Println(resp)

		jsonParsed, _ := gabs.ParseJSON([]byte(resp.String()))
		accessToken := jsonParsed.Path("access_token").Data()
		//fmt.Println(access_token)

		filename := "./.devnagri.yaml"

		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		accessTokenString := "AccessToken: " + accessToken.(string)
		fmt.Println(accessTokenString)
		// TODO: Add this to the .devnagri.yaml file as a string
		f.WriteString(accessTokenString)

		if err != nil {
			panic(err)
		}

	}


// Copyright © 2017 Abhinav Sharma <abhinav@fourtek.com>
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
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "This command fetches the status of the current project.",
	Long:  `A long description of status command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The current status of your project from Devnagri")

		fetchStatus()

		//fmt.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	//fmt.Println("Working on fetchStatusDev")
	//fetchStatusDev()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fetchStatus() {

	var ClientID = config.FetchAndValidate("ClientID") // returns string

	var ClientSecret = config.FetchAndValidate("ClientSecret") // returns string

	var ProjectKey = config.FetchAndValidate("ProjectKey") // returns string

	var AccessToken = config.FetchAndValidate("AccessToken") // returns string

	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(AccessToken).
		SetFormData(map[string]string{
			"client_id":     ClientID,
			"client_secret": ClientSecret,
			"project_key":   ProjectKey}).
		Post("http://dev.devnagri.co.in/api/project/status")

	if err != nil {
		panic(err)
	}

	//fmt.Println(resp)

	/*
		resJson, _ := gabs.ParseJSON([]byte(resp.String()))
		val := resJson.S("translatedWordsCount")
		fmt.Println(val)
		//TODO: Print the properly formatted output of the status command

	*/

	jsonParsed, _ := gabs.ParseJSON([]byte(resp.String()))

	projectStatus := jsonParsed.Path("project_status").Data().(string)

	fmt.Println("\nProject Status : ", projectStatus, "\n")

	childrenMap, _ := jsonParsed.Path("languages_status").ChildrenMap()
	for key, val := range childrenMap {
		fmt.Println("\n", key, val)
	}
	//TODO: add a suboption to print out the raw json
}

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
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	//"github.com/FourtekIT/devnagri-cli/utils"
	"github.com/FourtekIT/devnagri-cli/config"
	"github.com/spf13/cobra"
	"gopkg.in/resty.v1"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "This command pushes the untranslated files to Devnagri",
	Long:  `This command transfers all the untranslated local files to the Devnagri platform on a language basis.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pushing the local files to Devnagri for localization")
		listAllFilesAndPush()
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listAllFilesAndPush() {
	var ClientID = config.FetchAndValidate("ClientID") // returns string

	var ClientSecret = config.FetchAndValidate("ClientSecret") // returns string

	var ProjectKey = config.FetchAndValidate("ProjectKey") // returns string

	var AccessToken = config.FetchAndValidate("AccessToken") // returns string

	filename := "./en/strings.xml"

	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "multipart/form-data").
		SetAuthToken(AccessToken).
		SetFile("file[0][file]", filename).
		SetFormData(map[string]string{
			"client_id":          ClientID,
			"client_secret":      ClientSecret,
			"project_key":        ProjectKey,
			"file[0][hash]":      sha256Hash(filename),
			"file[0][extension]": "xml",
			"file[0][file_type]": "xml",
			"file[0][location]":  filename,
		}).
		Post("http://dev.devnagri.co.in/api/project/push")

	if err != nil {
		panic(err)
	}

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

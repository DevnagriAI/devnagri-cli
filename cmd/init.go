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
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var data = `
ClientID: ""

ClientSecret: ""

ProjectKey: ""

RootDir: ""

LocalizationFolder: ""

SourceLanguage: "en"

TargetLanguages:
    - "pn"
    - "hi"

GlobalPreferenceInCaseOfMergeConflict: ""
`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "This command creates a .devnagri.yaml file in the repo.",
	Long:  `A longer description of init command`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	createConfigFile()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	initCmd.PersistentFlags().String("clientid", "", "Enter your clientid.")
	initCmd.PersistentFlags().String("clientsecret", "", "Enter your clientsecret")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// TODO: create this using viper
func createConfigFile() {

	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//		fmt.Printf("--- m:\n%v\n\n", m)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("%s\n\n", string(d))

	file, err := os.Create(".devnagri.yaml")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, string(d))
}

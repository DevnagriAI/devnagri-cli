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

//TODO: change this package name
//package cmd

package main

import (
	"fmt"

	"github.com/FourtekIT/devnagri-cli/config"
	"gopkg.in/resty.v1"
)

func main() {

	var ClientID = config.FetchAndValidate("ClientID") // returns string

	var AccessToken = config.FetchAndValidate("AccessToken") // returns string

	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(AccessToken).
		//SetHeader("Content-Type", "multipart/form-data").
		SetFormData(map[string]string{
			"client_id":     ClientID,
			"client_secret": "3WnUqVSP7Vhs8DU7FInIrwHIVMg9twGshcpswlJW",
			"project_key":   "2e12635aca73c7d39ec76a514d7490a6"}).
		Post("http://dev.devnagri.co.in/api/project/status")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

}

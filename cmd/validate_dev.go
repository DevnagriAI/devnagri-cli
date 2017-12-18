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
	"gopkg.in/resty.v1"
)

func main() {

	resp, err := resty.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFormData(map[string]string{
			"client_id":     "4",
			"client_secret": "kBBhqrpeEjJCGLYUCQw9rJuZ6MjLT6C7iUeVISWd",
			"project_key":   "d41d8cd98f00b204e9800998ecf8427e"}).
		Post("https://requestb.in/16pfgiv1")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

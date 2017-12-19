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
		SetHeader("Accept", "application/json").
		SetAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjQ2ODhlOGIwNTNjZmEwZDlmNWZjZjA5NDBjYWI1MTRiNzM2NDgyZDYxNWM0ZWZiY2RkYzUzOTc5Yzk2YjUxZGE4MmJjZDQ0MGI5YjdmMmMwIn0.eyJhdWQiOiIxIiwianRpIjoiNDY4OGU4YjA1M2NmYTBkOWY1ZmNmMDk0MGNhYjUxNGI3MzY0ODJkNjE1YzRlZmJjZGRjNTM5NzljOTZiNTFkYTgyYmNkNDQwYjliN2YyYzAiLCJpYXQiOjE1MTM2NzgyMjUsIm5iZiI6MTUxMzY3ODIyNSwiZXhwIjoxNTQ1MjE0MjI1LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.IAuBFik92MfyTmJExDi7AuWPWNxFiCEjHU080PpYWNVafORYaydVJbizeY31ZNzquUFMND8shjvJxK_mmRKE3lcFqcnWnJDvSsSgFdavJ9z_v1tyq0e6hMPtqncRpRRkxmYhZsVVYjP4uLdUzdNhnJC9vDJtU2icv2OzR2zDgzvMhphfWlDtTaWNywJB2EHmun-hkbk4TLjx7rjAGy-yhskh_-3rEkvll_IhG4dSv87u6l-QCiZwS5OEBqCRcOWWqZtqghIDyktSZc1WNJVsvxbIdCkB8v44qGRMzkWm3LxRbqRl6WwcC5OZcqy-Uw9PLqlEEDMsVlO0GG9sEVPHZ5Oheh8ftNf_AiqQxamTzIwli_0Yjjwgj6U1tp37rDDYcgRCyYbwoBubWr7Maw45fLZvrrH8khndSRfRpY3o7S42UoCa3AVEWmHpuiRaqYAsSbqGsU_HD1urIsZTa19t8x6ocw8vmcQS9cJgp116QxMWYz4qdUqjpPdoxRDRehkh4ooY6q9e0DyRBbq1aXHGyqkhXWlqYUotd08GXS-nMGUwMOfrgalYvCycHSZdMAEgsfeZmKSaDRNNrRj39_JimsU6VcTSwdC99wRIPU6Yz70aIDJO0xhQA0S6j6KbwUUebiKtHoT3fDewEnLVctF7HCntnELqTIVDIlcwzu5ESxs").

		//SetHeader("Content-Type", "multipart/form-data").
		SetFormData(map[string]string{
			"client_id":     "4",
			"client_secret": " uTBn7fxknAPpmw9AiEXyIro7X8mP0JhkqPtvBS28",
			"project_key":   "d41d8cd98f00b204e9800998ecf8427e"}).
		Post("http://192.168.60.10/api/project/status")
		//Post("https://requestb.in/16pfgiv1")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

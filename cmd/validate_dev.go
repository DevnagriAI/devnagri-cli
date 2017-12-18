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
	//	"gopkg.in/resty.v1"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	// Dummy credentials during the dev
	//	var ProductKey = "d41d8cd98f00b204e9800998ecf8427e"

	//	var ClientID = "kBBhqrpeEjJCGLYUCQw9rJuZ6MjLT6C7iUeVISWd"

	//	var ClientSecret = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjEwNDM4ODNkNmQ4OWVlMjQ2NzkyNDNhMDNhODFlNTczMzg5MTQ3MGY0ZGRiNTFhMjAxYzQ5YTk3MGUyM2QzY2ZmMTUyMjFlYzY5ZjgwNWEwIn0.eyJhdWQiOiIxIiwianRpIjoiMTA0Mzg4M2Q2ZDg5ZWUyNDY3OTI0M2EwM2E4MWU1NzMzODkxNDcwZjRkZGI1MWEyMDFjNDlhOTcwZTIzZDNjZmYxNTIyMWVjNjlmODA1YTAiLCJpYXQiOjE1MTM1ODc1ODUsIm5iZiI6MTUxMzU4NzU4NSwiZXhwIjoxNTQ1MTIzNTg0LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.wzZPNKVj5PvnM-4tPHtvqRR4izkrEMTk12uuBtugqfbCBpVVQrUdJiMmqqYyavDm3WXxnGA71J58YFzp6v0gQLGRsHlmh6GM3vW47xRgtYwSjWf5A1mn-A9Jpfc56K8_R0TTQMLcmicKPqB7ttrqsArYU1kB-UomWvpVylJJ3jlK-a1LNdbfr19isFwv6mFgREbkgmVhpn7iHxItI0D7ta8CuPa9GUl2SFfr1Y4SqZxACqBqTo1-irvkMz7OZLoJJeSyEwluFFGmhCFFQ1dOl7Y9qFBowJjS2nD_u-TG59L1ic3RRculaDd_-hKEInacahaJwHDsu8-UFG2P1np_xDqAlegyDUBp895TBjcy1kgjeg07LteUUOnZmnnSIjj2q6bZQUMTpdhb_CpJszYtEKrLp33k8LPvjg0AfkGcnLMwJFVTUBrEC6V5fg94INHR1qROD4ZSVR6d0BiUTxxrj2jJ-5vvjwY7ZNf11Bh5ivApoLovoRUUSoWexrt-mbeIs6rBVG456sCtEhYPem-xaLBwa8u5fVVVc987rPLOEZSUaRb5MFRmF43-XoUDxEls0jZEroBCOrcCqaemwGwXh3o3tCLW_kay9aH645Hdufm8kZMvRGLKAC5q1k7YQhcL9AShNqr9buMXp6bQ9o9HF7G8n15_cWfua3j9IqKbHfE"

	url := "http://192.168.60.10/api/key/validations"
	//	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"client_id\"\r\n\r\n3\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"client_secret\"\r\n\r\nHfU2hKMpArJaZVUljxNBoBYgWlWDK0D7Rmh6u3wn\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"project_key\"\r\n\r\nd41d8cd98f00b204e9800998ecf8427e\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")

	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"client_id\"\r\n\r\n3\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"client_secret\"\r\n\r\nHfU2hKMpArJaZVUljxNBoBYgWlWDK0D7Rmh6u3wn\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"project_key\"\r\n\r\nd41d8cd98f00b204e9800998ecf8427e\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

/*
	// Based on resty
	resp, err := resty.R().Get("http://httpbin.org/get")

	// explore response object
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())

*/

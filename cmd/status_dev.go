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
		SetAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjZlZTc2YjMzOWY4NjFjNGZjODViZGZhMjEwMjkxODFlM2Y5ZTdhZjkwMWM4MjhmYWRmZDE5YTlmMTVlNmMxM2I2ODIxNjM4Yzg0OGUyZDg0In0.eyJhdWQiOiIxIiwianRpIjoiNmVlNzZiMzM5Zjg2MWM0ZmM4NWJkZmEyMTAyOTE4MWUzZjllN2FmOTAxYzgyOGZhZGZkMTlhOWYxNWU2YzEzYjY4MjE2MzhjODQ4ZTJkODQiLCJpYXQiOjE1MTM2NzY2MTQsIm5iZiI6MTUxMzY3NjYxNCwiZXhwIjoxNTQ1MjEyNjE0LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.h1BpMf0Ak244I-pKsBNNjvEavEhhCzQwgLo-wk3UCjVY9DTJpwV6FzCOCIPzemyGnkiLneaBmaKQ1ExUxZtDWxqmYojiWB73KDL2RMKuhs3odZVGNp73JJeAdDB8DJWw_SHBlVdsp2I45xnHfU6ho3RdLhbO0AEeUDsxRcpnw_HIYgv47-r-uwqHZOLiA4h5BeC15kt4kKKGVceRTYXEY6-AhMFByxvEV6dRNILUQ91Bjt0Rn0dLsLEM05ywWISHrUrkKFKcklzHfHiz31nfYBlx05X6xXX5XyeIdFWSMkyWI33tHThTbiaUkuix_9wZ_385izoUne0vFnKVByAocqc7fv1jtsQfP_2Ca-Hp4rfXgbKirVjkM7ZU2KyD3VsGbNn83irCPTUcxLLLB2r4hh9wUf9Ddc1twRkhpV0jksqnfh4mUEQ-RC6PjWMRFjMi5LZS6Dxu05aqpBMD0Gbi8vE7IuJlwGEhem_nGThm1zJd0U1_0ScUC2UJ-8HFFxR_FMjlOXcA3QH6HT04wkkG5gT5i8LFo3MgxFptuYptpOAE5Vfz-em07gOJCh8pNB9tSvuoB55fSvJVb5Wt1skO9AUvV7_2m34HnGml433pPD717WVGSZOPbBaYDQ0PFSBD1YgLf50Wo7ho8hSdjmGZtpcR3nc5PjsIA0oqKrFiujU").
		//SetHeader("Content-Type", "multipart/form-data").
		SetFormData(map[string]string{
			"client_id":     "4",
			"client_secret": " uTBn7fxknAPpmw9AiEXyIro7X8mP0JhkqPtvBS28",
			"project_key":   "d41d8cd98f00b204e9800998ecf8427e"}).
		//Post("http://192.168.60.10/api/key/status")
		Post("https://requestb.in/16pfgiv1")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

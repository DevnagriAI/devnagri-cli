package main

import (
	"encoding/base64"
	//	"reflect"

	"gopkg.in/resty.v1"
)

func main() {
	saveResponseAndConvert()
}

func decodeBase64(cypher string) string {

	data, _ := base64.StdEncoding.DecodeString(cypher)

	return string(data)

}

func saveResponseAndConvert() {
	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "multipart/form-data").
		SetAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjY2MTY0MzEzNGYyNDRlNjEyMWE1MzAwMjAyZDFhMDVkYmY3MTFjYjQxZDg2NjUxNmU5ZjRiNDI4OTBkNGE0NDk5NDU5ODc2M2M4ZGNjYTI5In0.eyJhdWQiOiIzIiwianRpIjoiNjYxNjQzMTM0ZjI0NGU2MTIxYTUzMDAyMDJkMWEwNWRiZjcxMWNiNDFkODY2NTE2ZTlmNGI0Mjg5MGQ0YTQ0OTk0NTk4NzYzYzhkY2NhMjkiLCJpYXQiOjE1MTYxNzIwNDQsIm5iZiI6MTUxNjE3MjA0NCwiZXhwIjoxNTQ3NzA4MDQ0LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.HqVfac3VEcTFT79WuQ_9CH1Q3KWeRO9i6IyQYT7iLqz7S6JmvY0mmLo34awViYN_7jssUZCOXo_esz7juzO3mPXMS9ej_rbUQUH3CHw-JRnBOJ8sPgsg43yF79omiZpxWjn5-Brlq7cyqu28mpXENqkIV_1mXE214hIK0ap0BZHrQ8iQmbYuKpWlZMJnmNzbD2FJ2GtcatkEtSjZsnglXpWHnYSuVcSaFc59KE4TOqFeNh4Xln7puWKB2MJrjSGFMvtKlh5u2RlWDVOwM1tAs9hrrE6nOq6ncoF6ur51P26_BjByjZtlhEXvKKHjLJkYdWipB-yjyjl8uS5JciINzZkj6AnnmRHWFNU0dOKYb8E8FJv9Z5yDBz4EitDqa-LSOoHfYidyrBpJU_c4n04Fjo2RbzyMdRphk1NT9CdgteR973DUBNRCkI0z-OlORRTz8mh62SlshanA4jtFme3qYXph-J5mzVyRInVB0HtJh8LF9yIPnlX3ge0jjvpZSgrP_r7XCtPlYeih6KFK3sQc24vo-wPKTE6-UBLZSxsH3mJeJGzsbR9F59vbAaJ518-WgbQfMeAo2lUz_ebsB24Lvrc7rVMqdA1yEdwPepcCYuGO7-Afb3XzOFjSVlRCtFTHLxWWvJ221g1a_lS0ifxxjbVetFYxqw8xaWnWMwoI2Vk").
		SetFormData(map[string]string{
			"client_id":     "2",
			"client_secret": "y8umxMS54nUBc1ak7cxod6mjYiAbht2rCNAKsW7c",
			"project_key":   "ded491851ec93a79ee89460902aff582"}).
		Post("http://dev.devnagri.co.in/api/project/pull")
	if err != nil {
		panic(err)
	}
}

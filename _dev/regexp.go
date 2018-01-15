package main

import (
	"fmt"
	"regexp"
)

func main() {

	val, _ := regexp.MatchString(".*xyz", "abcd.xyz")
	fmt.Println(val)

}

// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// Description of the Build task
func Build() {

	fmt.Println("Executing Build!")
}

// Description of the Fmt task
func Fmt() error {

	cmdString := `find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done`
	fmt.Println(cmdString)

	sh.RunCmd(cmdString)

	return nil
}

// Runs the cover tool for the devnagri-cli tool
func Cover() {

	//	cover := sh.RunCmd("go tool cover")

	//	cover("-html=coverage.txt")

	cmdString := `go tool cover -html=coverage.txt`

	fmt.Println(cmdString)

	sh.RunCmd(cmdString)

}

// Run the linter

func Lint() {
	cmdString := ` gometalinter --vendor --disable-all \
		--enable=deadcode \
		--enable=ineffassign \
		--enable=gosimple \
		--enable=staticcheck \
		--enable=gofmt \
		--enable=goimports \
		--enable=dupl \
		--enable=misspell \
		--enable=errcheck \
		--enable=vet \
		--enable=vetshadow \
		--deadline=10m \
		./... `

	fmt.Println(cmdString)

	sh.RunCmd(cmdString)

}

// Creates the beta build
func Beta() {

	cmdString := `go build -o beta ./main.go`

	fmt.Println(cmdString)

	sh.RunCmd(cmdString)

}

// Removes the beta binary
func Rmbeta() {

	cmdString := `rm ./beta`

	fmt.Println(cmdString)

	sh.RunCmd(cmdString)

}

// TODO: https://godoc.org/github.com/magefile/mage/sh#Exec
// TODO: https://magefile.org/libraries/

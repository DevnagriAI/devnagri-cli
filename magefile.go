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
	return nil
}

// Runs the cover tool for the devnagri-cli tool
func Cover() {

	cover := sh.RunCmd("go tool cover")

	cover("-html=coverage.txt")

}

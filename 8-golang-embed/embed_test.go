package embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

//go:embed char.png
var char []byte

func TestString(t *testing.T) {
	fmt.Println(version)
}

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("char_next.png", char, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

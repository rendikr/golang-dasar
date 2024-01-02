package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed char.png
var char []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("char_next.png", char, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed a.txt
//go:embed b.txt
//go:embed c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, err := files.ReadFile("a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(a))
	b, err := files.ReadFile("b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	c, err := files.ReadFile("c.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(c))
}

//go:embed *.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content:", string(content))
		}
	}
}

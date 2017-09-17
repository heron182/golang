/*
***EXPORTING SYNTAX***
An identifier may be exported to permit access to it from another package. An identifier is exported if both:

the first character of the identifier's name is a Unicode upper case letter (Unicode class "Lu"); and
the identifier is declared in the package block or it is a field name or method name.
*/
package practice

import (
	"container/list"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CustomError() {
	err := errors.New("Custom error message")
	panic(err)
}

type Containers struct {
}

func (c *Containers) LinkedList() {
	x := list.New()
	x.PushBack(40)
	x.PushBack(10)
	x.PushBack("MixedTypes")
	for v := x.Front(); v != nil; v = v.Next() {
		fmt.Println(v.Value)
	}
}

type IoOperations struct {
}

func (iop *IoOperations) ReadFile() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// create a buffer with the filesize and read it
	if stat, err := file.Stat(); err == nil {
		bs := make([]byte, stat.Size())
		if _, err := file.Read(bs); err == nil {
			str := string(bs)
			fmt.Println(str)
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}

}

func (iop *IoOperations) SimplerReadFile() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	str := string(bs)
	fmt.Println(str)
}

func (iop *IoOperations) ReadDirectory() {
	dir, _ := os.Open(".")
	defer dir.Close()
	fileInfos, _ := dir.Readdir(-1) // all files
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func (iop *IoOperations) WalkDir() {
	filepath.Walk(".", func(path string, fi os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
func (iop *IoOperations) CreateFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("this is a test\n")

}

func StringsPkg() {
	// Contains(s, substr string) bool
	fmt.Println(strings.Contains("test", "es")) // true

	// Count(s, sep string) int
	fmt.Println(strings.Count("test", "t")) // 2

	// HasPrefix(string, prefix string) bool
	fmt.Println(strings.HasPrefix("test", "es")) // false

	// HasSuffix(string, suffix string) bool
	fmt.Println(strings.HasSuffix("test", "st")) // true

	// Index(string, sep string) int (-1 if not found)
	fmt.Println(strings.Index("test", "s")) // 2

	// func Join(a []string, sep string) string
	fmt.Println(strings.Join([]string{"a", "b"}, "~"))
	// => "a~b"

	// func Repeat(string, count int) string
	fmt.Println(strings.Repeat("t", 4)) // tttt

	// func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	// => "bbaa"

	// func Split(s, sep string) []string
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	// => []string{"a","b","c","d","e"}

	// func ToLower(s string) string
	fmt.Println(strings.ToLower("TEST"))
	// => "test"

	// func ToUpper(s string) string
	fmt.Println(strings.ToUpper("test"))
	// => "TEST"

	// Convert a string to a slice of bytes and vice-versa
	arr := []byte("test")
	str := string([]byte{'T', 'E', 'S', 'T'})
	fmt.Println(arr)
	fmt.Println(str)

}

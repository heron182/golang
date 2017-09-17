package main

import (
	"fmt"
    "chap8/practice"
    "flag"
)

func main() {
	practice.StringsPkg()
	// iop := practice.IoOperations{}
	// iop.CreateFile()
	// iop.ReadFile()
	// iop.SimplerReadFile()
	// iop.ReadDirectory()
	// iop.WalkDir()
	practice.CustomError()
	c := practice.Containers{}
	c.LinkedList()
    flag.Parse()
    wordPtr := flag.String("word", "foo", "a string")
    fmt.Println("Command line args passed: ", *wordPtr, flag.Args())
}

package main

import (
	"fmt"
	"os"
)

// user input:
// 1. stdin -> a read only file handle
// 2. os.Args -> []string (a list of variables of string type)
func main() {

	//x:=fmt.Sprintf()
	// print to stdout
	// \n is called a newline; it is the same as "enter"
	fmt.Printf("program=%s first arg=%s\n", os.Args[0], os.Args[1])
}

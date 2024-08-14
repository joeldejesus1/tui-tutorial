package main

import (
	"os"

	"github.com/sascha33/tui-tutorial/timezonecalc"
)

// This is the entry point from the OS into this program
// user input:
// 1. stdin -> a read only file handle
// 2. os.Args -> []string (a list of variables of string type)
// 3. environmental arguments (learn later)
func main() {

	// import the timezonecalc function into main
	timezonecalc.PrintCLIArgs(os.Args)
}

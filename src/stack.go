package main

import "C"
import (
	_ "github.com/benesch/cgosymbolizer"
	"runtime/debug"
)

func init() {
	debug.SetTraceback("crash")
}

//export goPrintStack
func goPrintStack() {
	debug.PrintStack()
}

func main() {}

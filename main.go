package main

import (
	"fmt"
	"os"

	"github.com/healthy-tiger/scalc/parser"
	"github.com/healthy-tiger/scalc/runtime"
)

func main() {
	st := parser.NewSymbolTable()
	ns := runtime.NewRootNamespace(st)
	runtime.MakeDefaultNamespace(ns)

	lst, err := parser.Parse("stdin", st, os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	for _, l := range lst {
		_, err := runtime.EvalList(l, ns)
		if err != nil {
			fmt.Println(err)
		}
	}
}

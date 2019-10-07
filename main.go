package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/healthy-tiger/scalc/parser"
	"github.com/healthy-tiger/scalc/runtime"
)

func main() {
	st := parser.NewSymbolTable()
	ns := runtime.NewRootNamespace(st)
	runtime.MakeDefaultNamespace(ns)

	flag.Parse()

	for _, v := range flag.Args() {
		f, err := os.Open(v)
		if err != nil {
			fmt.Println(err)
		}
		lst, err := parser.Parse(v, st, f)
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
}

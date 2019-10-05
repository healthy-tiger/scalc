package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/healthy-tiger/scalc/parser"
	"github.com/healthy-tiger/scalc/runtime"
)

func runEval(st *parser.SymbolTable, ns *runtime.Namespace) {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := ">> "
	fmt.Fprint(os.Stdout, prompt)
	for scanner.Scan() {
		line := scanner.Text()
		lists, err := parser.ParseString("<stdin>", st, line)
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			for i := 0; i < len(lists); i++ {
				result, err := runtime.EvalList(lists[i], ns)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: %v\n", err)
				} else {
					fmt.Printf("%v\n", result)
				}
			}
		}
		fmt.Fprint(os.Stdout, prompt)
	}
}

func main() {
	interactive := false
	flag.BoolVar(&interactive, "i", false, "interactive mode")
	flag.Parse()

	st := parser.NewSymbolTable()
	ns := runtime.NewNamespace(nil)
	runtime.DefaultNamespace(st, ns)

	args := flag.Args()
	if len(args) == 0 {
		interactive = true
	} else {
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

	if interactive {
		runEval(st, ns)
	}
}

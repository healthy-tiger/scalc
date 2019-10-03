package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/healthy-tiger/scalc/parser"
	"github.com/healthy-tiger/scalc/runtime"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	st := parser.NewSymbolTable()
	ns := runtime.NewNamespace(nil)
	runtime.DefaultNamespace(st, ns)
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

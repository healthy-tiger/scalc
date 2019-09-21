package scalc

import (
	"fmt"
	"testing"

	"github.com/healthy-tiger/gostree"
)

type optest struct {
	src      string
	expected interface{}
}

var addtests []optest = []optest{
	{`(+ 1 2 3)`, int64(6)},
	{`(+ 1 2 (+ 10 20) 3)`, int64(36)},
	{`(+ 1 2 (+ -10 -20) 3)`, int64(-24)},
	{`(+ 1.0 2.0 3.0)`, float64(6.0)},
	{`(+ 1 2 (+ 10.0 20) 3)`, int64(36)},
	{`(+ 1 2 (+ -10.0 -20.0) 3)`, int64(-24)},
	{`(+ 1.0 2 3)`, float64(6.0)},
	{`(+ 1 2.0 3.0)`, int64(6)},
}

func TestAdd(t *testing.T) {
	for i, tst := range addtests {
		st, err := gostree.ParseString(fmt.Sprintf("TestAdd%d", i), tst.src)
		if err != nil {
			t.Errorf("Parse error: %v\n", err)
		} else {
			globals := DefaultNamespace(st)
			result, err := EvalList(st.Lists[0], globals, globals)
			if err != nil {
				t.Fatalf("Eval error: %v\n", err)
			}
			success := false
			if ir, ok := result.(int64); ok {
				if ie, ok := tst.expected.(int64); ok && ir == ie {
					success = true
				}
			} else if fr, ok := result.(float64); ok {
				if fe, ok := tst.expected.(float64); ok && fr == fe {
					success = true
				}
			}
			if !success {
				t.Errorf("Result is %v, not %v", result, tst.expected)
			}
		}
	}
}

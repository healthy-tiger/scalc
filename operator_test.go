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
	{`(+ 1 2 3)`, int64(1) + int64(2) + int64(3)},
	{`(+ 1 2 (+ 10 20) 3)`, int64(1) + int64(2) + (int64(10) + int64(20)) + int64(3)},
	{`(+ 1 2 (+ -10 -20) 3)`, int64(1) + int64(2) + (int64(-10) + int64(-20)) + int64(3)},
	{`(+ 1.0 2.0 3.0)`, float64(1.0) + float64(2.0) + float64(3.0)},
	{`(+ 1 2 (+ 10.0 20) 3)`, float64(int64(1)+int64(2)) + (float64(10.0) + float64(20.0)) + float64(int64(3))},
	{`(+ 1 2 (+ -10.0 -20.0) 3)`, float64(int64(1)+int64(2)) + (float64(-10.0) + float64(-20.0)) + float64(int64(3))},
	{`(+ 1.0 2 3)`, (float64(1.0) + float64(int64(2))) + float64(int64(3))},
	{`(+ 1 2.1 3.0)`, float64(int64(1)) + float64(2.1) + float64(3.0)},
}

var subtests []optest = []optest{
	{`(- 1 2 3)`, int64(1) - int64(2) - int64(3)},
	{`(- 1 2 (- 10 20) 3)`, int64(1) - int64(2) - (int64(10) - int64(20)) - int64(3)},
	{`(- 1 2 (- -10 -20) 3)`, int64(1) - int64(2) - (int64(-10) - int64(-20)) - int64(3)},
	{`(- 1.0 2.0 3.0)`, float64(1.0) - float64(2.0) - float64(3.0)},
	{`(- 1 2 (- 10.0 20) 3)`, float64(int64(1)-int64(2)) - (float64(10.0) - float64(20.0)) - float64(int64(3))},
	{`(- 1 2 (- -10.0 -20.0) 3)`, float64(int64(1)-int64(2)) - (float64(-10.0) - float64(-20.0)) - float64(int64(3))},
	{`(- 1.0 2 3)`, (float64(1.0) - float64(int64(2))) - float64(int64(3))},
	{`(- 1 2.1 3.0)`, float64(int64(1)) - float64(2.1) - float64(3.0)},
}

var multests []optest = []optest{
	{`(* 1 2 3)`, int64(1) * int64(2) * int64(3)},
	{`(* 1 2 (* 10 20) 3)`, int64(1) * int64(2) * (int64(10) * int64(20)) * int64(3)},
	{`(* 1 2 (* -10 -20) 3)`, int64(1) * int64(2) * (int64(-10) * int64(-20)) * int64(3)},
	{`(* 1.0 2.0 3.0)`, float64(1.0) * float64(2.0) * float64(3.0)},
	{`(* 1 2 (* 10.0 20) 3)`, float64(int64(1)*int64(2)) * (float64(10.0) * float64(20.0)) * float64(int64(3))},
	{`(* 1 2 (* -10.0 -20.0) 3)`, float64(int64(1)*int64(2)) * (float64(-10.0) * float64(-20.0)) * float64(int64(3))},
	{`(* 1.0 2 3)`, (float64(1.0) * float64(int64(2))) * float64(int64(3))},
	{`(* 1 2.1 3.0)`, float64(int64(1)) * float64(2.1) * float64(3.0)},
}

var divtests []optest = []optest{
	{`(/ 1 2 3)`, int64(1) / int64(2) / int64(3)},
	{`(/ 1 2 (/ 10 20) 3)`, nil},
	{`(/ 1 2 (/ -10 -20) 3)`, nil},
	{`(/ 1.0 2.0 3.0)`, float64(1.0) / float64(2.0) / float64(3.0)},
	{`(/ 1 2 (/ 10.0 20) 3)`, float64(int64(1)/int64(2)) / (float64(10.0) / float64(20.0)) / float64(int64(3))},
	{`(/ 1 2 (/ -10.0 -20.0) 3)`, float64(int64(1)/int64(2)) / (float64(-10.0) / float64(-20.0)) / float64(int64(3))},
	{`(/ 1.0 2 3)`, (float64(1.0) / float64(int64(2))) / float64(int64(3))},
	{`(/ 1 2.1 3.0)`, float64(int64(1)) / float64(2.1) / float64(3.0)},
}

func doOpTests(name string, t *testing.T, tests []optest) {
	for i, tst := range tests {
		st, err := gostree.ParseString(fmt.Sprintf("%v%d", name, i), tst.src)
		if err != nil {
			t.Errorf("[%d]Parse error: %v\n", i, err)
		} else {
			ns := DefaultNamespace(st)
			result, err := EvalList(st.Lists[0], ns)
			if err != nil {
				t.Fatalf("[%d]Eval error: %v\n", i, err)
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
			} else if result == nil && tst.expected == nil {
				success = true
			}
			if !success {
				t.Errorf("[%d]The expected value was %v, but the result was %v.", i, tst.expected, result)
			}
		}
	}
}

func TestAdd(t *testing.T) {
	doOpTests("TestAdd", t, addtests)
}

func TestSub(t *testing.T) {
	doOpTests("TestSub", t, subtests)
}

func TestMul(t *testing.T) {
	doOpTests("TestMul", t, multests)
}

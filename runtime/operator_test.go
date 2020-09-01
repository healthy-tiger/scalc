package runtime_test

import (
	"fmt"
	"testing"

	"github.com/healthy-tiger/scalc/parser"
	"github.com/healthy-tiger/scalc/runtime"
)

type optest struct {
	src        string
	parseError bool // パースエラーが期待値ならtrue
	evalError  bool // 評価時エラーが期待値ならtrue
	expected   interface{}
}

var addtests []optest = []optest{
	{`(+ 1 2 3)`, false, false, int64(1) + int64(2) + int64(3)},
	{`(+ 1 2 (+ 10 20) 3)`, false, false, int64(1) + int64(2) + (int64(10) + int64(20)) + int64(3)},
	{`(+ 1 2 (+ -10 -20) 3)`, false, false, int64(1) + int64(2) + (int64(-10) + int64(-20)) + int64(3)},
	{`(+ 1.0 2.0 3.0)`, false, false, float64(1.0) + float64(2.0) + float64(3.0)},
	{`(+ 1 2 (+ 10.0 20) 3)`, false, true, nil},
	{`(+ 1 2 (+ -10.0 -20.0) 3)`, false, true, nil},
	{`(+ 1.0 2 3)`, false, true, nil},
	{`(+ 1 2.1 3.0)`, false, true, nil},
	{`(+ "abc" "123")`, false, true, nil},
	{`(+ "abc" 123)`, false, true, nil},
	{`(+ "abc" 1.0)`, false, true, nil},
	{`(+ 1 + "abc")`, false, true, nil},
	{`(+ "" 123)`, false, true, nil},
	{`(+ true 1)`, false, false, int64(2)},
	{`(+ 1 false)`, false, false, int64(1)},
	{`(+ true false)`, false, false, int64(1)},
	{`(+ true "true")`, false, true, nil},
	{`(+)`, false, true, nil},
}

var subtests []optest = []optest{
	{`(- 1 2 3)`, false, false, int64(1) - int64(2) - int64(3)},
	{`(- 1 2 (- 10 20) 3)`, false, false, int64(1) - int64(2) - (int64(10) - int64(20)) - int64(3)},
	{`(- 1 2 (- -10 -20) 3)`, false, false, int64(1) - int64(2) - (int64(-10) - int64(-20)) - int64(3)},
	{`(- 1.0 2.0 3.0)`, false, false, float64(1.0) - float64(2.0) - float64(3.0)},
	{`(- 1 2 (- 10.0 20) 3)`, false, true, nil},
	{`(- 1 2 (- -10.0 -20.0) 3)`, false, true, nil},
	{`(- 1.0 2 3)`, false, true, nil},
	{`(- 1 2.1 3.0)`, false, true, nil},
	{`(- true 1)`, false, false, int64(0)},
	{`(- 1 false)`, false, false, int64(1)},
	{`(- true false)`, false, false, int64(1)},
	{`(- true "true")`, false, true, nil},
	{`(-)`, false, true, nil},
}

var multests []optest = []optest{
	{`(* 1 2 3)`, false, false, int64(1) * int64(2) * int64(3)},
	{`(* 1 2 (* 10 20) 3)`, false, false, int64(1) * int64(2) * (int64(10) * int64(20)) * int64(3)},
	{`(* 1 2 (* -10 -20) 3)`, false, false, int64(1) * int64(2) * (int64(-10) * int64(-20)) * int64(3)},
	{`(* 1.0 2.0 3.0)`, false, false, float64(1.0) * float64(2.0) * float64(3.0)},
	{`(* 1 2 (* 10.0 20) 3)`, false, true, nil},
	{`(* 1 2 (* -10.0 -20.0) 3)`, false, true, nil},
	{`(* 1.0 2 3)`, false, true, nil},
	{`(* 1 2.1 3.0)`, false, true, nil},
	{`(* true 1)`, false, false, int64(1)},
	{`(* 1 false)`, false, false, int64(0)},
	{`(* true false)`, false, false, int64(0)},
	{`(* true "true")`, false, true, nil},
	{`(*)`, false, true, nil},
}

var divtests []optest = []optest{
	{`(/ 1 2 3)`, false, false, int64(1) / int64(2) / int64(3)},
	{`(/ 4 2)`, false, false, int64(4) / int64(2)},
	{`(/ 3 2)`, false, false, int64(3) / int64(2)},
	{`(/ 1 0)`, false, true, nil},
	{`(/ 1 2 (/ 10 20) 3)`, false, true, nil},
	{`(/ 1 2 (/ -10 -20) 3)`, false, true, nil},
	{`(/ 1.0 2.0 3.0)`, false, false, float64(1.0) / float64(2.0) / float64(3.0)},
	{`(/ 1 2 (/ 10.0 20) 3)`, false, true, nil},
	{`(/ 1 2 (/ -10.0 -20.0) 3)`, false, true, nil},
	{`(/ 1.0 2 3)`, false, true, nil},
	{`(/ 1 2.1 3.0)`, false, true, nil},
	{`(/ 1 0.0)`, false, true, nil},
	{`(/ true 1)`, false, false, int64(1)},
	{`(/ 1 false)`, false, true, nil},
	{`(/ true false)`, false, true, nil},
	{`(/ true "true")`, false, true, nil},
	{`(/)`, false, true, nil},
}

var remtests []optest = []optest{
	{`(% 1 2)`, false, false, int64(1)},
	{`(% 1 0)`, false, true, nil},
	{`(% 1 2 3)`, false, true, nil},
	{`(% 1.0 2)`, false, true, nil},
	{`(% 1 2.0)`, false, true, nil},
	{`(% "1" 2)`, false, true, nil},
	{`(% 1 "2")`, false, true, nil},
	{`(% true 1)`, false, false, int64(0)},
	{`(% 1 false)`, false, true, nil},
	{`(% true false)`, false, true, nil},
	{`(% true "true")`, false, true, nil},
	{`(% 1)`, false, true, nil},
	{`(%)`, false, true, nil},
}

var eqtests []optest = []optest{
	{`(eq 1 1)`, false, false, int64(1)},
	{`(eq 2 1)`, false, false, int64(0)},
	{`(eq 1.0 1.0)`, false, false, int64(1)},
	{`(eq 2.0 1.0)`, false, false, int64(0)},
	{`(eq "abc" "abc")`, false, false, int64(1)},
	{`(eq "abc" "123")`, false, false, int64(0)},
	{`(eq 1 1.0)`, false, true, nil},
	{`(eq 1 "1")`, false, true, nil},
	{`(eq 1 true)`, false, false, int64(1)},
	{`(eq 1)`, false, true, false},
}

var lttests []optest = []optest{
	{`(< 1 2)`, false, false, int64(1)},
	{`(< 2 1)`, false, false, int64(0)},
	{`(< 2 2)`, false, false, int64(0)},
	{`(< 1.0 2.0)`, false, false, int64(1)},
	{`(< 2.0 1.0)`, false, false, int64(0)},
	{`(< 2.0 2.0)`, false, false, int64(0)},
	{`(< 1 2.0)`, false, true, nil},
	{`(< 2 1.0)`, false, true, nil},
	{`(< 2 2.0)`, false, true, nil},
	{`(< 1.0 2)`, false, true, nil},
	{`(< 2.0 1)`, false, true, nil},
	{`(< 2.0 2)`, false, true, nil},
	{`(< 1 2 3)`, false, true, nil},
}

var ltetests []optest = []optest{
	{`(<= 1 2)`, false, false, int64(1)},
	{`(<= 2 1)`, false, false, int64(0)},
	{`(<= 2 2)`, false, false, int64(1)},
	{`(<= 1.0 2.0)`, false, false, int64(1)},
	{`(<= 2.0 1.0)`, false, false, int64(0)},
	{`(<= 2.0 2.0)`, false, false, int64(1)},
	{`(<= 1 2.0)`, false, true, nil},
	{`(<= 2 1.0)`, false, true, nil},
	{`(<= 2 2.0)`, false, true, nil},
	{`(<= 1.0 2)`, false, true, nil},
	{`(<= 2.0 1)`, false, true, nil},
	{`(<= 2.0 2)`, false, true, nil},
	{`(<= 1 2 3)`, false, true, nil},
}

var gttests []optest = []optest{
	{`(> 1 2)`, false, false, int64(0)},
	{`(> 2 1)`, false, false, int64(1)},
	{`(> 2 2)`, false, false, int64(0)},
	{`(> 1.0 2.0)`, false, false, int64(0)},
	{`(> 2.0 1.0)`, false, false, int64(1)},
	{`(> 2.0 2.0)`, false, false, int64(0)},
	{`(> 1 2.0)`, false, true, nil},
	{`(> 2 1.0)`, false, true, nil},
	{`(> 2 2.0)`, false, true, nil},
	{`(> 1.0 2)`, false, true, nil},
	{`(> 2.0 1)`, false, true, nil},
	{`(> 2.0 2)`, false, true, nil},
	{`(> 1 2 3)`, false, true, nil},
}

var gtetests []optest = []optest{
	{`(>= 1 2)`, false, false, int64(0)},
	{`(>= 2 1)`, false, false, int64(1)},
	{`(>= 2 2)`, false, false, int64(1)},
	{`(>= 1.0 2.0)`, false, false, int64(0)},
	{`(>= 2.0 1.0)`, false, false, int64(1)},
	{`(>= 2.0 2.0)`, false, false, int64(1)},
	{`(>= 1 2.0)`, false, true, nil},
	{`(>= 2 1.0)`, false, true, nil},
	{`(>= 2 2.0)`, false, true, nil},
	{`(>= 1.0 2)`, false, true, nil},
	{`(>= 2.0 1)`, false, true, nil},
	{`(>= 2.0 2)`, false, true, nil},
	{`(>= 1 2 3)`, false, true, nil},
}

var strtests = []optest{
	{`(str 1)`, false, false, fmt.Sprint(int64(1))},
	{`(str -10)`, false, false, fmt.Sprint(int64(-10))},
	{`(str 1.0)`, false, false, fmt.Sprint(float64(1.0))},
	{`(str 3.1415926535)`, false, false, fmt.Sprint(float64(3.1415926535))},
	{`(str true)`, false, false, fmt.Sprint(int64(1))},
	{`(str false)`, false, false, fmt.Sprint(int64(0))},
}

var inttests = []optest{
	{`(int 10)`, false, false, int64(10)},
	{`(int 10.0)`, false, false, int64(10.0)},
	{`(int "10")`, false, false, int64(10)},
	{`(int "10.0")`, false, true, nil},
	{`(int true)`, false, false, int64(1)},
	{`(int false)`, false, false, int64(0)},
}

var floattests = []optest{
	{`(float 10)`, false, false, float64(int64(10))},
	{`(float 10.0)`, false, false, float64(10.0)},
	{`(float "10")`, false, false, float64(10)},
	{`(float "10.0")`, false, false, float64(10)},
	{`(float true)`, false, false, float64(int64(1))},
	{`(float false)`, false, false, float64(int64(0))},
}

func doOpTests(name string, t *testing.T, tests []optest) {
	for i, tst := range tests {
		st := parser.NewSymbolTable()
		lists, err := parser.ParseString(fmt.Sprintf("%v%d", name, i), st, tst.src)
		if err != nil {
			if !tst.parseError {
				t.Errorf("[%d]Parse error: %v\n", i, err)
			}
		} else {
			ns := runtime.NewRootNamespace(st)
			runtime.MakeDefaultNamespace(ns)
			result, err := runtime.EvalList(lists[0], ns)
			if err != nil {
				if !tst.evalError {
					t.Errorf("[%d]Eval error: %v\n", i, err)
				}
			} else {
				success := false
				if ir, ok := result.(int64); ok {
					if ie, ok := tst.expected.(int64); ok && ir == ie {
						success = true
					}
				} else if fr, ok := result.(float64); ok {
					if fe, ok := tst.expected.(float64); ok && fr == fe {
						success = true
					}
				} else if sr, ok := result.(string); ok {
					if se, ok := tst.expected.(string); ok && sr == se {
						success = true
					}
				} else if br, ok := result.(bool); ok {
					if be, ok := tst.expected.(bool); ok && br == be {
						success = true
					}
				}
				if !success {
					t.Errorf("[%d]The expected value was %v, but the result was %v.", i, tst.expected, result)
				}
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

func TestDiv(t *testing.T) {
	doOpTests("TestDiv", t, divtests)
}

func TestRem(t *testing.T) {
	doOpTests("TestRem", t, remtests)
}

func TestEq(t *testing.T) {
	doOpTests("TestEq", t, eqtests)
}

func TestGt(t *testing.T) {
	doOpTests("TestGt", t, gttests)
}

func TestGte(t *testing.T) {
	doOpTests("TestGte", t, gtetests)
}

func TestLt(t *testing.T) {
	doOpTests("TestLt", t, lttests)
}

func TestLte(t *testing.T) {
	doOpTests("TestLte", t, ltetests)
}

func TestStr(t *testing.T) {
	doOpTests("TestStr", t, strtests)
}

func TestInt(t *testing.T) {
	doOpTests("TestInt", t, inttests)
}

func TestFloat(t *testing.T) {
	doOpTests("TestFloat", t, floattests)
}

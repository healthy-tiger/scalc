package runtime_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/healthy-tiger/gostree"
	"github.com/healthy-tiger/scalc"
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
	{`(+ 1 2 (+ 10.0 20) 3)`, false, false, float64(int64(1)+int64(2)) + (float64(10.0) + float64(20.0)) + float64(int64(3))},
	{`(+ 1 2 (+ -10.0 -20.0) 3)`, false, false, float64(int64(1)+int64(2)) + (float64(-10.0) + float64(-20.0)) + float64(int64(3))},
	{`(+ 1.0 2 3)`, false, false, (float64(1.0) + float64(int64(2))) + float64(int64(3))},
	{`(+ 1 2.1 3.0)`, false, false, float64(int64(1)) + float64(2.1) + float64(3.0)},
	{`(+ "abc" "123")`, false, false, "abc123"},
	{`(+ "abc" 123)`, false, false, "abc123"},
	{`(+ "abc" 1.0)`, false, false, "abc" + strconv.FormatFloat(1.0, 'e', -1, 64)},
	//{`(+ 1 + "abc")`, false, false, nil}, // 算術データ型が必要な胸のエラーがでるはず
	{`(+ "" 123)`, false, false, "123"},
}

var subtests []optest = []optest{
	{`(- 1 2 3)`, false, false, int64(1) - int64(2) - int64(3)},
	{`(- 1 2 (- 10 20) 3)`, false, false, int64(1) - int64(2) - (int64(10) - int64(20)) - int64(3)},
	{`(- 1 2 (- -10 -20) 3)`, false, false, int64(1) - int64(2) - (int64(-10) - int64(-20)) - int64(3)},
	{`(- 1.0 2.0 3.0)`, false, false, float64(1.0) - float64(2.0) - float64(3.0)},
	{`(- 1 2 (- 10.0 20) 3)`, false, false, float64(int64(1)-int64(2)) - (float64(10.0) - float64(20.0)) - float64(int64(3))},
	{`(- 1 2 (- -10.0 -20.0) 3)`, false, false, float64(int64(1)-int64(2)) - (float64(-10.0) - float64(-20.0)) - float64(int64(3))},
	{`(- 1.0 2 3)`, false, false, (float64(1.0) - float64(int64(2))) - float64(int64(3))},
	{`(- 1 2.1 3.0)`, false, false, float64(int64(1)) - float64(2.1) - float64(3.0)},
}

var multests []optest = []optest{
	{`(* 1 2 3)`, false, false, int64(1) * int64(2) * int64(3)},
	{`(* 1 2 (* 10 20) 3)`, false, false, int64(1) * int64(2) * (int64(10) * int64(20)) * int64(3)},
	{`(* 1 2 (* -10 -20) 3)`, false, false, int64(1) * int64(2) * (int64(-10) * int64(-20)) * int64(3)},
	{`(* 1.0 2.0 3.0)`, false, false, float64(1.0) * float64(2.0) * float64(3.0)},
	{`(* 1 2 (* 10.0 20) 3)`, false, false, float64(int64(1)*int64(2)) * (float64(10.0) * float64(20.0)) * float64(int64(3))},
	{`(* 1 2 (* -10.0 -20.0) 3)`, false, false, float64(int64(1)*int64(2)) * (float64(-10.0) * float64(-20.0)) * float64(int64(3))},
	{`(* 1.0 2 3)`, false, false, (float64(1.0) * float64(int64(2))) * float64(int64(3))},
	{`(* 1 2.1 3.0)`, false, false, float64(int64(1)) * float64(2.1) * float64(3.0)},
}

var divtests []optest = []optest{
	{`(/ 1 2 3)`, false, false, int64(1) / int64(2) / int64(3)},
	{`(/ 1 0)`, false, true, nil},
	{`(/ 1 2 (/ 10 20) 3)`, false, true, nil},
	{`(/ 1 2 (/ -10 -20) 3)`, false, true, nil},
	{`(/ 1.0 2.0 3.0)`, false, false, float64(1.0) / float64(2.0) / float64(3.0)},
	{`(/ 1 2 (/ 10.0 20) 3)`, false, false, float64(int64(1)/int64(2)) / (float64(10.0) / float64(20.0)) / float64(int64(3))},
	{`(/ 1 2 (/ -10.0 -20.0) 3)`, false, false, float64(int64(1)/int64(2)) / (float64(-10.0) / float64(-20.0)) / float64(int64(3))},
	{`(/ 1.0 2 3)`, false, false, (float64(1.0) / float64(int64(2))) / float64(int64(3))},
	{`(/ 1 2.1 3.0)`, false, false, float64(int64(1)) / float64(2.1) / float64(3.0)},
	{`(/ 1 0.0)`, false, false, math.Inf(1)},
}

var remtests []optest = []optest{
	{`(% 1 2)`, false, false, int64(1)},
	{`(% 1 0)`, false, true, nil},
	{`(% 1 2 3)`, false, true, nil},
	{`(% 1.0 2)`, false, true, nil},
	{`(% 1 2.0)`, false, true, nil},
	{`(% "1" 2)`, false, true, nil},
	{`(% 1 "2")`, false, true, nil},
}

func doOpTests(name string, t *testing.T, tests []optest) {
	for i, tst := range tests {
		st := gostree.NewSymbolTable()
		lists, err := gostree.ParseString(fmt.Sprintf("%v%d", name, i), st, tst.src)
		if err != nil {
			if !tst.parseError {
				t.Errorf("[%d]Parse error: %v\n", i, err)
			}
		} else {
			ns := scalc.DefaultNamespace(st)
			result, err := scalc.EvalList(lists[0], ns)
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

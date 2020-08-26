package runtime

import (
	"github.com/healthy-tiger/scalc/parser"
)

const (
	idivSYmbol = "div" // 整数同士の除算、引数をすべて整数型にしてから除算する。結果も必ず整数
	remSymbol  = "rem"
)

// sqrt, log, pow, exp, sin, cos, tan, atan, acos, asin, ??????

// idivBody 引数をすべて整数型に変換してから除算を行う。
func idivBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len(), 3)
	}
	// 引数をすべて評価する。
	params := make([]interface{}, lst.Len())
	for i := 1; i < lst.Len(); i++ {
		ev, err := EvalElement(lst.ElementAt(i), ns)
		if err != nil {
			return nil, err
		}
		params[i] = ev
	}

	// 最初の引数を整数型に変換する。
	var result int64
	switch ir := params[1].(type) {
	case int64:
		result = ir
	case float64:
		result = int64(ir)
	default:
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric, params[1])
	}
	// ２番目以降の引数を整数型に変換しながら剰余を求めていく。
	for i := 2; i < lst.Len(); i++ {
		var t int64
		switch v := params[i].(type) {
		case int64:
			t = v
		case float64:
			t = int64(v)
		default:
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i])
		}
		if t == 0 {
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorDivisionByZero)
		}
		result = result / t
	}
	return result, nil
}

// remBody 引数をすべて整数型に変換してから剰余を求める。
func remBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len(), 3)
	}
	// 引数をすべて評価する。
	params := make([]interface{}, lst.Len())
	for i := 1; i < lst.Len(); i++ {
		ev, err := EvalElement(lst.ElementAt(i), ns)
		if err != nil {
			return nil, err
		}
		params[i] = ev
	}

	// 最初の引数を整数型に変換する。
	var result int64
	switch ir := params[1].(type) {
	case int64:
		result = ir
	case float64:
		result = int64(ir)
	default:
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric, params[1])
	}
	// ２番目以降の引数を整数型に変換しながら剰余を求めていく。
	for i := 2; i < lst.Len(); i++ {
		var t int64
		switch v := params[i].(type) {
		case int64:
			t = v
		case float64:
			t = int64(v)
		default:
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i])
		}
		if t == 0 {
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorDivisionByZero)
		}
		result = result % t
	}
	return result, nil
}

// RegisterMath stに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterMath(ns *Namespace) {
	ns.RegisterExtension(idivSYmbol, nil, idivBody)
	ns.RegisterExtension(remSymbol, nil, remBody)
}

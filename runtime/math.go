package runtime

import (
	"github.com/healthy-tiger/scalc/parser"
)

const (
	remSymbol = "rem" // 整数同士の剰余
)

// sqrt, log, pow, exp, sin, cos, tan, atan, acos, asin, ??????

// remBody 整数同士の剰余。整数でない引数が含まれる場合はエラー
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

	v, ok := params[1].(parser.SInt)
	if ok {
		for i := 2; i < lst.Len(); i++ {
			if iv, ok := params[i].(parser.SInt); ok {
				if iv == 0 {
					return nil, newEvalError(lst.ElementAt(i).Position(), ErrorDivisionByZero)
				}
				v = v % iv
			} else {
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeOfIntegerType, params[i])
			}
		}
		return v, nil
	}
	return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, v)
}

// RegisterMath stに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterMath(ns *Namespace) {
	ns.RegisterExtension(remSymbol, nil, remBody)
}

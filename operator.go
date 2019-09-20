package scalc

import (
	"errors"

	"github.com/healthy-tiger/gostree"
)

const (
	addSymbol = "+"
)

var (
	errorOperantIsATypeOfDataThatCannotBeAdded = errors.New("Operant is a type of data that cannot be added")
)

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらの合計を返す。
func addBody(_ interface{}, lst *gostree.List, locals *Namespace, globals *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, errorInsufficientNumberOfArguments
	}
	// 引数をすべて評価する。
	params := make([]interface{}, lst.Len())
	for i := 1; i < lst.Len(); i++ {
		ev, err := EvalElement(lst.ElementAt(i), locals, globals)
		if err != nil {
			return nil, err
		}
		params[i] = ev
	}

	// 引数を最初の引数の型に合わせながらすべて加算する。
	result := params[1]
	for i := 2; i < lst.Len(); i++ {
		switch v := result.(type) {
		case int64:
			if iv, ok := params[i].(int64); ok {
				result = v + iv
			} else if fv, ok := params[i].(float64); ok {
				result = v + int64(fv)
			} else {
				return nil, errorOperantIsATypeOfDataThatCannotBeAdded
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v + float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v + fv
			} else {
				return nil, errorOperantIsATypeOfDataThatCannotBeAdded
			}
		default:
			return nil, errorOperantIsATypeOfDataThatCannotBeAdded
		}
	}
	return result, nil
}

var addOperator = Extension{nil, addBody}

// RegisterOperators streeに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterOperators(stree *gostree.STree, ns *Namespace) {
	addid := stree.GetSymbolID(addSymbol)
	ns.Set(addid, &addOperator)
}

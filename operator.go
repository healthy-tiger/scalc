package scalc

import (
	"errors"

	"github.com/healthy-tiger/gostree"
)

type operatorAdd struct{}

const (
	addSymbol = "+"
)

var (
	errorOperantIsATypeOfDataThatCannotBeAdded = errors.New("Operant is a type of data that cannot be added")
)

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらの合計を返す。
func (op *operatorAdd) Eval(lst *gostree.List, locals *Namespace, globals *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, errorInsufficientNumberOfArguments
	}
	result, err := EvalElement(lst.ElementAt(1), locals, globals)
	if err != nil {
		return nil, err
	}
	for i := 2; i < lst.Len(); i++ {
		e := lst.ElementAt(i)
		switch v := result.(type) {
		case int64:
			if iv, ok := e.IntValue(); ok {
				result = v + iv
			} else if fv, ok := e.FloatValue(); ok {
				result = v + int64(fv)
			} else {
				return nil, errorOperantIsATypeOfDataThatCannotBeAdded
			}
		case float64:
			if iv, ok := e.IntValue(); ok {
				result = v + float64(iv)
			} else if fv, ok := e.FloatValue(); ok {
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

var addOperator operatorAdd

// RegisterOperators streeに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterOperators(stree *gostree.STree, ns *Namespace) {
	addid := stree.GetSymbolID(addSymbol)
	ns.Set(addid, addOperator)
}

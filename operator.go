package scalc

import (
	"errors"
	"strconv"

	"github.com/healthy-tiger/gostree"
)

const (
	addSymbol = "+"
	subSymbol = "-"
	mulSymbol = "*"
	divSymbol = "/"
	remSymbol = "%"
)

var (
	errorOperantMustHaveArithmeticDataType                     = errors.New("Operant must have arithmetic data type")
	errorTheRemainderCannotBeCalculatedUnlessItIsAnIntegerType = errors.New("The remainder cannot be calculated unless it is an integer type")
)

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらのすべてを加算した結果を返す。stringの場合にはオペラントすべてを既定の形式で文字列に変換したものをすべて連結した文字列を返す。
func addBody(_ interface{}, lst *gostree.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, errorInsufficientNumberOfArguments
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
				return nil, errorOperantMustHaveArithmeticDataType
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v + float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v + fv
			} else {
				return nil, errorOperantMustHaveArithmeticDataType
			}
		case string:
			if sv, ok := params[i].(string); ok {
				result = v + sv
			} else if iv, ok := params[i].(int64); ok {
				result = v + strconv.FormatInt(iv, 10)
			} else if fv, ok := params[i].(float64); ok {
				result = v + strconv.FormatFloat(fv, 'e', -1, 64)
			} else if bv, ok := params[i].(bool); ok {
				if bv {
					result = v + trueSymbol
				} else {
					result = v + falseSymbol
				}
			} else {
				return nil, errorOperantMustHaveArithmeticDataType
			}
		default:
			return nil, errorOperantMustHaveArithmeticDataType
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを減算した結果を返す。
func subBody(_ interface{}, lst *gostree.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, errorInsufficientNumberOfArguments
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

	// 引数を最初の引数の型に合わせながらすべて加算する。
	result := params[1]
	for i := 2; i < lst.Len(); i++ {
		switch v := result.(type) {
		case int64:
			if iv, ok := params[i].(int64); ok {
				result = v - iv
			} else if fv, ok := params[i].(float64); ok {
				result = v - int64(fv)
			} else {
				return nil, errorOperantMustHaveArithmeticDataType
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v - float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v - fv
			} else {
				return nil, errorOperantMustHaveArithmeticDataType
			}
		default:
			return nil, errorOperantMustHaveArithmeticDataType
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを乗算した結果を返す。
func mulBody(_ interface{}, lst *gostree.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, errorInsufficientNumberOfArguments
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

	// 引数を最初の引数の型に合わせながらすべて加算する。
	result := params[1]
	for i := 2; i < lst.Len(); i++ {
		switch v := result.(type) {
		case int64:
			if iv, ok := params[i].(int64); ok {
				result = v * iv
			} else if fv, ok := params[i].(float64); ok {
				result = v * int64(fv)
			} else {
				return nil, errorOperantMustHaveArithmeticDataType
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v * float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v * fv
			} else {
				return nil, errorOperantMustHaveArithmeticDataType
			}
		default:
			return nil, errorOperantMustHaveArithmeticDataType
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを除算した結果を返す。
func divBody(_ interface{}, lst *gostree.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, errorInsufficientNumberOfArguments
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

	// 引数を最初の引数の型に合わせながらすべて加算する。
	result := params[1]
	for i := 2; i < lst.Len(); i++ {
		switch v := result.(type) {
		case int64:
			if iv, ok := params[i].(int64); ok {
				result = v / iv
			} else if fv, ok := params[i].(float64); ok {
				result = v / int64(fv)
			} else {
				return nil, errorOperantMustHaveArithmeticDataType
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v / float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v / fv
			} else {
				return nil, errorOperantMustHaveArithmeticDataType
			}
		default:
			return nil, errorOperantMustHaveArithmeticDataType
		}
	}
	return result, nil
}

// Eval 2つのオペラントの評価結果がすべてint64の場合に最初の引数を次の引数で割った余りを返す。
func remBody(_ interface{}, lst *gostree.List, ns *Namespace) (interface{}, error) {
	// 要するにオペラントは2つしか許容しない
	if lst.Len() < 3 {
		return nil, errorInsufficientNumberOfArguments
	} else if lst.Len() > 3 {
		return nil, errorTooManyArguments
	}
	// 引数をすべて評価する。
	pa, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	pb, err := EvalElement(lst.ElementAt(2), ns)
	if err != nil {
		return nil, err
	}
	if a, ok := pa.(int64); ok {
		if b, ok := pb.(int64); ok {
			return a % b, nil
		}
	}
	return nil, errorTheRemainderCannotBeCalculatedUnlessItIsAnIntegerType
}

// RegisterOperators streeに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterOperators(stree *gostree.STree, ns *Namespace) {
	RegisterExtension(stree, ns, addSymbol, nil, addBody)
	RegisterExtension(stree, ns, subSymbol, nil, subBody)
	RegisterExtension(stree, ns, mulSymbol, nil, mulBody)
	RegisterExtension(stree, ns, divSymbol, nil, divBody)
	RegisterExtension(stree, ns, remSymbol, nil, remBody)
}

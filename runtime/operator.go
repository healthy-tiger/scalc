package runtime

import (
	"strconv"

	"github.com/healthy-tiger/scalc/parser"
)

const (
	addSymbol = "+"
	subSymbol = "-"
	mulSymbol = "*"
	divSymbol = "/"
	remSymbol = "%"
)

// 演算子に関するエラーコード
const (
	ErrorOperantMustHaveArithmeticDataType                     = iota
	ErrorTheRemainderCannotBeCalculatedUnlessItIsAnIntegerType = iota
	ErrorIntegerDivideByZero                                   = iota
)

var operatorExtName = "operators"
var operatorErrorMessages map[int]string

func init() {
	operatorErrorMessages = map[int]string{
		ErrorOperantMustHaveArithmeticDataType:                     "Operant must have arithmetic data type",
		ErrorTheRemainderCannotBeCalculatedUnlessItIsAnIntegerType: "The remainder cannot be calculated unless it is an integer type",
		ErrorIntegerDivideByZero:                                   "integer divide by zero",
	}
}

func newOpError(loc parser.Position, id int) *EvalError {
	if _, ok := operatorErrorMessages[id]; !ok {
		panic("Undefined error id")
	}
	return &EvalError{loc, operatorExtName, id, operatorErrorMessages}
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらのすべてを加算した結果を返す。stringの場合にはオペラントすべてを既定の形式で文字列に変換したものをすべて連結した文字列を返す。
func addBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, newOpError(lst.Position(), ErrorInsufficientNumberOfArguments)
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
				result = float64(v) + fv
			} else {
				return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType)
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v + float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v + fv
			} else {
				return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType)
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
				return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType) // TODO ここは違うエラーにすべき
			}
		default:
			return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType) // TODO ここは違うエラーにすべき
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを減算した結果を返す。
func subBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, newOpError(lst.Position(), ErrorInsufficientNumberOfArguments)
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
				result = float64(v) - fv
			} else {
				return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType)
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v - float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v - fv
			} else {
				return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType)
			}
		default:
			return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType)
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを乗算した結果を返す。
func mulBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, newOpError(lst.Position(), ErrorInsufficientNumberOfArguments)
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
				result = float64(v) * fv
			} else {
				return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType)
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v * float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v * fv
			} else {
				return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType)
			}
		default:
			return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType)
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを除算した結果を返す。
func divBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, newOpError(lst.Position(), ErrorInsufficientNumberOfArguments)
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
				if iv == 0 { // ゼロ割のチェック
					return nil, newOpError(lst.ElementAt(i).Position(), ErrorIntegerDivideByZero)
				}
				result = v / iv
			} else if fv, ok := params[i].(float64); ok {
				result = float64(v) / fv
			} else {
				return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType)
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v / float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v / fv
			} else {
				return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType)
			}
		default:
			return nil, newOpError(lst.ElementAt(i).Position(), ErrorOperantMustHaveArithmeticDataType)
		}
	}
	return result, nil
}

// Eval 2つのオペラントの評価結果がすべてint64の場合に最初の引数を次の引数で割った余りを返す。
func remBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// 要するにオペラントは2つしか許容しない
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
	} else if lst.Len() > 3 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments)
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
			if b == 0 {
				return nil, newOpError(lst.ElementAt(2).Position(), ErrorIntegerDivideByZero)
			}
			return a % b, nil
		}
		return nil, newOpError(lst.ElementAt(2).Position(), ErrorTheRemainderCannotBeCalculatedUnlessItIsAnIntegerType)
	}
	return nil, newOpError(lst.ElementAt(1).Position(), ErrorTheRemainderCannotBeCalculatedUnlessItIsAnIntegerType)
}

// RegisterOperators streeに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterOperators(st *parser.SymbolTable, ns *Namespace) {
	RegisterExtension(st, ns, addSymbol, nil, addBody)
	RegisterExtension(st, ns, subSymbol, nil, subBody)
	RegisterExtension(st, ns, mulSymbol, nil, mulBody)
	RegisterExtension(st, ns, divSymbol, nil, divBody)
	RegisterExtension(st, ns, remSymbol, nil, remBody)
}

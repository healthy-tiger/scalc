package runtime

import (
	"math"
	"strconv"

	"github.com/healthy-tiger/scalc/parser"
)

const (
	addSymbol        = "+"
	subSymbol        = "-"
	mulSymbol        = "*"
	divSymbol        = "/"
	bitwiseANDSymbol = "band"
	bitwiseORSymbol  = "bor"
	bitwiseXORSymbol = "bxor"
	lShiftSymbol     = "lshift"
	rShiftSymbol     = "rshift"
	eqSymbol         = "eq" // すべての引数の型と値が一致した場合にtrueになる
	ltSymbol         = "<"
	lteSymbol        = "<="
	gtSymbol         = ">"
	gteSymbol        = ">="
	notSymbol        = "not"
	andSymbol        = "and"
	orSymbol         = "or"
)

// 演算子に関するエラーコード
var (
	ErrorOperantsMustBeNumeric          int
	ErrorOperantsMustBeOfIntegerType    int
	ErrorIntegerDivideByZero            int
	ErrorAllOperantsMustBeOfTheSameType int
	ErrorOperantsMustBeBoolean          int
)

func init() {
	ErrorOperantsMustBeNumeric = RegisterEvalError("Operants must be numeric")
	ErrorOperantsMustBeOfIntegerType = RegisterEvalError("Operants must be of integer type")
	ErrorIntegerDivideByZero = RegisterEvalError("integer divide by zero")
	ErrorAllOperantsMustBeOfTheSameType = RegisterEvalError("All operants must be of the same type")
	ErrorOperantsMustBeBoolean = RegisterEvalError("Operants must be boolen")
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらのすべてを加算した結果を返す。stringの場合にはオペラントすべてを既定の形式で文字列に変換したものをすべて連結した文字列を返す。
func addBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
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
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric)
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v + float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v + fv
			} else {
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric)
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
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric) // TODO ここは違うエラーにすべき
			}
		default:
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric) // TODO ここは違うエラーにすべき
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを減算した結果を返す。
func subBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
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
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric)
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v - float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v - fv
			} else {
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric)
			}
		default:
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric)
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを乗算した結果を返す。
func mulBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
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
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric)
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v * float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v * fv
			} else {
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric)
			}
		default:
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric)
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを除算した結果を返す。
func divBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
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
	var result float64
	switch lv := params[1].(type) {
	case int64:
		result = float64(lv)
	case float64:
		result = lv
	default:
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric)
	}
	for i := 2; i < lst.Len(); i++ {
		switch rv := params[i].(type) {
		case int64:
			result = result / float64(rv)
		case float64:
			result = result / rv
		default:
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric)
		}
	}

	// 整数型に変換できそうならしてから返す。
	if math.Abs(result-math.Trunc(result)) > 0 {
		return result, nil
	}
	return int64(result), nil
}

func eqBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
	}
	l, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	for i := 2; i < lst.Len(); i++ {
		r, err := EvalElement(lst.ElementAt(i), ns)
		if err != nil {
			return nil, err
		} else if l != r {
			return false, nil
		}
	}
	return true, nil
}

func bitwiseANDbody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
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

	result, ok := params[1].(int64)
	if !ok {
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType)
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType)
		}
		result = result & ip
	}
	return result, nil
}

func bitwiseORbody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
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

	result, ok := params[1].(int64)
	if !ok {
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType)
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType)
		}
		result = result | ip
	}
	return result, nil
}

func bitwiseXORbody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
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

	result, ok := params[1].(int64)
	if !ok {
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType)
	}
	if lst.Len() == 2 { // 引数が一つのときはビットを反転させて返す。
		return ^result, nil
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType)
		}
		result = result ^ ip
	}
	return result, nil
}

func lShiftBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
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

	result, ok := params[1].(int64)
	if !ok {
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType)
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType)
		}
		result = result << ip
	}
	return result, nil
}

func rShiftBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
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

	result, ok := params[1].(int64)
	if !ok {
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType)
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType)
		}
		result = result >> ip
	}
	return result, nil
}

func ltBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	switch a := pa.(type) {
	case int64:
		switch b := pb.(type) {
		case int64:
			return a < b, nil
		case float64:
			return float64(a) < b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric)
		}
	case float64:
		switch b := pb.(type) {
		case int64:
			return a < float64(b), nil
		case float64:
			return a < b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric)
		}
	default:
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric)
	}
}

func lteBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	switch a := pa.(type) {
	case int64:
		switch b := pb.(type) {
		case int64:
			return a <= b, nil
		case float64:
			return float64(a) <= b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric)
		}
	case float64:
		switch b := pb.(type) {
		case int64:
			return a <= float64(b), nil
		case float64:
			return a <= b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric)
		}
	default:
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric)
	}
}

func gtBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	switch a := pa.(type) {
	case int64:
		switch b := pb.(type) {
		case int64:
			return a > b, nil
		case float64:
			return float64(a) > b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric)
		}
	case float64:
		switch b := pb.(type) {
		case int64:
			return a > float64(b), nil
		case float64:
			return a > b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric)
		}
	default:
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric)
	}
}

func gteBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	switch a := pa.(type) {
	case int64:
		switch b := pb.(type) {
		case int64:
			return a >= b, nil
		case float64:
			return float64(a) >= b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric)
		}
	case float64:
		switch b := pb.(type) {
		case int64:
			return a >= float64(b), nil
		case float64:
			return a >= b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric)
		}
	default:
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric)
	}
}

func notBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// 要するに引数は必ず一つ
	if lst.Len() < 2 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
	} else if lst.Len() > 2 {
		return nil, newEvalError(lst.ElementAt(2).Position(), ErrorTooManyArguments)
	}
	p, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if b, ok := p.(bool); ok {
		return !b, nil
	}
	return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeBoolean)
}

func andBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
	}
	// 引数を順に評価し、評価結果がfalseになったところで止めてfalseを返す。
	for i := 1; i < lst.Len(); i++ {
		ev, err := EvalElement(lst.ElementAt(i), ns)
		if err != nil {
			return nil, err
		}
		bv, ok := ev.(bool)
		if !ok {
			// 評価結果がboolに変換できない場合はエラーになる。
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeBoolean)
		}
		if !bv {
			return false, nil
		}
	}
	return true, nil
}

func orBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
	}
	// 引数を順に評価し、評価結果がtrueになったところで止めてtrueを返す。
	for i := 1; i < lst.Len(); i++ {
		ev, err := EvalElement(lst.ElementAt(i), ns)
		if err != nil {
			return nil, err
		}
		bv, ok := ev.(bool)
		if !ok {
			// 評価結果がboolに変換できない場合はエラーになる。
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeBoolean)
		}
		if bv {
			return true, nil
		}
	}
	return false, nil
}

// RegisterOperators stに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterOperators(st *parser.SymbolTable, ns *Namespace) {
	RegisterExtension(st, ns, addSymbol, nil, addBody)
	RegisterExtension(st, ns, subSymbol, nil, subBody)
	RegisterExtension(st, ns, mulSymbol, nil, mulBody)
	RegisterExtension(st, ns, divSymbol, nil, divBody)
	RegisterExtension(st, ns, eqSymbol, nil, eqBody)
	RegisterExtension(st, ns, bitwiseANDSymbol, nil, bitwiseANDbody)
	RegisterExtension(st, ns, bitwiseORSymbol, nil, bitwiseORbody)
	RegisterExtension(st, ns, bitwiseXORSymbol, nil, bitwiseXORbody)
	RegisterExtension(st, ns, lShiftSymbol, nil, lShiftBody)
	RegisterExtension(st, ns, rShiftSymbol, nil, rShiftBody)
	RegisterExtension(st, ns, ltSymbol, nil, ltBody)
	RegisterExtension(st, ns, lteSymbol, nil, lteBody)
	RegisterExtension(st, ns, gtSymbol, nil, gtBody)
	RegisterExtension(st, ns, gteSymbol, nil, gteBody)
	RegisterExtension(st, ns, notSymbol, nil, notBody)
	RegisterExtension(st, ns, andSymbol, nil, andBody)
	RegisterExtension(st, ns, orSymbol, nil, orBody)
}

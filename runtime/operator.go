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
	ErrorDivisionByZero                 int
	ErrorAllOperantsMustBeOfTheSameType int
	ErrorOperantsMustBeBoolean          int
)

func init() {
	ErrorOperantsMustBeNumeric = RegisterEvalError("Operants must be numeric: \"%v\"")
	ErrorOperantsMustBeOfIntegerType = RegisterEvalError("Operants must be of integer type: \"%v\"")
	ErrorDivisionByZero = RegisterEvalError("Division by zero")
	ErrorAllOperantsMustBeOfTheSameType = RegisterEvalError("All operants must be of the same type")
	ErrorOperantsMustBeBoolean = RegisterEvalError("Operants must be boolen: \"%v\"")
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらのすべてを加算した結果を返す。stringの場合にはオペラントすべてを既定の形式で文字列に変換したものをすべて連結した文字列を返す。
func addBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i])
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v + float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v + fv
			} else {
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i])
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
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i]) // TODO ここは違うエラーにすべき
			}
		default:
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i]) // TODO ここは違うエラーにすべき
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを減算した結果を返す。
func subBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i])
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v - float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v - fv
			} else {
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i])
			}
		default:
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i])
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを乗算した結果を返す。
func mulBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i])
			}
		case float64:
			if iv, ok := params[i].(int64); ok {
				result = v * float64(iv)
			} else if fv, ok := params[i].(float64); ok {
				result = v * fv
			} else {
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i])
			}
		default:
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i])
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはfloat64の値の場合にそれらすべてを除算した結果を返す。
func divBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric, params[1])
	}
	for i := 2; i < lst.Len(); i++ {
		switch rv := params[i].(type) {
		case int64:
			if rv == 0 {
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorDivisionByZero, nil)
			}
			result = result / float64(rv)
		case float64:
			if rv == 0.0 {
				return nil, newEvalError(lst.ElementAt(i).Position(), ErrorDivisionByZero, nil)
			}
			result = result / rv
		default:
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeNumeric, params[i])
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
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[1])
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[i])
		}
		result = result & ip
	}
	return result, nil
}

func bitwiseORbody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[1])
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[i])
		}
		result = result | ip
	}
	return result, nil
}

func bitwiseXORbody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[1])
	}
	if lst.Len() == 2 { // 引数が一つのときはビットを反転させて返す。
		return ^result, nil
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[i])
		}
		result = result ^ ip
	}
	return result, nil
}

func lShiftBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[1])
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[i])
		}
		result = result << ip
	}
	return result, nil
}

func rShiftBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[1])
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[i])
		}
		result = result >> ip
	}
	return result, nil
}

func ltBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// 要するにオペラントは2つしか許容しない
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
	} else if lst.Len() > 3 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments, nil)
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
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric, pb)
		}
	case float64:
		switch b := pb.(type) {
		case int64:
			return a < float64(b), nil
		case float64:
			return a < b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric, pb)
		}
	default:
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric, pa)
	}
}

func lteBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// 要するにオペラントは2つしか許容しない
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
	} else if lst.Len() > 3 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments, nil)
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
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric, pb)
		}
	case float64:
		switch b := pb.(type) {
		case int64:
			return a <= float64(b), nil
		case float64:
			return a <= b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric, pb)
		}
	default:
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric, pa)
	}
}

func gtBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// 要するにオペラントは2つしか許容しない
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
	} else if lst.Len() > 3 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments, nil)
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
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric, pb)
		}
	case float64:
		switch b := pb.(type) {
		case int64:
			return a > float64(b), nil
		case float64:
			return a > b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric, pb)
		}
	default:
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric, pa)
	}
}

func gteBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// 要するにオペラントは2つしか許容しない
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
	} else if lst.Len() > 3 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments, nil)
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
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric, pb)
		}
	case float64:
		switch b := pb.(type) {
		case int64:
			return a >= float64(b), nil
		case float64:
			return a >= b, nil
		default:
			return nil, newEvalError(lst.ElementAt(2).Position(), ErrorOperantsMustBeNumeric, pb)
		}
	default:
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeNumeric, pa)
	}
}

func notBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// 要するに引数は必ず一つ
	if lst.Len() < 2 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
	} else if lst.Len() > 2 {
		return nil, newEvalError(lst.ElementAt(2).Position(), ErrorTooManyArguments, nil)
	}
	p, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if b, ok := p.(bool); ok {
		return !b, nil
	}
	return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeBoolean, p)
}

func andBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeBoolean, ev)
		}
		if !bv {
			return false, nil
		}
	}
	return true, nil
}

func orBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
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
			return nil, newEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeBoolean, ev)
		}
		if bv {
			return true, nil
		}
	}
	return false, nil
}

// RegisterOperators stに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterOperators(ns *Namespace) {
	ns.RegisterExtension(addSymbol, nil, addBody)
	ns.RegisterExtension(subSymbol, nil, subBody)
	ns.RegisterExtension(mulSymbol, nil, mulBody)
	ns.RegisterExtension(divSymbol, nil, divBody)
	ns.RegisterExtension(eqSymbol, nil, eqBody)
	ns.RegisterExtension(bitwiseANDSymbol, nil, bitwiseANDbody)
	ns.RegisterExtension(bitwiseORSymbol, nil, bitwiseORbody)
	ns.RegisterExtension(bitwiseXORSymbol, nil, bitwiseXORbody)
	ns.RegisterExtension(lShiftSymbol, nil, lShiftBody)
	ns.RegisterExtension(rShiftSymbol, nil, rShiftBody)
	ns.RegisterExtension(ltSymbol, nil, ltBody)
	ns.RegisterExtension(lteSymbol, nil, lteBody)
	ns.RegisterExtension(gtSymbol, nil, gtBody)
	ns.RegisterExtension(gteSymbol, nil, gteBody)
	ns.RegisterExtension(notSymbol, nil, notBody)
	ns.RegisterExtension(andSymbol, nil, andBody)
	ns.RegisterExtension(orSymbol, nil, orBody)
}

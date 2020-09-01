package runtime

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/healthy-tiger/scalc/parser"
)

const (
	addSymbol        = "+"
	subSymbol        = "-"
	mulSymbol        = "*"
	divSymbol        = "/"
	remSymbol        = "%" // 整数同士の剰余
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
	strSymbol        = "str"
	intSymbol        = "int"
	floatSymbol      = "float"
)

// 演算子に関するエラーコード
var (
	ErrorTypeMissmatch                  int
	ErrorOperantsMustBeNumeric          int
	ErrorOperantsMustBeOfIntegerType    int
	ErrorOperantsMustBeOfFloatType      int
	ErrorDivisionByZero                 int
	ErrorAllOperantsMustBeOfTheSameType int
	ErrorNonArithmeticDataType          int
)

func init() {
	ErrorTypeMissmatch = RegisterEvalError("Type missmatch (%v, %v)")
	ErrorOperantsMustBeNumeric = RegisterEvalError("Operants must be numeric: %v")
	ErrorOperantsMustBeOfIntegerType = RegisterEvalError("Operants must be of integer type: %v")
	ErrorOperantsMustBeOfFloatType = RegisterEvalError("Operants must be of float type: %v")
	ErrorDivisionByZero = RegisterEvalError("Division by zero")
	ErrorAllOperantsMustBeOfTheSameType = RegisterEvalError("All operants must be of the same type")
	ErrorNonArithmeticDataType = RegisterEvalError("Non-arithmetic data type: '%v)")
}

func isArithmeticDataType(v *interface{}) bool {
	switch (*v).(type) {
	case int64, float64:
		return true
	default:
		return false
	}
}

func isSameType(a *interface{}, b *interface{}) bool {
	switch (*a).(type) {
	case int64:
		if _, ok := (*b).(int64); ok {
			return true
		}
	case float64:
		if _, ok := (*b).(float64); ok {
			return true
		}
	case *Function:
		if _, ok := (*b).(*Function); ok {
			return true
		}
	case string:
		if _, ok := (*b).(string); ok {
			return true
		}
	default:
		panic("Unexpected Data Type.")
	}
	return false
}

// Eval オペラントの評価結果がすべてint64、すべてfloat64の場合にそれらのすべてを加算（または連結）した結果を返す。
func addBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 1)
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

	result := params[1]
	if !isArithmeticDataType(&result) {
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorNonArithmeticDataType, reflect.TypeOf(result))
	}
	for i := 2; i < lst.Len(); i++ {
		b := params[i]
		if !isArithmeticDataType(&b) {
			return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorNonArithmeticDataType, reflect.TypeOf(b))
		}
		if !isSameType(&result, &b) {
			return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorTypeMissmatch, reflect.TypeOf(result), reflect.TypeOf(b))
		}
		switch v := result.(type) {
		case int64:
			result = v + b.(int64)
		case float64:
			result = v + b.(float64)
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはすべてfloat64の値の場合にそれらすべてを減算した結果を返す。
func subBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 1)
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

	result := params[1]
	if !isArithmeticDataType(&result) {
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorNonArithmeticDataType, reflect.TypeOf(result))
	}
	for i := 2; i < lst.Len(); i++ {
		b := params[i]
		if !isArithmeticDataType(&b) {
			return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorNonArithmeticDataType, reflect.TypeOf(b))
		}
		if !isSameType(&result, &b) {
			return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorTypeMissmatch, reflect.TypeOf(result), reflect.TypeOf(b))
		}
		switch v := result.(type) {
		case int64:
			result = v - b.(int64)
		case float64:
			result = v - b.(float64)
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはすべてfloat64の値の場合にそれらすべてを乗算した結果を返す。
func mulBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 1)
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

	result := params[1]
	if !isArithmeticDataType(&result) {
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorNonArithmeticDataType, reflect.TypeOf(result))
	}
	for i := 2; i < lst.Len(); i++ {
		b := params[i]
		if !isArithmeticDataType(&b) {
			return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorNonArithmeticDataType, reflect.TypeOf(b))
		}
		if !isSameType(&result, &b) {
			return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorTypeMissmatch, reflect.TypeOf(result), reflect.TypeOf(b))
		}
		switch v := result.(type) {
		case int64:
			result = v * b.(int64)
		case float64:
			result = v * b.(float64)
		}
	}
	return result, nil
}

// Eval オペラントの評価結果がすべてint64またはすべてfloat64の値の場合にそれらすべてを除算した結果を返す。
func divBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 1)
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

	result := params[1]
	if !isArithmeticDataType(&result) {
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorNonArithmeticDataType, reflect.TypeOf(result))
	}
	for i := 2; i < lst.Len(); i++ {
		b := params[i]
		if !isArithmeticDataType(&b) {
			return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorNonArithmeticDataType, reflect.TypeOf(b))
		}
		if !isSameType(&result, &b) {
			return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorTypeMissmatch, reflect.TypeOf(result), reflect.TypeOf(b))
		}
		switch v := result.(type) {
		case int64:
			bi := b.(int64)
			if bi == 0 {
				return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorDivisionByZero)
			}
			result = v / bi
		case float64:
			bf := b.(float64)
			if bf == 0.0 {
				return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorDivisionByZero)
			}
			result = v / bf
		}
	}
	return result, nil
}

// remBody 整数同士の剰余。整数でない引数が含まれる場合はエラー
func remBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 3 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
	}

	a, aerr := EvalAsInt(lst.ElementAt(1), ns)
	b, berr := EvalAsInt(lst.ElementAt(2), ns)
	if aerr != nil {
		return nil, aerr
	} else if berr != nil {
		return nil, berr
	}
	if b == 0 {
		return nil, NewEvalError(lst.ElementAt(2).Position(), ErrorDivisionByZero)
	}
	return a % b, nil
}

func eqBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 2)
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

	fst := params[1]
	for i := 2; i < lst.Len(); i++ {
		b := params[i]
		if !isSameType(&fst, &b) {
			return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorTypeMissmatch, reflect.TypeOf(fst), reflect.TypeOf(b))
		}
		if fst != b {
			return int64(0), nil
		}
	}
	return int64(1), nil
}

func bitwiseANDbody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 2)
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
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[1])
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[i])
		}
		result = result & ip
	}
	return result, nil
}

func bitwiseORbody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 2)
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
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[1])
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[i])
		}
		result = result | ip
	}
	return result, nil
}

func bitwiseXORbody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 2)
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
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[1])
	}
	if lst.Len() == 2 { // 引数が一つのときはビットを反転させて返す。
		return ^result, nil
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[i])
		}
		result = result ^ ip
	}
	return result, nil
}

func lShiftBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 2)
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
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[1])
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[i])
		}
		result = result << ip
	}
	return result, nil
}

func rShiftBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 2)
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
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[1])
	}
	for i := 2; i < lst.Len(); i++ {
		ip, ok := params[i].(int64)
		if !ok {
			return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, params[i])
		}
		result = result >> ip
	}
	return result, nil
}

func ltBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// オペラントは2つしか許容しない
	if lst.Len() != 3 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
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
		if b, ok := pb.(int64); ok {
			return BoolToInt(a < b), nil
		}
		return nil, NewEvalError(lst.ElementAt(2).Position(), ErrorTypeMissmatch, reflect.TypeOf(a), reflect.TypeOf(pb))
	case float64:
		if b, ok := pb.(float64); ok {
			return BoolToInt(a < b), nil
		}
		return nil, NewEvalError(lst.ElementAt(2).Position(), ErrorTypeMissmatch, reflect.TypeOf(a), reflect.TypeOf(pb))
	default:
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorNonArithmeticDataType, pa)
	}
}

func lteBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// オペラントは2つしか許容しない
	if lst.Len() != 3 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
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
		if b, ok := pb.(int64); ok {
			return BoolToInt(a <= b), nil
		}
		return nil, NewEvalError(lst.ElementAt(2).Position(), ErrorTypeMissmatch, reflect.TypeOf(a), reflect.TypeOf(pb))
	case float64:
		if b, ok := pb.(float64); ok {
			return BoolToInt(a <= b), nil
		}
		return nil, NewEvalError(lst.ElementAt(2).Position(), ErrorTypeMissmatch, reflect.TypeOf(a), reflect.TypeOf(pb))
	default:
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorNonArithmeticDataType, pa)
	}
}

func gtBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// オペラントは2つしか許容しない
	if lst.Len() != 3 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
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
		if b, ok := pb.(int64); ok {
			return BoolToInt(a > b), nil
		}
		return nil, NewEvalError(lst.ElementAt(2).Position(), ErrorTypeMissmatch, reflect.TypeOf(a), reflect.TypeOf(pb))
	case float64:
		if b, ok := pb.(float64); ok {
			return BoolToInt(a > b), nil
		}
		return nil, NewEvalError(lst.ElementAt(2).Position(), ErrorTypeMissmatch, a, pb)
	default:
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorNonArithmeticDataType, pa)
	}
}

func gteBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// オペラントは2つしか許容しない
	if lst.Len() != 3 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
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
		if b, ok := pb.(int64); ok {
			return BoolToInt(a >= b), nil
		}
		return nil, NewEvalError(lst.ElementAt(2).Position(), ErrorTypeMissmatch, reflect.TypeOf(a), reflect.TypeOf(pb))
	case float64:
		if b, ok := pb.(float64); ok {
			return BoolToInt(a >= b), nil
		}
		return nil, NewEvalError(lst.ElementAt(2).Position(), ErrorTypeMissmatch, reflect.TypeOf(a), reflect.TypeOf(pb))
	default:
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorNonArithmeticDataType, pa)
	}
}

func notBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	// 要するに引数は必ず一つ
	if lst.Len() != 2 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	p, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if b, ok := p.(int64); ok {
		if b != 0 {
			return 0, nil
		}
		return 1, nil
	}
	return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, p)
}

func andBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 2)
	}
	// 引数を順に評価し、評価結果がfalseになったところで止めてfalseを返す。
	for i := 1; i < lst.Len(); i++ {
		ev, err := EvalElement(lst.ElementAt(i), ns)
		if err != nil {
			return nil, err
		}
		bv, ok := ev.(int64)
		if !ok {
			// 評価結果がint64に変換できない場合はエラーになる。
			return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeOfIntegerType, ev)
		}
		if bv == 0 {
			return 0, nil
		}
	}
	return 1, nil
}

func orBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 2)
	}
	// 引数を順に評価し、評価結果がtrueになったところで止めてtrueを返す。
	for i := 1; i < lst.Len(); i++ {
		ev, err := EvalElement(lst.ElementAt(i), ns)
		if err != nil {
			return nil, err
		}
		bv, ok := ev.(int64)
		if !ok {
			// 評価結果がint64に変換できない場合はエラーになる。
			return nil, NewEvalError(lst.ElementAt(i).Position(), ErrorOperantsMustBeOfIntegerType, ev)
		}
		if bv != 0 {
			return 1, nil
		}
	}
	return 0, nil
}

func strBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, lst.Len()-1, 1)
	}
	result := ""
	for i := 1; i < lst.Len(); i++ {
		ev, err := EvalElement(lst.ElementAt(i), ns)
		if err != nil {
			return nil, err
		}
		switch v := ev.(type) {
		case int64:
			result += fmt.Sprint(v)
		case float64:
			result += fmt.Sprint(v)
		case string:
			result += v
		default:
			return nil, NewEvalError(lst.Position(), ErrorInvalidOperation)
		}
	}
	return result, nil
}

func intBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	ev, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	switch v := ev.(type) {
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	case string:
		iv, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		return int64(iv), nil
	default:
		return nil, NewEvalError(lst.Position(), ErrorInvalidOperation)
	}
}

func floatBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	ev, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	switch v := ev.(type) {
	case int64:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		fv, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		return fv, nil
	default:
		return nil, NewEvalError(lst.Position(), ErrorInvalidOperation)
	}
}

// RegisterOperators stに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterOperators(ns *Namespace) {
	ns.RegisterExtension(addSymbol, nil, addBody)
	ns.RegisterExtension(subSymbol, nil, subBody)
	ns.RegisterExtension(mulSymbol, nil, mulBody)
	ns.RegisterExtension(divSymbol, nil, divBody)
	ns.RegisterExtension(remSymbol, nil, remBody)
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
	ns.RegisterExtension(strSymbol, nil, strBody)
	ns.RegisterExtension(intSymbol, nil, intBody)
	ns.RegisterExtension(floatSymbol, nil, floatBody)
}

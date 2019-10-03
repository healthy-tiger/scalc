package runtime

import (
	"fmt"

	"github.com/healthy-tiger/scalc/parser"
)

const (
	setSymbol   = "set"
	ifSymbol    = "if"
	whileSymbol = "while"
	printSymbol = "print"
)

// set組み込み関数に関するエラーコード
var (
	ErrorYouCannotBindAValueToAnythingOtherThanASymbol int
	ErrorYouCannotBindMoreThanOneValueToASymbol        int
	ErrorYouMustSpecifyTheValueToBind                  int
)

func init() {
	ErrorYouCannotBindAValueToAnythingOtherThanASymbol = RegisterEvalError("You cannot bind a value to anything other than a symbol.")
	ErrorYouCannotBindMoreThanOneValueToASymbol = RegisterEvalError("You cannot bind more than one value to a symbol.")
	ErrorYouMustSpecifyTheValueToBind = RegisterEvalError("You must specify the value to bind.")
}

func setBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	sid, ok := lst.SymbolAt(1)
	if !ok {
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorYouCannotBindAValueToAnythingOtherThanASymbol)
	}
	if lst.Len() < 3 {
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorYouMustSpecifyTheValueToBind)
	} else if lst.Len() > 3 {
		return nil, newEvalError(lst.ElementAt(3).Position(), ErrorYouCannotBindMoreThanOneValueToASymbol)
	}
	v, err := EvalElement(lst.ElementAt(2), ns)
	if err != nil {
		return nil, err
	}
	ns.Set(sid, v)
	return v, nil
}

func ifBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 4 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
	} else if lst.Len() > 4 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments)
	}
	p, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if cond, ok := p.(bool); ok {
		if cond {
			return EvalElement(lst.ElementAt(2), ns)
		}
		return EvalElement(lst.ElementAt(3), ns)
	}
	return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeBoolean)
}

func evalAsBool(elm parser.SyntaxElement, ns *Namespace) (bool, error) {
	r, err := EvalElement(elm, ns)
	if err != nil {
		return false, err
	}
	c, ok := r.(bool)
	if ok {
		return c, nil
	}
	return false, newEvalError(elm.Position(), ErrorOperantsMustBeBoolean)
}

func whileBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
	} else if lst.Len() > 3 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments)
	}

	condelm := lst.ElementAt(1)
	bodyelm := lst.ElementAt(2)
	cond, err := evalAsBool(condelm, ns)
	var result interface{} = nil
	for err == nil && cond {
		result, err = EvalElement(bodyelm, ns)
		if err == nil { // bodyelmを評価してエラーがなければ再度、ループの条件を確認する。
			cond, err = evalAsBool(condelm, ns)
		}
	}
	if err != nil { // エラーで抜けた場合はEvalError
		return nil, err
	}
	return result, nil
}

func printBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	params := make([]interface{}, lst.Len()-1)
	for i := 1; i < lst.Len(); i++ {
		v, err := EvalElement(lst.ElementAt(i), ns)
		if err != nil {
			return nil, err
		}
		params[i-1] = v
	}
	return fmt.Println(params...)
}

// RegisterStmt 文に関する拡張関数を登録する。
func RegisterStmt(st *parser.SymbolTable, ns *Namespace) {
	RegisterExtension(st, ns, setSymbol, nil, setBody)
	RegisterExtension(st, ns, ifSymbol, nil, ifBody)
	RegisterExtension(st, ns, printSymbol, nil, printBody)
	RegisterExtension(st, ns, whileSymbol, nil, whileBody)
}

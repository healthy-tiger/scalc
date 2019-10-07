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
	beginSymbol = "begin"
	funcSymbol  = "func"
)

// set組み込み関数に関するエラーコード
var (
	ErrorYouCannotBindAValueToAnythingOtherThanASymbol      int
	ErrorYouCannotBindMoreThanOneValueToASymbol             int
	ErrorYouMustSpecifyTheValueToBind                       int
	ErrorAFunctionDefinitionRequiresAnArgumentList          int
	ErrorAFunctionDefinitionRequiresAFunctionBodyDefinition int
	ErrorTheArgumentListMustConsistOfSymbolsOnly            int
)

func init() {
	ErrorYouCannotBindAValueToAnythingOtherThanASymbol = RegisterEvalError("You cannot bind a value to anything other than a symbol.")
	ErrorYouCannotBindMoreThanOneValueToASymbol = RegisterEvalError("You cannot bind more than one value to a symbol.")
	ErrorYouMustSpecifyTheValueToBind = RegisterEvalError("You must specify the value to bind.")
	ErrorAFunctionDefinitionRequiresAnArgumentList = RegisterEvalError("A function definition requires an argument list.")
	ErrorAFunctionDefinitionRequiresAFunctionBodyDefinition = RegisterEvalError("A function definition requires a function body definition.")
	ErrorTheArgumentListMustConsistOfSymbolsOnly = RegisterEvalError("The argument list must consist of symbols only.")
}

func setBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	sid, ok := lst.SymbolAt(1)
	if !ok {
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorYouCannotBindAValueToAnythingOtherThanASymbol, nil)
	}
	if lst.Len() < 3 {
		return nil, newEvalError(lst.ElementAt(1).Position(), ErrorYouMustSpecifyTheValueToBind, nil)
	} else if lst.Len() > 3 {
		return nil, newEvalError(lst.ElementAt(3).Position(), ErrorYouCannotBindMoreThanOneValueToASymbol, nil)
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
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
	} else if lst.Len() > 4 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments, nil)
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
	return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeBoolean, p)
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
	return false, newEvalError(elm.Position(), ErrorOperantsMustBeBoolean, r)
}

func whileBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
	} else if lst.Len() > 3 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments, nil)
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

// printBody fmt.Printlnを呼び出して結果をそのまま返す。
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

func beginBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
	}
	var result interface{} = nil
	var err error = nil
	for i := 1; i < lst.Len(); i++ {
		result, err = EvalElement(lst.ElementAt(i), ns)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func funcBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 3 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments, nil)
	} else if lst.Len() > 3 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments, nil)
	}

	e1 := lst.ElementAt(1)
	if !e1.IsList() {
		return nil, newEvalError(e1.Position(), ErrorAFunctionDefinitionRequiresAnArgumentList, nil)
	}
	body := lst.ElementAt(2)
	if !body.IsList() {
		return nil, newEvalError(body.Position(), ErrorAFunctionDefinitionRequiresAFunctionBodyDefinition, nil)
	}
	// e1の中身が全てシンボルであることをチェックする。
	argdefs := e1.(*parser.List)
	args := make([]parser.SymbolID, argdefs.Len())
	for i := 0; i < argdefs.Len(); i++ {
		s, ok := argdefs.SymbolAt(i)
		if !ok {
			return nil, newEvalError(argdefs.ElementAt(i).Position(), ErrorTheArgumentListMustConsistOfSymbolsOnly, nil)
		}
		args[i] = s
	}
	return &Function{args, body.(*parser.List)}, nil
}

// RegisterStmt 文に関する拡張関数を登録する。
func RegisterStmt(ns *Namespace) {
	ns.RegisterExtension(setSymbol, nil, setBody)
	ns.RegisterExtension(ifSymbol, nil, ifBody)
	ns.RegisterExtension(printSymbol, nil, printBody)
	ns.RegisterExtension(whileSymbol, nil, whileBody)
	ns.RegisterExtension(beginSymbol, nil, beginBody)
	ns.RegisterExtension(funcSymbol, nil, funcBody)
}

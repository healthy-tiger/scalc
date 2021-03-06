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
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorYouCannotBindAValueToAnythingOtherThanASymbol)
	}
	if lst.Len() < 3 {
		return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorYouMustSpecifyTheValueToBind)
	} else if lst.Len() > 3 {
		return nil, NewEvalError(lst.ElementAt(3).Position(), ErrorYouCannotBindMoreThanOneValueToASymbol)
	}
	v, err := EvalElement(lst.ElementAt(2), ns)
	if err != nil {
		return nil, err
	}
	ns.Set(sid, v)
	return v, nil
}

func ifBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 4 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 4-1)
	}
	p, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if cond, ok := p.(int64); ok {
		if cond != 0 {
			return EvalElement(lst.ElementAt(2), ns)
		}
		return EvalElement(lst.ElementAt(3), ns)
	}
	return nil, NewEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeOfIntegerType, p)
}

func whileBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 3 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 3-1)
	}
	condelm := lst.ElementAt(1)
	bodyelm := lst.ElementAt(2)
	cond, err := EvalAsInt(condelm, ns)
	count := int64(0)
	for err == nil && cond != 0 {
		count++
		_, err = EvalElement(bodyelm, ns)
		if err == nil { // bodyelmを評価してエラーがなければ再度、ループの条件を確認する。
			cond, err = EvalAsInt(condelm, ns)
		}
	}
	if err != nil { // エラーで抜けた場合はEvalError
		return nil, err
	}
	return count, nil // bodyelmを評価した回数を返す。
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
	np, err := fmt.Println(params...)
	return int64(np), err
}

func beginBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 2 {
		return nil, NewEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
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
	if lst.Len() != 3 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 3-1)
	}
	e1 := lst.ElementAt(1)
	if !e1.IsList() {
		return nil, NewEvalError(e1.Position(), ErrorAFunctionDefinitionRequiresAnArgumentList)
	}
	body := lst.ElementAt(2)
	if !body.IsList() {
		return nil, NewEvalError(body.Position(), ErrorAFunctionDefinitionRequiresAFunctionBodyDefinition)
	}
	// e1の中身が全てシンボルであることをチェックする。
	argdefs := e1.(*parser.List)
	args := make([]parser.SymbolID, argdefs.Len())
	for i := 0; i < argdefs.Len(); i++ {
		s, ok := argdefs.SymbolAt(i)
		if !ok {
			return nil, NewEvalError(argdefs.ElementAt(i).Position(), ErrorTheArgumentListMustConsistOfSymbolsOnly)
		}
		args[i] = s
	}
	return &Function{args, body.(*parser.List), nil, nil}, nil
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

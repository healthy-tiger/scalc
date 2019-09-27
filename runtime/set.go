package runtime

import "github.com/healthy-tiger/scalc/parser"

const (
	setSymbol = "set"
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

// RegisterSetFunc set組み込み関数を登録する。
func RegisterSetFunc(st *parser.SymbolTable, ns *Namespace) {
	RegisterExtension(st, ns, setSymbol, nil, setBody)
}

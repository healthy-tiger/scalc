package runtime

import (
	"time"

	"github.com/healthy-tiger/scalc/parser"
)

// Callable 呼び出し可能なオブジェクトの呼び出し用
type Callable interface {
	Eval(lst *parser.List, ns *Namespace) (interface{}, error)
}

// Function func関数で定義されたユーザー定義関数を表す。
type Function struct {
	params []parser.SymbolID
	body   *parser.List
}

// Eval 関数fを引数agrsと、グローバルの名前空間globalsで評価し、その結果を返す。
func (f *Function) Eval(lst *parser.List, ns *Namespace) (interface{}, error) {
	if len(f.params) != lst.Len()-1 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, len(f.params), lst.Len()-1)
	}
	// 呼び出し先（の関数を実行する際）の名前空間を定義。最上位の名前空間以外は呼び出し元と名前空間を共有しない。
	lns := NewNamespace(ns.Root())
	// 引数を呼び出し元の名前空間で評価して、その結果を呼び出し先の名前空間にセット
	for i := 1; i < lst.Len(); i++ {
		a, err := EvalElement(lst.ElementAt(i), ns)
		if err != nil {
			return nil, err
		}
		lns.Set(f.params[i-1], a)
	}
	return EvalList(f.body, lns)
}

// EvalAsBool 名前空間nsでelmを評価し、その結果をboolとして返す。boolでない結果の場合はエラーを返す。
func EvalAsBool(elm parser.SyntaxElement, ns *Namespace) (bool, error) {
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

// EvalAsInt 名前空間nsでelmを評価し、その結果をint64として返す。int64でない結果の場合はエラーを返す。
func EvalAsInt(elm parser.SyntaxElement, ns *Namespace) (int64, error) {
	r, err := EvalElement(elm, ns)
	if err != nil {
		return -1, err
	}
	c, ok := r.(int64)
	if ok {
		return c, nil
	}
	return -1, newEvalError(elm.Position(), ErrorOperantsMustBeOfIntegerType, r)
}

// Extension scalcの拡張関数の構造体
type Extension struct {
	object interface{}
	body   func(obj interface{}, lst *parser.List, ns *Namespace) (interface{}, error) // lstは関数のシンボルを最初の要素に含んだ状態で渡される。
}

// Eval 拡張関数を呼び出す。
func (ex *Extension) Eval(lst *parser.List, ns *Namespace) (interface{}, error) {
	return ex.body(ex.object, lst, ns)
}

// EvalElement 構文要素を指定された名前空間で評価する。
func EvalElement(st parser.SyntaxElement, ns *Namespace) (interface{}, error) {
	if st.IsList() {
		return EvalList(st.(*parser.List), ns)
	}
	if sid, ok := st.SymbolValue(); ok {
		sv, ok := ns.Get(sid)
		if !ok {
			sn, err := ns.GetSymbolName(sid)
			if err != nil {
				panic(err)
			}
			return nil, newEvalError(st.Position(), ErrorUndefinedSymbol, sn)
		}
		switch ev := sv.(type) {
		case int64, float64, string, bool, Callable, time.Time:
			return ev, nil
		default:
			panic("Unexpected evaluation result type")
		}
	} else if ss, ok := st.StringValue(); ok {
		return ss, nil
	} else if si, ok := st.IntValue(); ok {
		return si, nil
	} else if sf, ok := st.FloatValue(); ok {
		return sf, nil
	} else {
		panic("Illegal syntax tree element")
	}
}

// EvalList リストlstを名前空間のもとで評価する。
func EvalList(lst *parser.List, ns *Namespace) (interface{}, error) {
	// 空のリストは評価できないのでエラー(Excentionがリストを評価する場合はExtentionsによる）
	if lst.Len() == 0 {
		return nil, newEvalError(lst.Position(), ErrorAnEmptyListIsNotAllowed)
	}
	// 最初の要素は必ずシンボルで、呼び出し可能なオブジェクト（*FunctionかExtensionにバインドされていなければならない）
	first := lst.ElementAt(0)
	callable, err := EvalElement(first, ns)
	if err != nil {
		return nil, err
	}
	if c, ok := callable.(Callable); ok {
		return c.Eval(lst, ns)
	}
	return nil, newEvalError(first.Position(), ErrorTheFirstElementOfTheListToBeEvaluatedMustBeACallableObject, callable)
}

// MakeDefaultNamespace 予約済みのシンボルをシンボルテーブに登録し、その値を登録済みの名前空間を作る。
func MakeDefaultNamespace(ns *Namespace) {
	RegisterBoolType(ns)
	RegisterOperators(ns)
	RegisterMath(ns)
	RegisterStmt(ns)
	RegisterTImeFunc(ns)
}

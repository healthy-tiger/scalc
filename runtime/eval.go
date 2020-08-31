package runtime

import (
	"fmt"
	"reflect"

	"github.com/healthy-tiger/scalc/parser"
)

// Function func関数で定義されたユーザー定義関数を表す。
type Function struct {
	params      []parser.SymbolID                                                           // ユーザー定義関数の引数リスト
	body        *parser.List                                                                // ユーザー定義関数の本体
	nativeparam interface{}                                                                 // ネイティブ関数の内部パラメータ
	native      func(obj interface{}, lst *parser.List, ns *Namespace) (interface{}, error) // ネイティブ関数の本体
}

// Eval 関数fをlstの第2要素以降を引数に、グローバルの名前空間globalsで評価し、その結果を返す。
func (f *Function) Eval(lst *parser.List, ns *Namespace) (interface{}, error) {
	if f.body != nil && f.params != nil {
		return f.EvalAsFunction(lst, ns)
	} else if f.native != nil {
		return f.EvalAsNative(lst, ns)
	}
	panic("Nil function.")
}

// EvalAsFunction 関数fをユーザー定義関数として、lstの第2要素以降を引数に、グローバルの名前空間globalsで評価し、その結果を返す。
func (f *Function) EvalAsFunction(lst *parser.List, ns *Namespace) (interface{}, error) {
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

// EvalAsNative 関数fをネイティブ関数として、lstの第2要素以降を引数に、グローバルの名前空間globalsで評価し、その結果を返す。
func (f *Function) EvalAsNative(lst *parser.List, ns *Namespace) (interface{}, error) {
	return f.native(f.nativeparam, lst, ns)
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
		return sv, nil
	} else if ss, ok := st.StringValue(); ok {
		return ss, nil
	} else if si, ok := st.IntValue(); ok {
		return si, nil
	} else if sf, ok := st.FloatValue(); ok {
		return sf, nil
	} else {
		panic(fmt.Sprintf("Illegal syntax tree element %v", reflect.TypeOf(st)))
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
	funcobj, err := EvalElement(first, ns)
	if err != nil {
		return nil, err
	}
	if c, ok := funcobj.(*Function); ok {
		return c.Eval(lst, ns)
	}
	return nil, newEvalError(first.Position(), ErrorTheFirstElementOfTheListToBeEvaluatedMustBeACallableObject, funcobj)
}

// MakeDefaultNamespace 予約済みのシンボルをシンボルテーブに登録し、その値を登録済みの名前空間を作る。
func MakeDefaultNamespace(ns *Namespace) {
	RegisterBoolType(ns)
	RegisterOperators(ns)
	RegisterMath(ns)
	RegisterStmt(ns)
	RegisterTimeFunc(ns)
}

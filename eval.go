package scalc

import (
	"github.com/healthy-tiger/gostree"
)

// TODO エラーは行と列を返せるようにする。

// Namespace シンボルと値のマップ
type Namespace struct {
	root     *Namespace
	parent   *Namespace
	bindings map[gostree.SymbolID]interface{} // string, int64, float64, bool, *Function, Extensionのいれずれか
}

// Get nsからシンボルID idに対応する値を取得する。
func (ns *Namespace) Get(id gostree.SymbolID) (interface{}, bool) {
	n := ns
	for n != nil {
		v, ok := n.bindings[id]
		if ok {
			return v, true
		}
		n = n.parent
	}
	return nil, false
}

// Set nsにシンボルID idに対応する値を格納する。
func (ns *Namespace) Set(id gostree.SymbolID, value interface{}) {
	ns.bindings[id] = value
}

// Parent nsの親の名前空間を返す。
func (ns *Namespace) Parent() *Namespace {
	return ns.parent
}

// Root nsの最上位の名前空間を返す。
func (ns *Namespace) Root() *Namespace {
	return ns.root
}

// NewNamespace 新しい名前空間を生成する。
func NewNamespace(parent *Namespace) *Namespace {
	// 最上位の名前空間を探しておく
	var p *Namespace = nil
	if parent != nil {
		p := parent
		for p.parent != nil {
			p = p.parent
		}
	}
	return &Namespace{p, parent, make(map[gostree.SymbolID]interface{})}
}

// Callable 呼び出し可能なオブジェクトの呼び出し用
type Callable interface {
	Eval(lst *gostree.List, ns *Namespace) (interface{}, error)
}

// Function func関数で定義されたユーザー定義関数を表す。
type Function struct {
	params []gostree.SymbolID
	body   *gostree.List
}

// Eval 関数fを引数agrsと、グローバルの名前空間globalsで評価し、その結果を返す。
func (f *Function) Eval(lst *gostree.List, ns *Namespace) (interface{}, error) {
	if len(f.params) != lst.Len()-1 {
		return nil, errorTheNumberOfArgumentsDoesNotMatch
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

// Extension scalcの拡張関数の構造体
type Extension struct {
	object interface{}
	body   func(obj interface{}, lst *gostree.List, ns *Namespace) (interface{}, error) // lstは関数のシンボルを最初の要素に含んだ状態で渡される。
}

// Eval 拡張関数を呼び出す。
func (ex *Extension) Eval(lst *gostree.List, ns *Namespace) (interface{}, error) {
	return ex.body(ex.object, lst, ns)
}

// EvalElement 構文要素を指定された名前空間で評価する。
func EvalElement(st gostree.SyntaxElement, ns *Namespace) (interface{}, error) {
	if st.IsList() {
		return EvalList(st.(*gostree.List), ns)
	}
	if sid, ok := st.SymbolValue(); ok {
		sv, ok := ns.Get(sid)
		if !ok {
			return nil, errorUndefinedSymbol
		}
		// 関数の引数に渡せるのは値のみ。シンボルや関数は渡せない。
		switch sarg := sv.(type) {
		case int64, float64, string, bool, Callable:
			return sarg, nil
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
func EvalList(lst *gostree.List, ns *Namespace) (interface{}, error) {
	// 空のリストは評価できないのでエラー(Excentionがリストを評価する場合はExtentionsによる）
	if lst.Len() == 0 {
		return nil, errorAnEmptyListIsNotAllowed
	}
	// 最初の要素は必ずシンボルで、呼び出し可能なオブジェクト（*FunctionかExtensionにバインドされていなければならない）
	callable, err := EvalElement(lst.ElementAt(0), ns)
	if err != nil {
		return nil, err
	}
	if c, ok := callable.(Callable); ok {
		return c.Eval(lst, ns)
	}
	return nil, errorTheFirstElementOfTheListToBeEvaluatedMustBeACallableObject
}

// RegisterExtension 拡張関数を登録する。
func RegisterExtension(stree *gostree.STree, ns *Namespace, symbolName string, extobj interface{}, extbody func(interface{}, *gostree.List, *Namespace) (interface{}, error)) {
	sid := stree.GetSymbolID(symbolName)
	ns.Set(sid, &Extension{extobj, extbody})
}

// DefaultNamespace 予約済みのシンボルをシンボルテーブに登録し、その値を登録済みの名前空間を作る。
func DefaultNamespace(stree *gostree.STree) *Namespace {
	ns := NewNamespace(nil)
	RegisterBoolType(stree, ns)
	RegisterOperators(stree, ns)
	return ns
}

// EvalSTree gostreeを評価する関数。与えられたデフォルトの名前空間のもとで、トップレベルのリストを順に評価する。
func EvalSTree(stree *gostree.STree, ns *Namespace, resultHandler func(interface{}), errorHandler func(error)) {
	for _, t := range stree.Lists {
		result, err := EvalList(t, ns)
		if err != nil {
			errorHandler(err)
		} else {
			resultHandler(result)
		}
	}
}

package scalc

import (
	"github.com/healthy-tiger/gostree"
)

const (
	true_symbol  = "true"
	false_symbol = "false"
)

// Namespace シンボルと値のマップ
type Namespace struct {
	bindings map[gostree.SymbolID]interface{} // string, int64, float64, bool, Function, Extensionのいれずれか
}

// Get nsからシンボルID idに対応する値を取得する。
func (ns *Namespace) Get(id gostree.SymbolID) (interface{}, bool) {
	v, ok := ns.bindings[id]
	return v, ok
}

// Set nsにシンボルID idに対応する値を格納する。
func (ns *Namespace) Set(id gostree.SymbolID, value interface{}) {
	ns.bindings[id] = value
}

// NewNamespace 新しい名前空間を生成する。
func NewNamespace() *Namespace {
	return &Namespace{make(map[gostree.SymbolID]interface{})}
}

// Function func関数で定義されたユーザー定義関数を表す。
type Function struct {
	params []gostree.SymbolID
	body   *gostree.List
}

// Call 関数fを引数agrsと、グローバルの名前空間globalsで評価し、その結果を返す。
func (f *Function) Call(params []interface{}, globals *Namespace) (interface{}, error) {
	if len(f.params) != len(params) {
		return nil, errorTheNumberOfArgumentsDoesNotMatch
	}
	lns := NewNamespace()
	// 引数の値を名前空間にセット
	for i := 0; i < len(params); i++ {
		lns.Set(f.params[i], params[i])
	}
	return EvalList(f.body, lns, globals)
}

// Extension scalcの拡張関数のインターフェース
type Extension interface {
	Eval(lst *gostree.List, locals *Namespace, globals *Namespace) (interface{}, error) // lstは関数のシンボルを最初の要素に含んだ状態で渡される。
}

func getSymbolValue(id gostree.SymbolID, ns *Namespace, globals *Namespace) interface{} {
	if ns != nil {
		v, ok := ns.Get(id)
		if ok {
			return v
		}
	}
	if globals != nil {
		v, ok := globals.Get(id)
		if ok {
			return v
		}
	}
	return nil
}

// EvalList リストlstを名前空間のもとで評価する。
func EvalList(lst *gostree.List, ns *Namespace, globals *Namespace) (interface{}, error) {
	// 空のリストは評価できないのでエラー(Excentionがリストを評価する場合はExtentionsによる）
	if lst.Len() == 0 {
		return nil, errorAnEmptyListIsNotAllowed
	}
	// 最初の要素は必ずシンボルで、呼び出し可能なオブジェクト（FunctionかExtensionにバインドされていなければならない）
	callableid, ok := lst.SymbolAt(0)
	if !ok {
		return nil, errorTheFirstElementOfTheListToBeEvaluatedMustBeASymbol
	}
	callable := getSymbolValue(callableid, ns, globals) // 名前空間から値を取得
	if callable == nil {
		return nil, errorUndefinedSymbol
	}
	// callableが拡張関数の場合は引数を事前に評価せずに渡す。
	if ex, ok := callable.(Extension); ok {
		return ex.Eval(lst, ns, globals)
	}
	fn, ok := callable.(Function)
	if !ok {
		return nil, errorTheFirstElementOfTheListToBeEvaluatedMustBeACallableObject
	}

	// callableがユーザー定義関数の場合は、2番目以降のリストの要素を評価し、その結果を引数にしてcallable(=fn)を呼び出して、その結果を返す。
	args := make([]interface{}, lst.Len()-1)
	for i := 1; i < lst.Len(); i++ {
		t := lst.ElementAt(i)
		if t.IsList() {
			sl, err := EvalList(t.(*gostree.List), ns, globals)
			if err != nil {
				args[i-1] = sl
			} else {
				return nil, err
			}
		} else {
			if sid, ok := t.SymbolValue(); ok {
				sv := getSymbolValue(sid, ns, globals)
				if sv == nil {
					return nil, errorUndefinedSymbol
				}
				// 関数の引数に渡せるのは値のみ。シンボルや関数は渡せない。
				switch sarg := sv.(type) {
				case int64, float64, string, bool:
					args[i-1] = sarg
				default:
					return nil, errorFunctionCannotBePassedAsFunctionArgument
				}
			} else if ss, ok := t.StringValue(); ok {
				args[i-1] = ss
			} else if si, ok := t.IntValue(); ok {
				args[i-1] = si
			} else if sf, ok := t.FloatValue(); ok {
				args[i-1] = sf
			} else {
				panic("Illegal syntax tree element")
			}
		}
	}
	return fn.Call(args, globals)
}

// defaultNamespace 予約済みのシンボルをシンボルテーブに登録し、その値を登録済みの名前空間を作る。
func defaultNamespace(stree *gostree.STree) *Namespace {
	ns := NewNamespace()
	// 今の所、予約されているのはtrueとfalseだけ
	trueid := stree.GetSymbolID(true_symbol)
	falseid := stree.GetSymbolID(false_symbol)
	ns.Set(trueid, true)
	ns.Set(falseid, false)
	return ns
}

// EvalSTree gostreeを評価する関数。globalsを初期化し、トップレベルのリストを順に評価する。
func EvalSTree(stree *gostree.STree, resultHandler func(interface{}), errorHandler func(error)) {
	globals := defaultNamespace(stree)
	for _, t := range stree.Lists {
		result, err := EvalList(t, globals, globals)
		if err != nil {
			errorHandler(err)
		} else {
			resultHandler(result)
		}
	}
}

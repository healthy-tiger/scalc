package runtime

import "github.com/healthy-tiger/scalc/parser"

// Namespace シンボルと値のマップ
type Namespace struct {
	symtbl   *parser.SymbolTable // ルートの名前空間の場合のみ非nilになる。
	root     *Namespace
	parent   *Namespace
	bindings map[parser.SymbolID]interface{} // string, int64, float64, bool, time.Time, *Function, *Extensionのいれずれか
}

// Get nsからシンボルID idに対応する値を取得する。
func (ns *Namespace) Get(id parser.SymbolID) (interface{}, bool) {
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
func (ns *Namespace) Set(id parser.SymbolID, value interface{}) {
	ns.bindings[id] = value
}

// Parent nsの親の名前空間を返す。
func (ns *Namespace) Parent() *Namespace {
	return ns.parent
}

// GetSymbolID シンボルを新たに登録する。
func (ns *Namespace) GetSymbolID(name string) parser.SymbolID {
	return ns.Root().symtbl.GetSymbolID(name)
}

// GetSymbolName シンボルIDに対応する名前を返す。
func (ns *Namespace) GetSymbolName(id parser.SymbolID) (string, error) {
	return ns.Root().symtbl.GetSymbolName(id)
}

// Root nsの最上位の名前空間を返す。
func (ns *Namespace) Root() *Namespace {
	if ns.parent == nil { // 親の名前空間がない＝自分自身が最上位
		return ns
	}
	return ns.root
}

// RegisterExtension 拡張関数を登録する。必ず名前空間のルートに対して登録を行う。
func (ns *Namespace) RegisterExtension(symbolName string, extobj interface{}, extbody func(interface{}, *parser.List, *Namespace) (interface{}, error)) parser.SymbolID {
	root := ns.Root()
	sid := root.symtbl.GetSymbolID(symbolName)
	root.Set(sid, &Function{nil, nil, extobj, extbody})
	return sid
}

// NewNamespace 新しい名前空間を生成する。
func NewNamespace(parent *Namespace) *Namespace {
	// 最上位の名前空間を探しておく
	var p *Namespace = nil
	if parent != nil {
		p = parent
		for p.parent != nil {
			p = p.parent
		}
	}
	return &Namespace{nil, p, parent, make(map[parser.SymbolID]interface{})}
}

// NewRootNamespace 新しく最上位の名前空間を作る
func NewRootNamespace(st *parser.SymbolTable) *Namespace {
	r := NewNamespace(nil)
	r.symtbl = st
	return r
}

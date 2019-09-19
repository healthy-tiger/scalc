package scalc

import (
	"github.com/healthy-tiger/gostree"
)

const (
	trueSymbol  = "true"
	falseSymbol = "false"
)

// RegisterBoolTyoe streeにbool型のシンボルを、nsにシンボルに対応する値を登録する。
func RegisterBoolTyoe(stree *gostree.STree, ns *Namespace) {
	// 今の所、予約されているのはtrueとfalseだけ
	trueid := stree.GetSymbolID(trueSymbol)
	falseid := stree.GetSymbolID(falseSymbol)
	ns.Set(trueid, true)
	ns.Set(falseid, false)
}

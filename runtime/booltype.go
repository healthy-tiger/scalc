package runtime

import (
	"github.com/healthy-tiger/scalc/parser"
)

const (
	trueSymbol  = "true"
	falseSymbol = "false"
)

// RegisterBoolType streeにbool型のシンボルを、nsにシンボルに対応する値を登録する。
func RegisterBoolType(st *parser.SymbolTable, ns *Namespace) {
	// 今の所、予約されているのはtrueとfalseだけ
	trueid := st.GetSymbolID(trueSymbol)
	falseid := st.GetSymbolID(falseSymbol)
	ns.Set(trueid, true)
	ns.Set(falseid, false)
}

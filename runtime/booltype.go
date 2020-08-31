package runtime

import "github.com/healthy-tiger/scalc/parser"

const (
	trueSymbol  = "true"
	falseSymbol = "false"
)

// RegisterBoolType streeにbool型のシンボルを、nsにシンボルに対応する値を登録する。
func RegisterBoolType(ns *Namespace) {
	trueid := ns.GetSymbolID(trueSymbol)
	falseid := ns.GetSymbolID(falseSymbol)
	ns.Set(trueid, parser.SInt(1))
	ns.Set(falseid, parser.SInt(0))
}

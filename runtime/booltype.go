package runtime

const (
	trueSymbol  = "true"
	falseSymbol = "false"
)

// RegisterBoolType streeにbool型のシンボルを、nsにシンボルに対応する値を登録する。
func RegisterBoolType(ns *Namespace) {
	trueid := ns.GetSymbolID(trueSymbol)
	falseid := ns.GetSymbolID(falseSymbol)
	ns.Set(trueid, int64(1))
	ns.Set(falseid, int64(0))
}

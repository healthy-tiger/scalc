package runtime

const (
	trueSymbol  = "true"
	falseSymbol = "false"
)

// BoolToInt Goのbool型をscalcの内部表現の整数型に変換する。
func BoolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

// RegisterBoolType streeにbool型のシンボルを、nsにシンボルに対応する値を登録する。
func RegisterBoolType(ns *Namespace) {
	trueid := ns.GetSymbolID(trueSymbol)
	falseid := ns.GetSymbolID(falseSymbol)
	ns.Set(trueid, BoolToInt(true))
	ns.Set(falseid, BoolToInt(false))
}

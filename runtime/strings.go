package runtime

import (
	"github.com/healthy-tiger/scalc/parser"
)

const (
	strCmpSymbol        = "strcmp"
	strCmpNaturalSymbol = "strcmp-natural"
)

func toNaturalString(s string) []rune {
	as := []rune(s)
	bs := make([]rune, len(as))
	bi := 0
	i := 0
	n := len(as)
	for i < n {
		c := as[i]
		if c >= '0' && c <= '9' {
			bs[bi] = c - '0'
			j := i + 1
			for j < n {
				d := as[j]
				if d >= '0' && d <= '9' {
					bs[bi] = bs[bi]*10 + (d - '0')
					j++
				} else {
					break
				}
			}
			bi++
			i = j
		} else {
			bs[bi] = c
			bi++
			i++
		}
	}
	return bs[:bi]
}

func compareRunes(a []rune, b []rune) int {
	al, bl := len(a), len(b)
	r := 0
	i := 0
	for ; i < al && i < bl; i++ {
		r = int(a[i] - b[i])
		if r != 0 {
			break
		}
	}
	if r == 0 {
		r = al - bl
	}
	if r < 0 {
		return -1
	} else if r > 0 {
		return 1
	}
	return r
}

func strCmpBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 3 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
	}
	a, aerr := EvalAsString(lst.ElementAt(1), ns)
	b, berr := EvalAsString(lst.ElementAt(2), ns)
	if aerr != nil {
		return nil, aerr
	}
	if berr != nil {
		return nil, berr
	}
	return int64(compareRunes([]rune(a), []rune(b))), nil
}

func strCmpNaturalBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 3 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
	}
	a, aerr := EvalAsString(lst.ElementAt(1), ns)
	b, berr := EvalAsString(lst.ElementAt(2), ns)
	if aerr != nil {
		return nil, aerr
	}
	if berr != nil {
		return nil, berr
	}
	an := toNaturalString(a)
	bn := toNaturalString(b)
	return int64(compareRunes(an, bn)), nil
}

// RegisterStrings stに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterStrings(ns *Namespace) {
	ns.RegisterExtension(strCmpSymbol, nil, strCmpBody)
	ns.RegisterExtension(strCmpNaturalSymbol, nil, strCmpNaturalBody)
}

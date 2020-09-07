package runtime

import (
	"strings"

	"github.com/healthy-tiger/scalc/parser"
)

const (
	strCmpSymbol        = "str-cmp"
	strCmpNaturalSymbol = "str-cmp-natural"
	containsSymbol      = "str-contains"
	containsAnySymbol   = "str-contains-any"
	countSymbol         = "str-count"
	equalFoldSymbol     = "str-equal-fold"
	hasPrefixSymbol     = "str-has-prefix"
	hasSuffixSymbol     = "str-has-suffix"
	indexSymbol         = "str-index"
	indexAnySymbol      = "str-index-any"
	lastIndexSymbol     = "str-lastindex"
	lastIndexAnySymbol  = "str-lastindex-any"
	repeatSymbol        = "str-repeat"
	replaceSymbol       = "str-replace"
	titleSymbol         = "str-title"
	toLowerSymbol       = "str-to-lower"
	toTitleSymbol       = "str-to-title"
	toUpperSymbol       = "str-to-upper"
	trimSymbol          = "str-trim"
	trimLeftSymbol      = "str-trim-left"
	trimPrefixSymbol    = "str-trim-prefix"
	trimRightSymbol     = "str-trim-right"
	trimSpaceSymbol     = "str-trim-space"
	trimSuffixSymbol    = "str-trim-suffix"
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

func containsBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return BoolToInt(strings.Contains(a, b)), nil
}

func containsAnyBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return BoolToInt(strings.ContainsAny(a, b)), nil
}

func countBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return int64(strings.Count(a, b)), nil
}

func equalFoldBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return BoolToInt(strings.EqualFold(a, b)), nil
}

func hasPrefixBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return BoolToInt(strings.HasPrefix(a, b)), nil
}

func hasSuffixBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return BoolToInt(strings.HasSuffix(a, b)), nil
}

func indexBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	i := strings.Index(a, b)
	if i >= 0 {
		ss := a[:i]
		return int64(len(ss)), nil
	}
	return int64(-1), nil
}

func indexAnyBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	i := strings.IndexAny(a, b)
	if i >= 0 {
		ss := a[:i]
		return int64(len(ss)), nil
	}
	return int64(-1), nil
}

func lastIndexBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	i := strings.LastIndex(a, b)
	if i >= 0 {
		ss := a[:i]
		return int64(len(ss)), nil
	}
	return int64(-1), nil
}

func lastIndexAnyBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	i := strings.LastIndexAny(a, b)
	if i >= 0 {
		ss := a[:i]
		return int64(len(ss)), nil
	}
	return int64(-1), nil
}

func repeatBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 3 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
	}
	a, aerr := EvalAsString(lst.ElementAt(1), ns)
	b, berr := EvalAsInt(lst.ElementAt(2), ns)
	if aerr != nil {
		return nil, aerr
	}
	if berr != nil {
		return nil, berr
	}
	return strings.Repeat(a, int(b)), nil
}

func replaceBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() == 4 {
		a, aerr := EvalAsString(lst.ElementAt(1), ns)
		b, berr := EvalAsString(lst.ElementAt(2), ns)
		c, cerr := EvalAsString(lst.ElementAt(3), ns)
		if aerr != nil {
			return nil, aerr
		}
		if berr != nil {
			return nil, berr
		}
		if cerr != nil {
			return nil, cerr
		}
		return strings.ReplaceAll(a, b, c), nil
	} else if lst.Len() == 5 {
		a, aerr := EvalAsString(lst.ElementAt(1), ns)
		b, berr := EvalAsString(lst.ElementAt(2), ns)
		c, cerr := EvalAsString(lst.ElementAt(3), ns)
		d, derr := EvalAsInt(lst.ElementAt(4), ns)
		if aerr != nil {
			return nil, aerr
		}
		if berr != nil {
			return nil, berr
		}
		if cerr != nil {
			return nil, cerr
		}
		if derr != nil {
			return nil, derr
		}
		return strings.Replace(a, b, c, int(d)), nil
	} else {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
	}
}

func titleBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsString(lst.ElementAt(1), ns)
	if aerr != nil {
		return nil, aerr
	}
	return strings.Title(a), nil
}

func toLowerBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsString(lst.ElementAt(1), ns)
	if aerr != nil {
		return nil, aerr
	}
	return strings.ToLower(a), nil
}

func toTitleBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsString(lst.ElementAt(1), ns)
	if aerr != nil {
		return nil, aerr
	}
	return strings.ToTitle(a), nil
}

func toUpperBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsString(lst.ElementAt(1), ns)
	if aerr != nil {
		return nil, aerr
	}
	return strings.ToUpper(a), nil
}

func trimBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return strings.Trim(a, b), nil
}

func trimLeftBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return strings.TrimLeft(a, b), nil
}

func trimPrefixBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return strings.TrimPrefix(a, b), nil
}

func trimRightBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return strings.TrimRight(a, b), nil
}

func trimSpaceBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, NewEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsString(lst.ElementAt(1), ns)
	if aerr != nil {
		return nil, aerr
	}
	return strings.TrimSpace(a), nil
}

func trimSuffixBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return strings.TrimSuffix(a, b), nil
}

// RegisterStrings stに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterStrings(ns *Namespace) {
	ns.RegisterExtension(strCmpSymbol, nil, strCmpBody)
	ns.RegisterExtension(strCmpNaturalSymbol, nil, strCmpNaturalBody)
	ns.RegisterExtension(containsSymbol, nil, containsBody)
	ns.RegisterExtension(containsAnySymbol, nil, containsAnyBody)
	ns.RegisterExtension(countSymbol, nil, countBody)
	ns.RegisterExtension(equalFoldSymbol, nil, equalFoldBody)
	ns.RegisterExtension(hasPrefixSymbol, nil, hasPrefixBody)
	ns.RegisterExtension(hasSuffixSymbol, nil, hasSuffixBody)
	ns.RegisterExtension(indexSymbol, nil, indexBody)
	ns.RegisterExtension(indexAnySymbol, nil, indexAnyBody)
	ns.RegisterExtension(lastIndexSymbol, nil, lastIndexBody)
	ns.RegisterExtension(lastIndexAnySymbol, nil, lastIndexAnyBody)
	ns.RegisterExtension(repeatSymbol, nil, repeatBody)
	ns.RegisterExtension(replaceSymbol, nil, replaceBody)
	ns.RegisterExtension(titleSymbol, nil, titleBody)
	ns.RegisterExtension(toLowerSymbol, nil, toLowerBody)
	ns.RegisterExtension(toTitleSymbol, nil, toTitleBody)
	ns.RegisterExtension(toUpperSymbol, nil, toUpperBody)
	ns.RegisterExtension(trimSymbol, nil, trimBody)
	ns.RegisterExtension(trimLeftSymbol, nil, trimLeftBody)
	ns.RegisterExtension(trimPrefixSymbol, nil, trimPrefixBody)
	ns.RegisterExtension(trimRightSymbol, nil, trimRightBody)
	ns.RegisterExtension(trimSpaceSymbol, nil, trimSpaceBody)
	ns.RegisterExtension(trimSuffixSymbol, nil, trimSuffixBody)
}

package parser

// SyntaxElement 構文要素を表す。
type SyntaxElement interface {
	Position() Position
	IsList() bool
	IntValue() (int64, bool)
	FloatValue() (float64, bool)
	StringValue() (string, bool)
	SymbolValue() (SymbolID, bool)
	ElementAt(int) SyntaxElement
}

// element 値は*ListかSymbolIDかIntかFloatかStringのいずれか
type element struct {
	value interface{}
	pos   Position
}

// List ListまたはValueを0個以上含む
type List struct {
	openchar rune
	elements []SyntaxElement
	pos      Position
}

// SymbolID シンボルのSTreeにおける一意な識別番号
type SymbolID int

// InvalidSymbolID 無効なシンボルID
const InvalidSymbolID = -1

func (lst *List) isMatchingParen(close rune) bool {
	if (lst.openchar == leftParenthesis && close == rightParenthesis) ||
		(lst.openchar == leftSquareBracket && close == rightSquareBracket) ||
		(lst.openchar == leftCurlyBracket && close == rightCurlyBracket) {
		return true
	}
	return false
}

// Len lstの子要素の数を返す。
func (lst *List) Len() int {
	return len(lst.elements)
}

// Position lstのソースコード上の位置を返す。
func (lst *List) Position() Position {
	return lst.pos
}

// IsList lstがリストの場合はtrueを返す。
func (lst *List) IsList() bool {
	return true
}

// IntValue lstは整数型の値を持たない。
func (lst *List) IntValue() (int64, bool) {
	return 0, false
}

// FloatValue lstは浮動小数点数型の値を持たない。
func (lst *List) FloatValue() (float64, bool) {
	return 0, false
}

// StringValue lstは文字列型の値を持たない。
func (lst *List) StringValue() (string, bool) {
	return "", false
}

// SymbolValue lstはシンボルではない。
func (lst *List) SymbolValue() (SymbolID, bool) {
	return InvalidSymbolID, false
}

// ElementAt lstのindex番目の要素を返す。
func (lst *List) ElementAt(index int) SyntaxElement {
	if index < 0 || index >= len(lst.elements) {
		return nil
	}
	return lst.elements[index]
}

// IntAt lstのindex番目の要素がint64ならその値を返す。
func (lst *List) IntAt(index int) (int64, bool) {
	se := lst.ElementAt(index)
	if se != nil {
		return se.IntValue()
	}
	return 0, false
}

// FloatAt lstのindex番目の要素がfloat64ならその値を返す。
func (lst *List) FloatAt(index int) (float64, bool) {
	se := lst.ElementAt(index)
	if se != nil {
		return se.FloatValue()
	}
	return 0.0, false
}

// StringAt lstのindex番目の要素がstringならその値を返す。
func (lst *List) StringAt(index int) (string, bool) {
	se := lst.ElementAt(index)
	if se != nil {
		return se.StringValue()
	}
	return "", false
}

// SymbolAt lstのindex番目の要素がint64ならその値を返す。
func (lst *List) SymbolAt(index int) (SymbolID, bool) {
	se := lst.ElementAt(index)
	if se != nil {
		return se.SymbolValue()
	}
	return InvalidSymbolID, false
}

func newLiteral(value interface{}, filename string, line int, column int) *element {
	switch value.(type) {
	case int64:
		return &element{value, Position{filename, line, column}}
	case float64:
		return &element{value, Position{filename, line, column}}
	case SymbolID:
		return &element{value, Position{filename, line, column}}
	case string:
		return &element{value, Position{filename, line, column}}
	}
	panic("Unexpected value type")
}

// IsList eがリストならtrueを返す。
func (e *element) IsList() bool {
	return false
}

// Position eのソースコード上の位置を返す。
func (e *element) Position() Position {
	return e.pos
}

// IntValue eが整数リテラルなら、整数リテラルのint64型の値を返す。
func (e *element) IntValue() (int64, bool) {
	v, ok := e.value.(int64)
	return v, ok
}

// FloatValue eが浮動小数点数リテラルなら、浮動小数点数リテラルのfloat64の値を返す。
func (e *element) FloatValue() (float64, bool) {
	v, ok := e.value.(float64)
	return v, ok
}

// StringValue eが文字列リテラルなら、文字列リテラルのstringの値を返す。
func (e *element) StringValue() (string, bool) {
	v, ok := e.value.(string)
	return v, ok
}

// SymbolValue eがシンボルなら、リテラルのSymbolIDを返す。
func (e *element) SymbolValue() (SymbolID, bool) {
	v, ok := e.value.(SymbolID)
	return v, ok
}

func (e *element) ElementAt(_ int) SyntaxElement {
	return nil
}

package runtime

import (
	"math"

	"github.com/healthy-tiger/scalc/parser"
)

const (
	absSymbol           = "abs"
	acosSymbol          = "acos"
	acoshSymbol         = "acosh"
	asinSymbol          = "asin"
	asinhSymbol         = "asinh"
	atanSymbol          = "atan"
	atan2Symbol         = "atan2"
	atanhSymbol         = "atanh"
	cbrtSymbol          = "cbrt"
	ceilSymbol          = "ceil"
	copysignSymbol      = "copy-sign"
	cosSymbol           = "cos"
	coshSymbol          = "cosh"
	dimSymbol           = "dim"
	erfSymbol           = "erf"
	erfcSymbol          = "erfc"
	erfcinvSymbol       = "erfcinv"
	erfinvSymbol        = "erfinv"
	expSymbol           = "exp"
	exp2Symbol          = "exp2"
	expm1Symbol         = "expm1"
	fMASymbol           = "fma"
	floatbitsSymbol     = "float-bits"
	floatfrombitsSymbol = "float-frombits"
	floorSymbol         = "floor"
	gammaSymbol         = "gamma"
	hypotSymbol         = "hypot"
	ilogbSymbol         = "ilogb"
	infSymbol           = "inf"
	isInfSymbol         = "isinf"
	isNaNSymbol         = "isnan"
	j0Symbol            = "j0"
	j1Symbol            = "j1"
	jnSymbol            = "jn"
	ldexpSymbol         = "ldexp"
	logSymbol           = "log"
	log10Symbol         = "log10"
	log1pSymbol         = "log1p"
	log2Symbol          = "log2"
	logbSymbol          = "logb"
	maxSymbol           = "max"
	minSymbol           = "min"
	modSymbol           = "mod"
	naNSymbol           = "nan"
	nextafterSymbol     = "next-after"
	powSymbol           = "pow"
	pow10Symbol         = "pow10"
	remainderSymbol     = "remainder"
	roundSymbol         = "round"
	roundToEvenSymbol   = "round-to-even"
	signbitSymbol       = "signbit"
	sinSymbol           = "sin"
	sinhSymbol          = "sinh"
	sqrtSymbol          = "sqrt"
	tanSymbol           = "tan"
	tanhSymbol          = "tanh"
	truncSymbol         = "trunc"
	y0Symbol            = "y0"
	y1Symbol            = "y1"
	ynSymbol            = "yn"
)

// Goのmathパッケージから以下の関数を導入
// [*] func Abs(x float64) float64
// [ ] func Acos(x float64) float64
// [ ] func Acosh(x float64) float64
// [ ] func Asin(x float64) float64
// [ ] func Asinh(x float64) float64
// [ ] func Atan(x float64) float64
// [ ] func Atan2(y, x float64) float64
// [ ] func Atanh(x float64) float64
// [ ] func Cbrt(x float64) float64
// [ ] func Ceil(x float64) float64
// [ ] func Copysign(x, y float64) float64
// [ ] func Cos(x float64) float64
// [ ] func Cosh(x float64) float64
// [ ] func Dim(x, y float64) float64
// [ ] func Erf(x float64) float64
// [ ] func Erfc(x float64) float64
// [ ] func Erfcinv(x float64) float64
// [ ] func Erfinv(x float64) float64
// [ ] func Exp(x float64) float64
// [ ] func Exp2(x float64) float64
// [ ] func Expm1(x float64) float64
// [ ] func FMA(x, y, z float64) float64
// [ ] func Float64bits(f float64) uint64
// [ ] func Float64frombits(b uint64) float64
// [ ] func Floor(x float64) float64
// [ ] func Gamma(x float64) float64
// [ ] func Hypot(p, q float64) float64
// [ ] func Ilogb(x float64) int
// [ ] func Inf(sign int) float64
// [ ] func IsInf(f float64, sign int) bool
// [ ] func IsNaN(f float64) (is bool)
// [ ] func J0(x float64) float64
// [ ] func J1(x float64) float64
// [ ] func Jn(n int, x float64) float64
// [ ] func Ldexp(frac float64, exp int) float64
// [ ] func Log(x float64) float64
// [ ] func Log10(x float64) float64
// [ ] func Log1p(x float64) float64
// [ ] func Log2(x float64) float64
// [ ] func Logb(x float64) float64
// [ ] func Max(x, y float64) float64
// [ ] func Min(x, y float64) float64
// [ ] func Mod(x, y float64) float64
// [ ] func NaN() float64
// [ ] func Nextafter(x, y float64) (r float64)
// [ ] func Nextafter32(x, y float32) (r float32)
// [ ] func Pow(x, y float64) float64
// [ ] func Pow10(n int) float64
// [ ] func Remainder(x, y float64) float64
// [ ] func Round(x float64) float64
// [ ] func RoundToEven(x float64) float64
// [ ] func Signbit(x float64) bool
// [ ] func Sin(x float64) float64
// [ ] func Sinh(x float64) float64
// [ ] func Sqrt(x float64) float64
// [ ] func Tan(x float64) float64
// [ ] func Tanh(x float64) float64
// [ ] func Trunc(x float64) float64
// [ ] func Y0(x float64) float64
// [ ] func Y1(x float64) float64
// [ ] func Yn(n int, x float64) float64

func absBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Abs(a), nil
}

func acosBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Acos(a), nil
}

func acoshBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Acosh(a), nil
}

func asinBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Asin(a), nil
}

func asinhBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Asinh(a), nil
}

func atanBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Atan(a), nil
}

func atan2Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 3 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	b, berr := EvalAsFloat(lst.ElementAt(2), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	if berr != nil {
		return math.NaN(), berr
	}
	return math.Atan2(a, b), nil
}

func atanhBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Atanh(a), nil
}

func cbrtBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Cbrt(a), nil
}

func ceilBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Ceil(a), nil
}

func copysignBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 3 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	b, berr := EvalAsFloat(lst.ElementAt(2), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	if berr != nil {
		return math.NaN(), berr
	}
	return math.Copysign(a, b), nil
}

func cosBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Cos(a), nil
}

func coshBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Cosh(a), nil
}

//func dimBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func erfBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func erfcBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)          {}
//func erfcinvBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)       {}
//func erfinvBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)        {}
//func expBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func exp2Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)          {}
//func expm1Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func fMABody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func floatbitsBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)     {}
//func floatfrombitsBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {}
//func floorBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func gammaBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func hypotBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func ilogbBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func infBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func isInfBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func isNaNBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func j0Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)            {}
//func j1Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)            {}
//func jnBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)            {}
//func ldexpBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func logBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func log10Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func log1pBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func log2Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)          {}
//func logbBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)          {}
//func maxBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func minBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func modBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func naNBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func nextafterBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)     {}
//func powBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func pow10Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func remainderBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)     {}
//func roundBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func roundToEvenBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)   {}
//func signbitBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)       {}
//func sinBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func sinhBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)          {}
//func sqrtBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)          {}
//func tanBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)           {}
//func tanhBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)          {}
//func truncBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)         {}
//func y0Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)            {}
//func y1Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)            {}
//func ynBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error)            {}

// RegisterMath stに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterMath(ns *Namespace) {
	ns.RegisterExtension(absSymbol, nil, absBody)
}

package runtime

import (
	"math"

	"github.com/healthy-tiger/scalc/parser"
)

const (
	eSymbol       = "E"
	piSymbol      = "Pi"
	phiSymbol     = "Phi"
	sqrt2Symbol   = "Sqrt2"
	sqrtESymbol   = "SqrtE"
	sqrtPiSymbol  = "SqrtPi"
	sqrtPhiSymbol = "SqrtPhi"
	ln2Symbol     = "Ln2"
	log2ESymbol   = "Log2E"
	ln10Symbol    = "Ln10"
	log10ESymbol  = "Log10E"
)

const (
	absSymbol         = "abs"
	acosSymbol        = "acos"
	acoshSymbol       = "acosh"
	asinSymbol        = "asin"
	asinhSymbol       = "asinh"
	atanSymbol        = "atan"
	atan2Symbol       = "atan2"
	atanhSymbol       = "atanh"
	cbrtSymbol        = "cbrt"
	ceilSymbol        = "ceil"
	copysignSymbol    = "copy-sign"
	cosSymbol         = "cos"
	coshSymbol        = "cosh"
	dimSymbol         = "dim"
	erfSymbol         = "erf"
	erfcSymbol        = "erfc"
	erfcinvSymbol     = "erfcinv"
	erfinvSymbol      = "erfinv"
	expSymbol         = "exp"
	exp2Symbol        = "exp2"
	expm1Symbol       = "expm1"
	fMASymbol         = "fma"
	floorSymbol       = "floor"
	gammaSymbol       = "gamma"
	hypotSymbol       = "hypot"
	ilogbSymbol       = "ilogb"
	infSymbol         = "inf"
	isInfSymbol       = "isinf"
	isNaNSymbol       = "isnan"
	j0Symbol          = "j0"
	j1Symbol          = "j1"
	jnSymbol          = "jn"
	ldexpSymbol       = "ldexp"
	logSymbol         = "log"
	log10Symbol       = "log10"
	log1pSymbol       = "log1p"
	log2Symbol        = "log2"
	logbSymbol        = "logb"
	maxSymbol         = "max"
	minSymbol         = "min"
	modSymbol         = "mod"
	naNSymbol         = "nan"
	nextafterSymbol   = "next-after"
	powSymbol         = "pow"
	pow10Symbol       = "pow10"
	remainderSymbol   = "remainder"
	roundSymbol       = "round"
	roundToEvenSymbol = "round-to-even"
	signbitSymbol     = "signbit"
	sinSymbol         = "sin"
	sinhSymbol        = "sinh"
	sqrtSymbol        = "sqrt"
	tanSymbol         = "tan"
	tanhSymbol        = "tanh"
	truncSymbol       = "trunc"
	y0Symbol          = "y0"
	y1Symbol          = "y1"
	ynSymbol          = "yn"
)

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

func dimBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return math.Dim(a, b), nil
}

func erfBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Erf(a), nil
}

func erfcBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Erfc(a), nil
}

func erfcinvBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Erfcinv(a), nil
}

func erfinvBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Erfinv(a), nil
}

func expBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Exp(a), nil
}

func exp2Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Exp2(a), nil
}

func expm1Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Expm1(a), nil
}

func fMABody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 4 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 3)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	b, berr := EvalAsFloat(lst.ElementAt(2), ns)
	c, cerr := EvalAsFloat(lst.ElementAt(3), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	if berr != nil {
		return math.NaN(), berr
	}
	if cerr != nil {
		return math.NaN(), cerr
	}

	return math.FMA(a, b, c), nil
}

func floorBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Floor(a), nil
}

func gammaBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Gamma(a), nil
}

func hypotBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return math.Hypot(a, b), nil
}

func ilogbBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Ilogb(a), nil
}

func infBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.Inf(int(a)), nil
}

func isInfBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 3 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	b, berr := EvalAsInt(lst.ElementAt(2), ns)
	if aerr != nil {
		return 0, aerr
	}
	if berr != nil {
		return 0, berr
	}
	r := math.IsInf(a, int(b))
	return boolToInt(r), nil
}

func isNaNBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	r := math.IsNaN(a)
	return boolToInt(r), nil
}

func j0Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.J0(a), nil
}

func j1Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, err := EvalAsFloat(lst.ElementAt(1), ns)
	if err != nil {
		return math.NaN(), err
	}
	return math.J1(a), nil
}

func jnBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 3 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
	}
	a, aerr := EvalAsInt(lst.ElementAt(1), ns)
	b, berr := EvalAsFloat(lst.ElementAt(2), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	if berr != nil {
		return math.NaN(), berr
	}
	return math.Jn(int(a), b), nil
}

func ldexpBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 3 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	b, berr := EvalAsInt(lst.ElementAt(2), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	if berr != nil {
		return math.NaN(), berr
	}
	return math.Ldexp(a, int(b)), nil
}

func logBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Log(a), nil
}

func log10Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Log10(a), nil
}

func log1pBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Log1p(a), nil
}

func log2Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Log2(a), nil
}

func logbBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Logb(a), nil
}

func maxBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return math.Max(a, b), nil
}

func minBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return math.Min(a, b), nil
}

func modBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return math.Mod(a, b), nil
}

func naNBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 1 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 0)
	}
	return math.NaN(), nil
}

func nextafterBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return math.Nextafter(a, b), nil
}

func powBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return math.Pow(a, b), nil
}

func pow10Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsInt(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Pow10(int(a)), nil
}

func remainderBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
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
	return math.Remainder(a, b), nil
}

func roundBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Round(a), nil
}

func roundToEvenBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.RoundToEven(a), nil
}

func signbitBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	r := math.Signbit(a)
	return boolToInt(r), nil
}

func sinBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Sin(a), nil
}

func sinhBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Sinh(a), nil
}

func sqrtBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Sqrt(a), nil
}

func tanBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Tan(a), nil
}

func tanhBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Tanh(a), nil
}

func truncBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Trunc(a), nil
}

func y0Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Y0(a), nil
}

func y1Body(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	a, aerr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	return math.Y1(a), nil
}

func ynBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 3 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 2)
	}
	a, aerr := EvalAsInt(lst.ElementAt(1), ns)
	b, berr := EvalAsFloat(lst.ElementAt(1), ns)
	if aerr != nil {
		return math.NaN(), aerr
	}
	if berr != nil {
		return math.NaN(), berr
	}
	return math.Yn(int(a), b), nil
}

// RegisterMath stに演算子のシンボルを、nsに演算子に対応する拡張関数をそれぞれ登録する。
func RegisterMath(ns *Namespace) {
	ns.Set(ns.GetSymbolID(eSymbol), float64(math.E))
	ns.Set(ns.GetSymbolID(piSymbol), float64(math.Pi))
	ns.Set(ns.GetSymbolID(phiSymbol), float64(math.Phi))
	ns.Set(ns.GetSymbolID(sqrt2Symbol), float64(math.Sqrt2))
	ns.Set(ns.GetSymbolID(sqrtESymbol), float64(math.SqrtE))
	ns.Set(ns.GetSymbolID(sqrtPiSymbol), float64(math.SqrtPi))
	ns.Set(ns.GetSymbolID(sqrtPhiSymbol), float64(math.SqrtPhi))
	ns.Set(ns.GetSymbolID(ln2Symbol), float64(math.Ln2))
	ns.Set(ns.GetSymbolID(log2ESymbol), float64(math.Log2E))
	ns.Set(ns.GetSymbolID(ln10Symbol), float64(math.Ln10))
	ns.Set(ns.GetSymbolID(log10ESymbol), float64(math.Log10E))
	ns.RegisterExtension(absSymbol, nil, absBody)
	ns.RegisterExtension(absSymbol, nil, absBody)
	ns.RegisterExtension(acosSymbol, nil, acosBody)
	ns.RegisterExtension(acoshSymbol, nil, acoshBody)
	ns.RegisterExtension(asinSymbol, nil, asinBody)
	ns.RegisterExtension(asinhSymbol, nil, asinhBody)
	ns.RegisterExtension(atanSymbol, nil, atanBody)
	ns.RegisterExtension(atan2Symbol, nil, atan2Body)
	ns.RegisterExtension(atanhSymbol, nil, atanhBody)
	ns.RegisterExtension(cbrtSymbol, nil, cbrtBody)
	ns.RegisterExtension(ceilSymbol, nil, ceilBody)
	ns.RegisterExtension(copysignSymbol, nil, copysignBody)
	ns.RegisterExtension(cosSymbol, nil, cosBody)
	ns.RegisterExtension(coshSymbol, nil, coshBody)
	ns.RegisterExtension(dimSymbol, nil, dimBody)
	ns.RegisterExtension(erfSymbol, nil, erfBody)
	ns.RegisterExtension(erfcSymbol, nil, erfcBody)
	ns.RegisterExtension(erfcinvSymbol, nil, erfcinvBody)
	ns.RegisterExtension(erfinvSymbol, nil, erfinvBody)
	ns.RegisterExtension(expSymbol, nil, expBody)
	ns.RegisterExtension(exp2Symbol, nil, exp2Body)
	ns.RegisterExtension(expm1Symbol, nil, expm1Body)
	ns.RegisterExtension(fMASymbol, nil, fMABody)
	ns.RegisterExtension(floorSymbol, nil, floorBody)
	ns.RegisterExtension(gammaSymbol, nil, gammaBody)
	ns.RegisterExtension(hypotSymbol, nil, hypotBody)
	ns.RegisterExtension(ilogbSymbol, nil, ilogbBody)
	ns.RegisterExtension(infSymbol, nil, infBody)
	ns.RegisterExtension(isInfSymbol, nil, isInfBody)
	ns.RegisterExtension(isNaNSymbol, nil, isNaNBody)
	ns.RegisterExtension(j0Symbol, nil, j0Body)
	ns.RegisterExtension(j1Symbol, nil, j1Body)
	ns.RegisterExtension(jnSymbol, nil, jnBody)
	ns.RegisterExtension(ldexpSymbol, nil, ldexpBody)
	ns.RegisterExtension(logSymbol, nil, logBody)
	ns.RegisterExtension(log10Symbol, nil, log10Body)
	ns.RegisterExtension(log1pSymbol, nil, log1pBody)
	ns.RegisterExtension(log2Symbol, nil, log2Body)
	ns.RegisterExtension(logbSymbol, nil, logbBody)
	ns.RegisterExtension(maxSymbol, nil, maxBody)
	ns.RegisterExtension(minSymbol, nil, minBody)
	ns.RegisterExtension(modSymbol, nil, modBody)
	ns.RegisterExtension(naNSymbol, nil, naNBody)
	ns.RegisterExtension(nextafterSymbol, nil, nextafterBody)
	ns.RegisterExtension(powSymbol, nil, powBody)
	ns.RegisterExtension(pow10Symbol, nil, pow10Body)
	ns.RegisterExtension(remainderSymbol, nil, remainderBody)
	ns.RegisterExtension(roundSymbol, nil, roundBody)
	ns.RegisterExtension(roundToEvenSymbol, nil, roundToEvenBody)
	ns.RegisterExtension(signbitSymbol, nil, signbitBody)
	ns.RegisterExtension(sinSymbol, nil, sinBody)
	ns.RegisterExtension(sinhSymbol, nil, sinhBody)
	ns.RegisterExtension(sqrtSymbol, nil, sqrtBody)
	ns.RegisterExtension(tanSymbol, nil, tanBody)
	ns.RegisterExtension(tanhSymbol, nil, tanhBody)
	ns.RegisterExtension(truncSymbol, nil, truncBody)
	ns.RegisterExtension(y0Symbol, nil, y0Body)
	ns.RegisterExtension(y1Symbol, nil, y1Body)
	ns.RegisterExtension(ynSymbol, nil, ynBody)
}

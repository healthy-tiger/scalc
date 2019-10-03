package runtime

import "github.com/healthy-tiger/scalc/parser"

const ifSymbol = "if"

func ifBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() < 4 {
		return nil, newEvalError(lst.Position(), ErrorInsufficientNumberOfArguments)
	} else if lst.Len() > 4 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments)
	}
	p, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if cond, ok := p.(bool); ok {
		if cond {
			return EvalElement(lst.ElementAt(2), ns)
		}
		return EvalElement(lst.ElementAt(3), ns)
	}
	return nil, newEvalError(lst.ElementAt(1).Position(), ErrorOperantsMustBeBoolean)
}

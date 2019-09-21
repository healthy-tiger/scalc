package scalc

import "errors"

var (
	errorTheNumberOfArgumentsDoesNotMatch                           = errors.New("The number of arguments does not match")
	errorUndefinedSymbol                                            = errors.New("Undefined symbol")
	errorAnEmptyListIsNotAllowed                                    = errors.New("An empty list is not allowed")
	errorTheFirstElementOfTheListToBeEvaluatedMustBeASymbol         = errors.New("The first element of the list to be evaluated must be a symbol")
	errorTheFirstElementOfTheListToBeEvaluatedMustBeACallableObject = errors.New("The first element of the list to be evaluated must be a callable object")
	errorFunctionCannotBePassedAsFunctionArgument                   = errors.New("Function cannot be passed as function argument")
	errorInsufficientNumberOfArguments                              = errors.New("Insufficient number of arguments")
	errorTooManyArguments                                           = errors.New("Too many arguments")
)

package runtime

import (
	"time"

	"github.com/healthy-tiger/scalc/parser"
)

const (
	dateSymbol       = "date"
	nowSymbol        = "now"
	daySymbol        = "day"
	hourSymbol       = "hour"
	iszeroSymbol     = "iszero"
	localSymbol      = "local"
	minuteSymbol     = "minute"
	monthSymbol      = "month"
	nanosecondSymbol = "nanosecond"
	secondSymbol     = "second"
	utcSymbol        = "utc"
	weekdaySymbol    = "weekday"
	yearSymbol       = "year"
	yeardaySymbol    = "yearday"
	zoneSymbol       = "zone"
	zoneoffsetSymbol = "zoneoffset"
)

func dateBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 8 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 7)
	}

	year, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	month, err := EvalAsInt(lst.ElementAt(2), ns)
	if err != nil {
		return nil, err
	}
	if month < 1 || month > 12 {
		return nil, newEvalError(lst.Position(), ErrorValueOutOfRange, month, 1, 12)
	}

	day, err := EvalAsInt(lst.ElementAt(3), ns)
	if err != nil {
		return nil, err
	}

	hour, err := EvalAsInt(lst.ElementAt(4), ns)
	if err != nil {
		return nil, err
	}

	min, err := EvalAsInt(lst.ElementAt(5), ns)
	if err != nil {
		return nil, err
	}
	sec, err := EvalAsInt(lst.ElementAt(6), ns)
	if err != nil {
		return nil, err
	}
	nsec, err := EvalAsInt(lst.ElementAt(7), ns)

	return time.Date(int(year), time.Month(month), int(day), int(hour), int(min), int(sec), int(nsec), time.Local), nil
}

func nowBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 1 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments, lst.Len()-1, 0)
	}
	return time.Now(), nil
}

func dayBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return int64(t.Day()), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func hourBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return int64(t.Hour()), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func iszeroBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return t.IsZero(), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func localBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return t.Local(), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func minuteBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return int64(t.Minute()), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func monthBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return int64(t.Month()), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func nanosecondBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return int64(t.Hour()), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func secondBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return int64(t.Second()), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func utcBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return t.UTC(), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func weekdayBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return int64(t.Weekday()), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func yearBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return int64(t.Year()), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func yeardayBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		return int64(t.YearDay()), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func zoneBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		z, _ := t.Zone()
		return z, nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

func zoneoffsetBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	e1, err := EvalElement(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	if t, ok := e1.(time.Time); ok {
		_, o := t.Zone()
		return int64(o), nil
	}
	return nil, newEvalError(lst.Position(), ErrorTypeMissmatch, "time", e1)
}

// RegisterTimeFunc 時刻に関する拡張関数を登録する。
func RegisterTimeFunc(ns *Namespace) {
	ns.RegisterExtension(dateSymbol, nil, dateBody)
	ns.RegisterExtension(nowSymbol, nil, nowBody)
	ns.RegisterExtension(daySymbol, nil, dayBody)
	ns.RegisterExtension(hourSymbol, nil, hourBody)
	ns.RegisterExtension(iszeroSymbol, nil, iszeroBody)
	ns.RegisterExtension(localSymbol, nil, localBody)
	ns.RegisterExtension(minuteSymbol, nil, minuteBody)
	ns.RegisterExtension(monthSymbol, nil, monthBody)
	ns.RegisterExtension(nanosecondSymbol, nil, nanosecondBody)
	ns.RegisterExtension(secondSymbol, nil, secondBody)
	ns.RegisterExtension(utcSymbol, nil, utcBody)
	ns.RegisterExtension(weekdaySymbol, nil, weekdayBody)
	ns.RegisterExtension(yearSymbol, nil, yearBody)
	ns.RegisterExtension(yeardaySymbol, nil, yeardayBody)
	ns.RegisterExtension(zoneSymbol, nil, zoneBody)
	ns.RegisterExtension(zoneoffsetSymbol, nil, zoneoffsetBody)
}

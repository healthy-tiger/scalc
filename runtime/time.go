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
	minuteSymbol     = "minute"
	monthSymbol      = "month"
	secondSymbol     = "second"
	weekdaySymbol    = "weekday"
	yearSymbol       = "year"
	yeardaySymbol    = "yearday"
	zoneSymbol       = "zone"
	zoneoffsetSymbol = "zoneoffset"
)

func dateBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 7 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 6)
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

	return time.Date(int(year), time.Month(month), int(day), int(hour), int(min), int(sec), 0, time.Local).Unix(), nil
}

func nowBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 1 {
		return nil, newEvalError(lst.Position(), ErrorTooManyArguments, lst.Len()-1, 0)
	}
	return time.Now().Unix(), nil
}

func dayBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	sec, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	return int64(time.Unix(sec, int64(0)).Day()), nil
}

func hourBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	sec, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	return int64(time.Unix(sec, int64(0)).Hour()), nil
}

func minuteBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	sec, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	return int64(time.Unix(sec, int64(0)).Minute()), nil
}

func monthBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	sec, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	return int64(time.Unix(sec, int64(0)).Month()), nil
}

func secondBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	sec, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	return int64(time.Unix(sec, int64(0)).Second()), nil
}

func weekdayBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	sec, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	return int64(time.Unix(sec, int64(0)).Weekday()), nil
}

func yearBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	sec, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	return int64(time.Unix(sec, int64(0)).Year()), nil
}

func yeardayBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	sec, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	return int64(time.Unix(sec, int64(0)).YearDay()), nil

}

func zoneBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	sec, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	t := time.Unix(sec, int64(0))
	z, _ := t.Zone()
	return z, nil
}

func zoneoffsetBody(_ interface{}, lst *parser.List, ns *Namespace) (interface{}, error) {
	if lst.Len() != 2 {
		return nil, newEvalError(lst.Position(), ErrorTheNumberOfArgumentsDoesNotMatch, lst.Len()-1, 1)
	}
	sec, err := EvalAsInt(lst.ElementAt(1), ns)
	if err != nil {
		return nil, err
	}
	t := time.Unix(sec, int64(0))
	_, o := t.Zone()
	return int64(o), nil
}

// RegisterTimeFunc 時刻に関する拡張関数を登録する。
func RegisterTimeFunc(ns *Namespace) {
	ns.RegisterExtension(dateSymbol, nil, dateBody)
	ns.RegisterExtension(nowSymbol, nil, nowBody)
	ns.RegisterExtension(daySymbol, nil, dayBody)
	ns.RegisterExtension(hourSymbol, nil, hourBody)
	ns.RegisterExtension(minuteSymbol, nil, minuteBody)
	ns.RegisterExtension(monthSymbol, nil, monthBody)
	ns.RegisterExtension(secondSymbol, nil, secondBody)
	ns.RegisterExtension(weekdaySymbol, nil, weekdayBody)
	ns.RegisterExtension(yearSymbol, nil, yearBody)
	ns.RegisterExtension(yeardaySymbol, nil, yeardayBody)
	ns.RegisterExtension(zoneSymbol, nil, zoneBody)
	ns.RegisterExtension(zoneoffsetSymbol, nil, zoneoffsetBody)
}

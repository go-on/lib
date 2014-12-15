package ractive

import (
	"gopkg.in/go-on/lib.v2/types"
)

func Disabled_(val string) types.Attribute {
	return types.Attribute{"disabled", val}
}

type Event string

func (e Event) OnClick(val string) types.Attribute {
	return types.Attribute{"on-click", string(e) + ":" + val}
}

func (e Event) OnClickEval(val string) types.Attribute {
	return types.Attribute{"on-click", string(e) + ":" + Eval(val)}
}

func (e Event) Trigger(val string) TriggeredEvent {
	return TriggeredEvent{e, val}
}

type TriggeredEvent struct {
	Event Event
	Value string
}

func Eval(s string) string {
	return "{{" + s + "}}"
}

func EvalHTML(s string) string {
	return "{{{" + s + "}}}"
}

func Property(s string) string {
	return "." + s
}

func PropertyEval(s string) string {
	return Eval("." + s)
}

var This = "."

var ThisEval = Eval(".")

func OnClick(events ...TriggeredEvent) types.Attribute {
	str := ""

	for _, ev := range events {
		str += string(ev.Event) + ":" + ev.Value + ";"
	}

	return types.Attribute{"on-click", str}
}

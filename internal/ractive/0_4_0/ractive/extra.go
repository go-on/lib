package ractive

import (
	"github.com/go-on/lib/internal/shared"
)

func Disabled_(val string) shared.Attribute {
	return shared.Attribute{"disabled", val}
}

type Event string

func (e Event) OnClick(val string) shared.Attribute {
	return shared.Attribute{"on-click", string(e) + ":" + val}
}

func (e Event) OnClickEval(val string) shared.Attribute {
	return shared.Attribute{"on-click", string(e) + ":" + Eval(val)}
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

func OnClick(events ...TriggeredEvent) shared.Attribute {
	str := ""

	for _, ev := range events {
		str += string(ev.Event) + ":" + ev.Value + ";"
	}

	return shared.Attribute{"on-click", str}
}

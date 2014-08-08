package ractivehtml

import (
	"github.com/go-on/lib/html"
	"github.com/go-on/lib/html/internal/element"
	"github.com/go-on/lib/types"
)

func context(name string, inner ...interface{}) *element.Element {
	el := element.Elements("{{#" + name + "}}")
	el.Add(inner...)
	el.Add("{{/" + name + "}}")
	return el
}

func ForAll(name string, inner ...interface{}) *element.Element {
	return context(name, inner...)
}

func With(name string, inner ...interface{}) *element.Element {
	return context(name, inner...)
}

func If(name string, inner ...interface{}) *element.Element {
	return context(name, inner...)
}

func Not(name string, inner ...interface{}) *element.Element {
	el := element.Elements("{{^" + name + "}}")
	el.Add(inner...)
	el.Add("{{/" + name + "}}")
	return el
}

func NewTemplate(id types.Id, inner ...interface{}) *element.Element {
	sc := html.SCRIPT(id, types.Attribute{"type", "text/ractive"})
	sc.Add(inner...)
	return sc
}

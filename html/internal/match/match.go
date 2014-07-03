package match

import (
	"fmt"
	"github.com/go-on/builtin"
	. "github.com/go-on/lib/html/internal/element"
	"github.com/go-on/lib/internal/shared"
	"regexp"
	"strings"
)

// something that matches an Element
type Matcher interface {
	Matches(*Element) bool
}

type or []Matcher

// matches if any of the Matchers matches
func (ø or) Matches(e *Element) bool {
	for _, m := range ø {
		if m.Matches(e) {
			return true
		}
	}
	return false
}

func Or(m ...Matcher) or {
	return or(m)
}

type and []Matcher

// matches if all of the Matchers matches
func (ø and) Matches(e *Element) bool {
	for _, m := range ø {
		if !m.Matches(e) {
			return false
		}
	}
	return true
}

func And(m ...Matcher) and {
	return and(m)
}

type not struct{ Matcher }

// matches if inner matcher does not match
func (ø *not) Matches(e *Element) bool { return !ø.Matcher.Matches(e) }

func Not(m Matcher) *not { return &not{m} }

type PositionMatcher struct {
	Element *Element
	Pos     int
	Found   bool
}

func (ø *PositionMatcher) Matches(e *Element) (f bool) {
	// no recursive findings
	if e.Parent != ø.Element.Parent {
		return
	}
	if ø.Element == e {
		ø.Found = true
		return true
	}
	if !ø.Found {
		ø.Pos += 1
	}
	return
}

type FieldMatcher int

func (ø FieldMatcher) Matches(t *Element) (m bool) { return Is(t, FormField) }

func New(thing interface{}) Matcher {
	switch th := thing.(type) {
	case shared.Class:
		return Class(th)
	case shared.Id:
		return Id(th)
	case shared.Tag:
		return Tag(th)
	case shared.HTMLString:
		return HTML(th)
	case shared.Style:
		return Style(th)
	case shared.Attribute:
		return Attribute(th)
	case string:
		return tag(th)
	case builtin.Stringer:
		return tag(th.String())
	default:
		return tag(fmt.Sprintf("%v", th))
	}
}

type class string

func Class(c shared.Class) Matcher {
	return class(string(c))
}

func (ø class) Matches(t *Element) bool {

	for _, c := range t.Classes {
		if string(c) == string(ø) {
			return true
		}
	}
	return false
}

// if Css has a Context, matching always fails
/*
func (ø *Css) Matches(t *Element) (m bool) {
	if ø.Context != "" {
		// if Css has a Context, matching always fails
		return false
	}
	if ø.class != "" {
		if !t.HasClass(ø.class) {
			return false
		}
	}

	if len(ø.Tags) > 0 {
		for _, tt := range ø.Tags {
			if ø.matchTag(t, Tag(tt)) {
				return true
			}
		}
	} else {
		return true
	}

	return
}
*/

type id string

func Id(i shared.Id) Matcher {
	return id(string(i))
}

func (ø id) Matches(t *Element) bool {
	if string(t.Id) == string(ø) {
		return true
	}
	return false
}

type html string

func HTML(h shared.HTMLString) Matcher {
	return html(h.String())
}

// matching an html string, ignoring whitespace
func (ø html) Matches(t *Element) bool {
	inner := removeWhiteSpace(InnerHtml(t))
	me := removeWhiteSpace(string(ø))
	if inner == me {
		return true
	}
	return false
}

type tag string

func Tag(t shared.Tag) Matcher {
	return tag(string(t))
}

func (ø tag) Matches(t *Element) bool {
	return string(ø) == t.Tag()
}

type style struct {
	key, value string
}

func Style(s shared.Style) Matcher {
	return style{s.Property, s.Value}
}

func (ø style) Matches(t *Element) bool {
	for _, st := range t.Styles {
		if st.Property == ø.key && st.Value == ø.value {
			return true
		}
	}
	return false
}

type attribute struct {
	key, value string
}

func Attribute(a shared.Attribute) Matcher {
	return attribute{a.Key, a.Value}
}

func (ø attribute) Matches(t *Element) bool {
	if ø.key == "id" {
		return id(ø.value).Matches(t)
	}
	if ø.key == "class" {
		return class(ø.value).Matches(t)
	}
	if ø.key == "style" {
		styles := strings.Split(ø.value, ";")
		m := true
		for _, st := range styles {
			a := strings.Split(st, ":")
			styl := style{a[0], a[1]}
			if !styl.Matches(t) {
				m = false
			}
		}
		return m
	}

	for _, a := range t.Attributes {
		if a.Key == ø.key && a.Value == ø.value {
			return true
		}
	}
	return false

}

func removeWhiteSpace(in string) string {
	reg := regexp.MustCompile(`\s`)
	return reg.ReplaceAllString(in, "")
}

// filter by anything that fullfills the matcher interface,
// e.g. Class, Id, Attr, Attrs, Css, Tag, Style, Styles
// recursive finds all tags from the children
func All(ø *Element, m Matcher) (r []*Element) {
	r = []*Element{}
	if len(ø.Children) == 0 {
		return
	}
	for _, in := range ø.Children {
		switch t := in.(type) {
		case *Element:
			if m.Matches(t) {
				r = append(r, t)
			}
			innerFound := All(t, m)
			for _, innerT := range innerFound {
				r = append(r, innerT)
			}
		}
	}
	return
}

// filter by anything that fullfills the matcher interface,
// e.g. Class, Id, Attr, Attrs, Css, Tag, Style, Styles
// returns the first tag in the children and the subchildren that matches
func Any(ø *Element, m Matcher) (r *Element) {
	if len(ø.Children) == 0 {
		return nil
	}
	for _, in := range ø.Children {
		switch t := in.(type) {
		case *Element:
			if m.Matches(t) {
				r = t
				return
			}
			r = Any(t, m)
			if r != nil {
				return
			}
		}
	}
	return nil
}
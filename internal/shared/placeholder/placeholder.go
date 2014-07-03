package placeholder

import (
	"fmt"
	"github.com/go-on/lib/internal/replacer"
	"github.com/go-on/lib/internal/shared"
	"github.com/go-on/lib/internal/template"
	"reflect"

	ph "github.com/go-on/lib/internal/template/placeholder"
)

var (
	commentType   = reflect.TypeOf(shared.Comment("")).Name()
	descrType     = reflect.TypeOf(shared.Descr("")).Name()
	classType     = reflect.TypeOf(shared.Class("")).Name()
	idType        = reflect.TypeOf(shared.Id("")).Name()
	htmlType      = reflect.TypeOf(shared.HTMLString("")).Name()
	textType      = reflect.TypeOf(shared.Text("")).Name()
	attributeType = reflect.TypeOf(shared.Attribute{}).Name()
	tagType       = reflect.TypeOf(shared.Tag("")).Name()
	styleType     = reflect.TypeOf(shared.Style{}).Name()
)

// TODO: extract the template.Setter interface in a extra subpackage and just include that subpackage
type Placeholder interface {
	ph.Setter
	Set(val interface{}) ph.Setter
	Setf(format string, val ...interface{}) ph.Setter
	String() string
	Type() interface{}
	// Handle(http.Handler) template.PlaceholderHandler
	// HandleFunc(func(http.ResponseWriter, *http.Request)) template.PlaceholderHandler
}

func New(thing interface{}) Placeholder {

	switch ø := thing.(type) {
	case shared.Comment:
		return newTPh(template.NewPlaceholder(commentType+"."+string(ø)), ø)
	case shared.Descr:
		return newTPh(template.NewPlaceholder(descrType+"."+string(ø)), ø)
	case shared.Class:
		return newTPh(template.NewPlaceholder(classType+"."+string(ø)), ø)
	case shared.Id:
		return newTPh(template.NewPlaceholder(idType+"."+string(ø)), ø)
	case shared.HTMLString:
		return newTPh(template.NewPlaceholder(htmlType+"."+string(ø)), ø)
	case shared.Text:
		t := template.NewPlaceholder(textType+"."+string(ø), handleStrings(shared.EscapeHTML))
		return newTPh(t, ø)
	case shared.Attribute:
		fn := func(in string) string {
			return (shared.Attribute{ø.Key, in}).String()
		}
		t := template.NewPlaceholder(attributeType+"."+ø.Value, handleStrings(fn))
		return newTPh(t, ø)
	case shared.Tag:
		return newTPh(template.NewPlaceholder(tagType+"."+string(ø)), ø)
	case shared.Style:
		fn := func(in string) string {
			return (shared.Style{ø.Property, in}).String()
		}
		t := template.NewPlaceholder(styleType+"."+ø.Value, handleStrings(fn))
		return newTPh(t, ø)
	case string:
		t := template.NewPlaceholder(textType+"."+ø, handleStrings(shared.EscapeHTML))
		return newTPh(t, shared.Text(ø))
	default:
		str := fmt.Sprintf("%v", ø)
		t := template.NewPlaceholder(textType+"."+str, handleStrings(shared.EscapeHTML))
		return newTPh(t, shared.Text(str))
	}

}

func newTPh(ph ph.Placeholder, i interface{}) typedPlaceholder {
	return typedPlaceholder{ph, i}
}

type typedPlaceholder struct {
	ph.Placeholder
	typ interface{}
}

func (ø typedPlaceholder) String() string {
	return replacer.Placeholder(ø.Name()).String()
}

func (ø typedPlaceholder) Type() interface{} {
	return ø.typ
}

func handleStrings(trafo func(string) string) func(interface{}) string {
	return func(in interface{}) (out string) {
		if in == nil {
			return ""
		}
		var s string
		switch v := in.(type) {
		case fmt.Stringer:
			s = v.String()
		case string:
			s = v
		default:
			s = fmt.Sprintf("%v", v)
		}
		return trafo(s)
	}
}

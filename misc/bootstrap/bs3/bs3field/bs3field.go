package bs3field

import (
	. "github.com/go-on/lib/html"
	. "github.com/go-on/lib/html/element"
	b "github.com/go-on/lib/misc/bootstrap/bs3"
)

type Field struct {
	Name, Label, Value string
}

func (f Field) Select(inner ...interface{}) *Element {
	inner = append(inner, b.Form_control, Attrs_(
		"name", f.Name,
		"value", f.Value,
	))

	return DIV(b.Form_group,
		LabelFor(f.Name, f.Label),
		SELECT(inner...),
	)
}

func (f Field) InputText(inner ...interface{}) *Element {
	inner = append(inner, b.Form_control, Attrs_(
		"value", f.Value,
	))
	return DIV(b.Form_group,
		LabelFor(f.Name, f.Label),
		InputText(f.Name, inner...),
	)
}

func (f Field) Textarea(inner ...interface{}) *Element {
	inner = append(inner, b.Form_control, Name_(f.Name), Value_(f.Value))
	return DIV(b.Form_group,
		LabelFor(f.Name, f.Label),
		TEXTAREA(inner...),
	)
}

func SelectButton(name string, inner ...*Element) *Element {
	ul := UL(b.Dropdown_menu, Role_("menu"))

	for _, el := range inner {
		ul.Add(
			LI(
				Role_("presentation"),
				el,
			),
		)
	}

	return DIV(b.Btn_group,
		BUTTON(b.Btn, b.Btn_primary, b.Dropdown_toggle, Type_("button"), DataToggle_("dropdown"),
			name+" ", SPAN(b.Caret),
		),
		ul,
	)
}

package bs3container

import (
	"gopkg.in/go-on/lib.v3/html"
	"gopkg.in/go-on/lib.v3/html/internal/element"
	"gopkg.in/go-on/lib.v3/internal/bootstrap/bs3"
	"gopkg.in/go-on/lib.v3/types"
)

func New(inner ...interface{}) *element.Element {
	inner = append(inner, bs3.Container)
	return html.DIV(inner...)
}

func Fluid(inner ...interface{}) *element.Element {
	inner = append(inner, bs3.Container_fluid)
	return html.DIV(inner...)
}

func Row(inner ...interface{}) *element.Element {
	inner = append(inner, bs3.Row)
	return html.DIV(inner...)
}

type panel struct {
	header []interface{}
	footer []interface{}
	body   []interface{}
	class  types.Class
}

func (p *panel) AddHeader(inner ...interface{}) *panel {
	p.header = append(p.header, inner...)
	return p
}

func (p *panel) AddFooter(inner ...interface{}) *panel {
	p.footer = append(p.footer, inner...)
	return p
}

func (p *panel) AddBody(inner ...interface{}) *panel {
	p.body = append(p.body, inner...)
	return p
}

func (p *panel) Element() *element.Element {
	el := html.DIV(bs3.Panel, bs3.Panel_default)
	if len(p.header) > 0 {
		el.Add(html.DIV(append([]interface{}{bs3.Panel_heading}, p.header...)...))
	}
	if len(p.body) > 0 {
		el.Add(html.DIV(append([]interface{}{bs3.Panel_body}, p.body...)...))
	}
	if len(p.footer) > 0 {
		el.Add(html.DIV(append([]interface{}{bs3.Panel_footer}, p.footer...)...))
	}
	return el
}

func newPanel(cl types.Class) *panel {
	return &panel{class: cl}
}

func PanelDefault() *panel {
	return newPanel(bs3.Panel_default)
}

func PanelDanger() *panel {
	return newPanel(bs3.Panel_danger)
}

func PanelWarning() *panel {
	return newPanel(bs3.Panel_warning)
}

func PanelSuccess() *panel {
	return newPanel(bs3.Panel_success)
}

func PanelPrimary() *panel {
	return newPanel(bs3.Panel_primary)
}

func PanelInfo() *panel {
	return newPanel(bs3.Panel_info)
}

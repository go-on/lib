package bs3

import (
	h "gopkg.in/go-on/lib.v3/html"
	"gopkg.in/go-on/lib.v3/html/internal/element"
	html "gopkg.in/go-on/lib.v3/types"
)

type VERSION string

func (v VERSION) cdnBase() string {
	return "//maxcdn.bootstrapcdn.com/bootstrap/" + string(v) + "/"
}

func (v VERSION) CSS() string {
	return v.cdnBase() + "css/bootstrap.css"
}

func (v VERSION) CSSMin() string {
	return v.cdnBase() + "css/bootstrap.min.css"
}

func (v VERSION) CSSTheme() string {
	return v.cdnBase() + "css/bootstrap-theme.css"
}

func (v VERSION) CSSThemeMin() string {
	return v.cdnBase() + "css/bootstrap-theme.min.css"
}

func (v VERSION) JS() string {
	return v.cdnBase() + "js/bootstrap.js"
}

func (v VERSION) JSMin() string {
	return v.cdnBase() + "js/bootstrap.min.js"
}

var html5shivMin = "//oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"
var respondMin = "//oss.maxcdn.com/respond/1.4.2/respond.min.js"
var jqueryMin = "//ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"

var V3_3 = VERSION("3.3.4")
var V3_1 = VERSION("3.1.1")

func (v VERSION) Head(cdnFn func(cdnURL string) string) *element.Element {
	return element.Elements(
		h.CharsetUtf8(),
		h.HttpEquiv("X-UA-Compatible", "IE=edge"),
		h.Viewport("width=device-width, initial-scale=1"),

		h.CssHref(cdnFn(v.CSSMin())),

		html.HTMLString(`<!--[if lt IE 9]>
      <script src="`+cdnFn(html5shivMin)+`"></script>
      <script src="`+cdnFn(respondMin)+`"></script>
    <![endif]-->`),
	)
}

func (v VERSION) Body(cdnFn func(cdnURL string) string) *element.Element {
	return element.Elements(
		h.JsSrc(cdnFn(jqueryMin)),
		h.JsSrc(cdnFn(v.JSMin())),
	)
}

func (v VERSION) New(htmlattrs ...interface{}) *document {
	return newDocument(v, htmlattrs...)
}

type document struct {
	VERSION
	attrs   []interface{}
	head    []interface{}
	body    []interface{}
	bodyEnd []interface{}
}

func newDocument(v VERSION, attrs ...interface{}) *document {
	return &document{
		VERSION: v,
		attrs:   attrs,
	}
}

func (d *document) AddToHead(v ...interface{}) *document {
	d.head = append(d.head, v...)
	return d
}

func (d *document) AddToBody(v ...interface{}) *document {
	d.body = append(d.body, v...)
	return d
}

func (d *document) AddToEndOfBody(v ...interface{}) *document {
	d.bodyEnd = append(d.bodyEnd, v...)
	return d
}

func (d *document) DocType(cdnFn func(cdnURL string) string) *h.DocType {
	head := append([]interface{}{d.VERSION.Head(cdnFn)}, d.head...)
	body := append(d.body, d.VERSION.Body(cdnFn))
	body = append(body, d.bodyEnd...)
	inner := append([]interface{}{}, d.attrs...)
	inner = append(inner, h.HEAD(head...), h.BODY(body...))
	return h.HTML5(h.HTML(inner...))
}

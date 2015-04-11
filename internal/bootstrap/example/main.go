package main

import (
	"encoding/json"
	"gopkg.in/go-on/cdncache.v1"
	. "gopkg.in/go-on/lib.v3/html"
	"gopkg.in/go-on/lib.v3/html/internal/element"
	bs "gopkg.in/go-on/lib.v3/internal/bootstrap/bs3"
	bsmenu "gopkg.in/go-on/lib.v3/internal/bootstrap/bs3/bs3menu"
	"gopkg.in/go-on/lib.v3/internal/menu"
	"gopkg.in/go-on/lib.v3/internal/menu/menuhandler"
	"gopkg.in/go-on/lib.v3/types"
	"net/http"
)

func main() {

	doc := bs.V3_1.New(Lang_("de"))
	doc.AddToHead(TITLE("Just a test"))
	doc.AddTheme()
	doc.AddToBody(mkBody(mkMenu(menuJson)))

	http.Handle("/", doc.DocType(cdncache.CDN("/cdn-cache/")))

	http.ListenAndServe(":8080", nil)
}

func mkMenu(mjson string) *menu.Node {
	m := &menu.Node{}

	if err := json.Unmarshal([]byte(mjson), &m); err != nil {
		panic(err.Error())
	}
	return m
}

func mkBody(m *menu.Node) *element.Element {
	return element.Elements(
		mkNav(m, 1),
		mkBreadcrumb(m, 2),

		DIV(
			bs.Container_fluid,
			DIV(bs.Col_sm_6, mkPanel("Dropdown Buttons", mkDropDownButtons(m))),
			DIV(bs.Col_sm_3, mkPanel("Pills stacked @ 1-2", mkPills(m))),
			DIV(bs.Col_sm_3, mkPanel("Listgroup @ 2", mkListGroup(m))),
			DIV(bs.Col_sm_6, mkPanel("Tabs @ 2", mkTabNav(m), tabContent)),
		),
	)
}

func mkPanel(title string, body ...interface{}) *element.Element {
	return DIV(bs.Panel, bs.Panel_default,
		DIV(bs.Panel_heading, title),
		DIV(bs.Panel_body, element.Elements(body...)),
	)
}

func mkNav(m *menu.Node, level int) *element.Element {
	return NAV(bs.Navbar, bs.Navbar_default,
		DIV(
			bs.Container_fluid,
			DIV(bs.Navbar_header, SPAN(bs.Navbar_brand, "NavBar Menu @ 0-1")),
			menuhandler.NewStatic(m, level, bsmenu.NavBar()),
		),
	)
}

func mkBreadcrumb(m *menu.Node, level int) *element.Element {
	return DIV(
		bs.Container_fluid,
		menuhandler.NewStatic(m, level, bsmenu.Breadcrumb()),
	)
}

func mkDropDownButtons(m *menu.Node) *element.Element {
	return DIV(bs.Btn_toolbar,
		DIV(
			bs.Btn_group,
			menuhandler.NewStaticSub(m, 0, 0, bsmenu.DropdownButton(bs.Btn_default, "", "Category")),
			menuhandler.NewStaticSub(m, 0, 0, bsmenu.Dropdown()),
		),
		DIV(
			bs.Btn_group,
			menuhandler.NewStaticSub(m, 1, 1, bsmenu.Button(bs.Btn_success, "%s", "»")),
			menuhandler.NewStaticSub(m, 1, 1, bsmenu.DropdownButton(bs.Btn_success, "", "")),
			menuhandler.NewStaticSub(m, 1, 1, bsmenu.Dropdown()),
		),
		DIV(
			bs.Btn_group,
			menuhandler.NewStaticSub(m, 2, 2, bsmenu.Button(bs.Btn_warning, "%s", "»")),
			menuhandler.NewStaticSub(m, 2, 2, bsmenu.DropdownButton(bs.Btn_warning, "", "")),
			menuhandler.NewStaticSub(m, 2, 2, bsmenu.Dropdown()),
		),
	)
}

func mkTabNav(m *menu.Node) http.Handler {
	return menuhandler.NewStaticSub(m, 2, 2, bsmenu.Tabs(true, true))
}

func mkPills(m *menu.Node) http.Handler {
	return menuhandler.NewStaticSub(m, 1, 2, bsmenu.Pills(true, bs.Nav_stacked))
}

func mkListGroup(m *menu.Node) http.Handler {
	return menuhandler.NewStaticSub(m, 2, 2, bsmenu.ListGroup())
}

var tabContent = DIV(bs.Tab_content,
	DIV(bs.Tab_pane, types.Id("uk"),
		AHref("http://en.wikipedia.org/wiki/United_Kingdom", "From Wikipedia:"),
		CITE(
			`The United Kingdom of Great Britain and Northern Ireland,`+
				`commonly known as the United Kingdom (UK) or Britain /ˈbrɪ.tən/, `+
				`is a sovereign state located off the north-western coast of `+
				`continental Europe.`,
		),
	),
	DIV(bs.Tab_pane, types.Id("france"),
		AHref("http://en.wikipedia.org/wiki/France", "From Wikipedia:"),
		CITE(
			`France (UK: /frɑːns/; US: Listeni/fræns/; French: [fʁɑ̃s], `+
				`officially the French Republic (French: République française [ʁepyblik fʁɑ̃sɛz]), `+
				`is a sovereign country in Western Europe that includes overseas `+
				`regions and territories.`,
		),
	),
)

var menuJson = `
{
  "Subs": [
    { "Text": "Languages", "Path": "/languages",
      "Subs": [
        { "Text": "english", "Path": "/english",
          "Subs": [
            { "Text": "american english", "Path": "/en_us"},
            { "Text": "british english", "Path": "/en_gb" }
          ]
        },
        {"Text": "---"},
        {"Text": "french", "Path": "/fr"}
      ]
    },
    { "Text": "Countries", "Path": "/countries",
      "Subs": [
        { "Text": "USA", "Path": "/usa" },
        { "Text": "Brazil", "Path": "_" },
        { "Text": "Europe", "Path": "/europe",
          "Subs": [
            { "Text": "UK", "Path": "#uk" },
            { "Text": "France", "Path": "#france"}
          ]
        }
      ]
    },
    { "Text": "Currencies", "Path": "/currencies" }
  ]
}
`

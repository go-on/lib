package main

import (
	"encoding/json"
	"gopkg.in/go-on/cdncache.v1"
	. "gopkg.in/go-on/lib.v3/html"
	"gopkg.in/go-on/lib.v3/html/internal/element"
	"gopkg.in/go-on/lib.v3/internal/bootstrap/bs3"
	"gopkg.in/go-on/lib.v3/internal/bootstrap/bs3/bs3menu"
	"gopkg.in/go-on/lib.v3/internal/menu"
	"gopkg.in/go-on/lib.v3/internal/menu/menuhandler"
	"gopkg.in/go-on/lib.v3/types"
	"net/http"
)

func main() {

	doc := bs3.V_3_1_1.Document(Lang_("de"))
	doc.AddToHead(Title_("Just a test"))

	m := &menu.Node{}

	if err := json.Unmarshal([]byte(menuJson), &m); err != nil {
		panic(err.Error())
	}
	doc.AddToBody(mkBody(m))

	var cdn = cdncache.CDN("/cdn-cache/")
	http.Handle("/", doc.DocType(cdn))

	http.ListenAndServe(":8080", nil)
}

func mkBody(m *menu.Node) *element.Element {
	return element.Elements(
		mkNav(m, 1),
		mkBreadcrumb(m, 2),

		DIV(
			bs3.Container_fluid,
			DIV(bs3.Col_sm_6, mkPanel("Dropdown Buttons", mkDropDownButtons(m))),
			DIV(bs3.Col_sm_3, mkPanel("Pills stacked @ 1-2", mkPills(m))),
			DIV(bs3.Col_sm_3, mkPanel("Listgroup @ 2", mkListGroup(m))),
			DIV(bs3.Col_sm_6, mkPanel("Tabs @ 2", mkTabNav(m), tabContent)),
		),
	)
}

func mkPanel(title string, body ...interface{}) *element.Element {
	return DIV(bs3.Panel, bs3.Panel_default,
		DIV(bs3.Panel_heading, title),
		DIV(bs3.Panel_body, element.Elements(body...)),
	)
}

func mkNav(m *menu.Node, level int) *element.Element {
	return NAV(bs3.Navbar, bs3.Navbar_default,
		DIV(
			bs3.Container_fluid,
			DIV(bs3.Navbar_header, SPAN(bs3.Navbar_brand, "NavBar Menu @ 0-1")),
			menuhandler.NewStatic(m, level, bs3menu.NavBar()),
		),
	)
}

func mkBreadcrumb(m *menu.Node, level int) *element.Element {
	return DIV(
		bs3.Container_fluid,
		menuhandler.NewStatic(m, level, bs3menu.Breadcrumb()),
	)
}

func mkDropDownButtons(m *menu.Node) *element.Element {
	return DIV(bs3.Btn_toolbar,
		DIV(
			bs3.Btn_group,
			menuhandler.NewStaticSub(m, 0, 0, bs3menu.DropdownButton(bs3.Btn_default, "", "Category")),
			menuhandler.NewStaticSub(m, 0, 0, bs3menu.Dropdown()),
		),
		DIV(
			bs3.Btn_group,
			menuhandler.NewStaticSub(m, 1, 1, bs3menu.Button(bs3.Btn_success, "%s", "»")),
			menuhandler.NewStaticSub(m, 1, 1, bs3menu.DropdownButton(bs3.Btn_success, "", "")),
			menuhandler.NewStaticSub(m, 1, 1, bs3menu.Dropdown()),
		),
		DIV(
			bs3.Btn_group,
			menuhandler.NewStaticSub(m, 2, 2, bs3menu.Button(bs3.Btn_warning, "%s", "»")),
			menuhandler.NewStaticSub(m, 2, 2, bs3menu.DropdownButton(bs3.Btn_warning, "", "")),
			menuhandler.NewStaticSub(m, 2, 2, bs3menu.Dropdown()),
		),
	)
}

func mkTabNav(m *menu.Node) http.Handler {
	return menuhandler.NewStaticSub(m, 2, 2, bs3menu.Tabs(true, true))
}

func mkPills(m *menu.Node) http.Handler {
	return menuhandler.NewStaticSub(m, 1, 2, bs3menu.Pills(true, bs3.Nav_stacked))
}

func mkListGroup(m *menu.Node) http.Handler {
	return menuhandler.NewStaticSub(m, 2, 2, bs3menu.ListGroup())
}

var tabContent = DIV(bs3.Tab_content,
	DIV(bs3.Tab_pane, types.Id("uk"),
		AHref("http://en.wikipedia.org/wiki/United_Kingdom", "From Wikipedia:"),
		CITE(
			`The United Kingdom of Great Britain and Northern Ireland,`+
				`commonly known as the United Kingdom (UK) or Britain /ˈbrɪ.tən/, `+
				`is a sovereign state located off the north-western coast of `+
				`continental Europe.`,
		),
	),
	DIV(bs3.Tab_pane, types.Id("france"),
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

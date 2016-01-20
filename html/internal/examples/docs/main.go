package main

import (
	"fmt"

	. "github.com/go-on/lib/html"
	. "github.com/go-on/lib/html/element"
	"github.com/go-on/lib/html/internal/match"
	. "github.com/go-on/lib/types"
	// . "github.com/go-on/lib/html/tag"
)

func main() {
	content := DIV(
		Id("content"), // sets the id attribute
		A(
			Class("button"), // sets the class attribute
			Attribute{"href", "#"},
			"something",                                     // sets the inner text
			Style{"color", "red"},                           // sets the style attribute
			STRONG(Id("sth"), HTMLString("not <escaped>!")), // sets some inner html.Element
		),
	)

	children := content.Children
	// el := children[0].(*Element).
	inner := InnerHtml(children[0].(*Element)) // everything inside
	buttons := match.All(content, match.New(Class("button")))
	sth := match.Any(content, match.New(Id("sth")))

	fmt.Printf(`
children[0].Classes(): %#v
inner: %#v
buttons[0].Tag(): %#v
sth: %#v
`, children[0].(*Element), inner, buttons[0].Tag(), sth.String())

	fmt.Println(content.String())
}

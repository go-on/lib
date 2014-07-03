package main

import (
	"os"

	. "github.com/go-on/lib/html"
	"github.com/go-on/lib/html/internal/element"
	. "github.com/go-on/lib/internal/shared"
	"github.com/go-on/lib/internal/shared/placeholder"
)

var (
	name    = placeholder.New(Text("name"))
	content = DIV(name, element.NewElement("p"), "hello world").Template("content")
	layout  = SECTION(content).Template("layout")
)

func main() {

	all := content.New()
	content.ReplaceTo(all.Buffer, name.Set("<heino>"))
	content.ReplaceTo(all.Buffer, name.Set("<hannelore>"))

	layout.Replace(all).WriteTo(os.Stdout)
}

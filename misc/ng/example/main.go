package main

import (
	"fmt"
	// "gopkg.in/metakeule/goh4.v5"
	. "github.com/go-on/lib/html"
	. "github.com/go-on/lib/types"
	// . "github.com/go-on/lib/html/tag"
	// . "gopkg.in/metakeule/goh4.v5/tag/short"
	"github.com/go-on/lib/misc/ng"
	// "strings"
)

func main() {
	fmt.Printf("%T: %#v\n", ng.Show("currentSection"), ng.Show("currentSection").String())
	fmt.Println(
		DIV(Class("col-xs-4"),
			ng.Show("currentSection"),
			DIV(
				Class("row"),
				"huhu",
			),
		).String(),
	)
}

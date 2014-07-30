package main

import (
	"fmt"
	// "github.com/metakeule/goh4"
	. "github.com/go-on/lib/html"
	. "github.com/go-on/lib/internal/shared"
	// . "github.com/go-on/html/tag"
	// . "github.com/metakeule/goh4/tag/short"
	"github.com/go-on/lib/internal/ng"
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

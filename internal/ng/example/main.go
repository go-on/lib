package main

import (
	"fmt"
	// "github.com/metakeule/goh4"
	. "gopkg.in/go-on/lib.v2/html"
	. "gopkg.in/go-on/lib.v2/types"
	// . "github.com/go-on/html/tag"
	// . "github.com/metakeule/goh4/tag/short"
	"gopkg.in/go-on/lib.v2/internal/ng"
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

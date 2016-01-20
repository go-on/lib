package main

import (
	"fmt"

	. "github.com/go-on/lib/html"
	"github.com/go-on/lib/html/element"
	// "github.com/go-on/lib/html/tag"
)

func main() {
	fmt.Println(element.Elements(LI("a"), LI("b")).String())
}

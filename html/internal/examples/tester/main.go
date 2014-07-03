package main

import (
	"fmt"

	. "github.com/go-on/lib/html"
	"github.com/go-on/lib/html/internal/element"
	// "github.com/go-on/html/tag"
)

func main() {
	fmt.Println(element.Elements(LI("a"), LI("b")).String())
}

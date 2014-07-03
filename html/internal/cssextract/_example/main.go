package main

import (
	"fmt"

	"github.com/go-on/html/cssextract"
)

var styles = `
  table.a { 
			display: none; 
			font-weight:bold; 
			background-image: url('image.jpg');
		}`

func main() {
	p := cssextract.Parse(styles)

	fmt.Printf("%#v\n", p)
}

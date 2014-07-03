package main

import (
	"github.com/go-on/html/types"
)

func main() {
	println(types.Attribute{"data", "<hiho>"}.String())
}

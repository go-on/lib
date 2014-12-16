package main

import (
	"gopkg.in/go-on/lib.v2/html/types"
)

func main() {
	println(types.Attribute{"data", "<hiho>"}.String())
}

package main

import (
	"github.com/go-on/html/tag"

	. "github.com/go-on/bootstrap/bs3"
)

func main() {
	tag.DIV(
		Success,
		BtnGroupXs,
		//Animated,
		"Success",

		tag.BUTTON(
			Btn, BtnDefault, BtnDanger,

			"Help!",
		),
	).Print()
}

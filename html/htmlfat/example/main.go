package main

import (
	"fmt"

	. "github.com/go-on/lib/html"
	. "github.com/go-on/lib/html/htmlfat"
	"github.com/go-on/lib/misc/fat"
)

type Person struct {
	FirstName *fat.Field `type:"string text"`
	LastName  *fat.Field `type:"string"`
	Vita      *fat.Field `type:"string html"`
}

var PERSON = fat.Proto(&Person{}).(*Person)

func NewPerson() *Person { return fat.New(PERSON, &Person{}).(*Person) }

func init() {
	Register(PERSON)
}

func main() {

	ul := UL("\n",
		LI(Placeholder(PERSON.FirstName)), "\n",
		LI(Placeholder(PERSON.LastName)), "\n",
		LI(Placeholder(PERSON.Vita)), "\n",
	)

	fmt.Println(ul.String())

	details := ul.Template("details")

	paul := NewPerson()
	paul.FirstName.Set("<Pa>ul")
	paul.LastName.Set("Pa<n>zer")
	paul.Vita.Set("<p>hier die vita</p>")

	details.Replace(Setters(paul)...).Println()

}

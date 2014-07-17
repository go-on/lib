package main

import (
	"fmt"

	. "github.com/go-on/lib/html"
	. "github.com/go-on/lib/html/internal/element"
	"github.com/go-on/lib/html/internal/element/compiler"
	. "github.com/go-on/lib/html/internal/htmlfat"
	"github.com/go-on/lib/internal/shared"
	"github.com/go-on/router"
	. "github.com/go-on/router/internal/routerfat"
	"github.com/go-on/router/route"
	"github.com/go-on/wrap"
	"github.com/go-on/wrap-contrib-testing/wrapstesting"
	// . "github.com/go-on/wrap-contrib/wraps"
	"net/http"
	"sync"

	"github.com/go-on/lib/internal/fat"
)

type Person struct {
	http.ResponseWriter
	Id        *fat.Field `type:"int" details:"id"`
	FirstName *fat.Field `type:"string"`
	LastName  *fat.Field `type:"string"`
	Vita      *fat.Field `type:"string html"`
}

var (
	PERSON  = fat.Proto(&Person{}).(*Person)
	id      = 0
	persons = map[string]*Person{}
	_       = NewPerson("P<ete>r", "Pan", Elements("A ", I("short"), " vita").String())
	_       = NewPerson("Paul", "S<imo>n", Elements("A ", I("long"), " vita").String())
	mutex   = &sync.Mutex{}
)

func init() {
	Register(PERSON)
}

func NewPerson(firstname, lastname, vita string) (p *Person) {
	id++
	p = fat.New(PERSON, &Person{}).(*Person)
	p.FirstName.Set(firstname)
	p.LastName.Set(lastname)
	p.Vita.Set(vita)
	p.Id.Set(id)
	persons[fmt.Sprintf("%d", id)] = p
	return
}

func FindPerson(wr http.ResponseWriter, req *http.Request) http.ResponseWriter {
	id := req.FormValue(":id")

	// prevent race condition for persons[]
	mutex.Lock()
	defer mutex.Unlock()
	pp, found := persons[id]

	if found {
		// make a copy of the person
		p := fat.New(PERSON, &Person{}).(*Person)
		p.FirstName.Set(pp.FirstName.String())
		p.LastName.Set(pp.LastName.String())
		p.Vita.Set(pp.Vita.String())
		// and wrap it to the responsewriter
		p.ResponseWriter = wr
		return p
	} else {
		wr.WriteHeader(404)
		fmt.Fprint(wr, "404 not found")
		return nil
	}
}

func (p *Person) NameView(wr http.ResponseWriter, req *http.Request) {
	H1(p.FirstName.String(), " ", p.LastName.String()).WriteTo(wr)
}

func ListView(wr http.ResponseWriter, req *http.Request) {
	for _, pers := range persons {
		LI(
			AHref(
				MustUrl(personDetails, pers, "details"),
				pers.FirstName.String(), " ", pers.LastName.String(),
			),
		).WriteTo(wr)
	}
}

func (p *Person) VitaView(wr http.ResponseWriter, req *http.Request) {
	DIV(
		shared.Class("vita"),
		H2("Vita"),
		P(
			shared.HTMLString(p.Vita.String()),
		),
	).WriteTo(wr)
}

var (
	yellow        = shared.Style{"background-color", "yellow"}
	green         = shared.Style{"background-color", "green"}
	personDetails *route.Route
	personList    *route.Route
)

func listLink(rw http.ResponseWriter, req *http.Request) {
	DIV(AHref(personList.MustURL(), "back to list")).WriteTo(rw)
}

func main() {
	personRouter := router.New()
	personDetails = personRouter.GET("/:id",
		wrap.New(
			wrapstesting.Context(FindPerson),
			wrap.Handler(
				Elements(
					listLink,
					DIV(yellow, wrapstesting.HandlerMethod((*Person).NameView)),
					DIV(green, wrapstesting.HandlerMethod((*Person).VitaView)),
				),
			),
		),
	)

	personList = personRouter.GET("/", UL(ListView))

	personRouter.Mount("/person", http.DefaultServeMux)

	handler := compiler.DocHandler(HTML5(HTML(BODY(personRouter))))

	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		panic(err.Error())
	}
}

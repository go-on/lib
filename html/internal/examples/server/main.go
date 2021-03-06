package main

import (
	"fmt"
	"net/http"
	"time"

	. "github.com/go-on/lib/html"
	"github.com/go-on/lib/html/element/compiler"
	"github.com/go-on/lib/types"
	"github.com/go-on/lib/types/placeholder"
	"github.com/metakeule/fmtdate"
)

var (
	_firstname_ = placeholder.New(types.Text("firstname"))
	_lastname_  = placeholder.New(types.Text("lastname"))
	person      = LI(_firstname_, " ", _lastname_).Template("person")
	links       = DIV(AHref("/", "simple"), E_nbsp, AHref("/optimized", "optimized"))
)

func list(wr http.ResponseWriter, req *http.Request) {
	for firstName, lastName := range map[string]string{"Peter": "Tosh", "Paul": "Simon"} {
		person.Replace(
			_firstname_.Set(firstName),
			_lastname_.Set(lastName),
		).WriteTo(wr)
	}
}

func printTime(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprint(wr, fmtdate.Format("ss.00000", time.Now())+" sec")
}

func handlerSimple() http.Handler {
	return HTML5(
		HTML(
			BODY(
				links,
				printTime,
				UL(list),
				printTime,
			),
		),
	)
}

func handlerOptimized() http.Handler {
	return compiler.DocHandler(
		HTML5(
			HTML(
				BODY(
					links,
					printTime,
					UL(list),
					printTime,
				),
			),
		),
	)
}

func main() {
	http.Handle("/", handlerSimple())
	http.Handle("/optimized", handlerOptimized())
	http.ListenAndServe(":8080", nil)
}

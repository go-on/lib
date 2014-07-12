package pseudoroutes

import (
	"github.com/go-on/method"

	"github.com/go-on/router/route"
)

var (
	item   = route.NewRoute("/:ressource/:uuid")
	GET    = item.AddMethod(method.GET)
	PATCH  = item.AddMethod(method.PATCH)
	DELETE = item.AddMethod(method.DELETE)

	list  = route.NewRoute("/:ressource/")
	POST  = list.AddMethod(method.POST)
	INDEX = list.AddMethod(method.GET)
)

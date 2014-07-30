package pseudoroutes

import (
	"github.com/go-on/method"
	"github.com/go-on/router/route"
)

var (
	Id_  = "uuid"
	Item = route.New("/:ressource/:"+Id_, method.GET, method.PATCH, method.DELETE)
	List = route.New("/:ressource/", method.POST, method.GET)
)

func Mount(mountPoint string) {
	route.Mount(mountPoint, Item, List)
}

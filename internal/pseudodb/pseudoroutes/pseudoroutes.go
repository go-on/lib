package pseudoroutes

import (
	"github.com/go-on/method"
	"github.com/go-on/router/route"
)

var (
	Id_        = "uuid"
	Ressource_ = "ressource"
	Item       = route.New("/:"+Ressource_+"/:"+Id_, method.GET, method.PATCH, method.DELETE)
	List       = route.New("/:"+Ressource_+"/", method.POST, method.GET)
)

func Mount(mountPoint string) {
	route.Mount(mountPoint, Item, List)
}

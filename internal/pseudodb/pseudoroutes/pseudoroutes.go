package pseudoroutes

import "github.com/go-on/router/route"

type mountpath string

func (mp mountpath) MountPath() string {
	return string(mp)
}

var (
	item   = route.New("/:ressource/:uuid")
	GET    = item
	PATCH  = item
	DELETE = item

	list  = route.New("/:ressource/")
	POST  = list
	INDEX = list
)

func Mount(mountpoint string) {
	mp := mountpath(mountpoint)
	GET.Router = mp
	POST.Router = mp
	PATCH.Router = mp
	DELETE.Router = mp
	INDEX.Router = mp
}

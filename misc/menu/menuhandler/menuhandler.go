package menuhandler

import (
	"net/http"

	"github.com/go-on/lib/misc/menu"
)

type RequestMenu interface {
	Menu(*http.Request) *menu.Node
}

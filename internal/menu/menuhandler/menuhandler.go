package menuhandler

import (
	"github.com/go-on/lib/internal/menu"
	"net/http"
)

type RequestMenu interface {
	Menu(*http.Request) *menu.Node
}

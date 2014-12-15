package main

import (
	"go/build"
	"net/http"
	"path/filepath"

	"gopkg.in/go-on/cdncache.v1"
	. "gopkg.in/go-on/lib.v2/html"
	"gopkg.in/go-on/lib.v2/internal/flight"
	"gopkg.in/go-on/lib.v2/types"
)

func main() {
	cdn := cdncache.CDN("/cdn/")

	layout := HTML5(
		HTML(
			HEAD(
				//JsSrc("/static/require_js_2_1_14/require.js"),
				JsSrc(cdn(flight.CDN_1_1_4_min)),
				JsSrc("/static/static.js"),
				/*
					SCRIPT(
						`require(['../static'], function(util) {
						    //This function is called when scripts/helper/util.js is loaded.
						    //If util.js calls define(), then this function is not fired until
						    //util's dependencies have loaded, and the util argument will hold
						    //the module value for "helper/util".
						});`,
					),
				*/
			),
			BODY(
				DIV(types.Id("start"), "click me"),
				DIV("hiho"),
			),
		),
	)

	gopath := filepath.SplitList(build.Default.GOPATH)[0]
	static := filepath.Join(gopath, `src/github.com/go-on/lib/internal/flight/1_1_4/flight/example/todo/static`)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(static))))
	http.Handle("/", layout)

	http.ListenAndServe(":8081", nil)

}

/*
<script src="flight.js"></script>
<script>
  var MyComponent = flight.component(function() {
    //...
  });
</script>
*/

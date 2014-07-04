package menuhandler

import (
	"github.com/go-on/lib/internal/menu"
	"github.com/go-on/lib/internal/menu/menuhtml"
	"github.com/go-on/lib/internal/shared"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var m = &menu.Node{
	Edges: []*menu.Node{
		&menu.Node{Leaf: menu.Item("B", "")},
		&menu.Node{
			Edges: []*menu.Node{
				&menu.Node{Leaf: menu.Item("repl", "~replacement")},
				&menu.Node{
					Edges: []*menu.Node{
						&menu.Node{Leaf: menu.Item("AAA", "/a/a/a")},
						&menu.Node{Leaf: menu.Item("AAB", "/a/a/b")},
					},
					Leaf: menu.Item("AA", "/a/a"),
				},
				&menu.Node{
					Edges: []*menu.Node{
						&menu.Node{Leaf: menu.Item("ABA", "/a/b/a")},
					},
					Leaf: menu.Item("AB", "$sub_a"),
				},
			},
			Leaf: menu.Item("A", "/a"),
		},
	},
}

var ul = menuhtml.NewUL(shared.Class("menu-open"), shared.Class("menu-active"), shared.Class("menu-sub"))

func stripWhiteSpace(in string) string {
	return strings.Replace(strings.Replace(strings.Replace(in, "\n", "", -1), "\t", "", -1), " ", "", -1)
}

func TestStaticSub(t *testing.T) {
	staticSub := NewStaticSub(m, 1, 2, ul)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/a/a/b", nil)

	staticSub.ServeHTTP(rec, req)

	expected1 := `
	<ul>
		<li>
			<span>repl</span>
		</li>
		<li class="menu-open">
			<a href="/a/a">AA</a>
			<ul class="menu-sub">
				<li><a href="/a/a/a">AAA</a></li>
				<li class="menu-active"><a href="/a/a/b">AAB</a></li>
			</ul>
		</li>
		<li>
			<span>AB</span>
			<ul class="menu-sub">
				<li><a href="/a/b/a">ABA</a></li>
			</ul>
		</li>
	</ul>
	`
	expected1 = stripWhiteSpace(expected1)

	got1 := stripWhiteSpace(rec.Body.String())

	if got1 != expected1 {
		t.Errorf("got: %#v, expected: %#v", got1, expected1)
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)

	staticSub.ServeHTTP(rec, req)

	expected2 := `
	<ul></ul>
	`
	expected2 = stripWhiteSpace(expected2)

	got2 := stripWhiteSpace(rec.Body.String())

	if got2 != expected2 {
		t.Errorf("got: %#v, expected: %#v", got2, expected2)
	}
}

func TestStatic(t *testing.T) {

	static := NewStatic(m, 3, ul)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/a/a/b", nil)

	static.ServeHTTP(rec, req)

	expected1 := `
<ul>
	<li><span>B</span></li>
	<li class="menu-open">
		<a href="/a">A</a>
		<ul class="menu-sub">
			<li><span>repl</span></li>
			<li class="menu-open">
				<a href="/a/a">AA</a>
				<ul class="menu-sub">
					<li><a href="/a/a/a">AAA</a></li>
					<li class="menu-active"><a href="/a/a/b">AAB</a></li>
				</ul>
			</li>
			<li>
				<span>AB</span>
				<ul class="menu-sub">
					<li>
						<a href="/a/b/a">ABA</a>
					</li>
				</ul>
			</li>
		</ul>
	</li>
</ul>
	`
	expected1 = stripWhiteSpace(expected1)

	got1 := stripWhiteSpace(rec.Body.String())

	if got1 != expected1 {
		t.Errorf("got: %#v, expected: %#v", got1, expected1)
	}
}

type reqMenu struct {
	*menu.Node
}

func (r *reqMenu) Menu(req *http.Request) *menu.Node {
	return r.Node
}

func TestDynamic(t *testing.T) {
	reqM := &reqMenu{m}

	dyn := NewDynamic(reqM, 3, ul)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/a/a/b", nil)

	dyn.ServeHTTP(rec, req)

	expected1 := `
<ul>
	<li><span>B</span></li>
	<li class="menu-open">
		<a href="/a">A</a>
		<ul class="menu-sub">
			<li><span>repl</span></li>
			<li class="menu-open">
				<a href="/a/a">AA</a>
				<ul class="menu-sub">
					<li><a href="/a/a/a">AAA</a></li>
					<li class="menu-active"><a href="/a/a/b">AAB</a></li>
				</ul>
			</li>
			<li>
				<span>AB</span>
				<ul class="menu-sub">
					<li>
						<a href="/a/b/a">ABA</a>
					</li>
				</ul>
			</li>
		</ul>
	</li>
</ul>
	`
	expected1 = stripWhiteSpace(expected1)

	got1 := stripWhiteSpace(rec.Body.String())

	if got1 != expected1 {
		t.Errorf("got: %#v, expected: %#v", got1, expected1)
	}
}

func TestDynamicSub(t *testing.T) {
	reqM := &reqMenu{m}

	dynSub := NewDynamicSub(reqM, 1, 2, ul)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/a/a/b", nil)

	dynSub.ServeHTTP(rec, req)

	expected1 := `
	<ul>
		<li>
			<span>repl</span>
		</li>
		<li class="menu-open">
			<a href="/a/a">AA</a>
			<ul class="menu-sub">
				<li><a href="/a/a/a">AAA</a></li>
				<li class="menu-active"><a href="/a/a/b">AAB</a></li>
			</ul>
		</li>
		<li>
			<span>AB</span>
			<ul class="menu-sub">
				<li><a href="/a/b/a">ABA</a></li>
			</ul>
		</li>
	</ul>
	`
	expected1 = stripWhiteSpace(expected1)

	got1 := stripWhiteSpace(rec.Body.String())

	if got1 != expected1 {
		t.Errorf("got: %#v, expected: %#v", got1, expected1)
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)

	dynSub.ServeHTTP(rec, req)

	expected2 := `
	<ul></ul>
	`
	expected2 = stripWhiteSpace(expected2)

	got2 := stripWhiteSpace(rec.Body.String())

	if got2 != expected2 {
		t.Errorf("got: %#v, expected: %#v", got2, expected2)
	}
}
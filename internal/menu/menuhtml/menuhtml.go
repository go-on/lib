package menuhtml

import (
	. "github.com/go-on/lib/html"
	"github.com/go-on/lib/html/internal/element"
	"github.com/go-on/lib/internal/shared"
	// "github.com/go-on/html/h"
	// "github.com/go-on/html/tag"
	"github.com/go-on/lib/internal/menu"
	"io"
)

type Formatter interface {
	Item(path, text string, isActive, hasChildren bool, depth int) *element.Element
	ClassOpen() shared.Class
	List(depth int, text string) *element.Element
}

func New(f Formatter) menu.WriterTo { return &formatterWriterTo{f} }

type formatterWriterTo struct{ Formatter }

func (d *formatterWriterTo) html(node *menu.Node, depth int, path string, currentDepth int) (elem *element.Element, parentOpened bool) {

	if node.Leaf != nil && currentDepth > 0 {
		p := node.Leaf.Path()
		parentOpened = path == p
		if p == "" || p[0] == '~' || p[0] == '$' {
			p = ""
		}
		elem = d.Item(p, node.Leaf.String(), parentOpened, len(node.Edges) > 0 && currentDepth <= depth, currentDepth)
	}

	//	if currentDepth > depth || len(node.Edges) == 0 {
	if len(node.Edges) == 0 {
		return
	}

	text := ""
	if node.Leaf != nil {
		text = node.Leaf.String()
	}

	list := d.List(currentDepth, text)

	var opened bool

	for _, m := range node.Edges {
		e, shouldOpen := d.html(m, depth, path, currentDepth+1)
		if shouldOpen {
			parentOpened = true
			opened = true
		}
		list.Add(e)
	}

	if elem != nil {
		if opened {
			elem.Add(d.ClassOpen())
		}

		if currentDepth <= depth {
			elem.Add(list)
		}
		return
	}

	if currentDepth > depth {
		return
	}

	elem = list
	return
}

// WriterTo provides a method to fullfill the menu.WriterTo interface
func (d *formatterWriterTo) WriterTo(menu *menu.Node, depth int, path string) io.WriterTo {
	if menu == nil {
		return d.List(0, "")
	}
	elem, _ := d.html(menu, depth, path, 0)
	return elem
}

type ul struct{ classOpen, classActive, classSub shared.Class }

func NewUL(classOpen, classActive, classSub shared.Class) menu.WriterTo {
	return New(&ul{classOpen, classActive, classSub})
}

func (u *ul) Item(path, text string, isActive, hasChildren bool, depth int) (li *element.Element) {
	li = LI()
	if path == "" {
		li.Add(SPAN(text))
		return
	}
	li.Add(AHref(path, text))
	if isActive {
		li.Add(u.classActive)
	}
	return
}

func (u *ul) ClassOpen() shared.Class { return u.classOpen }

func (u *ul) List(depth int, text string) (ul *element.Element) {
	ul = UL()
	if depth > 0 {
		ul.Add(u.classSub)
	}
	return
}

/*
var Ul = &Default_{
	ActiveClass:  shared.Class("menu-active"),
	OpenClass:    shared.Class("menu-open"),
	ListElement:  UL,
	EntryElement: LI,
}
*/

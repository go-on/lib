package element

import (
	"bytes"
	"fmt"
	"github.com/go-on/builtin"
	"github.com/go-on/lib/internal/shared"
	"github.com/go-on/lib/internal/shared/placeholder"
	"io"
)

type htmlerstring struct {
	shared.HTMLer
}

func (h *htmlerstring) String() string {
	return h.HTMLer.HTML()
}

// prepare the id attribute for output
func AttrsString(ø *Element) (res string) {
	var buffer bytes.Buffer
	if !Is(ø, IdForbidden) && ø.Id != "" {
		buffer.WriteString(" " + shared.Attribute{"id", string(ø.Id)}.String())
	}
	if !Is(ø, ClassForbidden) && len(ø.Classes) > 0 {
		buffer.WriteString(" " + shared.Attribute{"class", classAttrString(ø.Classes)}.String())
	}
	if !Is(ø, Invisible) && len(ø.Styles) > 0 {
		buffer.WriteString(" " + shared.Attribute{"style", css(ø.Styles)}.String())
	}

	for _, v := range ø.Attributes {
		// ignore id and class and style attributes, since they should be set via Id and Classes properties or Add
		if v.Key == "id" || v.Key == "class" || v.Key == "style" {
			continue
		}
		buffer.WriteString(" " + v.String())
	}
	return buffer.String()
}

func classAttrString(classes []shared.Class) (s string) {
	var buffer bytes.Buffer
	for _, cl := range classes {
		buffer.WriteString(" " + string(cl))
	}
	return buffer.String()[1:]
}

func css(fds []shared.Style) (s string) {
	var buffer bytes.Buffer
	for _, v := range fds {
		buffer.WriteString(v.String())
	}
	return buffer.String()
}

// adds css properties to the style attribute, same keys are overwritten
func addStyles(ø *Element, v []shared.Style) {
	ø.Styles = append(ø.Styles, v...)
}

func addStyle(ø *Element, v shared.Style) {
	ø.Styles = append(ø.Styles, v)
}

func addText(ø *Element, text string) {
	s := shared.Text(text)

	if !Is(ø, WithoutEscaping) {
		s = shared.Text(shared.EscapeHTML(text))
	}

	if Is(ø, JavascriptSpecialEscaping) {
		s = shared.Text(jsSpecialEscape(text))
	}
	ø.Children = append(ø.Children, s)
}

func addChild(ø *Element, child builtin.Stringer) {
	ø.Children = append(ø.Children, child)
}

func addClass(ø *Element, class shared.Class) {
	ø.Classes = append(ø.Classes, class)
}

func addClasses(ø *Element, classes []shared.Class) {
	ø.Classes = append(ø.Classes, classes...)
}

func addPlaceholder(ø *Element, v placeholder.Placeholder) {
	switch tp := v.Type().(type) {
	case shared.Descr:
		ø.Add(shared.Descr(v.String()))
	case shared.Id:
		ø.Add(shared.Id(v.String()))
	case shared.Class:
		ø.Add(shared.Class(v.String()))
	case shared.HTMLString:
		ø.Add(shared.HTMLString(v.String()))
	case shared.Text:
		ø.Add(shared.Text(v.String()))
	case shared.Attribute:
		ø.Add(shared.Attribute{tp.Key, v.String()})
	case shared.Tag:
		ø.tag = v.String()
	case shared.Style:
		ø.Add(shared.Style{tp.Property, v.String()})
	default:
		ø.Add(shared.Text(v.String()))
	}
}

func InnerHtmlBf(ø *Element, w io.Writer) (num int64, err error) {
	var n int64
	var nn int
	for _, in := range ø.Children {
		switch ch := in.(type) {
		case *Element:
			n, err = ch.WriteTo(w)
		case io.WriterTo:
			n, err = ch.WriteTo(w)
		default:
			nn, err = fmt.Fprint(w, in.String())
			n = int64(nn)
		}
		if err != nil {
			return
		}
		num += n
	}
	return
}

func InnerHtml(ø *Element) (res string) {
	var buffer bytes.Buffer
	InnerHtmlBf(ø, &buffer)
	return buffer.String()
}

/*
func InnerHtml(ø *Element) (res string) {
	var buffer bytes.Buffer
	for _, in := range ø.Children {
		buffer.WriteString(in.String())
	}
	return buffer.String()
}
*/

// Is checks if a given flag is set, e.g.
//
// 	Is(Inline)
//
// checks for the Inline flag
func Is(ø *Element, f flag) bool { return ø.flags&f != 0 }

// Pre writes properties of an element to w before the inner elements are written to w
func Pre(ø *Element, w io.Writer) (num int, err error) {
	var n int
	if ø.Descr != "" {
		n, err = fmt.Fprintf(w, "<!-- Begin: %s -->", ø.Descr)
		if err != nil {
			return
		}
		num += n
	}
	if Is(ø, WithoutDecoration) {
		return
	}
	if Is(ø, SelfClosing) {
		n, err = fmt.Fprintf(w, "<%s%s />", ø.Tag(), AttrsString(ø))
		if err != nil {
			return
		}
		num += n
		return
	}
	n, err = fmt.Fprintf(w, "<%s%s>", ø.Tag(), AttrsString(ø))
	num += n
	return
}

// Post writes properties of an element to w after the inner elements are written to w
func Post(ø *Element, w io.Writer) (num int, err error) {
	var n int
	if Is(ø, WithoutDecoration) || Is(ø, SelfClosing) {
		if ø.Descr != "" {
			n, err = fmt.Fprintf(w, "<!-- End: %s -->", ø.Descr)
			if err != nil {
				return
			}
			num += n
		}
		return
	}
	n, err = fmt.Fprintf(w, "</%s>", ø.Tag())
	if err != nil {
		return
	}
	num += n
	if ø.Descr != "" {
		n, err = fmt.Fprintf(w, "<!-- End: %s -->", ø.Descr)
		if err != nil {
			return
		}
		num += n
	}
	return
}

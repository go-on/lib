package css

import (

	// "fmt"

	. "github.com/go-on/lib/html"
	. "github.com/go-on/lib/misc/selector"
	. "github.com/go-on/lib/types"
	// . "github.com/go-on/lib/html/tag"
	"testing"
)

func err(t *testing.T, msg string, is interface{}, shouldbe interface{}) {
	t.Errorf(msg+": is %v, should be %v\n", is, shouldbe)
	//panic(fmt.Sprintf(msg+": is %v, should be %v\n", is, shouldbe))
}

func TestSimpleSelector(t *testing.T) {
	var bestClass = Class("best")
	var mainId = Id("main")
	var other = Class("other")

	css := Css(Each(DIV(mainId, bestClass), P(other)))

	if s := css.Selector.Selector(); s != "div#main.best,\np.other" {
		err(t, "incorrect selector", s, "div#main.best,\np.other")
	}
}

func TestRule(t *testing.T) {
	var expectedString string
	font, _ := Rule(Class("default"), Style{"font-size", "12px"})
	bg, _ := Rule(Class("bg"), Style{"background-color", "black"})
	bg = bg.Compose(font)
	special := Class("special")

	css, _ := Rule(
		Context(Id("myid"), UL(special), OL(special), LI(special)),
		Style{"color", "yellow"},
		Style{"width", "200"}, Style{"height", "300"},
		// Styles(Style{"width", "200"}, Style{"height", "300"}),
		Comment("no comment"))

	css = css.Compose(bg)

	expectedString = "#myid ol.special,\n#myid li.special,\n#myid ul.special"

	if css.Selector.Selector() != expectedString {
		err(t, "incorrect selector1", css.Selector.Selector(), expectedString)
	}

	// fmt.Printf("styles: %v", css.Styles)
	styles := css.Styles

	if len(css.Styles) == 0 {
		t.Errorf("no  styles")
		return
	}

	expectedString = "color:yellow;"
	if styles[0].Style() != expectedString {
		err(t, "incorrect style", styles[0].Style(), expectedString)
	}

	expectedString = "width:200;"
	if styles[1].Style() != expectedString {
		err(t, "incorrect style", styles[1].Style(), expectedString)
	}

	expectedString = "height:300;"
	if styles[2].Style() != expectedString {
		err(t, "incorrect style", styles[2].Style(), expectedString)
	}

	expectedString = "background-color:black;"
	if styles[3].Style() != expectedString {
		err(t, "incorrect style", styles[3].Style(), expectedString)
	}

	expectedString = "font-size:12px;"
	if styles[4].Style() != expectedString {
		err(t, "incorrect style", styles[4].Style(), expectedString)
	}

}

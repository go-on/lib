package html

import (
	"github.com/go-on/lib/internal/shared"
)

const (
	Hidden   = "hidden"
	Checkbox = "checkbox"
)

func Rel_(relation string) shared.Attribute { return shared.Attribute{"rel", relation} }
func Href_(url string) shared.Attribute     { return shared.Attribute{"href", url} }
func Type_(ty string) shared.Attribute      { return shared.Attribute{"type", ty} }
func Media_(md string) shared.Attribute     { return shared.Attribute{"media", md} }
func Src_(src string) shared.Attribute      { return shared.Attribute{"src", src} }
func Name_(n string) shared.Attribute       { return shared.Attribute{"name", n} }
func Content_(n string) shared.Attribute    { return shared.Attribute{"content", n} }
func Charset_(c string) shared.Attribute    { return shared.Attribute{"charset", c} }
func CharsetUtf8_() shared.Attribute        { return shared.Attribute{"charset", "utf-8"} }
func Lang_(l string) shared.Attribute       { return shared.Attribute{"lang", l} }
func Role_(l string) shared.Attribute       { return shared.Attribute{"role", l} }
func Value_(v string) shared.Attribute      { return shared.Attribute{"value", v} }
func Alt_(alttext string) shared.Attribute  { return shared.Attribute{"alt", alttext} }
func Title_(text string) shared.Attribute   { return shared.Attribute{"title", text} }
func Method_(m string) shared.Attribute     { return shared.Attribute{"method", m} }
func Action_(a string) shared.Attribute     { return shared.Attribute{"action", a} }
func For_(f string) shared.Attribute        { return shared.Attribute{"for", f} }
func Width_(f string) shared.Attribute      { return shared.Attribute{"width", f} }
func Height_(f string) shared.Attribute     { return shared.Attribute{"height", f} }
func OnSubmit_(js string) shared.Attribute  { return shared.Attribute{"onsubmit", "javascript:" + js} }
func OnClick_(js string) shared.Attribute   { return shared.Attribute{"onclick", "javascript:" + js} }
func Enctype_(f string) shared.Attribute    { return shared.Attribute{"enctype", f} }
func Target_(f string) shared.Attribute     { return shared.Attribute{"target", f} }
func DataToggle_(f string) shared.Attribute { return shared.Attribute{"data-toggle", f} }
func DataTarget_(f string) shared.Attribute { return shared.Attribute{"data-target", f} }
func DataId_(f string) shared.Attribute     { return shared.Attribute{"data-id", f} }
func Style_(f string) shared.Attribute      { return shared.Attribute{"style", f} }

// RDFa
func About_(a string) shared.Attribute    { return shared.Attribute{"about", a} }
func TypeOf_(a string) shared.Attribute   { return shared.Attribute{"typeof", a} }
func Property_(a string) shared.Attribute { return shared.Attribute{"property", a} }

// vars
var Checked_ = shared.Attribute{"checked", "checked"}
var Selected_ = shared.Attribute{"selected", "selected"}
var Disabled_ = shared.Attribute{"disabled", "disabled"}
var Required_ = shared.Attribute{"required", "required"}
var MultiPart_ = Enctype_("multipart/form-data")
var TargetBlank_ = Target_("_blank")

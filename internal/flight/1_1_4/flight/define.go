package flight

import (
	"github.com/gopherjs/gopherjs/js"
)

type Component struct {
	obj js.Object
}

func (c *Component) Define(mixins ...interface{}) js.Object {
	return c.obj.Invoke(mixins...)
}

func (c *Component) TeardownAll() {
	c.obj.Get("teardownAll").Invoke()
}

type Instance struct {
	obj js.Object
}

func (i *Instance) AttachTo(selector interface{}) {
	i.obj.Get("attachTo").Invoke(selector)
}

func (i *Instance) AttachToWithOptions(selector interface{}, options interface{}) {
	i.obj.Get("attachTo").Invoke(selector, options)
}

func (i *Instance) TeardownAll() {
	i.obj.Get("teardownAll").Invoke()
}

type Require func(string) js.Object

func (r Require) Component(s string) *Component {
	obj := r("flight/lib/component")
	return &Component{obj}
}

func (r Require) Instance(path string) *Instance {
	obj := r(path)
	return &Instance{obj}
}

func Define(fn func(r Require) js.Object) {

	js.Global.Get("flight").Get("define").Invoke(func(o js.Object) js.Object {
		var rr Require = func(s string) js.Object {
			return o.Invoke(s)
		}

		return fn(rr)
	})
}

//define(function(require) {

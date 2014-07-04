package promise

import (
	"github.com/gopherjs/gopherjs/js"
	"testing"
)

type jsObject struct {
	js.Object
	attrs map[string]interface{}
}

func newJsObject() *jsObject {
	return &jsObject{attrs: map[string]interface{}{}}
}

func (jso *jsObject) Get(key string) js.Object {
	return jso
}

func (jso *jsObject) Set(key string, i interface{}) {
	jso.attrs[key] = i
}

func (jso *jsObject) Invoke(i ...interface{}) js.Object {
	i[0].(func(js.Object) js.Object)(jso)
	i[1].(func(js.Object) js.Object)(jso)
	// Invoke(onFullFilled, onRejected)
	return jso
}

func (jso *jsObject) IsUndefined() bool {
	return false
}

func (jso *jsObject) IsNull() bool {
	return false
}

func onFullFilled(o js.Object) js.Object {
	o.Set("onFullFilled called", o)
	return o
}

func onRejected(o js.Object) js.Object {
	o.Set("onRejected called", o)
	return o
}

func TestPromiseNew(t *testing.T) {
	if New(nil) != nil {
		t.Errorf("nil js.Object should return nil, but does not")
	}
}

func TestPromiseThen(t *testing.T) {
	o := newJsObject()
	promise1 := New(o)
	_ = promise1.Then(onFullFilled, onRejected)

	_, fullfilledCalled := o.attrs["onFullFilled called"]

	if !fullfilledCalled {
		t.Errorf("onFullFilled not called")
	}

	_, rejectedCalled := o.attrs["onRejected called"]

	if !rejectedCalled {
		t.Errorf("onRejected not called")
	}
}

package flight

import (
	"github.com/gopherjs/gopherjs/js"
)

func Attributes(this js.Object, obj map[string]interface{}) {
	this.Get("attributes").Invoke(obj)
}

func AttachTo(comp js.Object, selector string, obj map[string]interface{}) {
	comp.Get("attachTo").Invoke(selector, obj)
}

func ReadAttr(this js.Object, attr string) string {
	return this.Get("attr").Get(attr).Str()
}

// attr must contain a selector
func SelectByAttr(this js.Object, attr string) js.Object {
	return this.Get("select").Invoke(attr)
}

func OnTrigger(this js.Object, eventType string, triggerEvent string) {
	this.Get("on").Invoke(eventType, triggerEvent)
}

func OnFunc(this js.Object, eventType string, fn func(event js.Object, data js.Object)) {
	this.Get("on").Invoke(eventType, fn)
}

func OnMap(this js.Object, eventType string, m map[string]interface{}) {
	this.Get("on").Invoke(eventType, m)
}

func SelectorOnTrigger(this js.Object, selector interface{}, eventType string, triggerEvent string) {
	this.Get("on").Invoke(selector, eventType, triggerEvent)
}

func SelectorOnFunc(this js.Object, selector interface{}, eventType string, fn func(event js.Object, data js.Object)) {
	this.Get("on").Invoke(selector, eventType, fn)
}

func SelectorOnMap(this js.Object, selector interface{}, eventType string, m map[string]interface{}) {
	this.Get("on").Invoke(selector, eventType, m)
}

func Off(this js.Object, eventType interface{}) {
	this.Get("off").Invoke(eventType)
}

func OffFunc(this js.Object, eventType interface{}, fn func(event js.Object, data js.Object)) {
	this.Get("off").Invoke(eventType, fn)
}

func OffMap(this js.Object, eventType interface{}, m map[string]interface{}) {
	this.Get("off").Invoke(eventType, m)
}

func SelectorOff(this js.Object, selector interface{}, eventType interface{}) {
	this.Get("off").Invoke(selector, eventType)
}

func SelectorOffFunc(this js.Object, selector interface{}, eventType interface{}, fn func(event js.Object, data js.Object)) {
	this.Get("off").Invoke(selector, eventType, fn)
}

func SelectorOffMap(this js.Object, selector interface{}, eventType interface{}, m map[string]interface{}) {
	this.Get("off").Invoke(selector, eventType, m)
}

func Trigger(this js.Object, eventType interface{}) {
	this.Get("trigger").Invoke(eventType)
}

func TriggerWith(this js.Object, eventType interface{}, eventPayload interface{}) {
	this.Get("trigger").Invoke(eventType, eventPayload)
}

func SelectorTrigger(this js.Object, selector interface{}, eventType interface{}) {
	this.Get("trigger").Invoke(selector, eventType)
}

func SelectorTriggerWith(this js.Object, selector interface{}, eventType interface{}, eventPayload interface{}) {
	this.Get("trigger").Invoke(selector, eventType, eventPayload)
}

func TearDown(this js.Object) {
	this.Get("teardown").Invoke()
}

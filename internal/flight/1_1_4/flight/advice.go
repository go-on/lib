package flight

import (
	"github.com/gopherjs/gopherjs/js"
)

func Before(this js.Object, existingfnName string, customFunc func()) {
	this.Get("before").Invoke(existingfnName, customFunc)
}

func After(this js.Object, existingfnName string, customFunc func()) {
	this.Get("after").Invoke(existingfnName, customFunc)
}

func Around(this js.Object, existingfnName string, customFunc func()) {
	this.Get("around").Invoke(existingfnName, customFunc)
}

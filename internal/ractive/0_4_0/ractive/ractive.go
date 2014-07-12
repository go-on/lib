package ractive

import (
	"github.com/go-on/lib/internal/promise"
	"github.com/gopherjs/gopherjs/js"
)

var ractive js.Object

type InitOptions struct {
	El   interface{}
	Data interface{} // js.Object
	//	Partials    js.Object
	Partials    map[string]string
	Transitions js.Object
	Complete    js.Object
	Adaptors    js.Object
	Computed    map[string]interface{}

	Magic              bool
	Debug              bool
	Lazy               bool
	Append             bool
	PreserveWhiteSpace bool
	Sanitize           bool

	NotModifyArrays bool
	NotTwoWay       bool

	// forceExplicit struct{}
}

func New(template string, options *InitOptions) *Ractive {
	obj := js.Global.Get("Object").New()
	obj.Set("template", template)
	if options != nil {
		if options.El != "" {
			obj.Set("el", options.El)
		}

		if options.Data != nil {
			obj.Set("data", options.Data)
		}

		if options.Partials != nil {
			obj.Set("partials", options.Partials)
		}

		if options.Transitions != nil {
			obj.Set("transitions", options.Transitions)
		}

		if options.Complete != nil {
			obj.Set("complete", options.Complete)
		}

		if options.Adaptors != nil {
			obj.Set("adaptors", options.Adaptors)
		}

		if options.Computed != nil {
			obj.Set("computed", options.Computed)
		}

		if options.Magic {
			obj.Set("magic", true)
		}

		if options.Debug {
			obj.Set("debug", true)
		}

		if options.Lazy {
			obj.Set("lazy", true)
		}

		if options.Append {
			obj.Set("append", true)
		}

		if options.PreserveWhiteSpace {
			obj.Set("preserveWhitespace", true)
		}

		if options.Sanitize {
			obj.Set("sanitize", true)
		}

		if options.NotModifyArrays {
			obj.Set("modifyArrays", false)
		}

		if options.NotTwoWay {
			obj.Set("twoway", false)
		}
	}
	return &Ractive{ractive.New(obj)}
}

func Defaults() js.Object {
	return ractive.Get("defaults")
}

func Extend(m map[string]interface{}) js.Object {
	return ractive.Get("extend").Invoke(m)
}

func Parse(template string, options *struct {
	PreserveWhitespace bool
	Sanitize           interface{}
}) js.Object {
	if options == nil {
		return ractive.Get("parse").Invoke(template)
	}
	if options.Sanitize == nil {
		options.Sanitize = false
	}
	return ractive.Get("parse").Invoke(template, options)
}

type Listener struct{ obj js.Object }

func (l *Listener) Cancel() { l.obj.Get("cancel").Invoke() }

type Ractive struct{ obj js.Object }

func (r *Ractive) Nodes() js.Object {
	return r.obj.Get("nodes")
}

func (r *Ractive) Transitions() js.Object {
	return r.obj.Get("transitions")
}

func (r *Ractive) Partials() js.Object {
	return r.obj.Get("partials")
}

func (r *Ractive) Add(keypath string) promise.Promise {
	p := r.obj.Get("add").Invoke(keypath)
	return promise.New(p)
}

func (r *Ractive) AddN(keypath string, number int) promise.Promise {
	p := r.obj.Get("add").Invoke(keypath, number)
	return promise.New(p)
}

type AnimationOptions struct {
	Duration int
	Easing   interface{}
	Step     func(t float32, value js.Object)
	Complete func(t float32, value js.Object)
	// forceExplicit struct{}
}

func (a *AnimationOptions) normalizeDefaults() {
	if a.Duration == 0 {
		a.Duration = 400
	}

	if a.Easing == nil {
		a.Easing = "linear"
	}
}

func (r *Ractive) Animate(keypath string, value js.Object, options *AnimationOptions) promise.Promise {
	if options == nil {
		p := r.obj.Get("animate").Invoke(keypath, value)
		return promise.New(p)
	}
	options.normalizeDefaults()
	p := r.obj.Get("animate").Invoke(keypath, value, options)
	return promise.New(p)
}

func (r *Ractive) AnimateMap(keypathValueMap map[string]js.Object, options *AnimationOptions) promise.Promise {
	if options == nil {
		p := r.obj.Get("animate").Invoke(keypathValueMap)
		return promise.New(p)
	}
	options.normalizeDefaults()
	p := r.obj.Get("animate").Invoke(keypathValueMap, options)
	return promise.New(p)
}

func (r *Ractive) Detach() js.Object {
	return r.obj.Get("detach").Invoke()
}

func (r *Ractive) Find(selector string) js.Object {
	return r.obj.Get("find").Invoke(selector)
}

// name is empty string, the first component instance of any kind is returned
func (r *Ractive) FindComponent(name string) js.Object {
	if name == "" {
		return r.obj.Get("findComponent").Invoke()
	}
	return r.obj.Get("findComponent").Invoke(name)
}

type FindOptions struct {
	Live bool
	// forceExplicit struct{}
}

func (r *Ractive) FindAll(selector string, options *FindOptions) js.Object {
	if options == nil {
		return r.obj.Get("findAll").Invoke(selector)
	}
	return r.obj.Get("findAll").Invoke(selector, options)
}

func (r *Ractive) FindAllComponents(name string, options *FindOptions) js.Object {
	if options == nil {
		return r.obj.Get("findAllComponents").Invoke(name)
	}
	return r.obj.Get("findAllComponents").Invoke(name, options)
}

func (r *Ractive) Fire(eventName string, arg interface{}) {
	r.obj.Get("fire").Invoke(eventName, arg)
}

func (r *Ractive) Get(key string) js.Object {
	return r.obj.Get("get").Invoke(key)
}

func (r *Ractive) Insert(target js.Object) {
	r.obj.Get("insert").Invoke(target)
}

func (r *Ractive) InsertAnchor(target js.Object, anchor js.Object) {
	r.obj.Get("insert").Invoke(target, anchor)
}

type MergeOptions struct {
	Compare js.Object
	// forceExplicit struct{}
}

func (r *Ractive) Merge(keypath string, value []js.Object, options *MergeOptions) promise.Promise {
	if options == nil {
		p := r.obj.Get("merge").Invoke(keypath, value)
		return promise.New(p)
	}
	p := r.obj.Get("merge").Invoke(keypath, value, options)
	return promise.New(p)
}

func (r *Ractive) Observe(placeholder string, fn func(newValue, oldValue js.Object)) (l *Listener) {
	o := r.obj.Get("observe").Invoke(placeholder, fn)
	return &Listener{o}
}

func (r *Ractive) Off(event string, fn func(js.Object)) {
	r.obj.Get("off").Invoke(event, fn)
}

func (r *Ractive) OffAll(event string) {
	r.obj.Get("off").Invoke(event)
}

func (r *Ractive) On(event string, fn func(o js.Object, eventData js.Object)) (l *Listener) {
	o := r.obj.Get("on").Invoke(event, fn)
	return &Listener{o}
}

func (r *Ractive) OnFire(event string, triggerEvent string) (l *Listener) {
	o := r.obj.Get("on").Invoke(event, func(ev js.Object, eventData js.Object) {
		r.Fire(triggerEvent, eventData)
	})
	return &Listener{o}
}

func (r *Ractive) OnMap(m map[string]func(o js.Object, eventData js.Object)) (l *Listener) {
	o := r.obj.Get("on").Invoke(m)
	return &Listener{o}
}

func (r *Ractive) Reset(data js.Object) promise.Promise {
	if data == nil {
		p := r.obj.Get("reset").Invoke()
		return promise.New(p)
	}
	p := r.obj.Get("reset").Invoke(data)
	return promise.New(p)
}

func (r *Ractive) Set(key string, value interface{}) promise.Promise {
	p := r.obj.Get("set").Invoke(key, value)
	return promise.New(p)
}

func (r *Ractive) SetMap(m map[string]interface{}) promise.Promise {
	p := r.obj.Get("set").Invoke(m)
	return promise.New(p)
}

func (r *Ractive) SetObject(key string, value js.Object) promise.Promise {
	p := r.obj.Get("set").Invoke(key, value)
	return promise.New(p)
}

func (r *Ractive) Substract(keypath string) promise.Promise {
	p := r.obj.Get("substract").Invoke(keypath)
	return promise.New(p)
}

func (r *Ractive) SubstractN(keypath string, number int) promise.Promise {
	p := r.obj.Get("substract").Invoke(keypath, number)
	return promise.New(p)
}

func (r *Ractive) TearDown() promise.Promise {
	p := r.obj.Get("teardown").Invoke()
	return promise.New(p)
}

func (r *Ractive) Toggle(keypath string) promise.Promise {
	p := r.obj.Get("toggle").Invoke(keypath)
	return promise.New(p)
}

func (r *Ractive) ToHTML() string {
	return r.obj.Get("toHTML").Invoke().Str()
}

func (r *Ractive) Update(key string) promise.Promise {
	p := r.obj.Get("update").Invoke(key)
	return promise.New(p)
}

func (r *Ractive) UpdateAll() promise.Promise {
	p := r.obj.Get("update").Invoke()
	return promise.New(p)
}

func (r *Ractive) UpdateModel(keypath string) {
	r.obj.Get("updateModel").Invoke(keypath, false)
}

func (r *Ractive) UpdateModelCascade(keypath string) {
	r.obj.Get("updateModel").Invoke(keypath, true)
}

func (r *Ractive) UpdateModelAll() {
	r.obj.Get("updateModel").Invoke()
}

func setupRactive() {
	ractive = js.Global.Get("Ractive")
}

func checkLoaded() {
	if ractive.IsUndefined() {
		panic("ractive not loaded")
	}
}

func checkVersion() {
	if ractive.Get("VERSION").Str() != "0.4.0" {
		panic("ractive v0.4.0 needed")
	}
}

func init() {
	if js.Global != nil {
		setupRactive()
		checkLoaded()
		checkVersion()
	}
}

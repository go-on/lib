package promise

import (
	"github.com/gopherjs/gopherjs/js"
)

// Promise is an A+ promise as specified in http://promises-aplus.github.io/promises-spec/
type Promise interface {
	Then(
		// optional (may be <nil>) onFulFilled function that may receive and may return an object
		onFullFilled func(js.Object) js.Object,
		// optional (may be <nil>) onRejected function that may receive and may return an object
		onRejected func(js.Object) js.Object,
	) Promise // Then must return a promise
}

type promise struct {
	js.Object
}

// NewPromise creates a Promise from a js.Object that is considered to be an A+ promise
func New(o js.Object) Promise {
	if o == nil || o.IsUndefined() || o.IsNull() {
		return nil
	}
	return &promise{o}
}

func (p *promise) Then(
	onFullFilled func(js.Object) js.Object,
	onRejected func(js.Object) js.Object,
) Promise {
	o := p.Object.Get("then").Invoke(onFullFilled, onRejected)
	return New(o)
}

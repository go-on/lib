package bacon

import (
	"github.com/gopherjs/gopherjs/js"
)

var Bacon js.Object

func init() {
	if js.Global != nil {
		Bacon = js.Global.Get("Bacon")
		if Bacon.IsUndefined() {
			panic("Bacon not loaded")
		}

		if Bacon.Get("version").Str() != "0.7.12" {
			panic("Bacon v0.7.12 needed")
		}
	}
}

type Bus struct {
	Object js.Object
}

func (b *Bus) Push(x js.Object)      { b.Object.Get("push").Invoke(x) }
func (b *Bus) End()                  { b.Object.Get("end").Invoke() }
func (b *Bus) Error(err error)       { b.Object.Get("error").Invoke(err) }
func (b *Bus) Plug(stream js.Object) { b.Object.Get("plug").Invoke(stream) }

func NewBus() *Bus {
	return &Bus{Bacon.Get("Bus").New()}
}

type EventStream struct {
	Object js.Object
}

func (s *EventStream) Subscribe(f func(js.Object) js.Object) func() {
	unsub := s.Object.Get("subscribe").Invoke(f)
	return func() { unsub.Invoke() }
}

func (s *EventStream) OnValue(f func(js.Object) js.Object) func() {
	unsub := s.Object.Get("onValue").Invoke(f)
	return func() { unsub.Invoke() }
}

/*
stream.skipDuplicates(isEqual) drops consecutive equal elements. So, from [1, 2, 2, 1] you'd get [1, 2, 1]. Uses the === operator for equality checking by default. If the isEqual argument is supplied, checks by calling isEqual(oldValue, newValue). For instance, to do a deep comparison,you can use the isEqual function from underscore.js like stream.skipDuplicates(_.isEqual).

stream.concat(otherStream) concatenates two streams into one stream so that it will deliver events from stream until it ends and then deliver events from otherStream. This means too that events from stream2, occurring before the end of stream will not be included in the result stream.

stream.merge(otherStream) merges two streams into one stream that delivers events from both

stream.holdWhen(valve) pauses and buffers the event stream if last event in valve is truthy. All buffered events are released when valve becomes falsy.


stream.startWith(value) adds a starting value to the stream, i.e. concats a single-element stream contains value with this stream.

stream.skipWhile(f) skips elements while given predicate function holds true. The Function Construction rules below apply here.

stream.skipWhile(property) skips elements while the value of the given Property is true.

stream.skipUntil(stream2) skips elements from stream until a Next event appears in stream2. In other words, starts delivering values from stream after first event appears in stream2.

stream.bufferWithTime(delay) buffers stream events with given delay. The buffer is flushed at most once in the given delay. So, if your input contains [1,2,3,4,5,6,7], then you might get two events containing [1,2,3,4] and [5,6,7] respectively, given that the flush occurs between numbers 4 and 5.

stream.bufferWithTime(f) works with a given "defer-function" instead of a delay. Here's a simple example, which is equivalent to stream.bufferWithTime(10):

stream.bufferWithTime(function(f) { setTimeout(f, 10) })

stream.bufferWithCount(count) buffers stream events with given count. The buffer is flushed when it contains the given number of elements. So, if you buffer a stream of [1, 2, 3, 4, 5] with count 2, you'll get output events with values [1, 2], [3, 4] and [5].

stream.bufferWithTimeOrCount(delay, count) buffers stream events and flushes when either the buffer contains the given number elements or the given amount of milliseconds has passed since last buffered event.

stream.toProperty() creates a Property based on the EventStream. Without arguments, you'll get a Property without an initial value. The Property will get its first actual value from the stream, and after that it'll always have a current value.

stream.toProperty(initialValue) creates a Property based on the EventStream with the given initial value that will be used as the current value until the first value comes from the stream.
*/

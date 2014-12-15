package main

import (
	// "gopkg.in/go-on/lib.v2/internal/flight/1_1_4/flight"
	"github.com/gopherjs/gopherjs/js"
)

/*
flight.component(
*/

func main() {
	/*
		js.Global.Get("require").Invoke(
			"../flight.js",
			func() {
				flight.Define(func(r flight.Require) js.Object {
					c := r.Component("testcomp")
					println("my component A")
					fn := func() {
						println("my component")
					}

					return c.Define(fn)
				})

			})
	*/
	///*

	var inbox = func() {
		//println("inbox invoked")
		js.This.Set("doSomething", func() {
			println("did something")
		})

		js.This.Get("after").Invoke("initialize", func() {
			js.This.Get("on").Invoke("click", js.This.Get("doSomething"))
		})

		println(js.This.Get("attachTo")) //.Invoke("#start")
	}

	/*
	   function inbox() {
	      // define custom functions here
	      this.doSomething = function() {
	        //...
	      }

	      this.doSomethingElse = function() {
	        //...
	      }

	      // now initialize the component
	      this.after('initialize', function() {
	        this.on('click', this.doSomething);
	        this.on('mouseover', this.doSomethingElse);
	      });
	    }
	*/

	var o = js.Global.Get("flight").Get("component").Invoke(inbox)

	o.Invoke()
	// println(js.Global.Get("flight"))

	//println(js.Global.Get("flight").Get("component").New())

	// println(j)

	/*
		.Invoke(func(o js.Object) {
			println(o)
		})
	*/
	/*
		flight.Define(func(r flight.Require) js.Object {
			c := r.Component("testcomp")
			println("my component A")
			fn := func() {
				println("my component")
			}

			return c.Define(fn)
		})
	*/

	/*
		comp := js.Global.Get("flight").Get("component").Invoke(func() func() {
			js.Global.Get("window").Call("alert", "hihohuhiui")
			return func() {
				// js.Global.Get("window").Call("alert", "hihohu")
			}
		})

		comp.Get("attachTo").Invoke("#start")
	*/
	//*/
}

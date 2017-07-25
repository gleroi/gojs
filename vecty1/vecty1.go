// main provides the smallest gopherjs/vecty examples
package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

type Main struct {
	vecty.Core
	counter int
}

func log(l string) {
	js.Global.Get("console").Call("log", l)
}

func (m *Main) Add(e *vecty.Event) {
	log("add")
	m.counter++
	vecty.Rerender(m)
}

func (m *Main) Remove(e *vecty.Event) {
	log("remove")
	m.counter--
	vecty.Rerender(m)
}

func (m *Main) Render() *vecty.HTML {
	return elem.Body(elem.Div(
		elem.Paragraph(
			vecty.Text(fmt.Sprintf("Counter : %d", m.counter)),
		),
		elem.Button(vecty.Text("ADD"), event.Click(m.Add)),
		elem.Button(vecty.Text("REMOVE"), event.Click(m.Remove)),
	))
}

func main() {
	vecty.SetTitle("Go JS!")
	vecty.RenderBody(&Main{})
}

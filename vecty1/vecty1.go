// main provides the smallest gopherjs/vecty examples
package main

import (
	"fmt"
	"log"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type Main struct {
	vecty.Core
	counter int
}

func (m *Main) Add(e *vecty.Event) {
	log.Println("add")
	m.counter++
	vecty.Rerender(m)
}

func (m *Main) Remove(e *vecty.Event) {
	log.Printf("remove")
	m.counter--
	vecty.Rerender(m)
}

func (m *Main) Render() *vecty.HTML {
	return elem.Body(elem.Div(
		elem.Paragraph(
			prop.Class("counter"),
			vecty.Text(fmt.Sprintf("Counter : %d", m.counter)),
		),
		elem.Button(
			prop.Class("action"),
			vecty.Text("Add"), event.Click(m.Add)),
		elem.Button(
			prop.Class("action"),

			vecty.Text("Remove"), event.Click(m.Remove)),
	))
}

func main() {
	vecty.SetTitle("Go JS!")
	vecty.AddStylesheet("index.css")
	vecty.RenderBody(&Main{})
}

package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"github.com/gopherjs/vecty/style"
)

type PageView struct {
	vecty.Core
}

// Render implements the vecty.Component interface
func (p *PageView) Render() *vecty.HTML {
	return elem.Body(
		elem.Form(
			vecty.Markup(
				style.Margin(style.Px(0)),
				event.Submit(p.onAdd).PreventDefault(),
			),
			elem.Input(
				vecty.Markup(
					prop.Placeholder("passwd"),
					event.Input(p.onLogin),
				),
			),
		),
	)
}

func (p *PageView) onAdd(event *vecty.Event) {

}

func (p *PageView) onLogin(event *vecty.Event) {

}

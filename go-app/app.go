// +build wasm

package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

var formContext map[string]string = make(map[string]string)

type input struct {
	app.Compo
	Value string
	Name string
}

func (h *input) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Value: "),
			app.Text(h.Value),
		),
		app.Input().
			Value(h.Value).
			OnChange(h.OnInputChange),
	)
}

func (h *input) OnInputChange(ctx app.Context, e app.Event) {
	h.Value = ctx.JSSrc.Get("value").String()
	formContext[h.Name] = h.Value
	h.Update()
}

type button struct {
	app.Compo
}

func (h *button) Render() app.UI {
	return app.Button().Body(app.Text("Submit")).OnClick(h.OnButtonClick)
}

func (h *button) OnButtonClick(ctx app.Context, e app.Event) {
	for k, v := range formContext {
		println(k)
		println(v)
	}
}

type container struct {
	app.Compo
}

func (h *container) Render() app.UI {
	return app.Div().Body(
		&input{ Name: "fname" },
		&button{},
	)
}

func main() {
	app.Route("/", &container{})
	app.Run()            
}

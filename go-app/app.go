// +build wasm
package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type input struct {
	app.Compo
	Value string
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
	h.Update()
}

func main() {
	app.Route("/", &input{})
	app.Run()                
}

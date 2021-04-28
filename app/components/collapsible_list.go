package components

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type collapsibleList struct {
	app.Compo

	Title        app.UI
	Elements     []app.UI
	ShowElements bool
}

func CollapsibleList() *collapsibleList {
	return &collapsibleList{}
}
func (cl *collapsibleList) WithTitle(title app.UI) *collapsibleList {
	cl.Title = title
	return cl
}

func (cl *collapsibleList) WithElements(elements []app.UI) *collapsibleList {
	cl.Elements = elements
	return cl
}

func (cl *collapsibleList) Render() app.UI {
	return app.Div().Class("collapsible").
		Body(
			app.Stack().Vertical().Content(
				app.Div().Class("collapsible-title").Class("row").Body(
					cl.Title,
					app.If(cl.ShowElements,
						app.I().
							Class("fa").
							Class("fa-chevron-circle-up").
							Style("color", "white"),
					).
						Else(
							app.I().
								Class("fa").
								Class("fa-chevron-circle-right").
								Style("color", "white"),
						),

				).OnClick(func(ctx app.Context, e app.Event) {
					ctx.Dispatch(func() {
						cl.ShowElements = !cl.ShowElements
						cl.Update()
					})
				}),
				app.If(cl.ShowElements,
					app.Stack().Class("collapsible-elements").Vertical().Content(
						app.Range(cl.Elements).Slice(func(i int) app.UI {
							return cl.Elements[i]
						}),
					),
				),
			),
		)
}

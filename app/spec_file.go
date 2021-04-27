package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/carousell/aggproto/app/models"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type specFile struct {
	app.Compo
	reg             *registry
	CurrentSpecInfo *models.SpecInfo
	ApiCtx          models.APIContext
}

func (sf *specFile) selectedInfoWatcher(ctx app.Context, selected *models.SpecInfo) {
	app.Log("selected info watcher")
	sf.CurrentSpecInfo = selected
	sf.reload(ctx)
}
func (sf *specFile) reload(ctx app.Context) {
	if sf.CurrentSpecInfo == nil {
		return
	}
	resp, err := http.Get(fmt.Sprintf("/api/specs/%s", sf.CurrentSpecInfo.Filename))
	if err != nil {
		app.Log(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		app.Log(err)
		return
	}
	var res models.APIContext
	err = json.Unmarshal(b, &res)
	if err != nil {
		app.Log(err)
		return
	}
	ctx.Dispatch(func() {
		sf.ApiCtx = res
		sf.Update()
	})
}

func (sf *specFile) OnMount(ctx app.Context) {
	ctx.Async(func() {
		app.Log("mounted")
		sf.reload(ctx)
	})
}

func (sf *specFile) Render() app.UI {
	return app.Div().Class("spec_file").
		Class("fill").
		Body(
			app.If(sf.CurrentSpecInfo != nil,
				app.Stack().Vertical().Content(
					app.Div().Body(
						app.Stack().Vertical().Content(
							app.Header().Text("API Descriptor"),
							app.Table().Body(
								app.TBody().Body(
									app.Tr().Body(
										app.Td().Text("Name"),
										app.Td().Text(sf.ApiCtx.API.Name),
									),
									app.Tr().Body(
										app.Td().Text("Group"),
										app.Td().Text(sf.ApiCtx.API.Group),
									),
									app.Tr().Body(
										app.Td().Text("Version"),
										app.Td().Text(sf.ApiCtx.API.Version),
									),
								),
							),
							//app.Stack().Content(app.H6().Text("Name"), app.Div().Text(sf.ApiCtx.API.Name)),
							//app.Stack().Content(app.H6().Text("Group"), app.Div().Text(sf.ApiCtx.API.Group)),
							//app.Stack().Content(app.H6().Text("Version"), app.Div().Text(sf.ApiCtx.API.Version)),
						),
					),
					app.Div().Body(
						app.Stack().Vertical().Content(
							app.Header().Text("Meta"),
							app.Table().Body(
								app.TBody().Body(
									app.Td().Text("GoPackage"),
									app.Td().Text(sf.ApiCtx.Meta.GoPackage),
								),
							),
						),
					),
					app.Div().Body(
						app.Stack().Vertical().Content(
							app.Header().Text("Output"),
							app.Table().Body(
								app.THead().Body(
									app.Td().Text("Alias"),
									app.Td().Text("Underlying"),
								),
								app.TBody().Body(
									app.Range(sf.ApiCtx.Outputs).Slice(func(i int) app.UI {
										o := sf.ApiCtx.Outputs[i]
										return app.Tr().Body(
											app.Td().Text(o.Output.Value),
											app.Td().Text(o.Input.Value),
										)
									}),
								),
							),
						),
					),
				),
			),
		)
}

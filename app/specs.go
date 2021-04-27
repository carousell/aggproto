package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/carousell/aggproto/app/models"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)


type selectedSpecWatcher func(ctx app.Context, info *models.SpecInfo)

type apiSpecList struct {
	app.Compo
	Specs        []*models.SpecInfo
	watchers     []selectedSpecWatcher
	SelectedSpec *models.SpecInfo
}

func (asl *apiSpecList) addWatcher(watcher selectedSpecWatcher) {
	asl.watchers = append(asl.watchers, watcher)
}

func (asl *apiSpecList) OnMount(ctx app.Context) {
	ctx.Async(func() {
		resp, err := http.Get("/api/specs/list")
		if err != nil {
			app.Log(err)
			return
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			app.Log(err)
			return
		}
		var res []*models.SpecInfo
		err = json.Unmarshal(b, &res)
		if err != nil {
			app.Log(err)
			return
		}
		ctx.Dispatch(func() {
			asl.Specs = res
			asl.Update()
		})
	})

}

func (asl *apiSpecList) Render() app.UI {

	return app.Div().
		Class("nav").
		Class("fill").
		Class("unselectable").
		Body(
			app.Stack().
				Class("app-title").
				Class("hspace-out").
				//Middle(),
				Content(
					app.Header().
						Body(
							app.A().
								Class("hApp").
								Class("focus").
								Class("glow").
								Href("/").
								Text("AggProto"),
						),
				),
			app.Nav().
				Class("nav-content").
				Body(app.Div().
					Class("nav-specs").
					Body(app.Stack().
						Class("nav-specs-stack").
						Content(
							app.Div().
								Class("hspace-out").
								Body(
									app.Range(asl.Specs).Slice(func(i int) app.UI {
										spec := asl.Specs[i]
										return newCard().
											ID(spec.Id()).
											Label(fmt.Sprintf("%s_v%d", spec.APIDescriptor.Group, spec.APIDescriptor.Version)).
											Help(spec.APIDescriptor.Name).
											OnClick(asl.onClick(spec)).
											Focus(asl.SelectedSpec != nil && asl.SelectedSpec.Id() == spec.Id())
									}),
								),
						),
					)),
		)

	//return app.Ul().Body(
	//	app.Range(asl.Specs).Slice(func(i int) app.UI {
	//		return app.Li().Text(asl.Specs[i].APIDescriptor.Name).OnClick(func(ctx app.Context, e app.Event) {
	//			app.Log("clicked ", asl.Specs[i].Filename)
	//			ctx.Async(func() {
	//				for _, watcher := range asl.watchers {
	//					app.Log("watcher called")
	//					watcher(ctx, asl.Specs[i])
	//				}
	//			})
	//		})
	//	}),
	//)
}
func (asl *apiSpecList) onClick(spec *models.SpecInfo) func(ctx app.Context) {
	return func(ctx app.Context) {
		ctx.Dispatch(func() {
			asl.SelectedSpec = spec
			asl.Update()
		})
		ctx.Async(func() {
			for _, w := range asl.watchers {
				w(ctx, spec)
			}
		})
	}
}

type card struct {
	app.Compo
	id      string
	classes []string
	label   string
	help    string
	href    string
	focus   bool
	onClick func(ui app.Context)
}

func newCard() *card {
	return &card{}
}
func (c *card) ID(v string) *card {
	c.id = v
	return c
}
func (c *card) Class(v string) *card {
	if v == "" {
		return c
	}
	c.classes = append(c.classes, v)
	return c
}
func (c *card) Label(v string) *card {
	c.label = v
	return c
}
func (c *card) Help(v string) *card {
	c.help = v
	return c
}
func (c *card) Focus(v bool) *card {
	c.focus = v
	return c
}
func (c *card) OnClick(v func(ctx app.Context)) *card {
	c.onClick = v
	return c
}

func (c *card) Render() app.UI {
	return app.Div().
		ID(c.id).
		Class("card").
		Class("heading").
		Class("fit").
		Class(strings.Join(c.classes, " ")).
		Title(c.help).
		OnClick(c.onClickInvoke).
		Body(
			app.Stack().Vertical().Content(
				app.H4().Class("H4").Text(c.help),
				app.Div().Text(c.label),
			),
		)
}

func (c *card) onClickInvoke(ctx app.Context, e app.Event) {
	if c.onClick == nil {
		return
	}
	e.PreventDefault()
	c.onClick(ctx)
}

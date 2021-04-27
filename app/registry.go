package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/carousell/aggproto/app/models"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type registry struct {
	app.Compo
	MessageList      []*models.Message
	searchTerm       string
	SelectedMessages []*models.Message
}

func (r *registry) OnMount(ctx app.Context) {
	ctx.Async(func() {
		msgs := r.loadMessages("*")
		if len(msgs) == 0 {
			return
		}
		app.Log("here", msgs)
		ctx.Dispatch(func() {
			r.MessageList = msgs
			r.Update()
		})
	})
}

func (r *registry) Render() app.UI {
	return app.Div().Class("registry").
		Body(
			app.Stack().Vertical().Content(
				app.Stack().Content(
					app.Text("Search Message:"),
					app.DataList().ID("messages").
						Body(
							app.Range(r.MessageList).Slice(func(i int) app.UI {
								msg := r.MessageList[i]
								return app.Option().Value(fmt.Sprintf("%s.%s", msg.PackageName, msg.Name))
							})),
					app.Input().AutoComplete(true).List("messages").OnInput(func(ctx app.Context, e app.Event) {
						r.searchTerm = ctx.JSSrc.Get("value").String()
					}),
					app.Button().Text("search").OnClick(func(ctx app.Context, e app.Event) {
						ctx.Async(func() {
							r.loadMessage(ctx)
						})
					}),
				),
				app.Range(r.SelectedMessages).Slice(func(i int) app.UI {
					msg := r.SelectedMessages[i]
				}),
			),
		)
}

func (r *registry) loadMessage(ctx app.Context) {
	msgs := r.loadMessages(r.searchTerm)
	if len(msgs) == 0 {
		return
	}
	ctx.Dispatch(func() {
		r.SelectedMessages = msgs
		r.Update()
	})
}

func (r *registry) loadMessages(prefix string) []*models.Message {
	var msgs []*models.Message
	resp, err := http.Get(fmt.Sprintf("/api/registry/search/message/%s", prefix))
	if err != nil {
		app.Log(err)
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		app.Log(err)
		return nil
	}
	defer resp.Body.Close()
	err = json.Unmarshal(body, &msgs)
	if err != nil {
		app.Log(err)
		return nil
	}
	return msgs
}

type messageCard struct {
	app.Compo
	msg *models.Message
}

func (mc *messageCard) Render() app.UI {
	return app.Div().Class("message-card").Body(
		app.Stack().Vertical().Content(
			app.Stack().Content(app.)
			),
	)
}

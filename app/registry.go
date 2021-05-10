//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/carousell/aggproto/app/components"
	registry2 "github.com/carousell/aggproto/pkg/registry"

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
					return &messageCard{msg: msg, fieldName: msg.Name, repeated: false}
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
	msg       *models.Message
	fieldName string
	repeated  bool
}

func (mc *messageCard) Render() app.UI {
	var fieldNames []app.UI
	for _, field := range mc.msg.Fields {
		if field.Type == registry2.FieldTypeMessage.GoTypeString() {
			fieldNames = append(fieldNames, &messageCard{msg: field.Message, fieldName: field.Name, repeated: field.Repeated})
		} else {
			fieldNames = append(fieldNames,
				app.Div().Class("row").Body(
					app.P().Text(field.Name),
					app.If(field.Repeated, app.P().Text("[]")),
					app.P().Text(field.Type),
				),
			)

		}
	}
	return app.Div().Class("message-card").Body(
		components.CollapsibleList().
			WithTitle(
				app.Div().Class("row").Body(
					app.P().Text(mc.fieldName),
					app.If(mc.repeated, app.P().Text("[]")),
					app.P().Text(fmt.Sprintf("%s.%s", mc.msg.PackageName, mc.msg.Name)),
				),
			).
			WithElements(fieldNames),
	)
}

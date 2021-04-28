package app

import (
	"net/http"

	"github.com/carousell/aggproto/app/apis"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

const backgroundColor = "#000000"

type aggProtoCanvas struct {
	app.Compo
}

func (s *aggProtoCanvas) Render() app.UI {
	reg := &registry{}
	asl := &apiSpecList{}
	sf := &specFile{reg: reg}
	asl.addWatcher(sf.selectedInfoWatcher)
	return app.Main().
		Class("aggProto").
		Body(
			app.Shell().
				MenuWidth(282).
				Menu(
					asl,
				).
				OverlayMenu().
				Content(
					app.Div().Class("content").Body(
						sf,
						reg,
					),
				),
		)
}

func Run(specDir string, registryDir string) error {
	sf := &specFile{}
	asl := &apiSpecList{}
	asl.addWatcher(sf.selectedInfoWatcher)
	app.Route("/", &aggProtoCanvas{
	})
	app.RunWhenOnBrowser()
	s := apis.Server(specDir, registryDir)
	http.Handle("/", &app.Handler{
		Name:            "Agg Proto",
		Description:     "A code generation tool",
		BackgroundColor: backgroundColor,
		ThemeColor:      backgroundColor,
		Styles: []string{
			"/web/app.css",
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/fontawesome.min.css",
		},
	})
	s.RegisterHttp()
	return http.ListenAndServe(":4891", nil)
}

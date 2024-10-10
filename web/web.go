package web

import "storality/app"

func Serve(app *app.Application) {
	app.Router.HandleFunc("GET /{$}", GetHome)
}
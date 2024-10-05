package app

import "net/http"

type Application struct {
	Router *http.ServeMux
}

func Init(r *http.ServeMux) *Application {
	app := &Application{
		Router: r,
	}
	return app
}
package admin

import (
	"storality/app"
)

func Serve(app *app.Application) {
	app.Router.HandleFunc("GET /admin/{$}", GetDashboard)

	app.Router.HandleFunc("GET /admin/page/new/{$}", GetNewPage)
	app.Router.HandleFunc("GET /admin/page/{id}/{$}", GetPageView)

	app.Router.HandleFunc("POST /admin/page/new/{$}", PostPageNew)
	app.Router.HandleFunc("POST /admin/page/{id}/{$}", PostPageNew)

	app.Router.HandleFunc("DELETE /admin/page/{id}/{$}", DeletePage)
}
package main

import (
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/app/drivers/web"
	"github.com/murlokswarm/app/examples/test"
)

func main() {
	app.DefaultLogger = app.ConcurrentLogger(app.NewConsole(true))

	app.Import(&test.Webview{})
	app.Import(&test.Menu{})

	app.Run(&web.Driver{
		DefaultURL: "/test.Webview",
	})
}

package main

import (
	"app/project/bootstrap"
	"app/project/web/middleware/identity"
	"app/project/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Awesome App", "kataras2006@hotmail.com")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}

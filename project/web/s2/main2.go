package main

import (
	"app/project/bootstrap"
	"app/project/web/middleware/identity"
	"app/project/web/routes"
)

func main() {
	app := bootstrap.New("Awesome App", "kataras2006@hotmail.com")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	app.Listen(":8081")
}

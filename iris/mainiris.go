package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	htmlEngine := iris.HTML("./iris", ".html")
	app.RegisterView(htmlEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("helllo world!")
	})
	//模板输出
	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "测试")
		ctx.ViewData("Content", "hello world!")
		ctx.View("hello.html")
	})


	app.Run(
		iris.Addr(":8080"),
		iris.WithCharset("UTF-8"),
		)
}

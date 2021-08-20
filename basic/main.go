package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()
    // load các template từ folder ".views"
    // Sẽ parse các file có đuôi html
    // sử dụng template HTML
    app.RegisterView(iris.HTML("./views", ".html"))

    // Method:    GET
    app.Get("/", func(ctx iris.Context) {
        // Gắn data biến {{.message}} bằng "Hello world!"
        ctx.ViewData("message", "Hello world!")
        // Render file: ./views/hello.html
        ctx.View("hello.html")
    })
    // Khởi động server ở cổng 8080
    app.Listen(":8080")
}

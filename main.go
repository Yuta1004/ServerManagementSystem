package main

import (
	"server-manage/view"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	baseURL := "/server-manage"

	router := gin.Default()
	router.Static(baseURL+"/static", "./static")
	router.HTMLRender = createHTMLRender()

	router.GET(baseURL+"/", view.TopView)
	router.GET(baseURL+"/crontab", view.CrontabView)

	router.Run(":19000")
}

func createHTMLRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("top", "templates/base.html", "templates/top.html", "templates/navbar.html")
	render.AddFromFiles("crontab", "templates/base.html", "templates/crontab.html", "templates/navbar.html")
	return render
}

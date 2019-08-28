package main

import (
	"os"
	"server-manage/controller"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/multitemplate"

)

func main() {
	baseURL := "/server-manage"

	router := gin.Default()
	router.Static(baseURL+"/static", "./static")
	router.HTMLRender = createHTMLRender()

	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET")))
	router.Use(sessions.Sessions("GoServerManageSession", store))

	router.GET(baseURL+"/", controller.TopPageController)
	router.GET(baseURL+"/login", controller.LoginPageController)
	router.POST(baseURL+"/login", controller.LoginPagePostController)

	menu := router.Group(baseURL+"/manage")
	menu.Use(sessionCheck())
	{
		menu.GET("/crontab", controller.CrontabPageController)
	}

	router.Run(":19000")
}

func createHTMLRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("top", "templates/base.html", "templates/top.html", "templates/navbar.html")
	render.AddFromFiles("crontab", "templates/base.html", "templates/crontab.html", "templates/navbar.html")
	render.AddFromFiles("login", "templates/base.html", "templates/login.html", "templates/navbar.html")
	return render
}

func sessionCheck() gin.HandlerFunc{
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("UserID")

		if userID == nil {
			c.Redirect(302, "../login")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
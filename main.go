package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/multitemplate"
	"server-manage/controller"
	"server-manage/authfunc"
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
	router.GET(baseURL+"/register", controller.RegisterPageController)
	router.POST(baseURL+"/register", controller.RegisterPagePostController)
	router.GET(baseURL+"/logout", controller.LogoutPageController)

	menu := router.Group(baseURL+"/manage")
	menu.Use(sessionCheck())
	{
		menu.Static("/static", "./static")
		menu.GET("/crontab", controller.CrontabPageController)
	}

	router.Run(":19000")
}

func createHTMLRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("top", "templates/base.html", "templates/top.html", "templates/navbar.html")
	render.AddFromFiles("crontab", "templates/base.html", "templates/crontab.html", "templates/navbar.html")
	render.AddFromFiles("login", "templates/base.html", "templates/login.html", "templates/navbar.html")
	render.AddFromFiles("register", "templates/base.html", "templates/register.html", "templates/navbar.html")
	render.AddFromFiles("logout", "templates/base.html", "templates/logout.html", "templates/navbar.html")
	return render
}

func sessionCheck() gin.HandlerFunc{
	return func(c *gin.Context) {
		checkResult, message := authfunc.CheckOKSession(c)
		if checkResult {
			c.Next()
		} else {
			c.Redirect(302, "../login?error=" + message)
			c.Abort()
		}
	}
}

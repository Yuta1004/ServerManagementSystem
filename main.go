package main

import (
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/multitemplate"
	"server-manage/controller"
	"server-manage/db"
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
	return render
}

func sessionCheck() gin.HandlerFunc{
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("UserID")
		passphrase := session.Get("Passphrase")

		if userID == nil {
			c.Redirect(302, "../login")
			c.Abort()
            return
		}

		// 有効期限チェック
		userIDStr := userID.(string)
		sessionDBData := (*db.GetSessionDataFromDB(userIDStr))[0]
		if time.Now().Unix() > int64(sessionDBData.ExpirationUnixTime) {
			c.Redirect(302, "../login?error=OverExpirationDate")
			c.Abort()
            return
		}

		// パスフレーズチェック
		passphraseStr := passphrase.(string)
		if passphraseStr != sessionDBData.Passphrase {
			c.Redirect(302, "../login?error=FraudSession")
			c.Abort()
            return
		}
		c.Next()
	}
}

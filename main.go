package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/irwantocia/go-template/controllers"
	"github.com/irwantocia/go-template/models"
	"github.com/irwantocia/go-template/pkg/logs"
	"github.com/irwantocia/go-template/pkg/validation"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	models.Init()
	logs.Init()
	validation.Init()
}

func main() {
	defer logs.Logger.Sync()
	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()

	r.SetTrustedProxies(nil)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET")))
	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24 * 7, // 1 week
		Path:     "/",
		Secure:   gin.Mode() == gin.ReleaseMode, // Secure only in production
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode, // Added SameSite attribute for better security
	})
	r.Use(sessions.Sessions("session", store))
	r.Use(logs.GinMiddlewareLogger())

	r.Static("/static", "./static")

	r.HTMLRender = func() multitemplate.Renderer {
		root := "templates/"
		r := multitemplate.NewRenderer()
		// ----------------------------
		// add here new route template
		// ----------------------------
		r.AddFromFiles("index", root+"index/index.html")
		r.AddFromFiles("404", root+"404.html")
		return r
	}()

	// -----------------------
	// add here new controller
	// -----------------------
	controllers.Index(r)

	// not found
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404", gin.H{
			"message": "The requested resource was not found",
		})
	})

	r.Run()
}

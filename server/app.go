package server

import (
	"db2/cms/ctrl"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func Init() {

	handleScripts()

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.HTMLRender = createRenderer()
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/img", "./static/img")
	r.Static("/fonts", "./static/fonts")
	r.Static("/uploads", "./uploads")

	// Render
	r.GET("/", ctrl.RenderIndex)
	r.GET("/login", ctrl.RenderLogin)
	r.GET("/register", ctrl.RenderRegister)
	r.GET("/events", ctrl.RenderEvents)

	// API
	r.POST("/login", ctrl.Login)
	r.POST("/register", ctrl.Register)
	r.GET("/logout", ctrl.Logout)

	r.POST("/api/posts", authenticationCheck(), ctrl.AddPost)
	r.PUT("/api/posts/:id", authenticationCheck(), ctrl.UpdatePost)
	r.DELETE("/api/posts/:id", authenticationCheck(), ctrl.RemovePost)

	r.POST("/api/events", authenticationCheck(), ctrl.AddEvent)
	r.PUT("/api/events/:id", authenticationCheck(), ctrl.UpdateEvent)
	r.DELETE("/api/events/:id", authenticationCheck(), ctrl.RemoveEvent)

	r.POST("/api/pages", authenticationCheck(), ctrl.AddPage)
	r.PUT("/api/pages/:id", authenticationCheck(), ctrl.UpdatePage)
	r.DELETE("/api/pages/:id", authenticationCheck(), ctrl.RemovePage)

	r.PUT("/api/sites/:id", authenticationCheck(), ctrl.UpdateSite)

	r.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("image")

		extension := filepath.Ext(file.Filename)
		name := uuid.NewString() + extension

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "./uploads/"+name)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", name))
	})

	// Render pages
	r.GET("/:page", ctrl.RenderPage)

	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}

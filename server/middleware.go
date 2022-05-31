package server

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// func fetchSiteAndPages() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		pages, err := rdg.ListPages(1)
// 		if err != nil {
// 			ctrl.RenderError(c, err)
// 			return
// 		}

// 		c.Set("pages", pages)

// 		site, err := rdg.GetSite(1)
// 		if err != nil {
// 			ctrl.RenderError(c, err)
// 			return
// 		}

// 		c.Set("site", site)

// 		c.Next()
// 	}
// }

func authenticationCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("id")
		if sessionID == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "unauthorized"})
			c.Abort()
		}
	}
}

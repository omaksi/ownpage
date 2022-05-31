package ctrl

import (
	"db2/cms/rdg"

	"github.com/gin-gonic/gin"
)

func UpdateSite(c *gin.Context) {
	var site rdg.Site
	c.BindJSON(&site)
	err := rdg.UpdateSite(site)
	if err != nil {
		RenderError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"id":          site.Id,
		"title":       site.Title,
		"description": site.Description,
	})
}

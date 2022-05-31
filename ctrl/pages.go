package ctrl

import (
	"db2/cms/rdg"

	"github.com/gin-gonic/gin"
)

func AddPage(c *gin.Context) {
	var page rdg.Page
	err := c.Bind(&page)
	if err != nil {
		RenderError(c, err)
		return
	}
	err = rdg.AddPage("1", page.Title, page.Body, page.Slug)
	if err != nil {
		RenderError(c, err)
		return
	}
	// c.JSON(200, gin.H{
	// 	"id":    page.Id,
	// 	"title": page.Title,
	// 	"body":  page.Body,
	// 	"slug":  page.Slug,
	// })
	c.Redirect(302, "/"+page.Slug)
}

func UpdatePage(c *gin.Context) {
	var page rdg.Page
	c.Bind(&page)
	err := rdg.UpdatePage(c.Param("id"), page.Title, page.Body)
	if err != nil {
		RenderError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"id":    c.Param("id"),
		"title": page.Title,
		"body":  page.Body,
		"slug":  page.Slug,
	})
}

func RemovePage(c *gin.Context) {
	err := rdg.DeletePage(c.Param("id"))
	if err != nil {
		RenderError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"id": c.Param("id"),
	})
}

package ctrl

import (
	"db2/cms/rdg"

	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {
	var post rdg.Post
	err := c.Bind(&post)
	if err != nil {
		RenderError(c, err)
		return
	}
	err = rdg.AddPost("1", post.Body)
	if err != nil {
		RenderError(c, err)
		return
	}
	// c.JSON(200, gin.H{
	// 	"id":   post.Id,
	// 	"body": post.Body,
	// })
	c.Redirect(302, "/")
}

func UpdatePost(c *gin.Context) {
	var post rdg.Post
	c.BindJSON(&post)
	err := rdg.UpdatePost(c.Param("id"), post.Body)
	if err != nil {
		RenderError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"id":   c.Param("id"),
		"body": post.Body,
	})
}

func RemovePost(c *gin.Context) {
	err := rdg.DeletePost(c.Param("id"))
	if err != nil {
		RenderError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"id": c.Param("id"),
	})
}

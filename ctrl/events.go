package ctrl

import (
	"db2/cms/rdg"

	"github.com/gin-gonic/gin"
)

func AddEvent(c *gin.Context) {
	var event rdg.Event
	err := c.Bind(&event)
	if err != nil {
		RenderError(c, err)
		return
	}
	err = rdg.AddEvent("1", event.Title, event.Body, event.EventDate)
	if err != nil {
		RenderError(c, err)
		return
	}
	// c.JSON(200, gin.H{
	// 	"id":        event.Id,
	// 	"title":     event.Title,
	// 	"body":      event.Body,
	// 	"eventDate": event.EventDate,
	// })
	c.Redirect(302, "/")
}

func UpdateEvent(c *gin.Context) {
	var event rdg.Event
	c.BindJSON(&event)
	err := rdg.UpdateEvent(c.Param("id"), event.Title, event.Body, event.EventDate)
	if err != nil {
		RenderError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"id":        c.Param("id"),
		"title":     event.Title,
		"body":      event.Body,
		"eventDate": event.EventDate,
	})
}

func RemoveEvent(c *gin.Context) {
	err := rdg.DeleteEvent(c.Param("id"))
	if err != nil {
		RenderError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"id": c.Param("id"),
	})
}

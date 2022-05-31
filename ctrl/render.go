package ctrl

import (
	"db2/cms/rdg"
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func withEssential(c *gin.Context, h gin.H) (gin.H, error) {

	pages, err := rdg.ListPages(1)
	if err != nil {
		return h, err
	}

	site, err := rdg.GetSite(1)
	if err != nil {
		return h, err
	}

	session := sessions.Default(c)
	sessionID := session.Get("id")
	sessionIDstring := ""
	if sessionID == nil {
		sessionIDstring = ""
	} else {
		sessionIDstring = sessionID.(string)
	}

	h["Pages"] = pages
	h["Site"] = site
	h["SessionId"] = sessionIDstring

	return h, nil
}

func RenderIndex(c *gin.Context) {
	posts, err := rdg.GetPosts(1)
	if err != nil {
		RenderError(c, err)
		return
	}

	res, err := withEssential(c, gin.H{
		"Posts": posts,
	})
	if err != nil {
		RenderError(c, err)
		return
	}

	c.HTML(200, "index", res)
}

func RenderEvents(c *gin.Context) {
	events, err := rdg.GetEvents(1)
	if err != nil {
		RenderError(c, err)
		return
	}

	res, err := withEssential(c, gin.H{
		"Events": events,
	})
	if err != nil {
		RenderError(c, err)
		return
	}

	c.HTML(200, "events", res)

}

func RenderPage(c *gin.Context) {
	// fmt.Println("RenderPage", c.Param("page"))
	page, err := rdg.GetPage(c.Param("page"))
	if err != nil {
		c.HTML(404, "error", gin.H{"error": errors.New("Page not found")})

		// RenderError(c, err)
		return
	}

	res, err := withEssential(c, gin.H{
		"Page": page,
	})
	if err != nil {
		RenderError(c, err)
		return
	}

	c.HTML(200, "page", res)

}

func RenderError(c *gin.Context, err error) {
	c.HTML(500, "error", gin.H{"error": err.Error()})
}

func RenderLogin(c *gin.Context) {

	res, err := withEssential(c, gin.H{})
	if err != nil {
		RenderError(c, err)
		return
	}

	c.HTML(200, "login", res)
}

func RenderRegister(c *gin.Context) {

	res, err := withEssential(c, gin.H{})
	if err != nil {
		RenderError(c, err)
		return
	}

	c.HTML(200, "register", res)
}

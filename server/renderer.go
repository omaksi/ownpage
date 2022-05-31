package server

import (
	"html/template"
	"math/rand"
	"time"

	"github.com/gin-contrib/multitemplate"
)

func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func toDisplayTime(t time.Time) string {
	return t.Format("02 Jan 2006 15:04")
}

func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	funcMap := template.FuncMap{
		"randomInt":     randomInt, // generates a random integer
		"toDisplayTime": toDisplayTime,
	}

	t := map[string][]string{
		"base": {
			"templates/base/layout.go.html",
			"templates/base/footer.go.html",
			"templates/base/title.go.html",
			"templates/base/header.go.html",
			"templates/base/main-menu.go.html",
			"templates/base/main-menu-item.go.html",
		},
		"index": {
			"templates/posts/posts.go.html",
			"templates/posts/post.go.html",
		},
		"events": {
			"templates/events/events.go.html",
			"templates/events/event.go.html",
		},
		"page": {
			"templates/page.go.html",
		},
		"login":    {"templates/login.go.html"},
		"register": {"templates/register.go.html"},
		"error":    {"templates/error.go.html"},
	}

	r.AddFromFilesFuncs("index", funcMap, append(t["base"], t["index"]...)...)
	r.AddFromFilesFuncs("login", funcMap, append(t["base"], t["login"]...)...)
	r.AddFromFilesFuncs("register", funcMap, append(t["base"], t["register"]...)...)
	r.AddFromFilesFuncs("error", funcMap, append(t["base"], t["error"]...)...)
	r.AddFromFilesFuncs("events", funcMap, append(t["base"], t["events"]...)...)
	r.AddFromFilesFuncs("page", funcMap, append(t["base"], t["page"]...)...)

	return r
}

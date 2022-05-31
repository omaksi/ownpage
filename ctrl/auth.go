package ctrl

import (
	"db2/cms/rdg"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordBytes), err
}

func doPasswordsMatch(hashedPassword string, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(currPassword),
	)
	return err == nil
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, err := rdg.GetUser(username)
	if err != nil {
		log.Println("Login couldn't get user")
		c.Redirect(302, "/login")
		return
	}

	if doPasswordsMatch(user.Password, password) {
		log.Println("Passwords match")
		session.Set("id", user.Email)
		session.Save()
		c.Redirect(302, "/")
		return
	}

	log.Println("Login failed")
	c.Redirect(302, "/login")

}

func Register(c *gin.Context) {
	// session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	hashedPassword, err := hashPassword(password)
	if err != nil {
		c.HTML(500, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	err = rdg.RegisterUser(username, hashedPassword)
	if err != nil {
		c.HTML(500, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	// session.Set("id", username)
	// session.Save()
	c.Redirect(302, "/login")

}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(302, "/")
}

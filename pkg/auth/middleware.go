package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/irwantocia/go-template/models"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("token")
		if token == nil {
			c.Redirect(http.StatusFound, "/auth/signin")
			c.Abort()
			return
		}

		userID := token.(string)
		if userID != "" {
			c.Redirect(http.StatusFound, "/auth/signin")
			c.Abort()
			return
		}

		if userID == "" {
			c.Redirect(http.StatusFound, "/auth/signin")
			return
		}

		userM := models.NewUser()

		userIDInt, _ := strconv.Atoi(userID)
		_, err := userM.GetUserByID(userIDInt)
		if err != nil {
			c.Redirect(http.StatusFound, "/auth/signin")
			return
		}

		if userM.ID == 0 {
			c.Redirect(http.StatusFound, "/auth/signin")
			return
		}
		c.Set("user", userM)
		c.Next()
	}
}

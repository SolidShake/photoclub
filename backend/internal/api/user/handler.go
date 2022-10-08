package user

import (
	"github.com/SolidShake/photoclub/internal/api/auth"
	coreUser "github.com/SolidShake/photoclub/internal/core/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(auth.IdentityKey)
	c.JSON(200, gin.H{
		"userID": claims[auth.IdentityKey],
		"email":  user.(*coreUser.User).Email,
		"text":   "Hello World.",
	})
}

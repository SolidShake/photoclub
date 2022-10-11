package user

import (
	"github.com/SolidShake/photoclub/internal/api/auth"
	coreUser "github.com/SolidShake/photoclub/internal/core/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *coreUser.Service
}

func NewHandler(service *coreUser.Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) UserHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(auth.IdentityKey)
	c.JSON(200, gin.H{
		"userID": claims[auth.IdentityKey],
		"email":  user.(*coreUser.User).Email,
	})
}

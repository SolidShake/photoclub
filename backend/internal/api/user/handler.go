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

// @TODO remove
type userResponse struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// User godoc
// @Summary      User
// @Description  get user info
// @Tags         Users
// @Produce      json
// @Success      200  {object}  userResponse
// @Failure      401  {object}  errorResponse
// @Security     ApiKeyAuth
// @Router       /user [get]
func (h Handler) UserHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(auth.IdentityKey)
	c.JSON(200, gin.H{
		"userID": claims[auth.IdentityKey],
		"email":  user.(*coreUser.User).Email,
	})
}

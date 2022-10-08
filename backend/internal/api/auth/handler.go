package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth godoc
// @Summary      Auth
// @Description  get user info
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  interface{}
// @Failure      400  {object}  interface{}
// @Security     ApiKeyAuth
// @Router       /auth/login [post]
func Login(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) == 0 {
		// return error
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": "test-token",
	})
}

// Auth godoc
// @Summary      Auth
// @Description  register user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      201  {object}  interface{}
// @Failure      400  {object}  interface{}
// @Security     ApiKeyAuth
// @Router       /auth/register [post]
func Register(ctx *gin.Context) {
	ctx.Status(http.StatusCreated)
}

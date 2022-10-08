package auth

import (
	"net/http"

	coreUser "github.com/SolidShake/photoclub/internal/core/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Auth godoc
// @Summary      Auth
// @Description  Login user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  interface{}
// @Failure      401  {object}  interface{}
// @Security     ApiKeyAuth
// @Router       /auth/login [post]
func Login(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	email := loginVals.Email
	password := loginVals.Password

	if (email == "admin" && password == "admin") || (email == "test" && password == "test") {
		return &coreUser.User{
			Email: email,
		}, nil
	}

	return nil, ErrFailedAuthentication
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

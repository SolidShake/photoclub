package auth

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup, jwtMiddleware *jwt.GinJWTMiddleware) {
	router.POST("/login", loginHandler(jwtMiddleware))
	router.GET("/refresh_token", refreshHandler(jwtMiddleware))
	router.GET("/register", register)
}

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
func loginHandler(ginJWT *jwt.GinJWTMiddleware) func(c *gin.Context) {
	return ginJWT.LoginHandler
}

// Auth godoc
// @Summary      Auth
// @Description  Refresh user login token
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  interface{}
// @Failure      401  {object}  interface{}
// @Security     ApiKeyAuth
// @Router       /auth/refresh_token [get]
func refreshHandler(ginJWT *jwt.GinJWTMiddleware) func(c *gin.Context) {
	return ginJWT.RefreshHandler
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
func register(ctx *gin.Context) {
	ctx.Status(http.StatusCreated)
}

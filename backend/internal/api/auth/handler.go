package auth

import (
	"fmt"
	"net/http"

	coreUser "github.com/SolidShake/photoclub/internal/core/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type registerForm struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Handler struct {
	service *coreUser.Service
}

func NewHandler(service *coreUser.Service) *Handler {
	return &Handler{service: service}
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
func (h Handler) loginHandler(ginJWT *jwt.GinJWTMiddleware) func(c *gin.Context) {
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
func (h Handler) refreshHandler(ginJWT *jwt.GinJWTMiddleware) func(c *gin.Context) {
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
func (h Handler) register(ctx *gin.Context) {
	var registerVals registerForm
	if err := ctx.ShouldBind(&registerVals); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid register values"})
		return
	}
	email := registerVals.Email
	password := registerVals.Password

	err := h.service.CreateUser(email, password)
	if err != nil {
		// log error
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"email":    email,
		"password": password,
	})
}

func (h Handler) Routes(router *gin.RouterGroup, jwtMiddleware *jwt.GinJWTMiddleware) {
	router.POST("/login", h.loginHandler(jwtMiddleware))
	router.GET("/refresh_token", h.refreshHandler(jwtMiddleware))
	router.POST("/register", h.register)
}

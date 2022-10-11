package auth

import (
	"net/http"

	coreUser "github.com/SolidShake/photoclub/internal/core/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type registerForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=50"`
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
	errs := validate(ctx, &registerVals)
	if errs != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	email := registerVals.Email
	password := registerVals.Password

	err := h.service.CreateUser(email, password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (h Handler) Routes(router *gin.RouterGroup, jwtMiddleware *jwt.GinJWTMiddleware) {
	router.POST("/login", h.loginHandler(jwtMiddleware))
	router.GET("/refresh_token", h.refreshHandler(jwtMiddleware))
	router.POST("/register", h.register)
}

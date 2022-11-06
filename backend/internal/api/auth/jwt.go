package auth

import (
	"errors"
	"time"

	coreUser "github.com/SolidShake/photoclub/internal/core/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	ErrMissingLoginValues   = errors.New("missing Email, Nickname or Password")
	ErrFailedAuthentication = errors.New("incorrect Email, Nickname or Password")
)

type loginForm struct {
	EmailOrNickname string `form:"email_or_nickname" json:"email_or_nickname" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required,min=3,max=50"`
}

var (
	IdentityKey = "id"
	EmailKey    = "email"
)

func AuthMiddleware(userService *coreUser.Service) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*coreUser.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.ID,
					EmailKey:    v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &coreUser.User{
				ID:    claims[IdentityKey].(string),
				Email: claims[EmailKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals loginForm
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", ErrMissingLoginValues
			}
			emailOrNickname := loginVals.EmailOrNickname
			password := loginVals.Password

			user, err := userService.GetUser(emailOrNickname, password)
			if err != nil {
				return "", ErrFailedAuthentication
			}
			return user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			_, ok := data.(*coreUser.User)
			return ok
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}

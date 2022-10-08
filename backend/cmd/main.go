package main

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/SolidShake/photoclub/docs"

	apiAuth "github.com/SolidShake/photoclub/internal/api/auth"
	apiUser "github.com/SolidShake/photoclub/internal/api/user"
)

// @title           Photoclub API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used

func main() {
	r := gin.Default()

	authMiddleware, err := apiAuth.AuthMiddleware()
	if err != nil {
		panic(err)
	}

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			apiAuth.Routes(auth, authMiddleware)
		}
		user := v1.Group("/user")
		user.Use(authMiddleware.MiddlewareFunc())
		{
			user.GET("", apiUser.UserHandler)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

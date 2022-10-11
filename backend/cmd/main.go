package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/SolidShake/photoclub/db"
	_ "github.com/SolidShake/photoclub/docs"

	apiAuth "github.com/SolidShake/photoclub/internal/api/auth"
	apiUser "github.com/SolidShake/photoclub/internal/api/user"

	coreUser "github.com/SolidShake/photoclub/internal/core/user"
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
	err := godotenv.Load(".env")
	if err != nil {
		err := godotenv.Load(".env.dist")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := db.Initialize(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()

	db.MigrationUp(database)

	userRepository := coreUser.NewRepository(database)
	userService := coreUser.NewService(userRepository)

	authHandler := apiAuth.NewHandler(userService)
	apiHandler := apiUser.NewHandler(userService)

	authMiddleware, err := apiAuth.AuthMiddleware(userService)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			authHandler.Routes(auth, authMiddleware)
		}
		user := v1.Group("/user")
		user.Use(authMiddleware.MiddlewareFunc())
		{
			user.GET("", apiHandler.UserHandler)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

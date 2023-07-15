package main

import (
	"gateway/cmd/docs"
	"gateway/pkg/handlers"
	"gateway/pkg/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func main() {

	docs.SwaggerInfo.Title = "Zog_festiv"
	docs.SwaggerInfo.Description = "Yo Yo Yo 148 3 to the 3 to the 6 to the 9 "
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	r, err := InitializeApp()
	if err != nil {
		panic(err)
	}
	r.Run(":8080")
}

func initializeAuthenticationHandler(cc *grpc.ClientConn) (*handlers.AuthenticationHandler, error) {
	return handlers.NewAuthenticationHandler(cc), nil
}

func InitializeApp() (*gin.Engine, error) {
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authConn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	authenticationHandler, err := initializeAuthenticationHandler(authConn)
	if err != nil {
		return nil, err
	}
	routes.AuthenticationRoutes(r, authenticationHandler)

	return r, nil
}

package main

import (
	"gateway/pkg/handlers"
	"gateway/pkg/routes"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
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

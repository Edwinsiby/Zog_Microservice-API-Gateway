package main

import (
	"gateway/pkg/handlers"
	"gateway/pkg/routes"
	"gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func main() {
	utils.NewSwaggerInfo()
	r, err := InitializeApp()
	if err != nil {
		panic(err)
	}
	r.Run(":8080")

}

func initializeAuthenticationHandler(cc *grpc.ClientConn) (*handlers.AuthenticationHandler, error) {
	return handlers.NewAuthenticationHandler(cc), nil
}

func initializeAdminDashboardHandler(cc *grpc.ClientConn) (*handlers.AdminDashboardHandler, error) {
	return handlers.NewAdminDashboardHandler(cc), nil
}

func initializeUserDashboardHandler(cc *grpc.ClientConn) (*handlers.UserDashboardHandler, error) {
	return handlers.NewUserDashboardHandler(cc), nil
}

func initializeOrderHandler(cc *grpc.ClientConn) (*handlers.OrderHandler, error) {
	return handlers.NewOrderHandler(cc), nil
}

func InitializeApp() (*gin.Engine, error) {
	r := gin.Default()

	// Service1
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authConn, err := grpc.Dial("service1:5050", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	authenticationHandler, err := initializeAuthenticationHandler(authConn)
	if err != nil {
		return nil, err
	}
	routes.AuthenticationRoutes(r, authenticationHandler)

	// Service2
	adminDashboardConn, err := grpc.Dial("service2:5051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	adminDashboardHandler, err := initializeAdminDashboardHandler(adminDashboardConn)
	if err != nil {
		return nil, err
	}
	routes.AdminDashboardRoutes(r, adminDashboardHandler)

	// Service3
	userDashboardConn, err := grpc.Dial("service3:5052", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	userDashboardHandler, err := initializeUserDashboardHandler(userDashboardConn)
	if err != nil {
		return nil, err
	}
	routes.UserDashboardRoutes(r, userDashboardHandler)

	// Service4
	orderConn, err := grpc.Dial("service4:5053", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	orderHandler, err := initializeOrderHandler(orderConn)
	if err != nil {
		return nil, err
	}
	routes.OrderRoutes(r, orderHandler)

	return r, nil
}

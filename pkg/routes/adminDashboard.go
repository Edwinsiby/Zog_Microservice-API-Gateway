package routes

import (
	"gateway/pkg/handlers"
	m "gateway/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func AdminDashboardRoutes(r *gin.Engine, handlers *handlers.AdminDashboardHandler) {
	adminRoutes := r.Group("", m.AdminRetriveCookie)

	r.GET("/service2/healthcheck", handlers.AdminDashboardHealthCheck)
	adminRoutes.GET("/admin/usermanagement", handlers.UserList)
	adminRoutes.GET("/admin/sortuser", handlers.SortUserByPermission)
	adminRoutes.GET("/admin/searchuser", handlers.SearchUser)
	adminRoutes.POST("/admin/userpermission/:id", handlers.TogglePermission)

	adminRoutes.POST("/admin/addapparel", handlers.CreateApparel)
	adminRoutes.PUT("/admin/editappaerl/:id", handlers.EditApparel)
	adminRoutes.DELETE("/admin/deleteapparel/:id", handlers.DeleteApparel)

	adminRoutes.POST("/admin/addcoupon", handlers.AddCoupon)
	adminRoutes.POST("/admin/addoffer", handlers.AddOffer)
}

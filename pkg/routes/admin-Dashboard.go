package routes

import (
	"gateway/pkg/handlers"
	m "gateway/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func AdminDashboardRoutes(r *gin.Engine, handlers *handlers.AdminDashboardHandler) {
	r.GET("/service2/healthcheck", handlers.AdminDashboardIndexHandler)
	r.GET("/admin/usermanagement", m.AdminRetriveCookie, handlers.UserList)
	r.GET("/admin/sortuser", m.AdminRetriveCookie, handlers.SortUserByPermission)
	r.GET("/admin/searchuser", m.AdminRetriveCookie, handlers.SearchUser)
	r.POST("/admin/userpermission/:id", m.AdminRetriveCookie, handlers.TogglePermission)

	r.POST("/admin/addapparel", m.AdminRetriveCookie, handlers.CreateApparel)
	r.PUT("/admin/editappaerl/:id", m.AdminRetriveCookie, handlers.EditApparel)
	r.DELETE("/admin/deleteapparel/:id", m.AdminRetriveCookie, handlers.DeleteApparel)

	r.POST("/admin/addcoupon", m.AdminRetriveCookie, handlers.AddCoupon)
	r.POST("/admin/addoffer", m.AdminRetriveCookie, handlers.AddOffer)
}

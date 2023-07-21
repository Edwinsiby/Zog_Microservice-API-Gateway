package routes

import (
	"gateway/pkg/handlers"
	m "gateway/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, handlers *handlers.OrderHandler) {
	userRoutes := r.Group("", m.UserRetriveCookie)
	adminRoutes := r.Group("", m.AdminRetriveCookie)
	r.GET("/service4/healthcheck", handlers.OrderHealthCheck)

	userRoutes.POST("/placeorder/:addressid/:payment", m.UserRetriveCookie, handlers.PlaceOrder)
	userRoutes.POST("/paymentverification/:sign/:razorid/:payid", m.UserRetriveCookie, handlers.PaymentVerification)
	userRoutes.PUT("/cancelorder/:orderid", m.UserRetriveCookie, handlers.CancelOrder)
	userRoutes.GET("/orderhistory", m.UserRetriveCookie, handlers.OrderHistory)
	userRoutes.POST("/orderreturn/:orderid", m.UserRetriveCookie, handlers.OrderReturn)

	adminRoutes.PUT("/updateorder/:orderid/:status", m.AdminRetriveCookie, handlers.AdminOrderUpdate)
	adminRoutes.POST("/updatereturn/:returnid/:status/:refund", m.AdminRetriveCookie, handlers.AdminReturnUpdate)
	adminRoutes.POST("/refund/:orderid", m.AdminRetriveCookie, handlers.AdminRefund)
	adminRoutes.GET("/salesreportbydate/:start/:end", m.AdminRetriveCookie, handlers.SalesReportByDate)
	adminRoutes.GET("/salesreportbyperiod/:period", m.AdminRetriveCookie, handlers.SalesReportByPeriod)
	adminRoutes.GET("/salesreportbycategory/:category/:period", m.AdminRetriveCookie, handlers.SalesReportByCategory)
	adminRoutes.POST("/sortorders", m.AdminRetriveCookie, handlers.SortOrderByStatus)
}

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

	userRoutes.POST("/placeorder/:addressid/:payment", handlers.PlaceOrder)
	userRoutes.POST("/paymentverification/:sign/:razorid/:payid", handlers.PaymentVerification)
	userRoutes.PUT("/cancelorder/:orderid", handlers.CancelOrder)
	userRoutes.GET("/orderhistory", handlers.OrderHistory)
	userRoutes.POST("/orderreturn/:orderid", handlers.OrderReturn)

	adminRoutes.PUT("/updateorder/:orderid/:status", handlers.AdminOrderUpdate)
	adminRoutes.POST("/updatereturn/:returnid/:status/:refund", handlers.AdminReturnUpdate)
	adminRoutes.POST("/refund/:orderid", handlers.AdminRefund)
	adminRoutes.GET("/salesreportbydate/:start/:end", handlers.SalesReportByDate)
	adminRoutes.GET("/salesreportbyperiod/:period", handlers.SalesReportByPeriod)
	adminRoutes.GET("/salesreportbycategory/:category/:period", handlers.SalesReportByCategory)
	adminRoutes.POST("/sortorders", handlers.SortOrderByStatus)
}

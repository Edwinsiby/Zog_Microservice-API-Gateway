package routes

import (
	"gateway/pkg/handlers"
	m "gateway/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, handlers *handlers.OrderHandler) {

	// user
	r.GET("/service4/healthcheck", handlers.OrderIndexHandler)
	r.POST("/placeorder/:addressid/:payment", m.UserRetriveCookie, handlers.PlaceOrder)
	r.POST("/paymentverification/:sign/:razorid/:payid", m.UserRetriveCookie, handlers.PaymentVerification)
	r.PUT("/cancelorder/:orderid", m.UserRetriveCookie, handlers.CancelOrder)
	r.GET("/orderhistory", m.UserRetriveCookie, handlers.OrderHistory)
	r.POST("/orderreturn/:orderid", m.UserRetriveCookie, handlers.OrderReturn)

	// admin
	r.PUT("/updateorder/:orderid/:status", m.AdminRetriveCookie, handlers.AdminOrderUpdate)
	r.POST("/updatereturn/:returnid/:status/:refund", m.AdminRetriveCookie, handlers.AdminReturnUpdate)
	r.POST("/refund/:orderid", m.AdminRetriveCookie, handlers.AdminRefund)
	r.GET("/salesreportbydate/:start/:end", m.AdminRetriveCookie, handlers.SalesReportByDate)
	r.GET("/salesreportbyperiod/:period", m.AdminRetriveCookie, handlers.SalesReportByPeriod)
	r.GET("/salesreportbycategory/:category/:period", m.AdminRetriveCookie, handlers.SalesReportByCategory)
	r.POST("/sortorders", m.AdminRetriveCookie, handlers.SortOrderByStatus)
}

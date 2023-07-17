package routes

import (
	"gateway/pkg/handlers"
	m "gateway/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserDashboardRoutes(r *gin.Engine, handlers *handlers.UserDashboardHandler) {
	r.GET("/service3/healthcheck", handlers.UserDashboardHealthCheck)
	r.GET("/home", m.UserRetriveCookie, handlers.Home)
	r.POST("/addaddress", m.UserRetriveCookie, handlers.AddAddress)
	r.GET("/userdetails", m.UserRetriveCookie, handlers.ShowUserDetails)
	r.GET("/apparels", m.UserRetriveCookie, handlers.Apparels)
	r.GET("/searchapparel", m.UserRetriveCookie, handlers.SearchApparels)
	r.GET("/apparelsdetails/:apparelid", m.UserRetriveCookie, handlers.ApparelDetails)

	r.POST("/addtocart/:category/:productid/:quantity", m.UserRetriveCookie, handlers.AddToCart)
	r.POST("/addtowishlist/:category/:productid", m.UserRetriveCookie, handlers.AddToWishlist)
	r.GET("/usercartlist", m.UserRetriveCookie, handlers.CartList)
	r.GET("/usercart", m.UserRetriveCookie, handlers.Cart)
	r.DELETE("/removefromcart/:product/:id", m.UserRetriveCookie, handlers.RemoveFromCart)
	r.DELETE("/removefromwishlist/:product/:id", m.UserRetriveCookie, handlers.RemoveFromWishlist)
	r.GET("/userwishlist", m.UserRetriveCookie, handlers.ViewWishlist)
	r.GET("/coupons", m.UserRetriveCookie, handlers.AvailableCoupons)
	r.POST("/applycoupon/:code", m.UserRetriveCookie, handlers.ApplyCoupon)
	r.GET("/offer", m.UserRetriveCookie, handlers.OfferCheck)
}

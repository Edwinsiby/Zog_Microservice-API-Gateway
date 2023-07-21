package routes

import (
	"gateway/pkg/handlers"
	m "gateway/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserDashboardRoutes(r *gin.Engine, handlers *handlers.UserDashboardHandler) {
	userRoutes := r.Group("", m.UserRetriveCookie)

	userRoutes.GET("/home", handlers.Home)
	userRoutes.POST("/addaddress", handlers.AddAddress)
	userRoutes.GET("/userdetails", handlers.ShowUserDetails)
	userRoutes.GET("/products", handlers.Apparels)
	userRoutes.GET("/searchapparel", handlers.SearchApparels)
	userRoutes.GET("/productdetails/:apparelid", handlers.ApparelDetails)

	userRoutes.POST("/addtocart/:productid/:quantity", handlers.AddToCart)
	userRoutes.POST("/addtowishlist/:productid", handlers.AddToWishlist)
	userRoutes.GET("/usercartlist", handlers.CartList)
	userRoutes.GET("/usercart", handlers.Cart)
	userRoutes.DELETE("/removefromcart/:product/:id", handlers.RemoveFromCart)
	userRoutes.DELETE("/removefromwishlist/:product/:id", handlers.RemoveFromWishlist)
	userRoutes.GET("/userwishlist", handlers.ViewWishlist)
	userRoutes.GET("/coupons", handlers.AvailableCoupons)
	userRoutes.POST("/applycoupon/:code", handlers.ApplyCoupon)
	userRoutes.GET("/offer", handlers.OfferCheck)
	r.GET("/service3/healthcheck", handlers.UserDashboardHealthCheck)
}

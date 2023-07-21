package routes

import (
	"gateway/pkg/handlers"
	m "gateway/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(r *gin.Engine, handlers *handlers.AuthenticationHandler) {
	userRoutes := r.Group("", m.UserRetriveCookie)
	adminRoutes := r.Group("", m.AdminRetriveCookie)
	userRoutes.GET("/service1/healthcheck", handlers.HealthCheck)
	userRoutes.POST("/user/signup", handlers.UserSignup)
	userRoutes.POST("/user/signupwithotp", handlers.SignupWithOtp)
	userRoutes.POST("/user/signupotpvalidation", handlers.SignupOtpValidation)
	userRoutes.POST("/user/loginwithotp", handlers.LoginWithOtp)
	userRoutes.POST("/user/otpvalidation", handlers.LoginOtpValidation)
	userRoutes.POST("/user/loginwithpassword", handlers.LoginWithPassword)

	adminRoutes.POST("/admin/registeradmin", m.AdminRetriveCookie, handlers.RegisterAdmin)
	adminRoutes.POST("/admin/loginpassword", handlers.AdminLoginWithPassword)
	adminRoutes.POST("/admin/otpvalidation", handlers.LoginOtpValidation)
	adminRoutes.GET("/admin/home", m.AdminRetriveCookie, handlers.AdminDashboard)
	adminRoutes.POST("/logout", handlers.Logout)
}

package routes

import (
	"gateway/pkg/handlers"
	m "gateway/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(r *gin.Engine, handlers *handlers.AuthenticationHandler) {
	adminRoutes := r.Group("", m.AdminRetriveCookie)
	r.GET("/service1/healthcheck", handlers.HealthCheck)
	r.POST("/user/signup", handlers.UserSignup)
	r.POST("/user/signupwithotp", handlers.SignupWithOtp)
	r.POST("/user/signupotpvalidation", handlers.SignupOtpValidation)
	r.POST("/user/loginwithotp", handlers.LoginWithOtp)
	r.POST("/user/otpvalidation", handlers.LoginOtpValidation)
	r.POST("/user/loginwithpassword", handlers.LoginWithPassword)
	r.POST("/logout", handlers.Logout)
	r.POST("/admin/loginpassword", handlers.AdminLoginWithPassword)
	r.POST("/admin/otpvalidation", handlers.LoginOtpValidation)
	r.POST("/admin/registeradmin", handlers.RegisterAdmin)
	adminRoutes.GET("/admin/home", handlers.AdminDashboard)

}

package handlers

import (
	"context"
	pb "gateway/pb"
	"gateway/pkg/entity"
	"gateway/pkg/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
)

type AuthenticationHandler struct {
	grpcClient pb.MyServiceClient
}

func NewAuthenticationHandler(cc *grpc.ClientConn) *AuthenticationHandler {
	client := pb.NewMyServiceClient(cc)
	return &AuthenticationHandler{
		grpcClient: client,
	}
}

func (a *AuthenticationHandler) IndexHandler(c *gin.Context) {
	req := &pb.Request{
		Data: "Mydata",
	}
	resp, err := a.grpcClient.MyMethod(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call MyMethod: %v", err)
	} else {
		c.JSON(http.StatusOK, resp.Result)
	}

}

// UserSignup  godoc
//
//	@Summary		signup
//	@Description	Adding new user to the database
//	@Tags			User&Admin Authentication
//	@Accept			json
//	@Produce		json
//	@Param			userInput	body		entity.Signup	true	"User Data"
//	@Success		200			{object}	entity.Signup
//	@Router			/user/signup [post]
func (a *AuthenticationHandler) UserSignup(c *gin.Context) {
	var userInput entity.Signup
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userInput.Email == "" || userInput.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Retry with valid credentials"})
		return
	}
	var user pb.CreateUserRequest
	copier.Copy(&user, &userInput)
	resp, err := a.grpcClient.CreateUser(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": resp})
	}

}

// UserSignup Otp  godoc
//
//	@Summary		signup with opt validation
//	@Description	Adding new user to the database
//	@Tags			User&Admin Authentication
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entity.Signup	true	"User Data"
//	@Success		200		{object}	entity.User
//	@Router			/user/signupwithotp [post]
func (a *AuthenticationHandler) SignupWithOtp(c *gin.Context) {
	var userInput entity.Signup
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user pb.CreateUserWithOtpRequest
	copier.Copy(&user, &userInput)
	resp, err := a.grpcClient.CreateUserWithOtp(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Otp send succesfuly to": resp.Phone, "Key": resp.Key})
	}
}

// SignupOtpValidation  godoc
//
//	@Summary		Sign Up Otp Validation
//	@Description	Validating user otp for signup
//	@Tags			User&Admin Authentication
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			key	formData	string	true	"Twilio Key"
//	@Param			otp	formData	string	true	"Otp"
//	@Success		200	{string}	string	"Success message"
//	@Router			/user/signupotpvalidation [post]
func (a *AuthenticationHandler) SignupOtpValidation(c *gin.Context) {
	key := c.PostForm("key")
	otp := c.PostForm("otp")
	req := &pb.OtpValidationRequest{
		Key: key,
		Otp: otp,
	}
	resp, err := a.grpcClient.SignupOtpValidation(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"massage": resp.Result})
	}
}

// UserLogin  godoc
//
//	@Summary		Login
//	@Description	Login for user with otp
//	@Tags			User&Admin Authentication
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			phone	formData	string	true	"Phone No"
//	@Success		200		{object}	entity.Login
//	@Router			/user/loginwithotp [post]
func (a *AuthenticationHandler) LoginWithOtp(c *gin.Context) {
	phone := c.PostForm("phone")
	req := &pb.LoginWithOtpRequest{
		Phone: phone,
	}
	resp, err := a.grpcClient.LoginWithOtp(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Otp send succesfuly to": resp.Phone, "Key": resp.Key})
	}
}

// UserOtpValidation  godoc
//
//	@Summary		Otp Validation
//	@Description	Validating user otp for login validation
//	@Tags			User&Admin Authentication
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			otp		formData	string	true	"Otp"
//	@Param			key		formData	string	true	"Key"
//	@Param			phone	formData	string	false	"phone"
//	@Param			resend	formData	string	false	"resend"
//	@Success		200		{object}	entity.Login
//	@Router			/user/otpvalidation [post]
func (a *AuthenticationHandler) LoginOtpValidation(c *gin.Context) {
	otp := c.PostForm("otp")
	key := c.PostForm("key")
	phone := c.PostForm("phone")
	req := &pb.OtpValidationRequest{
		Key: key,
		Otp: otp,
	}
	resp, err := a.grpcClient.LoginOtpValidation(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	middleware.CreateJwtCookie(int(resp.Userid), phone, "user", c)
	c.JSON(http.StatusOK, gin.H{"massage": "user loged in succesfully and cookie stored"})
}

// UserLogin  godoc
//
//	@Summary		Login
//	@Description	Login for user with password
//	@Tags			User&Admin Authentication
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			phone		formData	string	true	"Phone No"
//	@Param			password	formData	string	true	"Password"
//	@Success		200			{object}	entity.Login
//	@Router			/user/loginwithpassword [post]
func (a *AuthenticationHandler) LoginWithPassword(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	req := &pb.LoginWithPasswordRequest{
		Phone:    phone,
		Password: password,
	}
	resp, err := a.grpcClient.LoginWithPassword(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		middleware.CreateJwtCookie(int(resp.Userid), phone, "user", c)
		c.JSON(http.StatusOK, gin.H{"massage": "user loged in succesfully and cookie stored"})
	}
}

// Admin Register  godoc
//
//	@Summary		registering new admin
//	@Description	Adding new admin to the database
//	@Tags			User&Admin Authentication
//	@Accept			json
//	@Produce		json
//	@Param			admin	body		entity.Admin	true	"Admin Data"
//	@Success		200		{object}	entity.Admin
//	@Router			/admin/registeradmin [post]
func (a *AuthenticationHandler) RegisterAdmin(c *gin.Context) {
	var admin entity.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req := &pb.RegisterAdminRequest{
		Adminname: admin.AdminName,
		Email:     admin.Email,
		Phone:     admin.Phone,
		Password:  admin.Password,
		Role:      admin.Role,
	}
	resp, err := a.grpcClient.RegisterAdmin(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": resp.Result})
	}

}

// Admin Login With Password godoc
//
//	@Summary		Admin Login with password
//	@Description	Admin login with password and phone number
//	@Tags			User&Admin Authentication
//	@Accept			json
//	@Produce		json
//	@Param			admin	body		entity.Login	true	"Admin Data"
//	@Success		200		{object}	entity.Login
//	@Router			/admin/loginpassword [post]
func (a *AuthenticationHandler) AdminLoginWithPassword(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	phone, _ := payload["phone"].(string)
	password, _ := payload["password"].(string)
	req := &pb.LoginWithPasswordRequest{
		Phone:    phone,
		Password: password,
	}
	resp, err := a.grpcClient.AdminLoginWithPassword(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		middleware.CreateJwtCookie(int(resp.Userid), phone, "admin", c)
		c.JSON(http.StatusOK, gin.H{"massage": "admin loged in succesfully and cookie stored"})
	}
}

// Admin Home  godoc
//
//	@Summary		Admin dashbord
//	@Description	Admin dashbord
//	@Tags			User&Admin Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.AdminDashboard
//	@Router			/admin/home [get]
func (a *AuthenticationHandler) AdminDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"options": "Sales Report - User Management - Product Management - Order Management"})
	dashboardResponse := &entity.AdminDashboard{
		TotalUsers:        10050,
		NewUsers:          2000,
		TotalProducts:     200,
		StocklessCategory: "Men",
		TotalOrders:       687,
		AverageOrderValue: 1500,
		PendingOrders:     300,
		ReturnOrders:      15,
		TotalRevenue:      250000,
		TotalQuery:        125,
	}
	c.JSON(http.StatusOK, gin.H{"Dashboard": dashboardResponse})
}

// LogOut     godoc
//
//	@Summary		logout
//	@Description	Deleting cookie from the browser while logout
//	@Tags			User Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"Success message"
//	@Router			/logout [post]
func (a *AuthenticationHandler) Logout(c *gin.Context) {
	err := middleware.DeleteCookie(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user cookie deletion failed"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
	}
}

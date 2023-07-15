package handlers

import (
	"context"
	pb "gateway/pb"
	"gateway/pkg/entity"
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
//	@Tags			User Authentication
//	@Accept			json
//	@Produce		json
//	@Param			userInput	body		models.Signup	true	"User Data"
//	@Success		200			{object}	entity.User
//	@Router			/signup [post]
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

func (a *AuthenticationHandler) SignupWithOtp(c *gin.Context) {

}
func (a *AuthenticationHandler) SignupOtpValidation(c *gin.Context) {

}
func (a *AuthenticationHandler) LoginWithOtp(c *gin.Context) {

}
func (a *AuthenticationHandler) LoginOtpValidation(c *gin.Context) {

}
func (a *AuthenticationHandler) LoginWithPassword(c *gin.Context) {

}
func (a *AuthenticationHandler) RegisterAdmin(c *gin.Context) {

}
func (a *AuthenticationHandler) AdminLoginWithPassword(c *gin.Context) {

}
func (a *AuthenticationHandler) AdminOtpLogin(c *gin.Context) {

}
func (a *AuthenticationHandler) AdminDashboard(c *gin.Context) {

}

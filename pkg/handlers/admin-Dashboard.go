package handlers

import (
	"context"
	pb "gateway/pb"
	"gateway/pkg/entity"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type AdminDashboardHandler struct {
	grpcClient pb.AdminDashboardClient
}

func NewAdminDashboardHandler(cc *grpc.ClientConn) *AdminDashboardHandler {
	client := pb.NewAdminDashboardClient(cc)
	return &AdminDashboardHandler{
		grpcClient: client,
	}
}

// ServiceHealthCheck  godoc
//
//	@Summary		admin dashboard service health check
//	@Description	Service Health Check
//	@Tags			Admin Dashboard
//	@Accept			json
//	@Produce		json
//	@Success		200			string	message
//	@Router			/service2/healthcheck [get]
func (a *AdminDashboardHandler) AdminDashboardIndexHandler(c *gin.Context) {
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

// User Management  godoc
//
//	@Summary		User list
//	@Description	Showing user list for management by admin
//	@Tags			Admin Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			page	query		string	false	"page no"
//	@Param			limit	query		string	false	"limit no"
//	@Success		200		{object}	entity.User
//	@Router			/admin/usermanagement [get]
func (a *AdminDashboardHandler) UserList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limitStr := c.DefaultQuery("limit", "5")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req := &pb.UserListRequest{
		Page:  int32(page),
		Limit: int32(limit),
	}
	resp, err := a.grpcClient.UserList(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Users": resp})
}

// Sort User  godoc
//
//	@Summary		User list by permission
//	@Description	Showing user list for management by sorting with permisiion
//	@Tags			Admin Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			page		query		string	false	"page no"
//	@Param			limit		query		string	false	"limit no"
//	@Param			permission	query		string	true	"true/false"
//	@Success		200			{object}	entity.User
//	@Router			/admin/sortuser [get]
func (a *AdminDashboardHandler) SortUserByPermission(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limitStr := c.DefaultQuery("limit", "5")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	permission := c.Query("permission")
	req := &pb.SortUserRequest{
		Page:       int32(page),
		Limit:      int32(limit),
		Permission: permission,
	}
	resp, err := a.grpcClient.SortUserByPermission(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Users": resp})
}

// Search User  godoc
//
//	@Summary		Search user by id or name
//	@Description	Showing user list for management by searching with name or id
//	@Tags			Admin Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	false	"User Name"
//	@Param			id		query		string	false	"User Id"
//	@Success		200		{object}	entity.User
//	@Router			/admin/searchuser [get]
func (a *AdminDashboardHandler) SearchUser(c *gin.Context) {
	name := c.DefaultQuery("name", " ")
	userIdStr := c.DefaultQuery("id", "0")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userId != 0 {
		req := &pb.SearchUserByidRequest{
			Userid: int32(userId),
		}
		resp, err := a.grpcClient.SearchUserByid(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Users": resp})
	} else if name != " " {
		req := &pb.SearchUserBynameRequest{
			Name: name,
		}
		resp, err := a.grpcClient.SearchUserByname(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Users": resp})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{"Error": "invalid entry"})
	}
}

// Toggle User Permission  godoc
//
//	@Summary		block/unblock user
//	@Description	Toggling user permission for block/unblock
//	@Tags			Admin Dashboard
//	@Accept			json
//	@Produce		json
//	@param			id	path		string	true	"User ID"
//	@Success		200	{object}	entity.Admin
//	@Router			/admin/userpermission/{id} [post]
func (a *AdminDashboardHandler) TogglePermission(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := &pb.TogglePermissionRequest{
		Userid: int32(Id),
	}
	resp, err := a.grpcClient.TogglePermission(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"success": resp.Result})
}

// Add Apparel  godoc
//
//	@Summary		Adding new product
//	@Description	Adding new product of category apparel in database
//	@Tags			Admin Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			admin	body		entity.ApparelInput	true	"Apparel Data"
//	@Success		200		{object}	entity.Apparel
//	@Router			/admin/addapparel [post]
func (a *AdminDashboardHandler) CreateApparel(c *gin.Context) {
	var input entity.ApparelInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := &pb.CreateApparelRequest{
		Name:        input.Name,
		Price:       int32(input.Price),
		Image:       input.ImageURL,
		Category:    input.Category,
		Subcategory: input.SubCategory,
	}
	resp, err := a.grpcClient.CreateApparel(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": resp.Result})
	}
}

// Edit Apparel  godoc
//
//	@Summary		Edit existing product data
//	@Description	Edit data of a product in category apparel
//	@Tags			Admin Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			admin	body		entity.Apparel	true	"Apparel Data"
//	@Success		200		{object}	entity.Apparel
//	@Router			/admin/editapparel/{id} [put]
func (a *AdminDashboardHandler) EditApparel(c *gin.Context) {
	var input entity.Apparel
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := &pb.EditApparelResquest{
		Name:        input.Name,
		Price:       int32(input.Price),
		Image:       input.ImageURL,
		Category:    input.Category,
		Subcategory: input.SubCategory,
		Id:          int32(Id),
	}
	resp, err := a.grpcClient.EditApparel(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"success": resp.Result})
}

// Delete Apparel  godoc
//
//	@Summary		Delete existing product from database
//	@Description	Soft deleting the data of a product from database in category apparel
//	@Tags			Admin Dashboard
//	@Accept			json
//	@Produce		json
//	@param			ProductId	query		int	true	"product id"
//	@Success		200			{object}	entity.Apparel
//	@Router			/admin/deleteapparel/{id} [delete]
func (a *AdminDashboardHandler) DeleteApparel(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := &pb.DeleteApparelRequest{
		Id: int32(Id),
	}
	resp, err := a.grpcClient.DeleteApparel(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": resp.Result})
	}
}

// Add Coupon   godoc
//
//	@Summary		Adding coupon by admin
//	@Description	Addig coupon for users, with a unique code
//	@Tags			Admin Dashboard
//	@Accept			json
//	@Param			admin	body		entity.Coupon	true	"coupon"
//	@Success		200		{string}	string			"Success masage"
//	@Router			/admin/addcoupon  [post]
func (a *AdminDashboardHandler) AddCoupon(c *gin.Context) {
	var coupon entity.Coupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := &pb.AddCouponRequest{
		Code:     coupon.Code,
		Type:     coupon.Type,
		Amount:   int32(coupon.Amount),
		Limit:    int32(coupon.UsageLimit),
		Category: coupon.Category,
	}
	resp, err := a.grpcClient.AddCoupon(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": resp.Result})
	}
}

// Add Offer   godoc
//
//	@Summary		Adding offer by admin
//	@Description	Addig coupon for users, with a unique code
//	@Tags			Admin Dashboard
//	@Accept			json
//	@Param			admin	body		entity.Offer	true	"offer"
//	@Success		200		{string}	string			"Success masage"
//	@Router			/admin/addoffer  [post]
func (a *AdminDashboardHandler) AddOffer(c *gin.Context) {
	var offer entity.Offer
	if err := c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := &pb.AddOfferRequest{
		Code:     offer.Name,
		Type:     offer.Type,
		Amount:   int32(offer.Amount),
		Limit:    int32(offer.UsageLimit),
		Minprice: int32(offer.MinPrice),
		Category: offer.Category,
	}
	resp, err := a.grpcClient.AddOffer(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": resp.Result})
	}
}

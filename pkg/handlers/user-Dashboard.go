package handlers

import (
	"context"
	pb "gateway/pb"
	"gateway/pkg/entity"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
)

type UserDashboardHandler struct {
	grpcClient pb.UserDashboardClient
}

func NewUserDashboardHandler(cc *grpc.ClientConn) *UserDashboardHandler {
	client := pb.NewUserDashboardClient(cc)
	return &UserDashboardHandler{
		grpcClient: client,
	}
}

// ServiceHealthCheck  godoc
//
//	@Summary		user dashboard service health check
//	@Description	Service Health Check
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Success		200			string	message
//	@Router			/service3/healthcheck [get]
func (a *UserDashboardHandler) UserDashboardHealthCheck(c *gin.Context) {
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

// User Home  godoc
//
//	@Summary		User Home
//	@Description	User home with the next navigations
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"Success message"
//	@Router			/home [get]
func (a *UserDashboardHandler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"options": "Logout - Add_Address - Cart - Orders - Tickets - Apparels"})
}

// User Add Address  godoc
//
//	@Summary		Add Address
//	@Description	Add new address to the database with user id
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entity.Address	true	"User Address"
//	@Success		200		{string}	string			"Success message"
//	@Router			/addaddress [post]
func (a *UserDashboardHandler) AddAddress(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	var address entity.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	address.UserId = userId
	req := &pb.AddAddressRequest{
		House:   address.House,
		City:    address.City,
		Street:  address.Street,
		Pincode: int32(address.Pincode),
		Type:    address.Type,
	}
	resp, err := a.grpcClient.AddAddress(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"massage": resp.Result})
	}
}

// User Details    godoc
//
//	@Summary		User Details
//	@Description	User profile with adress and user details
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.User
//	@Router			/userdetails [get]
func (a *UserDashboardHandler) ShowUserDetails(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	req := &pb.UserDetailsRequset{
		Userid: int32(userId),
	}
	resp, err := a.grpcClient.UserDetails(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Details": resp.User, "address": resp.Address})
}

// Apparels       godoc
//
//	@Summary		Apparel List
//	@Description	Showing the available Apparels in the site
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int				false	"page no"
//	@Param			limit	query		int				false	"limit no"
//	@Param			sort	query		string			false	"Sort by Category"
//	@Success		200		{object}	entity.Apparel	"Apparel List"
//	@Router			/apparels [get]
func (a *UserDashboardHandler) Apparels(c *gin.Context) {
	category := c.Query("sort")
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}
	limitStr := c.DefaultQuery("limit", "5")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}
	req := &pb.ApparelsRequest{
		Page:     int32(page),
		Limit:    int32(limit),
		Category: category,
	}
	resp, err := a.grpcClient.Apparels(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Apparel list not found"})
		return
	}
	responseList := make([]entity.Apparel, len(resp.Apparels))
	for i, apparel := range resp.Apparels {
		responseList[i] = entity.Apparel{
			ID:          int(apparel.Id),
			Name:        apparel.Name,
			Price:       int(apparel.Price),
			ImageURL:    apparel.Image,
			SubCategory: apparel.Subcategory,
		}
	}
	c.JSON(http.StatusOK, gin.H{"Apperals": responseList})
}

// Apparel Details  godoc
//
//	@Summary		Details of a Apparel
//	@Description	Showing details of a single product and option to adding cart
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			apparelid	path		string					true	"Apparel ID"
//	@Success		200			{object}	entity.ApparelDetails	"Apparel Details"
//	@Router			/appareldetails/{apparelid} [get]
func (a *UserDashboardHandler) ApparelDetails(c *gin.Context) {
	id := c.Param("apparelid")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
	}
	req := &pb.ApparelDetailsRequest{
		Id: int32(Id),
	}
	resp, err := a.grpcClient.ApparelDetails(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Apparel": resp})
}

// Search Apparels       godoc
//
//	@Summary		Search Result
//	@Description	Showing the available apparels as per user search
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int				false	"page no"
//	@Param			limit	query		int				false	"limit no"
//	@Param			search	query		string			false	"Search By Name"
//	@Success		200		{object}	entity.Apparel	"Apparel Data"
//	@Router			/searchapparel [get]
func (a *UserDashboardHandler) SearchApparels(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "5")
	search := c.Query("search")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}
	req := &pb.SearchApparelsRequest{
		Page:   int32(page),
		Limit:  int32(limit),
		Search: search,
	}
	resp, err := a.grpcClient.SearchApparels(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Apparel list not found"})
		return
	}
	responseList := make([]entity.Apparel, len(resp.Apparels))
	for i, apparel := range resp.Apparels {
		responseList[i] = entity.Apparel{
			ID:          int(apparel.Id),
			Name:        apparel.Name,
			Price:       int(apparel.Price),
			ImageURL:    apparel.Image,
			SubCategory: apparel.Subcategory,
		}
	}
	c.JSON(http.StatusOK, gin.H{"Apparels": responseList})
}

// Add to cart  godoc
//
//	@Summary		Add product to cart
//	@Description	Adding product with quantity to cart with product id
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			productid	path		string	true	"Product ID"
//	@Param			quantity	path		string	true	"Product Quantity"
//	@Success		200			{string}	string	"Success message"
//	@Router			/addtocart/{category}/{productid}/{quantity} [post]
func (a *UserDashboardHandler) AddToCart(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	strId := c.Param("productid")
	strQuantity := c.Param("quantity")
	Id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	quantity, err := strconv.Atoi(strQuantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	req := &pb.AddToCartRequest{
		Productid: int32(Id),
		Quantity:  int32(quantity),
		Userid:    int32(userId),
	}
	resp, err := a.grpcClient.AddToCart(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp.Result})
}

// Add to wishlist  godoc
//
//	@Summary		Add product to wishlist
//	@Description	Adding single product to wishlist with product id
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			category	path		string	true	"Ticket/Apparel"
//	@Param			productid	path		string	true	"Product ID"
//	@Success		200			{string}	string	"Success message"
//	@Router			/addtowishlist/{category}/{productid} [post]
func (a *UserDashboardHandler) AddToWishlist(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	strId := c.Param("productid")
	Id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	req := &pb.AddToWishListRequest{
		Productid: int32(Id),
		Userid:    int32(userId),
	}
	resp, err := a.grpcClient.AddToWishList(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": resp.Result})
}

// Cart     godoc
//
//	@Summary		User Cart
//	@Description	Showing user cart
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.Cart	"User Cart"
//	@Router			/usercart [get]
func (a *UserDashboardHandler) Cart(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	var userCartResponse entity.Cart
	req := &pb.CartRequest{
		Userid: int32(userId),
	}
	resp, err := a.grpcClient.Cart(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	copier.Copy(&userCartResponse, &resp)
	c.JSON(http.StatusOK, gin.H{"User Cart": userCartResponse})
}

// Cart List    godoc
//
//	@Summary		Cart List
//	@Description	Showing the products in user cart
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.CartItem	"Cart List"
//	@Router			/usercartlist [get]
func (a *UserDashboardHandler) CartList(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	req := &pb.CartListRequest{
		Userid: int32(userId),
	}
	resp, err := a.grpcClient.CartList(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Cart List": resp})
}

// RemoveFromCart godoc
//
//	@Summary		Remove Product from cart
//	@Description	Removing product from the cart for unique and decrese quantity for existing product
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"Product ID"
//	@Success		200		{string}	string	"Success message"
//	@Router			/removefromcart/{product}/{id} [delete]
func (a *UserDashboardHandler) RemoveFromCart(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	req := &pb.RemoveFromCartRequest{
		Userid:    int32(userId),
		Productid: int32(Id),
	}
	resp, err := a.grpcClient.RemoveFromCart(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": resp.Result})
}

// RemoveFromWishlist godoc
//
//	@Summary		Remove Product from wishlist
//	@Description	Removing product from the user wishlist
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			product	path		string	true	"ticket/apparel"
//	@Param			id		path		int		true	"Product ID"
//	@Success		200		{string}	string	"Success message"
//	@Router			/removefromwishlist/{product}/{id} [delete]
func (a *UserDashboardHandler) RemoveFromWishlist(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	req := &pb.RemoveFromWishlistRequest{
		Userid:    int32(userId),
		Productid: int32(Id),
	}
	resp, err := a.grpcClient.RemoveFromWishlist(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": resp.Result})
}

// Wishlist    godoc
//
//	@Summary		Wish List
//	@Description	Showing the products in user wishlist
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.Wishlist	"Wishlist"
//	@Router			/userwishlist [get]
func (a *UserDashboardHandler) ViewWishlist(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	req := &pb.WishlistRequest{
		Userid: int32(userId),
	}
	resp, err := a.grpcClient.Wishlist(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Wish List": resp})
}

// Available Coupon  godoc
//
//	@Summary		checking coupon availability
//	@Description	showing the available coupons and eligibility
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.Coupon	"Available coupons"
//	@Router			/coupons [get]
func (a *UserDashboardHandler) AvailableCoupons(c *gin.Context) {
	resp, err := a.grpcClient.AvailableCoupons(context.Background(), nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Available coupons are ": resp})
}

// Apply Coupon  godoc
//
//	@Summary		checking coupon availability and adding offer amount
//	@Description	applying coupon offer for user cart amount
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string	true	"coupon code"
//	@Success		200		{string}	string	"total amount"
//	@Router			/applycoupon/{code} [post]
func (a *UserDashboardHandler) ApplyCoupon(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	code := c.Param("code")
	req := &pb.ApplyCouponRequest{
		Userid: int32(userId),
		Code:   code,
	}
	resp, err := a.grpcClient.ApplyCoupon(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Offer price for coupon is ": resp.Result, "Offer": "Applied succesfuly"})
	}
}

// Offer Check godoc
//
//	@Summary		checking offer availability
//	@Description	finding and showing offer for user with respect to user cart
//	@Tags			User Dashboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"offers"
//	@Router			/offer [get]
func (a *UserDashboardHandler) OfferCheck(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	req := &pb.OfferCheckRequest{
		Userid: int32(userId),
	}
	resp, err := a.grpcClient.OfferCheck(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if resp == nil {
		c.JSON(http.StatusOK, gin.H{"No offers": "Add few more products"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Available offers are ": resp})
	}
}

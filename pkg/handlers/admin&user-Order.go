package handlers

import (
	"context"
	pb "gateway/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type OrderHandler struct {
	grpcClient pb.OrderClient
}

func NewOrderHandler(cc *grpc.ClientConn) *OrderHandler {
	client := pb.NewOrderClient(cc)
	return &OrderHandler{
		grpcClient: client,
	}
}

// ServiceHealthCheck  godoc
//
//	@Summary		order service health check
//	@Description	Service Health Check
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@Success		200			string	message
//	@Router			/service4/healthcheck [get]
func (o *OrderHandler) OrderIndexHandler(c *gin.Context) {
	req := &pb.Request{
		Data: "Mydata",
	}
	resp, err := o.grpcClient.MyMethod(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call MyMethod: %v", err)
	} else {
		c.JSON(http.StatusOK, resp.Result)
	}
}

// Place Order   godoc
//
//	@Summary		Place Order
//	@Description	Placing order from user side with respect to the payment method
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@Param			addressid	path		string	true	"address id"
//	@Param			payment		path		string	true	"payment method"
//	@Success		200			{string}	string	"Success message"
//	@Router			/placeorder/{addressid}/{payment} [post]
func (o *OrderHandler) PlaceOrder(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	straddress := c.Param("addressid")
	paymentMethod := c.Param("payment")
	addressId, err := strconv.Atoi(straddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	if paymentMethod == "cod" {
		req := &pb.PlaceOrderRequest{
			Userid:        int32(userId),
			Addressid:     int32(addressId),
			Paymentmethod: paymentMethod,
		}
		resp, err := o.grpcClient.PlaceOrder(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"massage": "Order placed successfully", "Invoice": resp})
		}

	} else if paymentMethod == "paypal" {
		req := &pb.PlaceOrderRequest{
			Userid:        int32(userId),
			Addressid:     int32(addressId),
			Paymentmethod: paymentMethod,
		}
		resp, err := o.grpcClient.PlaceOrder(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {

			c.JSON(http.StatusOK, gin.H{"massage": "Order placed successfully", "Invoice": resp})
		}

	} else if paymentMethod == "razorpay" {
		req := &pb.PlaceOrderRequest{
			Userid:        int32(userId),
			Addressid:     int32(addressId),
			Paymentmethod: paymentMethod,
		}
		resp, err := o.grpcClient.PlaceOrder(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"massage": resp})
		}

	} else if paymentMethod == "wallet" {
		req := &pb.PlaceOrderRequest{
			Userid:        int32(userId),
			Addressid:     int32(addressId),
			Paymentmethod: paymentMethod,
		}
		resp, err := o.grpcClient.PlaceOrder(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"massage": "Order placed successfully", "Invoice": resp})
		}
	}
}

// Payment Verification   godoc
//
//	@Summary		Payment Verification
//	@Description	After placing order - checking the status of online payment
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	string	"Success message"
//	@Param			sign	path		string	true	"Payment signature"
//	@Param			razorid	path		string	true	"Razor Order Id"
//	@Param			payid	path		string	true	"Razor Payment Id"
//	@Router			/paymentverification/{sign}/{razorid}/{payid} [post]
func (o *OrderHandler) PaymentVerification(c *gin.Context) {
	Signature := c.Param("sign")
	razorId := c.Param("razorid")
	paymentId := c.Param("payid")
	req := &pb.PaymentVerificationRequest{
		Signature: Signature,
		Razorid:   razorId,
		Paymentid: paymentId,
	}
	resp, err := o.grpcClient.PaymentVerification(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		resp.Paymentid = paymentId
		c.JSON(http.StatusAccepted, gin.H{"massage": "Payment successful", "invoice": resp})
	}
}

// Order Cancelation   godoc
//
//	@Summary		Order Cancelation
//	@Description	canceling the order from user side and admin side
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@Param			orderid	path		string	true	"Order Id"
//	@Success		200		{string}	string	"Success message"
//	@Router			/cancelorder/{orderid} [put]
func (o *OrderHandler) CancelOrder(c *gin.Context) {
	strOrderId := c.Param("orderid")
	orderId, err := strconv.Atoi(strOrderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	req := &pb.CancelOrderRequest{
		Orderid: int32(orderId),
	}
	resp, err := o.grpcClient.CancelOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": resp})
}

// Order History   godoc
//
//	@Summary		Order History
//	@Description	showing the history of orders to the user
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"Success message"
//	@Router			/orderhistory [get]
func (o *OrderHandler) OrderHistory(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
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
	req := &pb.OrderHistoryRequest{
		Userid: int32(userId),
		Page:   int32(page),
		Limit:  int32(limit),
	}
	resp, err := o.grpcClient.OrderHistory(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Orders": resp})
}

// Order return godoc
//
//	@Summary		Return delivered order
//	@Description	Returning the orders which are delivered to the user
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@Param			orderid	path		string	true	"orderid"
//	@Param			reason	path		string	true	"wrong product-product quality-late delivery-other"
//	@Success		200		{string}	string	"Success massage"
//	@Router			/orderreturn/{orderid} [post]
func (o *OrderHandler) OrderReturn(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	strOrderId := c.Param("orderid")
	reason := c.Param("reason")
	orderId, err := strconv.Atoi(strOrderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}

	req := &pb.OrderReturnRequest{
		Userid:  int32(userId),
		Orderid: int32(orderId),
		Reason:  reason,
		Status:  "Initiated",
	}
	resp, err := o.grpcClient.OrderReturn(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": resp.Result})
}

// Order Update  godoc
//
//	@Summary		Update order status
//	@Description	Updating the order status by admin
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@Param			orderid	path		string	true	"Order Id"
//	@Param			status	path		string	true	"status"
//	@Success		200		{string}	string	"Success massage"
//	@Router			/updateorder/{orderid}/{status} [put]
func (o *OrderHandler) AdminOrderUpdate(c *gin.Context) {
	strOrderId := c.Param("orderid")
	orderId, err := strconv.Atoi(strOrderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	strStatus := c.Param("status")
	req := &pb.AdminOrderUpdateRequest{Orderid: int32(orderId), Status: strStatus}
	resp, err := o.grpcClient.AdminOrderUpdate(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": resp.Result})
}

// Update Return    godoc
//
//	@Summary		Updating return status and refund
//	@Description	Updating the retunr status by admin and implimenting refund
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@Param			returnid	path		string	true	"Return Id"
//	@Param			status		path		string	true	"status"
//	@Param			refund		path		string	true	"refund method -wallet -account"
//	@Success		200			{string}	string	"updated succesfuly"
//	@Router			/updatereturn/{returnid}/{status}/{refund} [post]
func (o *OrderHandler) AdminReturnUpdate(c *gin.Context) {
	status := c.Param("status")
	refund := c.Param("refund")
	strReturnId := c.Param("returnid")
	returnId, err := strconv.Atoi(strReturnId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	if refund == "wallet" {
		req := &pb.AdminReturnUpdateRequest{Status: status, Refund: refund, Returnid: int32(returnId)}
		resp, err := o.grpcClient.AdminReturnUpdate(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": resp.Result})
		}
	}
}

// Aproving Refund godoc
//
//	@Summary		Aproving refund by admin
//	@Description	Transfering the total amount of order to wallet or other methods
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@Param			orderid	path		string	true	"order return id"
//	@Success		200		{string}	string	"Success massage"
//	@Router			/refund/{orderid} [post]
func (o *OrderHandler) AdminRefund(c *gin.Context) {
	strReturnId := c.Param("orderid")
	orderId, err := strconv.Atoi(strReturnId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	req := &pb.AdminRefundRequest{Orderid: int32(orderId)}
	resp, err := o.grpcClient.AdminRefund(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"massage": resp.Result})
	}
}

// Sales report by date  godoc
//
//	@Summary		Sales report by date
//	@Description	Showing the sales report with respect to the given date
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@Param			start	path	string				true	"start date D-M-Y"
//	@Param			end		path	string				true	"end date D-M-Y"
//	@Success		200		body	entity.SalesReport	"report"
//	@Router			/salesreportbydate/{start}/{end} [get]
func (o *OrderHandler) SalesReportByDate(c *gin.Context) {
	// startDateStr := c.Param("start")
	// endDateStr := c.Param("end")
	// startDate, err := time.Parse("2-1-2006", startDateStr)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
	// 	return
	// }
	// endDate, err := time.Parse("2-1-2006", endDateStr)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
	// 	return
	// }
	// report, err := oh.OrderUsecase.ExecuteSalesReportByDate(startDate, endDate)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }
	c.JSON(http.StatusOK, gin.H{"report": "report"})
}

// Sales report by period godoc
//
//	@Summary		Sales report by time period
//	@Description	Showing the report of sales for last week,month and year
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@param			period	path		string	true	"Weekly Monthlt Yearly"
//	@Success		200		{string}	string	"Success message"
//	@Router			/salesreportbyperiod/{period} [get]
func (o *OrderHandler) SalesReportByPeriod(c *gin.Context) {
	period := c.Param("period")

	req := &pb.SalesReportByPeriodRequest{Period: period}
	resp, err := o.grpcClient.SalesReportByPeriod(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"report": resp})
}

// Sales report by category
//
//	@Summary		Sales report by product category
//	@Description	Showing the report of sales with respect to product category
//	@Tags			Admin&User Order Management
//	@Accept			json
//	@Produce		json
//	@Param			category	path	string				true	"Category"
//	@Param			period		path	string				true	"period"
//	@Success		200			body	entity.SalesReport	"report"
//	@Router			/salesreportbycategory/{category}/{period} [get]
func (o *OrderHandler) SalesReportByCategory(c *gin.Context) {
	category := c.Param("category")
	period := c.Param("period")
	req := &pb.SalesReportByCategoryRequest{Category: category, Period: period}
	resp, err := o.grpcClient.SalesReportByCategory(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"report": resp})
}

// Sort  order by status
//
//	@Summary		Sorting orders by order status
//	@Description	Showing the sorted list of orders in admin panel
//	@Tags			Admin&User Order Management
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			status	formData	string			true	"Status"
//	@Success		200		{object}	entity.Order	"sorted orders"
//	@Router			/sortorders [post]
func (o *OrderHandler) SortOrderByStatus(c *gin.Context) {
	status := c.PostForm("status")
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
	req := &pb.SortOrderByStatusRequest{Status: status, Page: int32(page), Limit: int32(limit)}
	resp, err := o.grpcClient.SortOrderByStatus(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"Orders": resp})
	}
}

package models

type Cart struct {
	UserId          int     `json:"-"`
	ApparelQuantity int     `json:"apparelquantity"`
	TicketQuantity  int     `json:"ticketquantity"`
	TotalPrice      float64 `json:"totalprice"`
	OfferPrice      int     `json:"offerprice"`
}

type CartItem struct {
	CartId      int     `json:"-"`
	Category    string  `json:"category"`
	ProductId   int     `json:"productid"`
	Quantity    int     `json:"quantity"`
	ProductName string  `json:"productname"`
	Price       float64 `json:"price"`
}

type Wishlist struct {
	UserId      int     `json:"-"`
	Category    string  `json:"category"`
	ProductId   int     `json:"productid"`
	ProductName string  `json:"productname"`
	Price       float64 `json:"price"`
}

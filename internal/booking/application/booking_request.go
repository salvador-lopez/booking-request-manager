package application

type BookingRequest struct {
	RequestId   string `json:"request_id" binding:"required"`
	CheckIn     string `json:"check_in" binding:"required"`
	Nights      int    `json:"nights" binding:"required"`
	SellingRate int    `json:"selling_rate" binding:"required"`
	Margin      int    `json:"margin" binding:"required"`
}

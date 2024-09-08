package model

type CallbackRequest struct {
	ID                   uint   `json:"id" gorm:"primaryKey"`
	MerchantRequestID    string `json:"MerchantRequestID"`
	CheckoutRequestID    string `json:"CheckoutRequestID"`
	ResultCode           string `json:"ResultCode"`
	ResultDesc           string `json:"ResultDesc"`
	TransactionType      string `json:"transactionType"`
	TransactionReference string `json:"transactionReference"`
	PhoneNumber          string `json:"phoneNumber"`
	Amount               string `json:"amount"`
	TransactionDate      string `json:"transactionDate"`
	BusinessID           string `json:"businessId"`
	CallbackString       string `json:"callbackString"`
	OrderID              string `json:"orderId"`
}
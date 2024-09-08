package interfacex

import "time"

type STKPushRequest struct {
	ID                  uint `gorm:"primaryKey"`
	Amount              string
	PhoneNumber         string
	AccountReference    string
	TransactionDesc     string
	TransactionType     string
	CallBackURL         string
	BusinessShortCode   string
	PartyA              string
	PartyB              string
	CheckoutRequestID   string
	MerchantRequestID   string
	ResponseCode        string
	ResponseDescription string
	CustomerMessage     string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type STKPushResponse struct {
	MerchantRequestID   string
	CheckoutRequestID   string
	ResponseCode        string
	ResponseDescription string
	CustomerMessage     string
}

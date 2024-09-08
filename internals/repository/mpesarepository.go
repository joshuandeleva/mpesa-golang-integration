package repository

import (
	"context"

	"github.com/go-mpesa-integration/interfacex"
	"github.com/go-mpesa-integration/internals/model"
	"gorm.io/gorm"
)

type MpesaRepository interface {
	SaveSTKPush(ctx context.Context, request *interfacex.STKPushRequest, response *interfacex.STKPushResponse) error
	SaveCallbackRequest(request *model.CallbackRequest) error

}

type mpesaRepository struct {
	db *gorm.DB
}

func NewMpesaRepository(db *gorm.DB) *mpesaRepository {
	return &mpesaRepository{
		db: db,
	}
}

func (r *mpesaRepository) SaveSTKPush(ctx context.Context, request *interfacex.STKPushRequest, response *interfacex.STKPushResponse) error {
	stkPushRequest := &interfacex.STKPushRequest{
		Amount:            request.Amount,
        PhoneNumber:       request.PhoneNumber,
        AccountReference:  request.AccountReference,
        TransactionDesc:   request.TransactionDesc,
        TransactionType:   request.TransactionType,
		CallBackURL:       request.CallBackURL,
        BusinessShortCode: request.BusinessShortCode,
        PartyA:            request.PartyA,
        PartyB:            request.PartyB,
		CheckoutRequestID: response.CheckoutRequestID,
		MerchantRequestID: response.MerchantRequestID,
		ResponseCode:      response.ResponseCode,
		ResponseDescription: response.ResponseDescription,
		CustomerMessage:   response.CustomerMessage,
	}

	if err := r.db.WithContext(ctx).Create(stkPushRequest).Error; err != nil {
		return err
	}
	return nil
}

func (r *mpesaRepository) SaveCallbackRequest(request *model.CallbackRequest) error {
	if err := r.db.Create(request).Error; err != nil {
		return err
	}
	return nil
}
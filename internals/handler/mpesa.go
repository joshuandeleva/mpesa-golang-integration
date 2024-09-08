package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-mpesa-integration/interfacex"
	"github.com/go-mpesa-integration/internals/model"
	"github.com/go-mpesa-integration/internals/services"
	"github.com/sirupsen/logrus"
)

type MpesaHandler struct {
	mpesaService services.MpesaService

}

func NewMpesaHandler(mpesaService services.MpesaService) *MpesaHandler {
	return &MpesaHandler{
		mpesaService: mpesaService,
	}
}

func (h *MpesaHandler) InitiateSTKPush(c *gin.Context) {
	var request interfacex.STKPushRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logrus.Errorf("error binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response , err := h.mpesaService.STKPush(c.Request.Context(), &request)
	if err != nil {
		logrus.Errorf("error initiating STK push: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *MpesaHandler) STKCallback(c *gin.Context) {
    // Get query parameters from URL
    businessNumber := c.Query("businessnumber") 
    orderID := c.Query("reference")

    logrus.Infof("Callback received - BusinessNumber: %s, OrderID: %s", businessNumber, orderID)

    var callbackResponse map[string]interface{}
    if err := c.ShouldBindJSON(&callbackResponse); err != nil {
        logrus.Errorf("Error binding request: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    body, ok := callbackResponse["Body"].(map[string]interface{})
    if !ok {
        logrus.Error("Invalid callback body format")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid callback body format"})
        return
    }

    stkCallback, ok := body["stkCallback"].(map[string]interface{})
    if !ok {
        logrus.Error("Invalid stkCallback format")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stkCallback format"})
        return
    }

    merchantRequestID, _ := stkCallback["MerchantRequestID"].(string)
    checkoutRequestID, _ := stkCallback["CheckoutRequestID"].(string)
    resultCode, _ := stkCallback["ResultCode"].(string)
    resultDesc, _ := stkCallback["ResultDesc"].(string)

    var amount, mpesaReceiptNumber, transactionDate, phoneNumber string

    callbackMetadata, ok := stkCallback["CallbackMetadata"].(map[string]interface{})
    if ok {
        items, _ := callbackMetadata["Item"].([]interface{})
        for _, item := range items {
            itemMap, _ := item.(map[string]interface{})
            name, _ := itemMap["Name"].(string)
            value := itemMap["Value"]
            switch name {
            case "Amount":
                amount = fmt.Sprintf("%v", value)
            case "MpesaReceiptNumber":
                mpesaReceiptNumber = fmt.Sprintf("%v", value)
            case "TransactionDate":
                transactionDate = fmt.Sprintf("%v", value)
            case "PhoneNumber":
                phoneNumber = fmt.Sprintf("%v", value)
            }
        }
    }

    // save the callback data to the database

    callbackRequest := &model.CallbackRequest{
        MerchantRequestID:    merchantRequestID,
        CheckoutRequestID:    checkoutRequestID,
        ResultCode:           resultCode,
        ResultDesc:           resultDesc,
        Amount:               amount,
        TransactionDate:      transactionDate,
        PhoneNumber:          phoneNumber,
        OrderID:              orderID,
        BusinessID:           businessNumber,
        CallbackString:       fmt.Sprintf("%v", stkCallback),
		TransactionReference: mpesaReceiptNumber,

    }
    h.mpesaService.SaveCallbackRequest(callbackRequest);

    if resultCode == "0" {
        logrus.Infof("STK Push successful. MerchantRequestID: %s, CheckoutRequestID: %s, Amount: %s, MpesaReceiptNumber: %s, TransactionDate: %s, PhoneNumber: %s, OrderID: %s, BusinessNumber: %s",
            merchantRequestID, checkoutRequestID, amount, mpesaReceiptNumber, transactionDate, phoneNumber, orderID, businessNumber)
        c.JSON(http.StatusOK, gin.H{"message": "STK Push successful"})
    } else {
        // Failure
        logrus.Errorf("STK Push failed. MerchantRequestID: %s, CheckoutRequestID: %s, ResultCode: %s, ResultDesc: %s",
            merchantRequestID, checkoutRequestID, resultCode, resultDesc)
        c.JSON(http.StatusBadRequest, gin.H{"message": "STK Push failed"})
    }
}
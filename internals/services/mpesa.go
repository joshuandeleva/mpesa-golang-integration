package services

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-mpesa-integration/config"
	"github.com/go-mpesa-integration/interfacex"
	"github.com/go-mpesa-integration/internals/model"
	"github.com/go-mpesa-integration/internals/repository"
	"github.com/go-mpesa-integration/utils"
	"github.com/sirupsen/logrus"
)

type MpesaService interface {
	STKPush(ctx context.Context, request *interfacex.STKPushRequest) (*interfacex.STKPushResponse, error)
	SaveCallbackRequest(request *model.CallbackRequest) 

}

type mpesaService struct {
	mpesaRepo repository.MpesaRepository
}

func NewMpesaService(mpesaRepo repository.MpesaRepository) MpesaService {
	return &mpesaService{
		mpesaRepo: mpesaRepo,
	}
}

func (s *mpesaService) STKPush(ctx context.Context, request *interfacex.STKPushRequest) (*interfacex.STKPushResponse, error) {
	enConfig := config.NewEnvConfig()
	token, err := utils.GenerateAccessToken(enConfig.ConsumerKey, enConfig.ConsumerSecret)

	if err != nil {
		logrus.Errorf("Error generating access token: %v", err)
		return nil, err
	}

	//generate timestamps && password
	timeStamp := utils.GenerateTimeStamp()
	password := utils.GeneratePassword(enConfig.ShortCode, enConfig.PassKey, timeStamp)



	// callback url -> /api/v1/stkcallback
	callbackURL := "https://touching-quagga-tightly.ngrok-free.app/api/v1/stkcallback?businessnumber=" + url.QueryEscape(strings.TrimSpace(enConfig.ShortCode)) + "&reference=" + url.QueryEscape(strings.TrimSpace(request.AccountReference))

	logrus.Infof("Generated callback URL: %s", callbackURL)

	
	// request body
	reqPayload := map[string]interface{}{
		"BusinessShortCode": enConfig.ShortCode,
		"Password":          password,
		"Timestamp":         timeStamp,
		"TransactionType":   "CustomerPayBillOnline",
		"Amount":            request.Amount,
		"PartyA":            request.PhoneNumber,
		"PartyB":            enConfig.ShortCode,
		"PhoneNumber":       request.PhoneNumber,
		"CallBackURL":    	 callbackURL,
		"AccountReference":  request.AccountReference,
		"TransactionDesc":   request.TransactionDesc,
	}
	// make a post request to the mpesa api

	url := enConfig.MpesaUrl
	reqBody, err := json.Marshal(reqPayload)

	if err != nil {
		logrus.Errorf("Error marshaling request body: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))

	if err != nil {
		logrus.Errorf("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res , err := client.Do(req)

	if err != nil {
		logrus.Errorf("Error making request: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	var response interfacex.STKPushResponse

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		logrus.Errorf("Error decoding response: %v", err)
		return nil, err
	}

	//save transaction to db using repository

	if err := s.mpesaRepo.SaveSTKPush(ctx, request, &response); err != nil {
		logrus.Errorf("Error saving transaction to db: %v", err)
		return nil, err
	}
	return &response, nil
}

func (s *mpesaService) SaveCallbackRequest(request *model.CallbackRequest)  {
	go func(){
		if err := s.mpesaRepo.SaveCallbackRequest(request); err != nil {
			logrus.Errorf("Error saving callback request: %v", err)
		}else {
            logrus.Info("Callback request saved successfully")
        }
	}()
}
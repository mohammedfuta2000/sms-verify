package api

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var accountSID, authToken, serviceID = env("cmd/.env")
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: accountSID,
	Password: authToken,
})


func (app *Config) twilioSendOTP(phone_number string) (string,error)  {
	params:= &twilioApi.CreateVerificationParams{}
	params.SetTo(phone_number)
	params.SetChannel("sms")

	resp, err:=client.VerifyV2.CreateVerification(serviceID,params)
	if err!=nil {
		return "",err
	} 
	return *resp.Sid,nil
}

func (app *Config) twilioVerifyOTP(phone_number string,code string) (error)  {
	params:= &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phone_number)
	params.SetCode(code)

	resp,err:=client.VerifyV2.CreateVerificationCheck(serviceID,params)
	if err!=nil {
		return err
	} 
	if *resp.Status != "approved"{
		return errors.New("not a valid code")
	}
	return nil
	
}
	
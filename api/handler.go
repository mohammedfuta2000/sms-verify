package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mohammedfuta2000/sms-verify/data"
)

func(app *Config) check()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,jsonResponse{Status: http.StatusOK,Message: "Success" })
	}
}
func(app *Config) sendSMS()gin.HandlerFunc{
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()
		var payload data.OTPData

		app.validateBody(c,&payload)
		fmt.Println("last",payload.PhoneNumber)
		newData:= data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}
		fmt.Println("last",newData.PhoneNumber)
		_, err:= app.twilioSendOTP(newData.PhoneNumber)
		if err!=nil {
			app.errorJson(c,err)
			return
		}

		app.writeJson(c, http.StatusAccepted, "OTP sent successfully")
		


	}
}

func(app *Config) verifySMS()gin.HandlerFunc{
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()
		var payload data.VerifyData

		app.validateBody(c,&payload)

		newData :=data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		err:= app.twilioVerifyOTP(newData.User.PhoneNumber,newData.Code)
		if err!=nil {
			app.errorJson(c,err)
			return
		}

		app.writeJson(c, http.StatusAccepted, "OTP verified successfully")
	}
}
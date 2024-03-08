package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
}

var validate = validator.New()
// validateBody is used to apply the validations u specify in your structs `` backticks
func (app *Config) validateBody(c *gin.Context, data any) error{
	if err:=c.BindJSON(&data); err!=nil{
		return err
	}
	fmt.Println("first",data)
	// ??
	if err:=validate.Struct(&data);  err!=nil{
		return err
	}
	return nil
}

func (app *Config) writeJson(c *gin.Context, status int, data any){
	c.JSON(status, jsonResponse{Status: status,Message: "Success",Data: data })
}

func (app *Config) errorJson(c *gin.Context,err error, status ...int){
	statusCode := http.StatusBadRequest
	if len(status)>0 {
		statusCode=status[0]
	}
	c.JSON(statusCode, jsonResponse{Status: statusCode,Message: err.Error()})
}
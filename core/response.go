package core

import (
	"github.com/agxmaster/atrz/adapter"
	"net/http"
)

type BusinessResponse struct {
	Code    ResponseCode `json:"code"`
	Data    interface{}  `json:"data"`
	Message string       `json:"message"`
}

func Success(c adapter.HertzCtxCore, data interface{}) {
	c.JSON(http.StatusOK, BusinessResponse{Code: Mp.SuccessCode, Data: data})
}

func Error(c adapter.HertzCtxCore, err error) {
	c.JSON(http.StatusOK, BusinessResponse{Code: Mp.ErrorCode, Message: err.Error()})
}

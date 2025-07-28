package handlers

import (
	"net/http"

	utils "github.com/Rus203/shop/util"
	"github.com/gin-gonic/gin"
)

type AppHandler struct {}

func (ap *AppHandler) HealCheck(ctx *gin.Context) {
	utils.WriteJSON(ctx, "Pizza Shop is open", http.StatusOK)
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}
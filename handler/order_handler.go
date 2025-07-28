package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Rus203/shop/constants"
	"github.com/Rus203/shop/service"
	"github.com/Rus203/shop/util"
)

type OrderHandler struct {
	messagePublisher services.IMessagePublisher
}

func (oh *OrderHandler) CreateOrder(ctx *gin.Context) {
	payload := make(map[string]any) // todo: add validation here too

	err := utils.ParseJSON(ctx, payload)

	if err != nil {
		utils.WriteErrorJSON(ctx, http.StatusBadRequest, err)
	}

	payload["order_status"] = constants.ORDER_ORDERED

	if err := oh.messagePublisher.PublishEvent(constants.KITCHEN_ORDER_QUEUE, payload); err != nil {
		utils.WriteErrorJSON(ctx, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(ctx, map[string]any{
		"message": "order accepted successfully",
		"data":    payload,
	}, http.StatusCreated)
}

func NewOrderHandler(messagePublisher services.IMessagePublisher) *OrderHandler {
	return &OrderHandler{
		messagePublisher: messagePublisher,
	}
}

package controller

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"paymentService/service"
)

type PaymentController struct {
	Service service.PaymentService
}

// @Summary Get Payment by OrderID
// @Description Get a payment
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path string true "order ID"
// @Success 200 {object} model.Payment
// @Router /payments/{order-id} [get]
func (oc *PaymentController) GetPaymentByOrderID(ctx iris.Context) {
	idParam := ctx.Params().Get("order-id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid order ID"})
		return
	}

	order, err := oc.Service.GetPaymentByOrderId(id)
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"message": fmt.Sprintf("payment not found with OrderID: %s", id)})
		return
	}

	ctx.JSON(order)
}

// @Summary Get all Payments
// @Description Get a list of model.Payment
// @Tags Payment
// @Accept json
// @Produce json
// @Success 200 {array} model.Payment
// @Router /payments [get]
func (oc *PaymentController) GetPayments(ctx iris.Context) {
	result, err := oc.Service.GetPayments()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Internal server error"})
	}

	if result == nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"message": "No payments found"})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "payments retrieved successfully", "payments": result})
}

package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	httpSwagger "github.com/swaggo/http-swagger"
	"paymentService/config"
	"paymentService/controller"
	_ "paymentService/docs"
	"paymentService/repository"
	"paymentService/service"
)

// @title Payment Service API
// @version 1.0
// @description This is a Simple application for checking payments status
// @BasePath /
func main() {
	app := iris.New()
	db, error := config.ConnectToDB()
	if error != nil {
		fmt.Printf("Connection Lost")
		return
	}

	repo := &repository.PaymentRepository{DB: db}
	service := &service.PaymentService{Repo: repo}
	paymentHandler := &controller.PaymentController{Service: *service}

	config.CheckForPublishedPayments(repo)

	app.Get("/swagger/{any:path}", iris.FromStd(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)))

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello, World from payments!")
	})
	app.Get("/payments", paymentHandler.GetPayments)
	app.Get("/paymentbyorderid/{order-id}", paymentHandler.GetPaymentByOrderID)
	app.Listen(":8080")
}

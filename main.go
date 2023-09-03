package main

import (
	"ticket-api/config"
	"ticket-api/modules/ticket/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	route := echo.New()

	apiv1 := route.Group("/api/v1")
	ticketController := controller.NewTicketController(db)
	apiv1.POST("/ticket/create", echo.HandlerFunc(ticketController.Create))
	apiv1.GET("/ticket/get_all", echo.HandlerFunc(ticketController.GetAll))
	apiv1.GET("/ticket/detail", echo.HandlerFunc(ticketController.GetById))
	apiv1.PUT("/ticket/update/:id", echo.HandlerFunc(ticketController.Update))
	apiv1.DELETE("/ticket/delete/:id", echo.HandlerFunc(ticketController.Delete))

	route.Start(":8084")
}

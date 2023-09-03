package controller

import (
	"net/http"
	"strconv"
	"ticket-api/modules/ticket/model"
	"ticket-api/modules/ticket/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TicketController struct {
	ticketService service.TicketService
	validate      validator.Validate
}

func (controller TicketController) Create(c echo.Context) error {
	type payload struct {
		AirlineName     string `json:"airline_name" validate:"required"`
		Price           int    `json:"price" validate:"required"`
		Amount          int    `json:"amount" validate:"required"`
		AvailableTicket int    `json:"available_tickets" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return err
	}

	result := controller.ticketService.Create(model.Ticket{AirlineName: payloadValidator.AirlineName, Price: payloadValidator.Price, Amount: payloadValidator.Amount, AvailableTicket: payloadValidator.AvailableTicket})

	return c.JSON(http.StatusOK, result)
}

func (controller TicketController) GetAll(c echo.Context) error {

	result := controller.ticketService.GetAll()

	return c.JSON(http.StatusOK, result)
}

func (controller TicketController) GetById(c echo.Context) error {

	id, _ := strconv.Atoi(c.QueryParam("id"))
	result := controller.ticketService.GetById(id)

	return c.JSON(http.StatusOK, result)
}

func (controller TicketController) Update(c echo.Context) error {
	type payload struct {
		AirlineName     string `json:"airline_name" validate:"required"`
		Price           int    `json:"price" validate:"required"`
		Amount          int    `json:"amount" validate:"required"`
		AvailableTicket int    `json:"available_tickets" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	result := controller.ticketService.Update(id, model.Ticket{AirlineName: payloadValidator.AirlineName, Price: payloadValidator.Price, Amount: payloadValidator.Amount, AvailableTicket: payloadValidator.AvailableTicket})

	return c.JSON(http.StatusOK, result)
}

func (controller TicketController) Delete(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	result := controller.ticketService.Delete(id)

	return c.JSON(http.StatusOK, result)
}

func NewTicketController(db *gorm.DB) TicketController {
	service := service.NewTicketService(db)

	controller := TicketController{
		ticketService: service,
		validate:      *validator.New(),
	}

	return controller
}

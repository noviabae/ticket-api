package service

import (
	"fmt"
	"ticket-api/modules/ticket/helper"
	"ticket-api/modules/ticket/model"
	"ticket-api/modules/ticket/repository"

	"gorm.io/gorm"
)

type ticketService struct {
	ticketRepository repository.TicketRepository
}

// Create implements TicketService.
func (service *ticketService) Create(ticket model.Ticket) helper.Response {
	var response helper.Response

	if err := service.ticketRepository.Create(ticket); err != nil {
		response.Status = 500
		response.Message = "Failed to create a new ticket"
	} else {
		response.Status = 200
		response.Message = "Success to create a new ticket!"
	}

	return response
}

// Delete implements TicketService.
func (service *ticketService) Delete(id int) helper.Response {
	var response helper.Response

	if err := service.ticketRepository.Delete(id); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to delete a new ticket: ", id)
	} else {
		response.Status = 200
		response.Message = "Success to delete a new ticket!"
	}

	return response
}

// GetAll implements TicketService.
func (service *ticketService) GetAll() helper.Response {
	var response helper.Response
	data, err := service.ticketRepository.GetAll()
	if err != nil {
		response.Status = 500
		response.Message = "Failed to get tickets"
	} else {
		response.Status = 200
		response.Message = "Success to get tickets"
		response.Data = data
	}
	return response

}

// GetById implements TicketService.
func (service *ticketService) GetById(id int) helper.Response {
	var response helper.Response
	data, err := service.ticketRepository.GetById(id)
	if err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to get ticket: ", id)
	} else {
		response.Status = 200
		response.Message = "Success to get tickets"
		response.Data = data
	}
	return response
}

// Update implements TicketService.
func (service *ticketService) Update(id int, ticket model.Ticket) helper.Response {
	var response helper.Response

	if err := service.ticketRepository.Update(id, ticket); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to update ticket ", id)
	} else {
		response.Status = 200
		response.Message = "Success to update ticket!"
	}

	return response
}

type TicketService interface {
	Create(ticket model.Ticket) helper.Response
	GetAll() helper.Response
	GetById(id int) helper.Response
	Update(id int, ticket model.Ticket) helper.Response
	Delete(id int) helper.Response
}

func NewTicketService(db *gorm.DB) TicketService {
	return &ticketService{ticketRepository: repository.NewTicketRepository(db)}
}

package repository

import (
	"ticket-api/modules/ticket/model"

	"gorm.io/gorm"
)

type databaseTicket struct {
	Conn *gorm.DB
}

// Create implements TicketRepository.
func (db *databaseTicket) Create(ticket model.Ticket) error {
	return db.Conn.Create(&ticket).Error
}

// Delete implements TicketRepository.
func (db *databaseTicket) Delete(id int) error {
	return db.Conn.Delete(&model.Ticket{Id: id}).Error
}

// GetAll implements TicketRepository.
func (db *databaseTicket) GetAll() ([]model.Ticket, error) {
	var data []model.Ticket
	result := db.Conn.Find(&data)

	return data, result.Error
}

// GetById implements TicketRepository.
func (db *databaseTicket) GetById(id int) (model.Ticket, error) {
	var data model.Ticket

	result := db.Conn.Where("id", id).First(&data)

	return data, result.Error
}

// Update implements TicketRepository.
func (db *databaseTicket) Update(id int, ticket model.Ticket) error {
	return db.Conn.Where("id", id).Updates(ticket).Error
}

type TicketRepository interface {
	Create(ticket model.Ticket) error
	GetAll() ([]model.Ticket, error)
	GetById(id int) (model.Ticket, error)
	Update(id int, ticket model.Ticket) error
	Delete(id int) error
}

func NewTicketRepository(Conn *gorm.DB) TicketRepository {
	return &databaseTicket{Conn: Conn}
}

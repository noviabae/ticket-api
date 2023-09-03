package model

type Ticket struct {
	Id              int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	AirlineName     string `json:"airline_name" gorm:"column:airline_name"`
	Price           int    `json:"price" gorm:"column:price"`
	Amount          int    `json:"amount" gorm:"column:amount"`
	AvailableTicket int    `json:"available_tickets" gorm:"column:available_tickets"`
}

func (Ticket) TableName() string {
	return "ticket"
}

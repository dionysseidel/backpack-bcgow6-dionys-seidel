package service

import "fmt"

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	ticketsLengthBeforeCreation := len(b.Tickets)
	b.Tickets = append(b.Tickets, t)
	if ticketsLengthAfterCreation := len(b.Tickets); ticketsLengthAfterCreation <= ticketsLengthBeforeCreation {
		return Ticket{}, fmt.Errorf("couldn't append new ticket %v to bookings list", t)
	}
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			return ticket, nil
		}
	}
	return Ticket{}, fmt.Errorf("couldn't find ticket by ID %d", id)
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	ticketFound, err := b.Read(id)
	if err != nil {
		return Ticket{}, fmt.Errorf("cannot update with ID %d: error in search", id)
	}
	b.Tickets[ticketFound.Id-1] = t
	return b.Tickets[ticketFound.Id-1], nil
}

func (b *bookings) Delete(id int) (int, error) {
	// ticketFound, err := b.Read(id)
	// if err != nil {
	// 	return id, fmt.Errorf("cannot delete ticket: error in search")
	// }
	newSlice := append(b.Tickets[:(id-1)], b.Tickets[id:]...)
	fmt.Println(newSlice)
	return 0, nil
}

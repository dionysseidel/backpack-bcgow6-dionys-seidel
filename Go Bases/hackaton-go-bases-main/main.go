package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	// "github.com/dionysseidel/backpack-bcgow6-dionys-seidel/Go Bases/hackaton-go-bases-main/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
	// "github.com/dionysseidel/backpack-bcgow6-dionys-seidel/Go Bases/hackaton-go-bases-main/internal/service"
)

func testSearchById(ticketsToFind []service.Ticket) {
	for _, ticket := range ticketsToFind {
		fmt.Println(ticket)
	}
}

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	file := file.File{}
	var err error
	tickets, err = file.Read()
	if err != nil {
		panic(err)
	}

	bookingsService := service.NewBookings(tickets)

	// Crear un nuevo ticket y añadir al registro
	newTicket := service.Ticket{
		Id:          1001,
		Names:       "Dionys Seidel",
		Email:       "dionys.seidel@mercadolibre.com",
		Destination: "Buenos Aires",
		Date:        "18:16",
		Price:       600,
	}
	ticketToWrite, err2 := service.Bookings.Create(bookingsService, newTicket)
	if err2 != nil {
		panic("couldn't add new ticket to tickets service")
	}
	fmt.Println(tickets)
	file.Write(ticketToWrite)

	// Actualizar los campos de un ticket
	ticketToBeUpdated := service.Ticket{
		Id:          1001,
		Names:       "Brian Heumann",
		Email:       "brianheumann@gmail.com",
		Destination: "Buenos Aires",
		Date:        "7:40",
		Price:       700,
	}
	service.Bookings.Update(bookingsService, 1001, ticketToBeUpdated)
	ticketToUpdateInFile, err7 := service.Bookings.Update(bookingsService, 1001, ticketToBeUpdated)
	if err7 != nil {
		panic("couldn't update new ticket in tickets service")
	}
	file.Write(ticketToUpdateInFile)

	// Eliminar un campo por su ID
	service.Bookings.Delete(bookingsService, 1001)

	// Traer un ticket a través de su campo Id
	ticketFound1, err3 := service.Bookings.Read(bookingsService, 640)
	ticketFound2, err4 := service.Bookings.Read(bookingsService, 1001)
	ticketFound3, err5 := service.Bookings.Read(bookingsService, 1002)
	ticketsToFind := []service.Ticket{ticketFound1, ticketFound2, ticketFound3}
	defer func() {
		err6 := recover()
		if err6 != nil {
			fmt.Println(err6)
			testSearchById(ticketsToFind)
		}
	}()
	if err3 != nil || err4 != nil || err5 != nil {
		panic("couldn't perform search")
	}
	testSearchById(ticketsToFind)
}

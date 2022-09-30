package file

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	data, err := os.ReadFile("tickets.csv")
	if err != nil {
		return nil, errors.New("an error ocurred trying to read the file")
	}
	var ticketsToReturn []service.Ticket
	records := strings.Split(string(data), "\n")
	for i, record := range records {
		fields := strings.Split(record, ",")
		iD, err2 := strconv.ParseInt(fields[0], 10, 64)
		if err2 != nil {
			return nil, fmt.Errorf("cannot parse string ID %d to int at record %d", iD, i)
		}
		price, err3 := strconv.ParseInt(fields[5], 10, 64)
		if err3 != nil {
			return nil, fmt.Errorf("cannot parse string Price %d to int at record %d", price, i)
		}
		ticket := service.Ticket{
			Id:          int(iD),
			Names:       fields[1],
			Email:       fields[2],
			Destination: fields[3],
			Date:        fields[4],
			Price:       int(price),
		}
		ticketsToReturn = append(ticketsToReturn, ticket)
	}
	return ticketsToReturn, nil
}

func (f *File) Write(ticket service.Ticket) error {
	file, err := os.OpenFile("tickets.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot open file")
	}
	data := []byte(fmt.Sprintf("\n%d,%s,%s,%s,%s,%d", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price))
	if _, err2 := file.Write(data); err2 != nil {
		return fmt.Errorf("cannot write in file for ticket %v", ticket)
	}
	return nil
}

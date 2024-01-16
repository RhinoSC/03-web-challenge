package loader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/rhinosc/03-web-challenge/internal"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (ticketFile *LoaderTicketCSV) Load() (t map[int]internal.TicketAttributes, err error) {
	// open the file
	f, err := os.Open(ticketFile.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	t = make(map[int]internal.TicketAttributes)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			err = fmt.Errorf("error reading record: %v", err)
			return t, err
		}

		// serialize the record
		price, err := strconv.Atoi(record[5])
		if err != nil {
			err = fmt.Errorf("error converting price: %v", err)
			return t, err
		}
		if err != nil {
			err = fmt.Errorf("error converting price: %v", err)
			return t, err
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			err = fmt.Errorf("error converting id: %v", err)
			return t, err
		}
		ticket := internal.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   float64(price),
		}

		// add the ticket to the map
		t[id] = ticket
	}

	return
}

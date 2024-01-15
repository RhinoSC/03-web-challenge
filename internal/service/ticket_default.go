package service

import (
	"context"
	"fmt"

	"github.com/RhinoSC/03-web-challenge/internal"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp internal.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	tickets, err := s.rp.Get(context.TODO())
	if err != nil {
		err = fmt.Errorf("error to get tickets: %w", err)
		return
	}
	total = len(tickets)
	return total, nil
}

// GetTotalTicketsByCountry returns the total number of tickets by country
func (s *ServiceTicketDefault) GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {
	tickets, err := s.rp.GetTicketsByDestinationCountry(ctx, country)
	if err != nil {
		err = fmt.Errorf("error to get tickets by country: %w", err)
		return
	}
	t = tickets
	return
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets by country
func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(ctx context.Context, country string) (percentage float64, err error) {
	tickets, err := s.rp.GetTicketsByDestinationCountry(ctx, country)
	if err != nil {
		err = fmt.Errorf("error to get tickets by country: %w", err)
		return
	}
	totalTickets, err := s.rp.Get(ctx)
	if err != nil {
		err = fmt.Errorf("error to get tickets: %w", err)
		return
	}
	percentage = float64(len(tickets)) / float64(len(totalTickets))
	return
}

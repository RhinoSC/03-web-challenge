package internal

import "context"

type ServiceTicket interface {
	// GetTotalAmountTickets returns the total amount of tickets
	GetTotalAmountTickets() (total int, err error)

	// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
	GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]TicketAttributes, err error)

	// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
	GetPercentageTicketsByDestinationCountry(ctx context.Context, country string) (percentage float64, err error)
}

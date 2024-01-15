package service_test

import (
	"context"
	"testing"

	"github.com/RhinoSC/03-web-challenge/internal"
	"github.com/RhinoSC/03-web-challenge/internal/repository"
	"github.com/RhinoSC/03-web-challenge/internal/service"
	"github.com/stretchr/testify/require"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		db := map[int]internal.TicketAttributes{
			1: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "USA",
				Hour:    "10:00",
				Price:   100,
			},
		}
		rp := repository.NewRepositoryTicketMap(db, 0)
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTotalAmountTickets()

		// assert

		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})

	t.Run("succes to get total tickets from destination country", func(t *testing.T) {
		// arrange
		db := map[int]internal.TicketAttributes{
			1: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "USA",
				Hour:    "10:00",
				Price:   100,
			},
			2: {
				Name:    "John",
				Email:   "johndoe@hotmail.com",
				Country: "USA",
				Hour:    "10:00",
				Price:   100,
			},
		}

		rp := repository.NewRepositoryTicketMap(db, 0)
		sv := service.NewServiceTicketDefault(rp)

		// act

		total, err := sv.GetTicketsByDestinationCountry(context.TODO(), "USA")

		// assert

		expectedTotal := 2
		require.NoError(t, err)
		require.Equal(t, expectedTotal, len(total))
	})

	t.Run("success to get percentage of tickets from destination country", func(t *testing.T) {
		// arrange
		db := map[int]internal.TicketAttributes{
			1: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "USA",
				Hour:    "10:00",
				Price:   100,
			},
			2: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "USA",
				Hour:    "10:00",
				Price:   100,
			},
			3: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "USA",
				Hour:    "10:00",
				Price:   100,
			},
			4: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "Brazil",
				Hour:    "10:00",
				Price:   100,
			},
		}

		rp := repository.NewRepositoryTicketMap(db, 0)
		sv := service.NewServiceTicketDefault(rp)

		// act

		percentage, err := sv.GetPercentageTicketsByDestinationCountry(context.TODO(), "USA")

		// assert

		expectedPercentage := 0.75

		require.NoError(t, err)
		require.Equal(t, expectedPercentage, percentage)
	})
}

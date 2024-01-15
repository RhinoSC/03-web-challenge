package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/RhinoSC/03-web-challenge/internal"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type TicketDefault struct {
	sv internal.ServiceTicket
}

func NewTicketDefault(sv internal.ServiceTicket) *TicketDefault {
	return &TicketDefault{
		sv: sv,
	}
}

func (h *TicketDefault) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// request

		// process

		total, err := h.sv.GetTotalAmountTickets()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// response

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success to get total tickets",
			"total":   total,
		})
	}
}

func (h *TicketDefault) GetByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// request

		dest := strings.ToLower(chi.URLParam(r, "dest"))

		// process

		total, err := h.sv.GetTicketsByDestinationCountry(context.TODO(), dest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// response

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success to get total tickets for " + dest,
			"total":   len(total),
		})
	}
}

func (h *TicketDefault) GetPercentageByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// request

		dest := strings.ToLower(chi.URLParam(r, "dest"))

		// process

		percentage, err := h.sv.GetPercentageTicketsByDestinationCountry(context.TODO(), dest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// response

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success to get percentage of tickets for " + dest,
			"total":   percentage,
		})
	}
}

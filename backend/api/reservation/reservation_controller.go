package reservation

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nikolaypleshkov/uni-api/api/reservation/dto"
)

type ReservationController struct {
	reservationService ReservationService
}

func NewReservationController(service ReservationService) *ReservationController {
	return &ReservationController{service}
}

func (c *ReservationController) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var createReservationDTO dto.CreateReservationDTO
	if err := json.NewDecoder(r.Body).Decode(&createReservationDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdReservation, err := c.reservationService.CreateReservation(createReservationDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdReservation)
}

func (c *ReservationController) DeleteReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reservationID, err := strconv.ParseInt(vars["reservationId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.reservationService.DeleteReservation(reservationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ReservationController) GetAllReservations(w http.ResponseWriter, r *http.Request) {
	reservations, err := c.reservationService.GetAllReservations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}

func (c *ReservationController) GetReservationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reservationID, err := strconv.ParseInt(vars["reservationId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reservation, err := c.reservationService.GetReservationByID(reservationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservation)
}

func (c *ReservationController) UpdateReservation(w http.ResponseWriter, r *http.Request) {
	var updateReservationDTO dto.UpdateReservationDTO
	if err := json.NewDecoder(r.Body).Decode(&updateReservationDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedReservation, err := c.reservationService.UpdateReservation(updateReservationDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedReservation)
}

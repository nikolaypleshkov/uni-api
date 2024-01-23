package holiday

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nikolaypleshkov/uni-api/api/holiday/dto"
)

type Controller struct {
	service *Service
}

func convertHolidayToDTO(holiday Holiday) dto.CreateHolidayDTO {
	return dto.CreateHolidayDTO{
		Title:     holiday.Title,
		StartDate: holiday.StartDate,
		Duration:  holiday.Duration,
		FreeSlots: holiday.FreeSlots,
		Price:     holiday.Price,
		Location:  holiday.LocationID,
	}
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) CreateHoliday(w http.ResponseWriter, r *http.Request) {
	var holiday Holiday
	err := json.NewDecoder(r.Body).Decode(&holiday)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createHolidayDTO := convertHolidayToDTO(holiday)

	createdHoliday, err := c.service.CreateHoliday(createHolidayDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdHoliday)
}

func (c *Controller) DeleteHoliday(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	holidayIDStr := params["holidayId"]

	holidayID, err := strconv.ParseInt(holidayIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid holiday ID", http.StatusBadRequest)
		return
	}

	err = c.service.DeleteHoliday(holidayID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *Controller) GetHolidays(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	holidays, err := c.service.GetHolidays(queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	holidaysJSON, err := json.Marshal(holidays)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(holidaysJSON)
}

func (c *Controller) GetHoliday(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	holidayIDStr := params["holidayId"]

	holidayID, err := strconv.ParseInt(holidayIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid holiday ID", http.StatusBadRequest)
		return
	}

	holiday, err := c.service.GetHoliday(holidayID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(holiday)
}

func (c *Controller) UpdateHoliday(w http.ResponseWriter, r *http.Request) {
	var updateDTO dto.UpdateHolidayDTO

	err := json.NewDecoder(r.Body).Decode(&updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.UpdateHoliday(updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateDTO)
}

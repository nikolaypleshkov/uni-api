package location

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nikolaypleshkov/uni-api/api/location/dto"
)

type LocationController struct {
	service LocationService
}

func NewLocationController(service LocationService) *LocationController {
	return &LocationController{service}
}

func (c *LocationController) CreateLocation(w http.ResponseWriter, r *http.Request) {
	var createLocationDTO dto.CreateLocationDTO
	err := json.NewDecoder(r.Body).Decode(&createLocationDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdLocation, err := c.service.CreateLocation(createLocationDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdLocation)
}

func (c *LocationController) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	locationID, err := strconv.ParseInt(params["locationId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.DeleteLocation(locationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *LocationController) GetAllLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := c.service.GetAllLocations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}

func (c *LocationController) GetLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	locationID, err := strconv.ParseInt(params["locationId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	location, err := c.service.GetLocation(locationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}

func (c *LocationController) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	var updateLocationDTO dto.UpdateLocationDTO
	err := json.NewDecoder(r.Body).Decode(&updateLocationDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedLocation, err := c.service.UpdateLocation(updateLocationDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedLocation)
}

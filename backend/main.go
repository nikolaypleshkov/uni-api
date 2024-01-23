package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nikolaypleshkov/uni-api/api/holiday"
	"github.com/nikolaypleshkov/uni-api/api/location"
	"github.com/nikolaypleshkov/uni-api/api/reservation"
)

func main() {
	db, err := sql.Open("postgres", "postgres://myuser:mypassword@mypostgres:5432/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	locationService := location.NewLocationService(db)
	holidayService := holiday.NewService(db, locationService)
	reservationService := reservation.NewReservationService(db, holidayService)

	holidayController := holiday.NewController(holidayService)
	locationController := location.NewLocationController(locationService)
	reservationController := reservation.NewReservationController(reservationService)

	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": "ok", "message": "Backend is healthy"}`)
	}).Methods("GET")

	router.HandleFunc("/travel-agency/holidays", holidayController.CreateHoliday).Methods("POST")
	router.HandleFunc("/travel-agency/holidays/{holidayId}", holidayController.DeleteHoliday).Methods("DELETE")
	router.HandleFunc("/travel-agency/holidays", holidayController.GetHolidays).Methods("GET")
	router.HandleFunc("/travel-agency/holidays/{holidayId}", holidayController.GetHoliday).Methods("GET")
	router.HandleFunc("/travel-agency/holidays", holidayController.UpdateHoliday).Methods("PUT")

	router.HandleFunc("/travel-agency/locations", locationController.CreateLocation).Methods("POST")
	router.HandleFunc("/travel-agency/locations/{locationId:[0-9]+}", locationController.DeleteLocation).Methods("DELETE")
	router.HandleFunc("/travel-agency/locations", locationController.GetAllLocations).Methods("GET")
	router.HandleFunc("/travel-agency/locations/{locationId:[0-9]+}", locationController.GetLocation).Methods("GET")
	router.HandleFunc("/travel-agency/locations", locationController.UpdateLocation).Methods("PUT")

	router.HandleFunc("/travel-agency/reservations", reservationController.CreateReservation).Methods("POST")
	router.HandleFunc("/travel-agency/reservations/{reservationId}", reservationController.GetReservationByID).Methods("GET")
	router.HandleFunc("/travel-agency/reservations", reservationController.GetAllReservations).Methods("GET")
	router.HandleFunc("/travel-agency/reservations/{reservationId}", reservationController.DeleteReservation).Methods("DELETE")

	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(router)

	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), corsHandler))
}

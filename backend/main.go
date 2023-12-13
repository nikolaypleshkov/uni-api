package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nikolaypleshkov/uni-api/api/holiday"
	"github.com/nikolaypleshkov/uni-api/api/location"
	"github.com/nikolaypleshkov/uni-api/api/reservation"
)

func main() {
	db, err := sql.Open("postgres", "user=myuser password=mypassword dbname=mydatabase sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	holidayService := holiday.NewService(db)
	locationService := location.NewLocationService(db)
	reservationService := reservation.NewReservationService(db)

	holidayController := holiday.NewController(holidayService)
	locationController := location.NewLocationController(locationService)
	reservationController := reservation.NewReservationController(reservationService)

	router := mux.NewRouter()

	router.HandleFunc("/holidays", holidayController.CreateHoliday).Methods("POST")
	router.HandleFunc("/holidays/{holidayId}", holidayController.DeleteHoliday).Methods("DELETE")
	router.HandleFunc("/holidays", holidayController.GetHolidays).Methods("GET")
	router.HandleFunc("/holidays/{holidayId}", holidayController.GetHoliday).Methods("GET")
	router.HandleFunc("/holidays", holidayController.UpdateHoliday).Methods("PUT")

	router.HandleFunc("/locations", locationController.CreateLocation).Methods("POST")
	router.HandleFunc("/locations/{locationId:[0-9]+}", locationController.DeleteLocation).Methods("DELETE")
	router.HandleFunc("/locations", locationController.GetAllLocations).Methods("GET")
	router.HandleFunc("/locations/{locationId:[0-9]+}", locationController.GetLocation).Methods("GET")
	router.HandleFunc("/locations", locationController.UpdateLocation).Methods("PUT")

	router.HandleFunc("/reservations", reservationController.CreateReservation).Methods("POST")
	router.HandleFunc("/reservations/{reservationId}", reservationController.GetReservationByID).Methods("GET")
	router.HandleFunc("/reservations", reservationController.GetAllReservations).Methods("GET")
	router.HandleFunc("/reservations/{reservationId}", reservationController.DeleteReservation).Methods("DELETE")

	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

package dto

import "github.com/nikolaypleshkov/uni-api/api/holiday"

type CreateReservationDTO struct {
	PhoneNumber string `json:"phone_number"`
	ContactName string `json:"contact_name"`
	HolidayID   int64  `json:"holiday"`
}

type UpdateReservationDTO struct {
	ID          int64  `json:"id"`
	PhoneNumber string `json:"phone_number"`
	ContactName string `json:"contact_name"`
	HolidayID   int64  `json:"holiday_id"`
}

type ResponseReservationDTO struct {
	ID          int64           `json:"id"`
	PhoneNumber string          `json:"phone_number"`
	ContactName string          `json:"contact_name"`
	Holiday     holiday.Holiday `json:"holiday"`
}

type GetAllResponseReservationDTO []ResponseReservationDTO

package dto

import location "github.com/nikolaypleshkov/uni-api/api/location/dto"

type CreateHolidayDTO struct {
	Title     string `json:"title"`
	StartDate string `json:"startDate"`
	Duration  int32  `json:"duration"`
	FreeSlots int32  `json:"freeSlots"`
	Price     string `json:"price"`
	Location  int64  `json:"location"`
}

type UpdateHolidayDTO struct {
	ID        int64   `json:"id"`
	Title     string  `json:"title"`
	StartDate string  `json:"startDate"`
	Duration  int32   `json:"duration"`
	FreeSlots int32   `json:"freeSlots"`
	Price     float64 `json:"price"`
	Location  int64   `json:"location"`
}

type ResponseHolidayDTO struct {
	ID         int64                        `json:"id"`
	Title      string                       `json:"title"`
	StartDate  string                       `json:"startDate"`
	Duration   int32                        `json:"duration"`
	FreeSlots  int32                        `json:"freeSlots"`
	Price      string                       `json:"price"`
	Location   location.ResponseLocationDTO `json:"location"`
	LocationID int64                        `json:"location_id"`
}

package dto

type CreateHolidayDTO struct {
	Title     string      `json:"title"`
	StartDate string      `json:"startDate"`
	Duration  int32       `json:"duration"`
	FreeSlots int32       `json:"freeSlots"`
	Price     float64     `json:"price"`
	Location  LocationDTO `json:"location"`
}

type UpdateHolidayDTO struct {
	ID        int64       `json:"id"`
	Title     string      `json:"title"`
	StartDate string      `json:"startDate"`
	Duration  int32       `json:"duration"`
	FreeSlots int32       `json:"freeSlots"`
	Price     float64     `json:"price"`
	Location  LocationDTO `json:"location"`
}

type ResponseHolidayDTO struct {
	ID        int64       `json:"id"`
	Title     string      `json:"title"`
	StartDate string      `json:"startDate"`
	Duration  int32       `json:"duration"`
	FreeSlots int32       `json:"freeSlots"`
	Price     float64     `json:"price"`
	Location  LocationDTO `json:"location"`
}

type LocationDTO struct {
	ID       int64  `json:"id"`
	Number   string `json:"number"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Street   string `json:"street"`
	ImageURL string `json:"imageUrl"`
}

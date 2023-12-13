package reservation

type Reservation struct {
	ID          int64  `json:"id"`
	PhoneNumber string `json:"phone_number"`
	ContactName string `json:"contact_name"`
	HolidayID   int64  `json:"holiday_id"`
}

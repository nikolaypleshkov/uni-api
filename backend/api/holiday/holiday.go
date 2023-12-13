package holiday

type Holiday struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	StartDate string   `json:"startDate"`
	Duration  int32    `json:"duration"`
	FreeSlots int32    `json:"freeSlots"`
	Price     float64  `json:"price"`
	Location  Location `json:"location"`
}

type Location struct {
	ID       int64  `json:"id"`
	Number   string `json:"number"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Street   string `json:"street"`
	ImageURL string `json:"imageUrl"`
}

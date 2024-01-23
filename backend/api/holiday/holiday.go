package holiday

type Holiday struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	StartDate  string `json:"startDate"`
	Duration   int32  `json:"duration"`
	FreeSlots  int32  `json:"freeSlots"`
	Price      string `json:"price"`
	LocationID int64  `json:"location"`
}

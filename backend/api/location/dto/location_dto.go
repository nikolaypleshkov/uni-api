package dto

type CreateLocationDTO struct {
	Number   string `json:"number"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Street   string `json:"street"`
	ImageURL string `json:"imageUrl"`
}

type UpdateLocationDTO struct {
	ID       int64  `json:"id"`
	Number   string `json:"number"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Street   string `json:"street"`
	ImageURL string `json:"imageUrl"`
}

type ResponseLocationDTO struct {
	ID       int64  `json:"id"`
	Number   string `json:"number"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Street   string `json:"street"`
	ImageURL string `json:"imageUrl"`
}

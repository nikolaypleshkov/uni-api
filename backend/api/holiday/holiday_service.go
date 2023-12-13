package holiday

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/nikolaypleshkov/uni-api/api/holiday/dto"
)

type Service struct {
	mu       sync.Mutex
	db       *sql.DB
	holidays []Holiday
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db:       db,
		holidays: make([]Holiday, 0),
	}
}

func (s *Service) CreateHoliday(holidayDTO dto.CreateHolidayDTO) (dto.ResponseHolidayDTO, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := `
		INSERT INTO holidays (title, start_date, duration, free_slots, price, location_number, location_country, location_city, location_street, location_image_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, title, start_date, duration, free_slots, price, location_number, location_country, location_city, location_street, location_image_url
	`
	row := s.db.QueryRow(
		query,
		holidayDTO.Title,
		holidayDTO.StartDate,
		holidayDTO.Duration,
		holidayDTO.FreeSlots,
		holidayDTO.Price,
		holidayDTO.Location.Number,
		holidayDTO.Location.Country,
		holidayDTO.Location.City,
		holidayDTO.Location.Street,
		holidayDTO.Location.ImageURL,
	)

	var createdHoliday Holiday
	err := row.Scan(
		&createdHoliday.ID,
		&createdHoliday.Title,
		&createdHoliday.StartDate,
		&createdHoliday.Duration,
		&createdHoliday.FreeSlots,
		&createdHoliday.Price,
		&createdHoliday.Location.Number,
		&createdHoliday.Location.Country,
		&createdHoliday.Location.City,
		&createdHoliday.Location.Street,
		&createdHoliday.Location.ImageURL,
	)

	if err != nil {
		return dto.ResponseHolidayDTO{}, err
	}

	responseDTO := dto.ResponseHolidayDTO{
		ID:        createdHoliday.ID,
		Title:     createdHoliday.Title,
		StartDate: createdHoliday.StartDate,
		Duration:  createdHoliday.Duration,
		FreeSlots: createdHoliday.FreeSlots,
		Price:     createdHoliday.Price,
		Location: dto.LocationDTO{
			ID:       createdHoliday.Location.ID,
			Number:   createdHoliday.Location.Number,
			Country:  createdHoliday.Location.Country,
			City:     createdHoliday.Location.City,
			Street:   createdHoliday.Location.Street,
			ImageURL: createdHoliday.Location.ImageURL,
		},
	}

	return responseDTO, nil
}

func (s *Service) DeleteHoliday(holidayID int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := "DELETE FROM holidays WHERE id = $1"

	_, err := s.db.Exec(query, holidayID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetHolidays(queryParams map[string]string) ([]dto.ResponseHolidayDTO, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := "SELECT * FROM holidays WHERE 1 = 1"

	if startDate, ok := queryParams["startDate"]; ok {
		query += fmt.Sprintf(" AND start_date = '%s'", startDate)
	}
	if duration, ok := queryParams["duration"]; ok {
		query += fmt.Sprintf(" AND duration = %s", duration)
	}

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultDTOs []dto.ResponseHolidayDTO
	for rows.Next() {
		var holiday Holiday
		err := rows.Scan(
			&holiday.ID,
			&holiday.Title,
			&holiday.StartDate,
			&holiday.Duration,
			&holiday.FreeSlots,
			&holiday.Price,
			&holiday.Location.Number,
			&holiday.Location.Country,
			&holiday.Location.City,
			&holiday.Location.Street,
			&holiday.Location.ImageURL,
		)
		if err != nil {
			return nil, err
		}

		resultDTO := dto.ResponseHolidayDTO{
			ID:        holiday.ID,
			Title:     holiday.Title,
			StartDate: holiday.StartDate,
			Duration:  holiday.Duration,
			FreeSlots: holiday.FreeSlots,
			Price:     holiday.Price,
			Location: dto.LocationDTO{
				ID:       holiday.Location.ID,
				Number:   holiday.Location.Number,
				Country:  holiday.Location.Country,
				City:     holiday.Location.City,
				Street:   holiday.Location.Street,
				ImageURL: holiday.Location.ImageURL,
			},
		}

		resultDTOs = append(resultDTOs, resultDTO)
	}

	return resultDTOs, nil
}

func (s *Service) GetHoliday(holidayID int64) (dto.ResponseHolidayDTO, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := "SELECT * FROM holidays WHERE id = $1"

	row := s.db.QueryRow(query, holidayID)

	var holiday Holiday
	err := row.Scan(
		&holiday.ID,
		&holiday.Title,
		&holiday.StartDate,
		&holiday.Duration,
		&holiday.FreeSlots,
		&holiday.Price,
		&holiday.Location.Number,
		&holiday.Location.Country,
		&holiday.Location.City,
		&holiday.Location.Street,
		&holiday.Location.ImageURL,
	)

	if err != nil {
		return dto.ResponseHolidayDTO{}, err
	}

	responseDTO := dto.ResponseHolidayDTO{
		ID:        holiday.ID,
		Title:     holiday.Title,
		StartDate: holiday.StartDate,
		Duration:  holiday.Duration,
		FreeSlots: holiday.FreeSlots,
		Price:     holiday.Price,
		Location: dto.LocationDTO{
			ID:       holiday.Location.ID,
			Number:   holiday.Location.Number,
			Country:  holiday.Location.Country,
			City:     holiday.Location.City,
			Street:   holiday.Location.Street,
			ImageURL: holiday.Location.ImageURL,
		},
	}

	return responseDTO, nil
}

func (s *Service) UpdateHoliday(updateDTO dto.UpdateHolidayDTO) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := `
		UPDATE holidays
		SET title = $1, start_date = $2, duration = $3, free_slots = $4, price = $5,
		    location_number = $6, location_country = $7, location_city = $8, location_street = $9, location_image_url = $10
		WHERE id = $11
	`
	_, err := s.db.Exec(
		query,
		updateDTO.Title,
		updateDTO.StartDate,
		updateDTO.Duration,
		updateDTO.FreeSlots,
		updateDTO.Price,
		updateDTO.Location.Number,
		updateDTO.Location.Country,
		updateDTO.Location.City,
		updateDTO.Location.Street,
		updateDTO.Location.ImageURL,
		updateDTO.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetHolidayDTO(holidayID int64) (Holiday, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := "SELECT * FROM holidays WHERE id = $1"

	row := s.db.QueryRow(query, holidayID)

	var holiday Holiday
	err := row.Scan(
		&holiday.ID,
		&holiday.Title,
		&holiday.StartDate,
		&holiday.Duration,
		&holiday.FreeSlots,
		&holiday.Price,
		&holiday.Location.Number,
		&holiday.Location.Country,
		&holiday.Location.City,
		&holiday.Location.Street,
		&holiday.Location.ImageURL,
	)

	if err != nil {
		return Holiday{}, err
	}

	getHolidayDTO := Holiday{
		ID:        holiday.ID,
		Title:     holiday.Title,
		StartDate: holiday.StartDate,
	}

	return getHolidayDTO, nil
}

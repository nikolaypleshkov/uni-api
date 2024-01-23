package holiday

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"sync"

	_ "github.com/lib/pq"
	"github.com/nikolaypleshkov/uni-api/api/holiday/dto"
	"github.com/nikolaypleshkov/uni-api/api/location"
)

type Service struct {
	mu              sync.Mutex
	db              *sql.DB
	holidays        []Holiday
	locationService *location.LocationServiceImpl
}

func NewService(db *sql.DB, locationService *location.LocationServiceImpl) *Service {
	return &Service{
		db:              db,
		holidays:        make([]Holiday, 0),
		locationService: locationService,
	}
}
func (s *Service) ensureTableExists() error {
	query := `
        CREATE TABLE IF NOT EXISTS holidays (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255),
            start_date DATE,
            duration INT,
            free_slots INT,
            price DECIMAL(10,2),
            location_id INT,
            FOREIGN KEY (location_id) REFERENCES locations(id)
        );
    `
	_, err := s.db.Exec(query)
	return err
}

func (s *Service) DropAndRecreateTable() error {
	dropQuery := "DROP TABLE IF EXISTS holidays"
	_, err := s.db.Exec(dropQuery)
	if err != nil {
		return fmt.Errorf("failed to drop holidays table: %v", err)
	}

	createQuery := `
		CREATE TABLE IF NOT EXISTS holidays (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255),
			start_date DATE,
			duration INT,
			free_slots INT,
			price DECIMAL(10,2),
			location_id INT,
			FOREIGN KEY (location_id) REFERENCES locations(id)
		);
	`
	_, err = s.db.Exec(createQuery)
	if err != nil {
		return fmt.Errorf("failed to recreate holidays table: %v", err)
	}

	return nil
}

func (s *Service) CreateHoliday(holidayDTO dto.CreateHolidayDTO) (dto.ResponseHolidayDTO, error) {
	if err := s.ensureTableExists(); err != nil {
		return dto.ResponseHolidayDTO{}, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	var locationID sql.NullInt64
	if holidayDTO.Location != -1 {
		locationID = sql.NullInt64{Int64: holidayDTO.Location, Valid: true}
	} else {
		locationID = sql.NullInt64{Valid: false}
	}

	query := `
        INSERT INTO holidays (title, start_date, duration, free_slots, price, location_id)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, title, start_date, duration, free_slots, price, location_id
    `
	row := s.db.QueryRow(
		query,
		holidayDTO.Title,
		holidayDTO.StartDate,
		holidayDTO.Duration,
		holidayDTO.FreeSlots,
		holidayDTO.Price,
		locationID,
	)

	var createdHoliday Holiday
	err := row.Scan(
		&createdHoliday.ID,
		&createdHoliday.Title,
		&createdHoliday.StartDate,
		&createdHoliday.Duration,
		&createdHoliday.FreeSlots,
		&createdHoliday.Price,
		&createdHoliday.LocationID,
	)

	if err != nil {
		return dto.ResponseHolidayDTO{}, err
	}

	responseDTO := dto.ResponseHolidayDTO{
		ID:         createdHoliday.ID,
		Title:      createdHoliday.Title,
		StartDate:  createdHoliday.StartDate,
		Duration:   createdHoliday.Duration,
		FreeSlots:  createdHoliday.FreeSlots,
		Price:      createdHoliday.Price,
		LocationID: createdHoliday.LocationID,
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

func (s *Service) GetHolidays(queryParams url.Values) ([]dto.ResponseHolidayDTO, error) {
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
			&holiday.LocationID,
		)
		if err != nil {
			return nil, err
		}

		locationDTO, err := s.locationService.GetLocation(holiday.LocationID)
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
			Location:  locationDTO,
		}

		resultDTOs = append(resultDTOs, resultDTO)
	}

	log.Printf("Successfully retrieved %d holidays", len(resultDTOs))
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
		&holiday.LocationID,
	)

	if err != nil {
		return dto.ResponseHolidayDTO{}, err
	}

	locationDTO, err := s.locationService.GetLocation(holiday.LocationID)
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
		Location:  locationDTO,
	}

	return responseDTO, nil
}
func (s *Service) UpdateHoliday(updateDTO dto.UpdateHolidayDTO) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := `
        UPDATE holidays
        SET title = $1, start_date = $2, duration = $3, free_slots = $4, price = $5,
            location_id = $6
        WHERE id = $7
    `

	priceString := strconv.FormatFloat(updateDTO.Price, 'f', -1, 64)

	_, err := s.db.Exec(
		query,
		updateDTO.Title,
		updateDTO.StartDate,
		updateDTO.Duration,
		updateDTO.FreeSlots,
		priceString,
		updateDTO.Location,
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
		&holiday.LocationID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return Holiday{}, fmt.Errorf("holiday with ID %d not found", holidayID)
		}
		return Holiday{}, err
	}

	return holiday, nil
}

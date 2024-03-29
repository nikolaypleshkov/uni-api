package reservation

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/nikolaypleshkov/uni-api/api/holiday"
	"github.com/nikolaypleshkov/uni-api/api/reservation/dto"
)

type ReservationService interface {
	GetAllReservations() ([]dto.ResponseReservationDTO, error)
	CreateReservation(createDTO dto.CreateReservationDTO) (dto.ResponseReservationDTO, error)
	UpdateReservation(updateDTO dto.UpdateReservationDTO) (dto.ResponseReservationDTO, error)
	GetReservation(reservationID int64) (dto.ResponseReservationDTO, error)
	DeleteReservation(reservationID int64) error
	GetReservationByID(reservationID int64) (dto.ResponseReservationDTO, error)
}

type HolidayDTO struct {
	ID int64 `json:"id"`
}

type GetHolidayDTO struct {
	ID int64 `json:"id"`
}

type ReservationServiceImpl struct {
	db             *sql.DB
	HolidayService *holiday.Service
	mu             sync.Mutex
}

func NewReservationService(db *sql.DB, holidayService *holiday.Service) *ReservationServiceImpl {
	return &ReservationServiceImpl{
		db:             db,
		HolidayService: holidayService,
	}
}

func (s *ReservationServiceImpl) ensureTableExists() error {
	query := `
        CREATE TABLE IF NOT EXISTS reservations (
            id SERIAL PRIMARY KEY,
            phone_number VARCHAR(255),
            contact_name VARCHAR(255),
            holiday_id INT
        );
    `
	_, err := s.db.Exec(query)
	return err
}

func (s *ReservationServiceImpl) GetAllReservations() ([]dto.ResponseReservationDTO, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := "SELECT * FROM reservations"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []dto.ResponseReservationDTO
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(
			&reservation.ID,
			&reservation.PhoneNumber,
			&reservation.ContactName,
			&reservation.HolidayID,
		)
		if err != nil {
			return nil, err
		}

		holidayDTO, err := s.HolidayService.GetHolidayDTO(reservation.HolidayID)
		if err != nil {
			return nil, err
		}

		responseDTO := dto.ResponseReservationDTO{
			ID:          reservation.ID,
			PhoneNumber: reservation.PhoneNumber,
			ContactName: reservation.ContactName,
			Holiday:     holidayDTO,
		}
		reservations = append(reservations, responseDTO)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (s *ReservationServiceImpl) CreateReservation(createDTO dto.CreateReservationDTO) (dto.ResponseReservationDTO, error) {
	if err := s.ensureTableExists(); err != nil {
		return dto.ResponseReservationDTO{}, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.HolidayService.GetHolidayDTO(createDTO.HolidayID)
	if err != nil {
		return dto.ResponseReservationDTO{}, fmt.Errorf("associated holiday not found: %v", err)
	}

	query := `
        INSERT INTO reservations (phone_number, contact_name, holiday_id)
        VALUES ($1, $2, $3)
        RETURNING id, phone_number, contact_name, holiday_id
    `

	row := s.db.QueryRow(
		query,
		createDTO.PhoneNumber,
		createDTO.ContactName,
		createDTO.HolidayID,
	)

	var createdReservation Reservation
	err = row.Scan(
		&createdReservation.ID,
		&createdReservation.PhoneNumber,
		&createdReservation.ContactName,
		&createdReservation.HolidayID,
	)

	if err != nil {
		return dto.ResponseReservationDTO{}, err
	}

	holidayDTO, err := s.HolidayService.GetHolidayDTO(createdReservation.HolidayID)
	if err != nil {
		return dto.ResponseReservationDTO{}, err
	}

	responseDTO := dto.ResponseReservationDTO{
		ID:          createdReservation.ID,
		PhoneNumber: createdReservation.PhoneNumber,
		ContactName: createdReservation.ContactName,
		Holiday:     holidayDTO,
	}

	return responseDTO, nil
}

func (s *ReservationServiceImpl) UpdateReservation(updateDTO dto.UpdateReservationDTO) (dto.ResponseReservationDTO, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := `
		UPDATE reservations
		SET phone_number = $1, contact_name = $2
		WHERE id = $3
		RETURNING id, phone_number, contact_name, holiday_id
	`

	row := s.db.QueryRow(
		query,
		updateDTO.PhoneNumber,
		updateDTO.ContactName,
		updateDTO.ID,
	)

	var updatedReservation Reservation
	err := row.Scan(
		&updatedReservation.ID,
		&updatedReservation.PhoneNumber,
		&updatedReservation.ContactName,
		&updatedReservation.HolidayID,
	)

	if err != nil {
		return dto.ResponseReservationDTO{}, err
	}

	holidayDTO, err := s.HolidayService.GetHolidayDTO(updatedReservation.HolidayID)
	if err != nil {
		return dto.ResponseReservationDTO{}, err
	}

	responseDTO := dto.ResponseReservationDTO{
		ID:          updatedReservation.ID,
		PhoneNumber: updatedReservation.PhoneNumber,
		ContactName: updatedReservation.ContactName,
		Holiday:     holidayDTO,
	}

	return responseDTO, nil
}

func (s *ReservationServiceImpl) GetReservation(reservationID int64) (dto.ResponseReservationDTO, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := "SELECT * FROM reservations WHERE id = $1"

	row := s.db.QueryRow(query, reservationID)

	var reservation Reservation
	err := row.Scan(
		&reservation.ID,
		&reservation.PhoneNumber,
		&reservation.ContactName,
		&reservation.HolidayID,
	)

	if err != nil {
		return dto.ResponseReservationDTO{}, err
	}

	holidayDTO, err := s.HolidayService.GetHolidayDTO(reservation.HolidayID)
	if err != nil {
		return dto.ResponseReservationDTO{}, err
	}

	responseDTO := dto.ResponseReservationDTO{
		ID:          reservation.ID,
		PhoneNumber: reservation.PhoneNumber,
		ContactName: reservation.ContactName,
		Holiday:     holidayDTO,
	}

	return responseDTO, nil
}

func (s *ReservationServiceImpl) DeleteReservation(reservationID int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := "DELETE FROM reservations WHERE id = $1"

	_, err := s.db.Exec(query, reservationID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ReservationServiceImpl) GetReservationByID(reservationID int64) (dto.ResponseReservationDTO, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := "SELECT * FROM reservations WHERE id = $1"

	row := s.db.QueryRow(query, reservationID)

	var reservation Reservation
	err := row.Scan(
		&reservation.ID,
		&reservation.PhoneNumber,
		&reservation.ContactName,
		&reservation.HolidayID,
	)

	if err != nil {
		return dto.ResponseReservationDTO{}, err
	}

	responseDTO := dto.ResponseReservationDTO{
		ID:          reservation.ID,
		PhoneNumber: reservation.PhoneNumber,
		ContactName: reservation.ContactName,
	}

	return responseDTO, nil
}

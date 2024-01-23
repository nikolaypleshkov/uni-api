package location

import (
	"database/sql"
	"errors"

	"github.com/nikolaypleshkov/uni-api/api/location/dto"
)

type LocationService interface {
	CreateLocation(createLocationDTO dto.CreateLocationDTO) (dto.ResponseLocationDTO, error)
	DeleteLocation(locationID int64) error
	GetAllLocations() ([]dto.ResponseLocationDTO, error)
	GetLocation(locationID int64) (dto.ResponseLocationDTO, error)
	UpdateLocation(updateLocationDTO dto.UpdateLocationDTO) (dto.ResponseLocationDTO, error)
}

type LocationServiceImpl struct {
	db *sql.DB
}

func NewLocationService(db *sql.DB) *LocationServiceImpl {
	return &LocationServiceImpl{db}
}

func (s *LocationServiceImpl) ensureTableExists() error {
	query := `
        CREATE TABLE IF NOT EXISTS locations (
            id SERIAL PRIMARY KEY,
            number VARCHAR(255),
            country VARCHAR(255),
            city VARCHAR(255),
            street VARCHAR(255),
            image_url VARCHAR(255)
        );
    `
	_, err := s.db.Exec(query)
	return err
}

func (s *LocationServiceImpl) CreateLocation(createLocationDTO dto.CreateLocationDTO) (dto.ResponseLocationDTO, error) {
	if err := s.ensureTableExists(); err != nil {
		return dto.ResponseLocationDTO{}, err
	}

	query := `
		INSERT INTO locations (number, country, city, street, image_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, number, country, city, street, image_url
	`

	row := s.db.QueryRow(
		query,
		createLocationDTO.Number,
		createLocationDTO.Country,
		createLocationDTO.City,
		createLocationDTO.Street,
		createLocationDTO.ImageURL,
	)

	var createdLocation dto.ResponseLocationDTO
	err := row.Scan(
		&createdLocation.ID,
		&createdLocation.Number,
		&createdLocation.Country,
		&createdLocation.City,
		&createdLocation.Street,
		&createdLocation.ImageURL,
	)

	if err != nil {
		return dto.ResponseLocationDTO{}, err
	}

	return createdLocation, nil
}

func (s *LocationServiceImpl) DeleteLocation(locationID int64) error {
	query := "DELETE FROM locations WHERE id = $1"

	result, err := s.db.Exec(query, locationID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("location not found")
	}

	return nil
}

func (s *LocationServiceImpl) GetAllLocations() ([]dto.ResponseLocationDTO, error) {
	query := "SELECT * FROM locations"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []dto.ResponseLocationDTO
	for rows.Next() {
		var location dto.ResponseLocationDTO
		err := rows.Scan(
			&location.ID,
			&location.Number,
			&location.Country,
			&location.City,
			&location.Street,
			&location.ImageURL,
		)
		if err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	return locations, nil
}

func (s *LocationServiceImpl) GetLocation(locationID int64) (dto.ResponseLocationDTO, error) {
	query := "SELECT * FROM locations WHERE id = $1"

	row := s.db.QueryRow(query, locationID)

	var location dto.ResponseLocationDTO
	err := row.Scan(
		&location.ID,
		&location.Number,
		&location.Country,
		&location.City,
		&location.Street,
		&location.ImageURL,
	)

	if err != nil {
		return dto.ResponseLocationDTO{}, err
	}

	return location, nil
}

func (s *LocationServiceImpl) UpdateLocation(updateLocationDTO dto.UpdateLocationDTO) (dto.ResponseLocationDTO, error) {
	query := `
		UPDATE locations
		SET number = $2, country = $3, city = $4, street = $5, image_url = $6
		WHERE id = $1
		RETURNING id, number, country, city, street, image_url
	`

	row := s.db.QueryRow(
		query,
		updateLocationDTO.ID,
		updateLocationDTO.Number,
		updateLocationDTO.Country,
		updateLocationDTO.City,
		updateLocationDTO.Street,
		updateLocationDTO.ImageURL,
	)

	var updatedLocation dto.ResponseLocationDTO
	err := row.Scan(
		&updatedLocation.ID,
		&updatedLocation.Number,
		&updatedLocation.Country,
		&updatedLocation.City,
		&updatedLocation.Street,
		&updatedLocation.ImageURL,
	)

	if err != nil {
		return dto.ResponseLocationDTO{}, err
	}

	return updatedLocation, nil
}

package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/constants"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
	"gorm.io/gorm"
)

type CountryService struct {
	database *gorm.DB
	logger   logging.Logger
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		database: db.GetDb(),
		logger:   logging.NewLogger(cfg),
	}
}

// CREATE
func (s *CountryService) CreateCountry(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {

	country := models.Country{Name: req.Name}
	country.CreatedBy = int(ctx.Value(constants.UserIDKey).(float64))
	country.CreatedAt = time.Now().UTC()

	tx := s.database.WithContext(ctx).Begin()
	err := tx.Create(&country).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return &dto.CountryResponse{Name: country.Name, ID: country.Id}, nil
}

// UPDATE
func (s *CountryService) UpdateCountry(ctx context.Context, id int, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {

	updateMap := map[string]interface{}{
		"name":        req.Name,
		"modified_by": &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIDKey).(float64))},
		"modified_at": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}
	tx := s.database.WithContext(ctx).Begin()

	err := tx.Model(&models.Country{}).Where("id = ?", id).Updates(updateMap).Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}
	country := &models.Country{}
	err = tx.Model(&models.Country{}).
		Where("id = ? AND deleted_by is null", id).
		First(&country).Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	tx.Commit()

	dto := &dto.CountryResponse{Name: country.Name, ID: country.Id}
	return dto, nil
}

// DELETE
func (s *CountryService) DeleteCountry(ctx context.Context, id int) error {
	tx := s.database.WithContext(ctx).Begin()
	DeleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIDKey).(float64))},
		"deleted_at": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	err := tx.Model(&models.Country{}).Where("id = ?", id).
		Updates(DeleteMap).Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Delete, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil
}

// GET

func (s *CountryService) GetCountryById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	country := &models.Country{}

	err := s.database.Where("id = ?", id).
		First(&country).Error

	if err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}

	return &dto.CountryResponse{Name: country.Name, ID: country.Id}, nil
}

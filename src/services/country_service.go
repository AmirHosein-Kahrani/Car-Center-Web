package services

import (
	"context"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
)

type CountryService struct {
	base *BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		base: &BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			// Preloads: []Preload{{string: "Cities.Region"}},
			Preloads: []preload{{string: "Cities"}},
		},
	}
}

// CREATE
func (s *CountryService) CreateCountry(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {

	return s.base.Create(ctx, req)

}

// UPDATE
func (s *CountryService) UpdateCountry(ctx context.Context, id int, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	return s.base.Update(ctx, id, req)

}

// DELETE
func (s *CountryService) DeleteCountry(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GET

func (s *CountryService) GetCountryById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CountryService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CountryResponse], error) {
	return s.base.GetByFilter(ctx, req)
}

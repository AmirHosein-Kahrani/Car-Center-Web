package services

import (
	"context"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
)

type PropertyService struct {
	base *BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]
}

func NewPropertyService(cfg *config.Config) *PropertyService {
	return &PropertyService{
		base: &BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			// Preloads: []Preload{{string: "Cities.Region"}},
			Preloads: []preload{{string: "Category"}},
		},
	}
}

// CREATE
func (s *PropertyService) CreateProperty(ctx context.Context, req *dto.CreatePropertyRequest) (*dto.PropertyResponse, error) {

	return s.base.Create(ctx, req)

}

// UPDATE
func (s *PropertyService) UpdateProperty(ctx context.Context, id int, req *dto.UpdatePropertyRequest) (*dto.PropertyResponse, error) {
	return s.base.Update(ctx, id, req)

}

// DELETE
func (s *PropertyService) DeleteProperty(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GET

func (s *PropertyService) GetPropertyById(ctx context.Context, id int) (*dto.PropertyResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *PropertyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PropertyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}

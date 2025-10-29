package services

import (
	"context"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
)

type PropertyCategoryService struct {
	base *BaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse]
}

func NewPropertyCategoryService(cfg *config.Config) *PropertyCategoryService {
	return &PropertyCategoryService{
		base: &BaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			// Preloads: []Preload{{string: "Cities.Region"}},
			Preloads: []preload{{string: "Properties"}},
		},
	}
}

// CREATE
func (s *PropertyCategoryService) CreatePropertyCategory(ctx context.Context, req *dto.CreatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {

	return s.base.Create(ctx, req)

}

// UPDATE
func (s *PropertyCategoryService) UpdatePropertyCategory(ctx context.Context, id int, req *dto.UpdatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {
	return s.base.Update(ctx, id, req)

}

// DELETE
func (s *PropertyCategoryService) DeletePropertyCategory(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GET

func (s *PropertyCategoryService) GetPropertyCategoryById(ctx context.Context, id int) (*dto.PropertyCategoryResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *PropertyCategoryService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PropertyCategoryResponse], error) {
	return s.base.GetByFilter(ctx, req)
}

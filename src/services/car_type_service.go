package services

import (
	"context"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
)

type CarTypeService struct {
	base *BaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse]
}

func NewCarTypeService(cfg *config.Config) *CarTypeService {
	return &CarTypeService{
		base: &BaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// CREATE
func (s *CarTypeService) CreateCarType(ctx context.Context, req *dto.CreateCarTypeRequest) (*dto.CarTypeResponse, error) {

	return s.base.Create(ctx, req)

}

// UPDATE
func (s *CarTypeService) UpdateCarType(ctx context.Context, id int, req *dto.UpdateCarTypeRequest) (*dto.CarTypeResponse, error) {
	return s.base.Update(ctx, id, req)

}

// DELETE
func (s *CarTypeService) DeleteCarType(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GET

func (s *CarTypeService) GetCarTypeById(ctx context.Context, id int) (*dto.CarTypeResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarTypeService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarTypeResponse], error) {
	return s.base.GetByFilter(ctx, req)
}

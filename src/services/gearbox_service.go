package services

import (
	"context"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]
}

func NewGearboxService(cfg *config.Config) *GearboxService {
	return &GearboxService{
		base: &BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// CREATE
func (s *GearboxService) CreateGearbox(ctx context.Context, req *dto.CreateGearboxRequest) (*dto.GearboxResponse, error) {

	return s.base.Create(ctx, req)

}

// UPDATE
func (s *GearboxService) UpdateGearbox(ctx context.Context, id int, req *dto.UpdateGearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Update(ctx, id, req)

}

// DELETE
func (s *GearboxService) DeleteGearbox(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GET

func (s *GearboxService) GetGearboxById(ctx context.Context, id int) (*dto.GearboxResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *GearboxService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GearboxResponse], error) {
	return s.base.GetByFilter(ctx, req)
}

package services

import (
	"context"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
)

type PersianYearService struct {
	base *BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]
}

func NewPersianYearService(cfg *config.Config) *PersianYearService {
	return &PersianYearService{
		base: &BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// CREATE
func (s *PersianYearService) CreatePersianYear(ctx context.Context, req *dto.CreatePersianYearRequest) (*dto.PersianYearResponse, error) {

	return s.base.Create(ctx, req)

}

// UPDATE
func (s *PersianYearService) UpdatePersianYear(ctx context.Context, id int, req *dto.UpdatePersianYearRequest) (*dto.PersianYearResponse, error) {
	return s.base.Update(ctx, id, req)

}

// DELETE
func (s *PersianYearService) DeletePersianYear(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GET

func (s *PersianYearService) GetPersianYearById(ctx context.Context, id int) (*dto.PersianYearResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *PersianYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PersianYearResponse], error) {
	return s.base.GetByFilter(ctx, req)
}

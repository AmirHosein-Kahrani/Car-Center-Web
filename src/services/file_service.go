package services

import (
	"context"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
)

type FileService struct {
	base *BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse]
}

func NewFileService(cfg *config.Config) *FileService {
	return &FileService{
		base: &BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// CREATE
func (s *FileService) CreateFile(ctx context.Context, req *dto.CreateFileRequest) (*dto.FileResponse, error) {

	return s.base.Create(ctx, req)

}

// UPDATE
func (s *FileService) UpdateFile(ctx context.Context, id int, req *dto.UpdateFileRequest) (*dto.FileResponse, error) {
	return s.base.Update(ctx, id, req)

}

// DELETE
func (s *FileService) DeleteFile(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GET

func (s *FileService) GetFileById(ctx context.Context, id int) (*dto.FileResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *FileService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.FileResponse], error) {
	return s.base.GetByFilter(ctx, req)
}

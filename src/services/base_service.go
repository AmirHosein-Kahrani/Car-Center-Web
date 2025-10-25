package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/common"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/constants"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/service_errors"
	"gorm.io/gorm"
)

// Generic 		country, create country, update country, Response country
type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Logger   logging.Logger
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {

	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDb(),
		Logger:   logging.NewLogger(cfg),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {

	model, _ := common.TypeConverter[T](req)
	tx := s.Database.WithContext(ctx).Begin()

	err := tx.Create(model).Error
	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return common.TypeConverter[Tr](model)
}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	// updateMap is refrence type
	updateMap, _ := common.TypeConverter[map[string]interface{}](req)
	// *DeRefrence The model
	(*updateMap)["modified_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIDKey).(float64))}
	(*updateMap)["modified_at"] = &sql.NullTime{Time: time.Now().UTC(), Valid: true}
	model := new(T)
	tx := s.Database.WithContext(ctx).Begin()
	err := tx.Model(model).
		Where("id = ? AND deleted_by is null", id).
		Updates(*updateMap).Error
	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return s.GetById(ctx, id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	tx := s.Database.WithContext(ctx).Begin()
	model := new(T)

	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIDKey).(float64))},
		"deleted_at": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}
	// *DeRefrence The model
	deleteMap["modified_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIDKey).(float64))}
	deleteMap["modified_at"] = &sql.NullTime{Time: time.Now().UTC(), Valid: true}

	if ctx.Value(constants.UserIDKey) == nil {
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}
	count := tx.Model(model).Where("id = ? and deleted_by is null", id).
		Updates(deleteMap).RowsAffected

	if count == 0 {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Update, service_errors.RecordNotFound, nil)
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}

	}
	tx.Commit()
	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {

	model := new(T)

	err := s.Database.Where("id = ? and deleted_by is null", id).
		First(model).Error

	if err != nil {
		return nil, err
	}

	return common.TypeConverter[Tr](model)
}

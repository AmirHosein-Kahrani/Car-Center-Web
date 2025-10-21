package services

import (
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/common"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/constants"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger          logging.Logger
	cfg             *config.Config
	otpService      *OtpService
	databaseService *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	databaseService := db.GetDb()
	logger := logging.NewLogger(cfg)

	return &UserService{
		cfg:             cfg,
		databaseService: databaseService,
		logger:          logger,
		otpService:      NewOtpService(cfg),
	}
}

func (s *UserService) SetOtp(req *dto.GetOtpRequest) error {

	otp := common.GenerateOtp()
	err := s.otpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil

}

func (s *UserService) existByEmail(email string) (bool, error) {
	var exists bool
	if err := s.databaseService.Model(&models.User{}).
		Select("count(*) > 0").Where("email = ?", email).
		Find(&exists).Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}

	return exists, nil
}

func (s *UserService) existByUsername(username string) (bool, error) {

	var exists bool
	if err := s.databaseService.Model(&models.User{}).Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existByMobileNumber(mobileNumber string) (bool, error) {
	var exists bool

	if err := s.databaseService.Model(&models.User{}).Select("count(*) > 0 ").
		Where("mobile_number = ?", mobileNumber).
		Find(&exists).Error; err != nil {

		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}

	return exists, nil
}

func (s *UserService) getDefaultRole() (roleId int, err error) {
	if err = s.databaseService.Model(&models.Role{}).
		Select("id").Where("name = ?", constants.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}

package services

import (
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/common"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
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


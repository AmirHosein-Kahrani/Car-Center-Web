package services

import (
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/common"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/constants"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	logger          logging.Logger
	cfg             *config.Config
	otpService      *OtpService
	databaseService *gorm.DB
	tokenservice    *TokenService
}

func NewUserService(cfg *config.Config) *UserService {
	databaseService := db.GetDb()
	logger := logging.NewLogger(cfg)

	return &UserService{
		cfg:             cfg,
		databaseService: databaseService,
		logger:          logger,
		otpService:      NewOtpService(cfg),
		tokenservice:    NewTokenService(cfg),
	}
}

// Login by UserName

func (s *UserService) LoginByUserName(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	var user models.User

	err := s.databaseService.Model(&models.User{}).
		Where("user_name = ?", req.Username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).Find(&user).Error

	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	tdto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
		Email: user.Email, MobileNumber: user.MobileNumber}
	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tdto.Roles = append(tdto.Roles, ur.Role.Name)
		}
	}

	token, err := s.tokenservice.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Register By Username
func (s *UserService) RegisterByUserName(req *dto.RegisterUserByUsernameRequest) error {
	u := models.User{UserName: req.Username, FirstName: req.FirstName,
		LastName: req.LastName, Email: req.Email}
	exists, err := s.existByEmail(req.Email)
	if err != nil {
		return nil
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}

	exists, err = s.existByUsername(req.Username)
	if err != nil {
		return nil
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UserNameExists}
	}

	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)

	if err != nil {
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	u.Password = string(hp)
	roleId, err := s.getDefaultRole()
	if err != nil {
		s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}
	// *Important
	tx := s.databaseService.Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil

}

// Register/Login by Mobile Number
func (s *UserService) RegisterLoginByMobileNumber(req *dto.RegisterLoginByMobileRequest) (*dto.TokenDetail, error) {

	err := s.otpService.ValidateOtp(req.MobileNumber, req.Otp)
	if err != nil {
		return nil, err
	}
	exists, err := s.existByMobileNumber(req.MobileNumber)
	if err != nil {
		return nil, err
	}

	u := models.User{MobileNumber: req.MobileNumber, UserName: req.MobileNumber}
	if exists {

		var user models.User

		err = s.databaseService.Model(&models.User{}).
			Where("user_name = ?", u.UserName).
			Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
				return tx.Preload("Role")
			}).Find(&user).Error

		if err != nil {
			return nil, err
		}
		tdto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
			Email: user.Email, MobileNumber: user.MobileNumber}
		if len(*user.UserRoles) > 0 {
			for _, ur := range *user.UserRoles {
				tdto.Roles = append(tdto.Roles, ur.Role.Name)
			}
		}
		token, err := s.tokenservice.GenerateToken(&tdto)
		if err != nil {
			return nil, err
		}
		return token, nil

	}

	bp := []byte(common.GeneratePassword())
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)

	if err != nil {
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return nil, err
	}
	u.Password = string(hp)
	roleId, err := s.getDefaultRole()
	if err != nil {
		s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return nil, err
	}
	// *Important
	tx := s.databaseService.Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return nil, err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return nil, err
	}
	tx.Commit()

	var user models.User
	err = s.databaseService.Model(&models.User{}).
		Where("user_name = ?", u.UserName).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).Find(&user).Error

	if err != nil {
		return nil, err
	}
	tdto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
		Email: user.Email, MobileNumber: user.MobileNumber}
	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tdto.Roles = append(tdto.Roles, ur.Role.Name)
		}
	}
	token, err := s.tokenservice.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}
	return token, nil

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
		Where("user_name = ?", username).
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

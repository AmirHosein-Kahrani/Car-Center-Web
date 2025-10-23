package services

import (
	"time"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/constants"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/service_errors"
	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	logger logging.Logger
	cfg    *config.Config
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	UserName     string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)

	return &TokenService{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *TokenService) GenerateToken(token *tokenDto) (*dto.TokenDetail, error) {

	tokenDetail := dto.TokenDetail{}
	tokenDetail.AccessTokenExpireTime = time.Now().Add(s.cfg.Jwt.AccessTokenExpireDuration * time.Minute).Unix()
	tokenDetail.RefreshTokenExpireTime = time.Now().Add(s.cfg.Jwt.RefreshTokenExpireDuration * time.Minute).Unix()

	accessTokenClaims := jwt.MapClaims{}

	accessTokenClaims[constants.UserIDKey] = token.UserId
	accessTokenClaims[constants.FirstNameKey] = token.FirstName
	accessTokenClaims[constants.LastNameKey] = token.LastName
	accessTokenClaims[constants.UserNameKey] = token.UserName
	accessTokenClaims[constants.EmailKey] = token.Email
	accessTokenClaims[constants.MobileNumberKey] = token.MobileNumber
	accessTokenClaims[constants.RolesKey] = token.Roles
	accessTokenClaims[constants.ExpireTimeKey] = tokenDetail.AccessTokenExpireTime

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	var err error
	tokenDetail.AccessToken, err = accessToken.SignedString([]byte(s.cfg.Jwt.Secret))
	if err != nil {
		return nil, err
	}
	refreshTokenClaims := jwt.MapClaims{}

	refreshTokenClaims[constants.UserIDKey] = token.UserId
	refreshTokenClaims[constants.ExpireTimeKey] = tokenDetail.RefreshTokenExpireTime
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenDetail.RefreshToken, err = refreshToken.SignedString([]byte(s.cfg.Jwt.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return &tokenDetail, nil
}

func (s *TokenService) VerifyToken(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnexpectedError}
		}
		return []byte(s.cfg.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *TokenService) GetClaims(t string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := s.VerifyToken(t)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}

}

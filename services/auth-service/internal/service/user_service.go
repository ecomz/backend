package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ecomz/backend/auth-service/internal/dto"
	"github.com/ecomz/backend/auth-service/internal/model"
	"github.com/ecomz/backend/auth-service/internal/repository"
	"github.com/ecomz/backend/libs/config"
	"github.com/ecomz/backend/libs/utils"
	"go.uber.org/zap"
)

type UserService interface {
	Login(email, password string) (res dto.LoginAndRegisiterResponse, err error)
	Register(user *model.User) (res dto.LoginAndRegisiterResponse, err error)
	CurrentUser(accessToken string) (res dto.UserResponse, err error)
}

type userService struct {
	logger   *zap.Logger
	cfg      *config.Config
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUserService(logger *zap.Logger, cfg *config.Config, userRepo repository.UserRepository, roleRepo repository.RoleRepository) UserService {
	return &userService{
		logger:   logger,
		cfg:      cfg,
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *userService) Register(user *model.User) (res dto.LoginAndRegisiterResponse, err error) {
	s.logger.Info("Register User ",
		zap.String("name", user.Name),
		zap.String("email", user.Email),
		zap.Int("role_id", user.RoleID),
	)
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		s.logger.Error("error hashing password", zap.Error(err))
		return res, err
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		s.logger.Error("error creating user", zap.Error(err))
		return res, err
	}

	role, err := s.roleRepo.GetRoleByID(user.RoleID)
	if err != nil {
		s.logger.Error("error getting role",
			zap.Error(err),
			zap.Int("role_id", user.RoleID),
		)
		return
	}

	accessToken, refreshToken, err := s.generateTokens(user)
	if err != nil {
		s.logger.Error("error generating tokens", zap.Error(err))
		return res, err
	}

	return dto.NewLoginResponse(user, role, accessToken, refreshToken), nil
}

func (s *userService) Login(email, password string) (res dto.LoginAndRegisiterResponse, err error) {
	s.logger.Info("Login User",
		zap.String("email", email),
	)

	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Error("user not found", zap.String("email", email))
			return res, fmt.Errorf("user not found")
		}
		s.logger.Error("error getting user by email", zap.Error(err), zap.String("email", email))
		return res, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		s.logger.Error("invalid password")
		return res, fmt.Errorf("invalid password")
	}

	role, err := s.roleRepo.GetRoleByID(user.RoleID)
	if err != nil {
		s.logger.Error("error getting role", zap.Error(err))
		return res, err
	}

	accessToken, refreshToken, err := s.generateTokens(user)
	if err != nil {
		s.logger.Error("error generating tokens", zap.Error(err))
		return res, err
	}

	return dto.NewLoginResponse(user, role, accessToken, refreshToken), nil
}

func (s *userService) CurrentUser(accessToken string) (res dto.UserResponse, err error) {
	claims, err := utils.ParseToken(accessToken, s.cfg.JWT.SecretKey)
	if err != nil {
		s.logger.Error("error parsing token", zap.Error(err))
		return res, err
	}

	user, err := s.userRepo.GetUserByEmail(claims.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Error("user not found", zap.String("email", claims.Email))
			return res, fmt.Errorf("user not found")
		}
		s.logger.Error("error getting user by email", zap.Error(err), zap.String("email", claims.Email))
		return res, err
	}

	role, err := s.roleRepo.GetRoleByID(user.RoleID)
	if err != nil {
		s.logger.Error("error getting role", zap.Error(err))
		return res, err
	}

	return dto.NewUserResponse(user, role), nil
}

func (s *userService) generateTokens(user *model.User) (accessToken, refreshToken string, err error) {
	s.logger.Info("Generate Tokens for User",
		zap.String("email", user.Email),
	)

	claimsAccessToken := utils.NewClaims(
		user.ID,
		user.Name,
		user.Email,
		s.cfg.App.Name,
		time.Now().Add(s.cfg.JWT.LoginExp),
	)
	accessToken, err = utils.GenerateToken(claimsAccessToken, s.cfg.JWT.SecretKey)
	if err != nil {
		s.logger.Error("error generating access token", zap.Error(err))
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	claimsRefreshToken := utils.NewClaims(
		user.ID,
		user.Name,
		user.Email,
		s.cfg.App.Name,
		time.Now().Add(s.cfg.JWT.RefreshExp),
	)
	refreshToken, err = utils.GenerateToken(claimsRefreshToken, s.cfg.JWT.SecretKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

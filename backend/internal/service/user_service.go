package service

import (
	"context"

	"github.com/maxzhirnov/formease/config"
	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/repository"
	"github.com/maxzhirnov/formease/internal/utils"
	"github.com/maxzhirnov/formease/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	config   *config.Config
	userRepo repository.UserRepository
	jwtUtil  *utils.JWTUtil
}

func NewUserService(config *config.Config, userRepo repository.UserRepository, jwtUtil *utils.JWTUtil) *UserService {
	return &UserService{
		config:   config,
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (s *UserService) Register(ctx context.Context, user *models.User) error {
	logger.Info("Registering user", zap.String("email", user.Email))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("error hashing password", zap.Error(err))
		return err
	}
	user.Password = string(hashedPassword)
	logger.Info("Hashed password", zap.String("hash", user.Password))
	return s.userRepo.Create(ctx, user)
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, string, *models.User, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	logger.Info("Login attempt", zap.String("email", email))
	if err != nil {
		logger.Error("error finding user", zap.Error(err))
		return "", "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logger.Error("invalid credentials", zap.Error(err))
		return "", "", nil, err
	}

	accessToken, err := s.jwtUtil.GenerateToken(user.ID.Hex(), user.Email, s.config.TokenExpirationHours)
	if err != nil {
		logger.Error("error generating access token", zap.Error(err))
		return "", "", nil, err
	}

	refreshToken, err := s.jwtUtil.GenerateRefreshToken(user.ID.Hex())
	if err != nil {
		logger.Error("error generating refresh token", zap.Error(err))
		return "", "", nil, err
	}

	// Store refresh token in the database
	if err := s.userRepo.UpdateRefreshToken(ctx, user.ID, refreshToken); err != nil {
		logger.Error("error updating refresh token", zap.Error(err))
		return "", "", nil, err
	}

	return accessToken, refreshToken, user, nil
}

func (s *UserService) ValidateToken(tokenString string) (string, string, error) {
	return s.jwtUtil.ValidateToken(tokenString)
}

func (s *UserService) RefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	user, err := s.userRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		logger.Error("Error finding user by refresh token", zap.Error(err))
		return "", "", err
	}
	logger.Info("Refresh token found", zap.String("refreshToken", refreshToken))

	accessToken, err := s.jwtUtil.GenerateToken(user.ID.Hex(), user.Email, s.config.TokenExpirationHours)
	if err != nil {
		logger.Error("Error generating access token", zap.Error(err))
		return "", "", err
	}
	logger.Info("Access token generated", zap.String("accessToken", accessToken))

	newRefreshToken, err := s.jwtUtil.GenerateRefreshToken(user.ID.Hex())
	if err != nil {
		logger.Error("Error generating new refresh token", zap.Error(err))
		return "", "", err
	}
	logger.Info("New refresh token generated", zap.String("refreshToken", newRefreshToken))

	// Update the stored refresh token in the database
	if err := s.userRepo.UpdateRefreshToken(ctx, user.ID, newRefreshToken); err != nil {
		logger.Error("Error updating refresh token", zap.Error(err))
		return "", "", err
	}
	logger.Info("Refresh token updated", zap.String("refreshToken", newRefreshToken))

	return accessToken, newRefreshToken, nil
}

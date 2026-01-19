package services

import (
	"errors"
	"tbo_backend/repositories"
	"tbo_backend/utils"
)

// AuthService defines the interface for authentication logic
type AuthService interface {
	Login(mobile string) (bool, error)
	ValidateOtp(mobile string, otp string) (string, bool, error)
}

// AuthServiceImpl implements AuthService
type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewAuthServiceImpl(userRepo repositories.UserRepository) AuthService {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

func (s *AuthServiceImpl) Login(mobile string) (bool, error) {
	// 1. Generate OTP
	otp, err := utils.GenerateOTP()
	if err != nil {
		return false, err
	}

	// 2. Upsert OTP to database
	if err := s.userRepo.UpsertLoginOtp(mobile, otp); err != nil {
		return false, err
	}

	// 3. Send OTP to mobile number
	if !utils.PushMobileOtp(mobile, otp) {
		return false, errors.New("failed to push mobile otp")
	}

	return true, nil
}

func (s *AuthServiceImpl) ValidateOtp(mobile string, otp string) (string, bool, error) {
	name, err := s.userRepo.ValidateOtp(mobile, otp)
	if err != nil {
		return "", false, err
	}
	if name == "" {
		// Name empty means not found/invalid in our repository logic
		return "", false, nil
	}
	return name, true, nil
}

package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"chekisvc/internal/domain/entity"
	repository "chekisvc/internal/domain/interface"
	"chekisvc/pkg/logger"
	"chekisvc/pkg/utils"
)

// UserUsecase represents the user's use cases
type UserUsecase struct {
	userRepo repository.UserRepository
	logger   logger.Logger
	timeout  time.Duration
}

// NewUserUsecase creates a new user use case instance
func NewUserUsecase(userRepo repository.UserRepository, logger logger.Logger) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		logger:   logger,
		timeout:  time.Second * 10,
	}
}

// AuthenticateUser authenticates a user by email and password
func (u *UserUsecase) AuthenticateUser(ctx context.Context, email, password string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := utils.CheckPassword(password, user.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

// CreateUser creates a new user
func (u *UserUsecase) CreateUser(ctx context.Context, user *entity.UserRegisterRequest) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	// Check if user with email already exists
	existingUser, err := u.userRepo.GetByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return errors.New("user with this email already exists")
	}
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	log.Println("Hashed Password:", hashPassword)
	user.Password = hashPassword

	// Create user
	return u.userRepo.Create(ctx, user)
}

// GetUserByID retrieves a user by ID
func (u *UserUsecase) GetUserByID(ctx context.Context, id uint) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepo.GetByID(ctx, id)
}

// UpdateUser updates an existing user
func (u *UserUsecase) UpdateUser(ctx context.Context, user *entity.UpdateUserRequest, userId uint) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	// Check if user exists
	existingUser, err := u.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return u.userRepo.Update(ctx, user, userId)
}

// DeleteUser deletes a user by ID
func (u *UserUsecase) DeleteUser(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	// Check if user exists
	existingUser, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return u.userRepo.Delete(ctx, id)
}

// ListUsers retrieves a list of users with pagination
func (u *UserUsecase) ListUsers(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepo.List(ctx, limit, offset)
}

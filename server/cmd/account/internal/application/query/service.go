package query

import (
	"context"
	"golang-ddd-cqrs/cmd/user/internal/domain"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID    uuid.UUID
	Email string
}

type UserQueryRepository interface {
	GetAll(ctx context.Context) ([]*domain.User, error)
	GetById(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetListings(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

type Service struct {
	repository UserQueryRepository
	// Pass in cache here, without defined methods. Cache is used directly
}

func NewUserQueryService(repository UserQueryRepository) *Service {
	return &Service{
		repository,
	}
}

// func (s *Service) GetUserListings(ctx context.Context, id uuid.UUID) (*GetUserListingsDto, error) {
// 	// Query cache for user listings
// 	data, err := redisClient.Get(ctx, key).Result()

// 	// Check for cache miss
// 	if err == redis.Nil {
// 		// On cache miss load from repository
// 		data, err := s.repository.GetListings(ctx, id)
// 		if err != nil {
// 			return nil, err
// 		}
// 		// On cache miss place data in cache
// 		redisClient.Set(ctx, "key", data, time.Hour)
// 	} else if err != nil {
// 		return &GetUserListingsDto{}, err

// 	}

// 	// Convert domain object to data transfer object
// 	dto := &GetUserListingsDto{mapToSingle(data)}

// 	// Return converted dto
// 	return dto, nil
// }

func (s *Service) GetUser(ctx context.Context, id uuid.UUID) (*GetUserDto, error) {
	user, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Convert domain object to data transfer object
	dto := &GetUserDto{mapToSingle(user)}

	return dto, nil
}

func (s *Service) GetUsers(ctx context.Context) (*GetUsersDto, error) {
	users, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Convert domain object to data transfer object
	dto := &GetUsersDto{mapToList(users)}

	return dto, nil
}

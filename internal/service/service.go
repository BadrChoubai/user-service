package service

import "context"

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) GetUserById(ctx context.Context, userId int) (*User, error) {
	return service.userRepository.SingleUserById(ctx, userId)
}

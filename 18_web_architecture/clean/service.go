package main

type Service interface {
	Get(uuid string) (*User, error)
}

type userService struct {
	repository Repository
}

func (s *userService) Get(uuid string) (*User, error) {
	err := ValidateUUID(uuid)
	if err != nil {
		return nil, err
	}

	user, err := s.repository.Get(uuid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func ValidateUUID(uuid string) error {
	return nil
}

func NewUserService(repository Repository) Service {
	return &userService{
		repository: repository,
	}
}

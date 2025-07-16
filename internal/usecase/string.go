package usecase

import "random-service/internal/entity"

type StringRepo interface {
	GetRandomString() (string, error)
}

type StringUseCase struct {
	repo StringRepo
}

func NewStringUseCase(r StringRepo) StringUseCase {
	return StringUseCase{repo: r}
}

func (uc StringUseCase) Get() (*entity.RandomString, error) {
	s, err := uc.repo.GetRandomString()
	if err != nil {
		return nil, err
	}
	return &entity.RandomString{Value: s}, nil
}

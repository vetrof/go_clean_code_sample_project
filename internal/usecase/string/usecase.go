package stringusecase

import "random-service/internal/entity"

type StringRepo interface {
	GetRandomString() (string, error)
}

type UseCase struct {
	repo StringRepo
}

func New(r StringRepo) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) Get() (*entity.RandomString, error) {
	s, err := uc.repo.GetRandomString()
	if err != nil {
		return nil, err
	}
	return &entity.RandomString{Value: s}, nil
}

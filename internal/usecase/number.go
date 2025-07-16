package usecase

import "random-service/internal/entity"

type NumberRepo interface {
	GetRandomNumber() (int, error)
}

type NumberUseCase struct {
	repo NumberRepo
}

func NewNumberUseCase(r NumberRepo) NumberUseCase {
	return NumberUseCase{repo: r}
}

func (uc NumberUseCase) Get() (*entity.RandomNumber, error) {
	n, err := uc.repo.GetRandomNumber()
	if err != nil {
		return nil, err
	}
	return &entity.RandomNumber{Value: n}, nil
}

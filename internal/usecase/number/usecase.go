package numberusecase

import "random-service/internal/entity"

type NumberRepo interface {
	GetRandomNumber() (int, error)
}

type UseCase struct {
	repo NumberRepo
}

func New(r NumberRepo) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) Get() (*entity.RandomNumber, error) {
	n, err := uc.repo.GetRandomNumber()
	if err != nil {
		return nil, err
	}
	return &entity.RandomNumber{Value: n}, nil
}

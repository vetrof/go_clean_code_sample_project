package string

type Repository interface {
	Get() string
}

type UseCase struct {
	repo Repository
}

func New(r Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) Do() string {
	return uc.repo.Get()
}

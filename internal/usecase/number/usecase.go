package number

type Repository interface {
	Get() int
}

type UseCase struct {
	repo Repository
}

func New(r Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) Do() int {
	return uc.repo.Get()
}

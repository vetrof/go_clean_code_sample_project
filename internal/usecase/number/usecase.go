package number

// Repository - интерфейс, реализуемый в репозитории.
type Repository interface {
	Get() int
}

// UseCase - реализация бизнес-логики, использует интерфейс репозитория.
type UseCase struct {
	repo Repository
}

func New(r Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) Do() int {
	return uc.repo.Get()
}

package number

import (
	"math/rand"
)

// Repo - конкретная реализация интерфейса Repository из usecase/number.
type Repo struct{}

func New() *Repo {
	return &Repo{}
}

// Get реализует интерфейс number.Repository
func (r *Repo) Get() int {
	return rand.Intn(1000)
}

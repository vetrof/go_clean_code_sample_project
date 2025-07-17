package number

import "math/rand"

type Repo struct{}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Get() int {
	return rand.Intn(1000)
}

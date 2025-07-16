package stringrepo

import "math/rand"

type MemoryRepo struct {
	values []string
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		values: []string{"clean", "arch", "go", "awesome"},
	}
}

func (r *MemoryRepo) GetRandomString() (string, error) {
	return r.values[rand.Intn(len(r.values))], nil
}

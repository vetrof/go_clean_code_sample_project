package numberrepo

import (
	"math/rand"
)

type MemoryRepo struct{}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{}
}

func (r *MemoryRepo) GetRandomNumber() (int, error) {
	return rand.Intn(1000), nil
}

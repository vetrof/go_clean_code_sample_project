package string

import (
	"math/rand"
)

type Repo struct{}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Get() string {
	chars := []rune("abcdefghijklmnopqrstuvwxyz")
	str := make([]rune, 8)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

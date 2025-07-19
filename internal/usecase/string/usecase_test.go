package string

import (
	"testing"
)

type mockRepo struct{}

func (m *mockRepo) Get() string {
	return "42"
}

func TestUseCase_Do(t *testing.T) {
	repo := &mockRepo{}
	uc := New(repo)

	result := uc.Do()
	expected := "42"

	if result != expected {
		t.Errorf("ожидалось %s, получено %s", expected, result)
	}
}

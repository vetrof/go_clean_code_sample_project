package number

import (
	"testing"
)

// Мок-реализация репозитория для теста
type mockRepo struct{}

func (m *mockRepo) Get() int {
	return 42
}

func TestUseCase_Do(t *testing.T) {
	repo := &mockRepo{}
	uc := New(repo)

	result := uc.Do()
	expected := 42

	if result != expected {
		t.Errorf("ожидалось %d, получено %d", expected, result)
	}
}

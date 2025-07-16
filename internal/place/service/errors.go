package service

import "fmt"

// ServiceError базовый интерфейс для всех ошибок сервисного слоя
type ServiceError interface {
	error
	Type() string
	Code() int
}

// ValidationError ошибка валидации входных данных
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.Message)
}

func (e *ValidationError) Type() string {
	return "validation"
}

func (e *ValidationError) Code() int {
	return 400 // Bad Request
}

// NewValidationError создает новую ошибку валидации
func NewValidationError(message string) *ValidationError {
	return &ValidationError{Message: message}
}

// RepositoryError ошибка при работе с репозиторием
type RepositoryError struct {
	Message string
	Cause   error
}

func (e *RepositoryError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("repository error: %s, cause: %v", e.Message, e.Cause)
	}
	return fmt.Sprintf("repository error: %s", e.Message)
}

func (e *RepositoryError) Type() string {
	return "repository"
}

func (e *RepositoryError) Code() int {
	return 500 // Internal Server Error
}

// NewRepositoryError создает новую ошибку репозитория
func NewRepositoryError(message string, cause error) *RepositoryError {
	return &RepositoryError{
		Message: message,
		Cause:   cause,
	}
}

// NotFoundError ошибка "не найдено"
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.Message)
}

func (e *NotFoundError) Type() string {
	return "not_found"
}

func (e *NotFoundError) Code() int {
	return 404 // Not Found
}

// NewNotFoundError создает новую ошибку "не найдено"
func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{Message: message}
}

// BusinessLogicError ошибка бизнес-логики
type BusinessLogicError struct {
	Message string
}

func (e *BusinessLogicError) Error() string {
	return fmt.Sprintf("business logic error: %s", e.Message)
}

func (e *BusinessLogicError) Type() string {
	return "business_logic"
}

func (e *BusinessLogicError) Code() int {
	return 422 // Unprocessable Entity
}

// NewBusinessLogicError создает новую ошибку бизнес-логики
func NewBusinessLogicError(message string) *BusinessLogicError {
	return &BusinessLogicError{Message: message}
}

// IsValidationError проверяет, является ли ошибка ValidationError
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

// IsRepositoryError проверяет, является ли ошибка RepositoryError
func IsRepositoryError(err error) bool {
	_, ok := err.(*RepositoryError)
	return ok
}

// IsNotFoundError проверяет, является ли ошибка NotFoundError
func IsNotFoundError(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

// GetErrorCode возвращает HTTP код ошибки
func GetErrorCode(err error) int {
	if serviceErr, ok := err.(ServiceError); ok {
		return serviceErr.Code()
	}
	return 500 // Internal Server Error по умолчанию
}

// GetErrorType возвращает тип ошибки
func GetErrorType(err error) string {
	if serviceErr, ok := err.(ServiceError); ok {
		return serviceErr.Type()
	}
	return "unknown"
}

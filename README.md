# Place Service - Учебный проект по чистой архитектуре

Этот проект демонстрирует принципы чистой архитектуры (Clean Architecture) на примере простого веб-сервиса для работы с местами.

## 🎯 Цели проекта

- Показать правильное разделение слоев в Go приложении
- Продемонстрировать использование интерфейсов для инверсии зависимостей
- Создать понятную и расширяемую архитектуру
- Показать, как каждый компонент (репозиторий, сервис, хендлер) выделен в отдельный файл

## 📋 Функциональность

- **GET /api/v1/nearby?lat=55.7&lng=37.6** - поиск ближайших мест по координатам
- **GET /api/v1/popular** - получение списка популярных мест
- **GET /api/v1/places/:id** - получение места по ID
- **GET /api/v1/places** - получение всех мест
- **GET /api/v1/health** - проверка состояния сервиса

## 🏗️ Архитектура

Проект построен по принципам чистой архитектуры с четким разделением на слои:

```
┌─────────────────────────────────────────────────────┐
│                 HTTP Handler Layer                  │
│  (Обработка HTTP запросов)                         │
├─────────────────────────────────────────────────────┤
│                 Service Layer                       │
│  (Бизнес-логика приложения)                        │
├─────────────────────────────────────────────────────┤
│                Repository Layer                     │
│  (Работа с данными)                                │
├─────────────────────────────────────────────────────┤
│                 Models Layer                        │
│  (Структуры данных)                                │
└─────────────────────────────────────────────────────┘
```

### Направление зависимостей:
- **Handler** → **Service** (через интерфейсы)
- **Service** → **Repository** (через интерфейсы)
- **Repository** → **Models** (напрямую)

## 📂 Структура проекта

```
cmd/
  main.go                         # Точка входа, сборка зависимостей
internal/
  place/
    models/
      place.go                    # Модель Place и методы
    repository/
      interfaces.go               # Интерфейсы репозиториев
      nearby_repository.go        # Репозиторий для ближайших мест
      popular_repository.go       # Репозиторий для популярных мест
      place_repository.go         # Репозиторий для работы с местами по ID
    service/
      interfaces.go               # Интерфейсы сервисов
      errors.go                   # Система ошибок
      nearby_service.go           # Сервис поиска ближайших мест
      popular_service.go          # Сервис популярных мест
      place_service.go            # Сервис работы с местами по ID
    handler/
      nearby_handler.go           # HTTP хендлер для ближайших мест
      popular_handler.go          # HTTP хендлер для популярных мест
      place_handler.go            # HTTP хендлер для работы с местами
    router/
      router.go                   # Настройка маршрутов
examples/
  api_examples.md                 # Примеры использования API
go.mod                            # Зависимости модуля
go.sum                            # Контрольные суммы зависимостей
```

## 🔧 Принципы SOLID в действии

### 1. **Single Responsibility Principle (SRP)**
- `nearby_repository.go` - только работа с ближайшими местами
- `popular_service.go` - только бизнес-логика для популярных мест
- `place_handler.go` - только HTTP обработка для мест

### 2. **Open/Closed Principle (OCP)**
- Легко добавить новый репозиторий (например, PostgreSQL)
- Легко добавить новый сервис для новой функциональности
- Легко добавить новый хендлер для новых endpoints

### 3. **Liskov Substitution Principle (LSP)**
- Любая реализация `NearbyRepository` может быть заменена на другую
- Mock репозитории легко заменяются на реальные

### 4. **Interface Segregation Principle (ISP)**
- `NearbyService`, `PopularService`, `PlaceService` - отдельные интерфейсы
- Каждый интерфейс содержит только нужные методы

### 5. **Dependency Inversion Principle (DIP)**
- Сервисы зависят от интерфейсов репозиториев
- Хендлеры зависят от интерфейсов сервисов
- Высокоуровневые модули не зависят от низкоуровневых

## 🚀 Запуск проекта

```bash
# Клонировать репозиторий
git clone <repo-url>
cd go_interface_sample_project

# Установить зависимости
go mod tidy

# Запустить проект
go run cmd/main.go

# Тестировать API
curl "http://localhost:8080/api/v1/health"
curl "http://localhost:8080/api/v1/nearby?lat=55.7&lng=37.6"
curl "http://localhost:8080/api/v1/popular"
curl "http://localhost:8080/api/v1/places/1"
```

## 🧪 Примеры расширения

### Добавление нового репозитория (PostgreSQL)

```go
// internal/place/repository/postgres_repository.go
type PostgresNearbyRepository struct {
    db *sql.DB
}

func NewPostgresNearbyRepository(db *sql.DB) NearbyRepository {
    return &PostgresNearbyRepository{db: db}
}

func (r *PostgresNearbyRepository) GetNearbyPlaces(lat, lng float64) ([]models.Place, error) {
    // Реализация работы с PostgreSQL
}
```

### Добавление нового сервиса

```go
// internal/place/service/review_service.go
type ReviewService interface {
    GetReviewsByPlaceID(placeID int) ([]models.Review, error)
    AddReview(review models.Review) error
}

type ReviewServiceImpl struct {
    reviewRepo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) ReviewService {
    return &ReviewServiceImpl{reviewRepo: repo}
}
```

### Добавление нового хендлера

```go
// internal/place/handler/review_handler.go
type ReviewHandler struct {
    reviewService service.ReviewService
}

func NewReviewHandler(service service.ReviewService) *ReviewHandler {
    return &ReviewHandler{reviewService: service}
}

func (h *ReviewHandler) GetReviews(c *gin.Context) {
    // Обработка HTTP запроса
}
```

## 📊 Преимущества такой архитектуры

### ✅ **Тестируемость**
- Каждый компонент можно тестировать независимо
- Легко создавать mock'и для интерфейсов
- Быстрые unit тесты без внешних зависимостей

### ✅ **Масштабируемость**
- Легко добавлять новые функции
- Каждый слой может развиваться независимо
- Простое разделение на микросервисы в будущем

### ✅ **Поддерживаемость**
- Четкое разделение ответственности
- Каждый файл имеет единую цель
- Легко найти и изменить нужную логику

### ✅ **Гибкость**
- Легко менять реализации (mock → database)
- Простое добавление новых транспортных слоев (gRPC, CLI)
- Независимость от внешних библиотек

## 🔍 Детали реализации

### Разделение по файлам

**Почему каждый репозиторий/сервис/хендлер в отдельном файле?**

1. **Читаемость** - легче найти нужную логику
2. **Поддержка** - изменения в одном компоненте не затрагивают другие
3. **Тестирование** - каждый компонент можно тестировать отдельно
4. **Командная разработка** - меньше конфликтов в git

### Система ошибок

```go
// Типизированные ошибки для лучшей обработки
type ValidationError struct {
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error: %s", e.Message)
}

func (e *ValidationError) Code() int {
    return 400 // Bad Request
}
```

### Интерфейсы

```go
// Интерфейсы определены в тех же пакетах, где используются
type NearbyService interface {
    FindNearby(lat, lng float64) ([]models.Place, error)
}

// Реализация может быть в любом месте
type NearbyServiceImpl struct {
    nearbyRepo repository.NearbyRepository
}
```

## 💡 Выводы

Этот проект демонстрирует:

1. **Как правильно разделить код на слои**
2. **Как использовать интерфейсы для инверсии зависимостей**
3. **Как каждый компонент выделить в отдельный файл**
4. **Как сделать код тестируемым и расширяемым**
5. **Как применить принципы SOLID на практике**

---

*Этот проект создан для обучения принципам чистой архитектуры в Go. Используйте его как основу для понимания правильного разделения кода на слои и компоненты.*
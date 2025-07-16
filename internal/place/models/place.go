package models

// Place представляет сущность "Место"
type Place struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
	Photo string  `json:"photo"`
}

// IsValid проверяет валидность места
func (p *Place) IsValid() bool {
	if p.Name == "" {
		return false
	}
	if p.Lat < -90 || p.Lat > 90 {
		return false
	}
	if p.Lng < -180 || p.Lng > 180 {
		return false
	}
	return true
}

// DistanceTo вычисляет простое расстояние до точки
func (p *Place) DistanceTo(lat, lng float64) float64 {
	dlat := p.Lat - lat
	dlng := p.Lng - lng
	return dlat*dlat + dlng*dlng
}

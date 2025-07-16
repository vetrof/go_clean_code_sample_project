package httpserver

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func New(app *fiber.App) *Server {
	return &Server{app: app}
}

func (s *Server) Start() error {
	return s.app.Listen(":8080")
}

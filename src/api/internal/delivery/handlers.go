package delivery_api

import (
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	ConcateLogic(a, b int) int
}

type BasicAuth interface {
}

type JWTAuth interface {
}

type Handler struct {
	service Service
	app     *fiber.App
}

func NewHandlers(service Service) *Handler {
	h := &Handler{
		service: service,
		app:     fiber.New(),
	}

	h.app.Get("api/concate", h.Concate)

	return h
}

func (h *Handler) Concate(c *fiber.Ctx) error {
	a := 1
	b := 2
	response := h.service.ConcateLogic(a, b)
	return c.JSON(response)
}

func (h *Handler) BasicAuthorization() {

}

func (h *Handler) JWTAuthorization() {

}

func (h *Handler) Listen(host string) error {
	return h.app.Listen(host)
}

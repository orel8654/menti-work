package delivery_api

import (
	"context"
	"menti/pkg/types"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	ConcateLogic(a, b int) int
	MiddleWareBasic(ctx context.Context, credentials string) error
	MiddleWareJWT(ctx context.Context, credentials string) error
	RegisterUserService(ctx context.Context, credentials types.UserPayload) error
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
	h.app.Post("api/register", h.RegisterUser)

	return h
}

func (h *Handler) Concate(c *fiber.Ctx) error {
	a := 1
	b := 2
	response := h.service.ConcateLogic(a, b)
	return c.JSON(response)
}

func (h *Handler) RegisterUser(c *fiber.Ctx) error {
	data := new(types.UserPayload)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(types.ResponseError{"error": err.Error()})
	}
	if err := h.service.RegisterUserService(c.Context(), *data); err != nil {
		return c.Status(500).JSON(types.ResponseError{"error": err.Error()})
	}
	return nil
}

func (h *Handler) BasicAuthorization(c *fiber.Ctx) error {
	token := ""
	if err := h.service.MiddleWareBasic(c.Context(), token); err != nil {
		return c.Status(403).JSON(types.ResponseError{"error": err.Error()})
	}
	return c.Next()
}

func (h *Handler) JWTAuthorization(c *fiber.Ctx) error {
	if err := h.service.MiddleWareJWT(c.Context(), "orel"); err != nil {
		return c.Status(403).JSON(types.ResponseError{"error": err.Error()})
	}
	return c.Next()
}

func (h *Handler) Listen(host string) error {
	return h.app.Listen(host)
}

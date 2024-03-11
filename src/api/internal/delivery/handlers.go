package delivery_api

import (
	"context"

	"menti/pkg/types"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	RegisterUserService(ctx context.Context, credentials types.UserPayload) error
	LoginUserService(ctx context.Context, credentials types.UserPayload) (res types.User, err error)
	LoginUserByBasicService(ctx context.Context, credentials types.UserPayload) (res types.UserToken, err error)
	LoginUserByBearerService(ctx context.Context, credentials types.UserPayload) (res types.UserToken, err error)
	AuthMiddlewareService(ctx context.Context, parts []string) (user types.User, err error)
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

	h.app.Post("api/register", h.RegisterUser)
	h.app.Post("api/login", h.LoginUser)
	h.app.Post("api/v1/login", h.LoginUserByBasic)
	h.app.Post("api/v2/login", h.LoginUserByBearer)
	h.app.Get("api/v1/test", h.AuthMiddleware, h.TestFunc)

	return h
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

func (h *Handler) LoginUser(c *fiber.Ctx) error {
	data := new(types.UserPayload)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(types.ResponseError{"error": err.Error()})
	}
	user, err := h.service.LoginUserService(c.Context(), *data)
	if err != nil {
		return c.Status(400).JSON(types.ResponseError{"error": err.Error()})
	}
	return c.Status(200).JSON(types.ResponseLogin{"uuid": user.UUID})
}

func(h *Handler) LoginUserByBasic(c *fiber.Ctx) error {
	data := new(types.UserPayload)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(types.ResponseError{"error": err.Error()})
	}
	token, err := h.service.LoginUserByBasicService(c.Context(), *data)
	if err != nil {
		return c.Status(500).JSON(types.ResponseError{"error": err.Error()})
	}
	return c.Status(200).JSON(types.ResponseLogin{"token": token.Token})
}

func (h *Handler) LoginUserByBearer(c *fiber.Ctx) error {
	data := new(types.UserPayload)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(types.ResponseError{"error": err.Error()})
	}
	token, err := h.service.LoginUserByBearerService(c.Context(), *data)
	if err != nil {
		return c.Status(500).JSON(types.ResponseError{"error": err.Error()})
	}
	return c.Status(200).JSON(types.ResponseLogin{"token": token.Token})
}

func (h *Handler) AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	parts := strings.Split(authHeader, " ")
	user, err := h.service.AuthMiddlewareService(c.Context(), parts)
	if err != nil {
		return c.Status(403).JSON(types.ResponseError{"error": err.Error()})
	}
	c.Locals("user", user)
	return c.Next()
}

func (h *Handler) TestFunc(c *fiber.Ctx) error {
	user := c.Locals("user").(types.User)
	_ = user
	return c.Status(200).JSON(types.ResponseLogin{"status": "ok"})
}

func (h *Handler) Listen(host string) error {
	return h.app.Listen(host)
}

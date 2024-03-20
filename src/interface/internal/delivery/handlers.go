package delivery

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

type Service interface {
    
}

type Handler struct {
	app *fiber.App
}

func NewHandlers() *Handler {
	h := &Handler{
//		service: service,
		app: fiber.New(),
	}
	
	h.app.Get("login", h.Login)
	h.app.Get("static/css/:filename", h.StaticCSS)
	h.app.Get("static/js/:filename", h.StaticJS)
	
	return h
}

func (h *Handler) Login(c *fiber.Ctx) error {
	filePath := "../../../pkg/templates/auth/login.html"
	return c.SendFile(filePath)
}

func (h *Handler) StaticCSS(c *fiber.Ctx) error {
	fileName := c.Params("filename")
	filePath := fmt.Sprintf("../../../pkg/static/css/%s", fileName)
	return c.SendFile(filePath)
}

func (h *Handler) StaticJS(c *fiber.Ctx) error{
	fileName := c.Params("filename")
	filePath := fmt.Sprintf("../../../pkg/static/js/%s", fileName)
	return c.SendFile(filePath)
}

func (h *Handler) Listen(host string) error {
	return h.app.Listen(host)
}
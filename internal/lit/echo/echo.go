package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/m-nny/go-lit/internal/lit"
)

type HttpServer struct {
	e *echo.Echo

	config *lit.HttpConfig

	userService lit.UserService
}

func NewHttp(config *lit.HttpConfig, userService *lit.UserService) *HttpServer {
	e := echo.New()
	e.Debug = config.Debug
	return &HttpServer{
		e:           e,
		config:      config,
		userService: *userService,
	}
}

func (h *HttpServer) Open() (err error) {
	listenAddress := h.config.Address()

	h.RegisterUserController()

	if err := h.e.Start(listenAddress); err != nil {
		return err
	}

	return nil
}

func (h *HttpServer) Close() (err error) {
	if err := h.e.Close(); err != nil {
		return err
	}

	return nil
}

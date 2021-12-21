package echo

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/m-nny/go-lit/internal/lit"
)

type HttpServer struct {
	e *echo.Echo

	config *lit.HttpConfig

	userService lit.UserService
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewHttp(config *lit.HttpConfig, userService *lit.UserService) *HttpServer {
	e := echo.New()
	e.Debug = config.Debug
	e.Validator = &CustomValidator{validator: validator.New()}
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

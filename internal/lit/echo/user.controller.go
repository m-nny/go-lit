package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/m-nny/go-lit/internal/lit"
)

func (h *HttpServer) RegisterUserController() *echo.Group {
	g := h.e.Group("/users")

	g.GET("/", h.handleGetUsers)

	return g
}

type GetUsersResult struct {
	Items []*lit.User `json:"users"`
	Count int         `json:"count"`
}

func (h *HttpServer) handleGetUsers(c echo.Context) (err error) {
	usersFilter := lit.UserFilter{}
	items, count, err := h.userService.FindUsers(c.Request().Context(), usersFilter)
	if err != nil {
		return err
	}
	result := &GetUsersResult{
		Items: items,
		Count: count,
	}

	return c.JSON(http.StatusOK, result)
}

package echo

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/m-nny/go-lit/internal/lit"
)

func (h *HttpServer) RegisterUserController() *echo.Group {
	g := h.e.Group("/users")

	g.GET("/", h.handleGetUsers)
	g.GET("/:userId", h.handleGetUserById)

	return g
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

func (h *HttpServer) handleGetUserById(c echo.Context) (err error) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 0, 0)
	if err != nil {
		return err
	}
	item, err := h.userService.FindUserById(c.Request().Context(), uint(userId))
	if err != nil {
		// TODO(m-nny): move to error hanlder middleware
		if lit.ErrorCode(err) == lit.ENOTFOUND {
			return c.JSON(http.StatusNotFound, err)
		}
		return err
	}
	return c.JSON(http.StatusOK, item)
}

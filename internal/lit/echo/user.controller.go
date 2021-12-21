package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/m-nny/go-lit/internal/lit"
)

func (h *HttpServer) RegisterUserController() *echo.Group {
	g := h.e.Group("/users")
	{
		g.GET("/", h.handleGetUsers)
		g.POST("/", h.handleCreateUser)
		g.GET("/:userId", h.handleGetUserById)
		g.POST("/:userId", h.handleUpdateUser)
	}

	return g
}

func (h *HttpServer) handleGetUsers(c echo.Context) (err error) {
	usersFilter := &lit.UserFilter{}
	items, count, err := h.userService.FindUsers(c.Request().Context(), usersFilter)
	if err != nil {
		return err
	}
	result := &echo.Map{
		"items": items,
		"count": count,
	}

	return c.JSON(http.StatusOK, result)
}

func (h *HttpServer) handleGetUserById(c echo.Context) (err error) {
	userId, err := h.bindUserId(c)
	if err != nil {
		return err
	}
	item, err := h.userService.FindUserById(c.Request().Context(), userId)
	if err != nil {
		// TODO(m-nny): move to error hanlder middleware
		if lit.ErrorCode(err) == lit.ENOTFOUND {
			return c.JSON(http.StatusNotFound, err)
		}
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func (h *HttpServer) handleCreateUser(c echo.Context) (err error) {
	createUserDto := new(CreateUserArg)
	if err := c.Bind(createUserDto); err != nil {
		return err
	}
	if err := c.Validate(createUserDto); err != nil {
		return err
	}
	user, err := createUserDto.User()
	if err != nil {
		return err
	}
	createdUser, err := h.userService.CreateUser(c.Request().Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func (h *HttpServer) handleUpdateUser(c echo.Context) (err error) {
	userId, err := h.bindUserId(c)
	if err != nil {
		return err
	}
	updateUserDto := new(UpdateUserArg)
	if err := c.Bind(updateUserDto); err != nil {
		return err
	}
	if err := c.Validate(updateUserDto); err != nil {
		return err
	}

	user, err := updateUserDto.User()
	if err != nil {
		return err
	}

	updatedUser, err := h.userService.UpdateUser(c.Request().Context(), userId, user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, updatedUser)
}

func (h *HttpServer) bindUserId(c echo.Context) (userId uint, err error) {
	userIdDto := new(UserIdArg)
	if err := c.Bind(userIdDto); err != nil {
		return userIdDto.UserId, err
	}
	if err := c.Validate(userIdDto); err != nil {
		return userIdDto.UserId, err
	}
	return userIdDto.UserId, nil
}

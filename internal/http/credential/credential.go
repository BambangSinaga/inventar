package credential

import (
	"context"
	"net/http"

	"github.com/West-Labs/inventar"
	"github.com/labstack/echo"
)

// ResponseError wraps json error response
type ResponseError struct {
	Message string `json:"message"`
}

type httpHandler struct {
	service inventar.UserService
}

func (h *httpHandler) HandleSignupUser(c echo.Context) error {

	var u *inventar.Credential

	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ResponseError{Message: "Can't Process Entity"})
	}

	ctx := context.Background()
	credNew, err := h.service.Signup(ctx, u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, credNew)
}

func (h *httpHandler) HandleSigninUser(c echo.Context) error {

	var u *inventar.Credential

	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ResponseError{Message: "Can't Process Entity"})
	}

	ctx := context.Background()
	credNew, err := h.service.Signin(ctx, u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, credNew)
}

func Init(e *echo.Echo, service inventar.UserService) {
	handler := httpHandler{service: service}
	e.POST("/user/signup", handler.HandleSignupUser)
	e.POST("/user/signin", handler.HandleSigninUser)
}

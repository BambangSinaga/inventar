package credential

import (
	"context"
	"net/http"

	"github.com/West-Labs/inventar"
	"github.com/labstack/echo"
)

// ResponseError wraps json error response
type Response struct {
	Message string `json:"message"`
}

type httpHandler struct {
	service inventar.UserService
}

func (h *httpHandler) HandleSignupUser(c echo.Context) error {

	var u *inventar.Credential

	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &Response{Message: err.Error()})
	}

	ctx := context.Background()
	credNew, err := h.service.Signup(ctx, u)
	if err != nil || !credNew {
		return c.JSON(http.StatusInternalServerError, &Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &Response{Message: "Sign Up Success"})
}

func (h *httpHandler) HandleSigninUser(c echo.Context) error {

	var u *inventar.Credential

	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &Response{Message: err.Error()})
	}

	ctx := context.Background()
	credNew, err := h.service.Signin(ctx, u)
	if err != nil || !credNew {
		return c.JSON(http.StatusInternalServerError, &Response{err.Error()})
	}

	return c.JSON(http.StatusOK, &Response{Message: "Login Success"})
}

func Init(e *echo.Echo, service inventar.UserService) {
	handler := httpHandler{service: service}
	e.POST("/user/signup", handler.HandleSignupUser)
	e.POST("/user/signin", handler.HandleSigninUser)
}

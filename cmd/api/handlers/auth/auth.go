package auth

import (
	"log"
	"net/http"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/service"
	"github.com/DitoAdriel99/go-monsterdex/pkg/response"
	"github.com/labstack/echo/v4"
)

type _Handlers struct {
	service *service.Service
}

func NewHandlers(service *service.Service) *_Handlers {
	return &_Handlers{service: service}
}

func (h *_Handlers) RegisterHandler(c echo.Context) error {
	var (
		payload      entity.RegisterPayload
		succResponse = response.NewResponse().WithStatus("success").WithMessage("success")
		errResponse  = response.NewResponse().WithStatus("error").WithMessage("error")
	)

	log.Println("register....")

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errResponse.WithError(err))
	}

	respData, err := h.service.AuthService.Register(&payload)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errResponse.WithError(err))
	}

	log.Println("register success...")

	return c.JSON(http.StatusOK, succResponse.WithData(respData))
}

func (h *_Handlers) LoginHandler(c echo.Context) error {
	var (
		payload      entity.Login
		succResponse = response.NewResponse().WithStatus("success").WithMessage("success")
		errResponse  = response.NewResponse().WithStatus("error").WithMessage("error")
	)

	log.Println("login....")

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errResponse.WithError(err))
	}

	respData, err := h.service.AuthService.Login(&payload)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errResponse.WithError(err))
	}

	log.Println("login success....")
	return c.JSON(http.StatusOK, succResponse.WithData(respData))

}

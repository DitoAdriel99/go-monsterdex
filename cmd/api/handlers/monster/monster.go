package monster

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/service"
	"github.com/DitoAdriel99/go-monsterdex/pkg/meta"
	"github.com/DitoAdriel99/go-monsterdex/pkg/response"
	"github.com/labstack/echo/v4"
)

type _Handlers struct {
	service *service.Service
}

func NewHandlers(service *service.Service) *_Handlers {
	return &_Handlers{service: service}
}

func (h *_Handlers) CreateMonsterHandler(c echo.Context) error {
	var (
		payload      entity.MonsterPayload
		succResponse = response.NewResponse().WithStatus("success").WithMessage("success")
		errResponse  = response.NewResponse().WithStatus("error").WithMessage("error")
	)

	log.Println("create monster....")

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
	}

	data, err := h.service.MonsterService.Create(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
	}

	log.Println("create monster success....")
	return c.JSON(http.StatusOK, succResponse.WithData(data))
}

func (h *_Handlers) GetMonstersHandler(c echo.Context) error {
	var (
		succResponse = response.NewResponse().WithStatus("success").WithMessage("success")
		errResponse  = response.NewResponse().WithStatus("error").WithMessage("error")
		query        = c.Request().URL.Query()
		metadata     = meta.MetadataFromURL(query)
		bearer       = c.Request().Header.Get("Authorization")
	)
	log.Println("get monster....")

	if metadata.MonsterID != 0 {
		data, err := h.service.MonsterService.GetID(bearer, metadata.MonsterID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
		}

		log.Println("get monster success....")
		return c.JSON(http.StatusOK, succResponse.WithData(data))

	} else {
		data, err := h.service.MonsterService.Get(bearer, &metadata)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
		}

		metadata.Total = int64(len(*data))

		log.Println("get monster success....")
		return c.JSON(http.StatusOK, succResponse.WithMeta(metadata).WithData(data))
	}
}

func (h *_Handlers) UpdateMonsterHandler(c echo.Context) error {
	var (
		payload      entity.MonsterPayload
		succResponse = response.NewResponse().WithStatus("success").WithMessage("success")
		errResponse  = response.NewResponse().WithStatus("error").WithMessage("error")
		rawId        = c.Param("id")
	)

	log.Println("update monster....")

	if rawId == ":id" {
		return c.JSON(http.StatusBadRequest, errResponse.WithError("id cannot be blank"))
	}

	intID, err := strconv.Atoi(rawId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
	}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
	}

	data, err := h.service.MonsterService.Update(intID, &payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
	}

	log.Println("update monster success....")
	return c.JSON(http.StatusOK, succResponse.WithData(data))
}

func (h *_Handlers) SetStatusMonsterHandler(c echo.Context) error {
	var (
		payload      entity.StatusPayload
		succResponse = response.NewResponse().WithStatus("success").WithMessage("success")
		errResponse  = response.NewResponse().WithStatus("error").WithMessage("error")
		rawId        = c.Param("id")
	)

	log.Println("set status monster....")

	if rawId == ":id" {
		return c.JSON(http.StatusBadRequest, errResponse.WithError("id cannot be blank"))
	}

	intID, err := strconv.Atoi(rawId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
	}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
	}

	if err := h.service.MonsterService.SetStatus(intID, &payload); err != nil {
		log.Printf("(handler) error set status %s", err)
		return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
	}

	log.Println("set status monster success....")
	return c.JSON(http.StatusOK, succResponse.WithData("Success Set Status!"))
}

func (h *_Handlers) CatchMonsterHandler(c echo.Context) error {
	var (
		succResponse = response.NewResponse().WithStatus("success").WithMessage("success")
		errResponse  = response.NewResponse().WithStatus("error").WithMessage("error")
		rawId        = c.Param("id")
		bearer       = c.Request().Header.Get("Authorization")
		msg          string
	)

	log.Println("catch monster....")
	if rawId == ":id" {
		return c.JSON(http.StatusBadRequest, errResponse.WithError("id cannot be blank"))
	}

	intID, err := strconv.Atoi(rawId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
	}

	isCatch, err := h.service.MonsterService.Catch(bearer, intID)
	if err != nil {
		log.Printf("(handler) error catch: %s", err)
		return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
	}

	if *isCatch {
		msg = "Success Release Monster!"
	} else {
		msg = "Success Catch Monster!"
	}

	log.Println("catch monster success....")
	return c.JSON(http.StatusOK, succResponse.WithData(msg))
}

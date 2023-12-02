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

// @Summary Create Monsters
// @Description create a Monsters
// @ID create-monsters
// @Produce json
// @Param Authorization header string false "Bearer token" default(Bearer your_token_here)
// @Param request body entity.MonsterPayload true "Monster Payload"
// @Success 201
// @Router /api/v1/monster [post]
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
	return c.JSON(http.StatusCreated, succResponse.WithData(data))
}

// @Summary Get Monsters
// @Description Get a list of Monsters Or Get Detail of Monster using Query Param "monster_id"
// @ID get-monsters
// @Produce json
// @Param Authorization header string false "Bearer token" default(Bearer your_token_here)
// @Param monster_id query int false "ID of the monster"
// @Param name query string false "Where monster name is (Charizard, Turquise)"
// @Param is_catched query string false "Where monster is Catched (true, false)"
// @Param type query []string false "Type of monsters" collectionFormat(multi)
// @Param order_by query string false "Order by field (e.g., name, id)"
// @Param order_type query string false "Order type (e.g., asc, desc)"
// @Param page query string false "Order type (e.g., asc, desc)"
// @Param per_page query string false "Order type (e.g., asc, desc)"
// @Router /api/v1/monsters [get]
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
			log.Println("get monster by id error", err)
			return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
		}

		log.Println("get monster success....")
		return c.JSON(http.StatusOK, succResponse.WithData(data))

	} else {
		data, err := h.service.MonsterService.Get(bearer, &metadata)
		if err != nil {
			log.Println("get monster error", err)
			return c.JSON(http.StatusBadRequest, errResponse.WithError(err))
		}

		metadata.Total = int64(len(*data))

		log.Println("get monster success....")
		return c.JSON(http.StatusOK, succResponse.WithMeta(metadata).WithData(data))
	}
}

// @Summary Update Monsters
// @Description Update a Monsters
// @ID update-monsters
// @Produce json
// @Param Authorization header string false "Bearer token" default(Bearer your_token_here)
// @Param id path int true "Monster ID"
// @Param request body entity.MonsterPayload true "Monster Payload"
// @Success 200
// @Router /api/v1/monster/{id} [put]
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

// @Summary Update Status Monsters
// @Description Update status of Monsters
// @ID update-status-monsters
// @Produce json
// @Param Authorization header string false "Bearer token" default(Bearer your_token_here)
// @Param id path int true "Monster ID"
// @Param request body entity.StatusPayload true "Status Payload"
// @Success 200
// @Router /api/v1/monster/status/{id} [put]
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

// @Summary Catch Monsters
// @Description Update status of Monsters
// @ID catch-monsters
// @Produce json
// @Param Authorization header string false "Bearer token" default(Bearer your_token_here)
// @Param id path int true "Monster ID"
// @Success 200
// @Router /api/v1/monster/catch/{id} [post]
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

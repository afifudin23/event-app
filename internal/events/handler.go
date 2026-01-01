package events

import (
	"event-app/internal/common"
	"event-app/internal/events/dto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	Service Service
}

type Params struct {
	ID string `uri:"id" binding:"required"`
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetAllEvents(c *gin.Context) {
	events, err := h.Service.FindAll()
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToListResponse(events)))
}

func (h *Handler) GetEventByID(c *gin.Context) {
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	event, err := h.Service.FindByID(uuid.MustParse(params.ID))
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToResponse(event)))
}

func (h *Handler) CreateEvent(c *gin.Context) {
	var payload dto.EventRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}
	uidStr := c.MustGet("uid").(string)
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		common.ErrorHandler(c, common.UnauthorizedError("invalid user id"))
		return
	}
	event, err := h.Service.Create(uid, payload)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	log.Println(12343333)
	c.JSON(http.StatusCreated, common.SuccessResponse(dto.ToSuccessResponse(event.ID)))
}

func (h *Handler) UpdateEvent(c *gin.Context) {
	var params Params
	var payload dto.EventRequest

	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	event, err := h.Service.Update(uuid.MustParse(params.ID), payload)

	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToSuccessResponse(event.ID)))
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	_, err := h.Service.Delete(uuid.MustParse(params.ID))
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToSuccessResponse(uuid.MustParse(params.ID))))
}

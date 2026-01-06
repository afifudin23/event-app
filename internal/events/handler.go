package events

import (
	"event-app/internal/common"
	"event-app/internal/events/dto"
	"net/http"

	"github.com/gin-gonic/gin"
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

	event, err := h.Service.FindByID(params.ID)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToDetailResponse(event)))
}

func (h *Handler) CreateEvent(c *gin.Context) {
	var payload dto.EventRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}
	uid, _ := c.Get("uid")
	event, err := h.Service.Create(uid.(string), payload)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
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

	event, err := h.Service.Update((params.ID), payload)

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

	_, err := h.Service.Delete(params.ID)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToSuccessResponse(params.ID)))
}

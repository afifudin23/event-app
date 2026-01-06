package event_participants

import (
	"event-app/internal/common"
	"event-app/internal/event_participants/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetAllEventParticipants(c *gin.Context) {
	var params dto.Params

	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	event_participants, err := h.Service.FindAll(params.ID)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToListResponse(event_participants)))

}

func (h *Handler) RegisterEventParticipant(c *gin.Context) {
	var params dto.Params

	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	uid, _ := c.Get("uid")

	event_participant, err := h.Service.Register(uid.(string), params.ID)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusCreated, common.SuccessResponse(dto.ToSuccessResponse(event_participant.ID)))
}

func (h *Handler) CancelEventParticipant(c *gin.Context) {
	var params dto.Params

	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	uid, _ := c.Get("uid")

	_, err := h.Service.Cancel(uid.(string), params.ID)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToSuccessResponse(params.ID)))
}

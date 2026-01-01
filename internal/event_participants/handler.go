package event_participants

import (
	"event-app/internal/common"
	"event-app/internal/event_participants/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	event_participants, err := h.Service.FindAll(uuid.MustParse(params.ID))
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

	uidStr := c.MustGet("uid").(string)
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		common.ErrorHandler(c, common.UnauthorizedError("Invalid user id"))
		return
	}

	event_participant, err := h.Service.Register(uid, uuid.MustParse(params.ID))
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusCreated, common.SuccessResponse(dto.ToSuccessResponse(event_participant.ID)))
}

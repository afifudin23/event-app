package dto

import (
	"event-app/internal/models"
	"time"
)

type EventParticipantResponse struct {
	ID      string `json:"id"`
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
	Status  string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EventParticipantListResponse struct {
	Participants []EventParticipantResponse `json:"participants"`
}

type SuccessResponse struct {
	ID string `json:"id"`
}

func ToResponse(participant models.EventParticipants) EventParticipantResponse {
	return EventParticipantResponse{
		ID:        participant.ID,
		EventID:   participant.EventID,
		UserID:    participant.UserID,
		Status:    string(participant.Status),
		CreatedAt: participant.CreatedAt,
		UpdatedAt: participant.UpdatedAt,
	}
}

func ToListResponse(participants []models.EventParticipants) EventParticipantListResponse {
	responses := make([]EventParticipantResponse, 0, len(participants))
	for _, participant := range participants {
		responses = append(responses, EventParticipantResponse{
			ID:        participant.ID,
			EventID:   participant.EventID,
			UserID:    participant.UserID,
			Status:    string(participant.Status),
			CreatedAt: participant.CreatedAt,
			UpdatedAt: participant.UpdatedAt,
		})
	}
	return EventParticipantListResponse{
		Participants: responses,
	}
}

func ToSuccessResponse(id string) SuccessResponse {
	return SuccessResponse{
		ID: id,
	}
}

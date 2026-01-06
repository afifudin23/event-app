package dto

type EventRequest struct {
	Title       string `json:"title" binding:"required,min=6,max=250"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required,min=6,max=250"`
	Capacity    int    `json:"capacity" binding:"required,gt=0"`

	StartDate string `json:"start_date" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	EndDate   string `json:"end_date" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

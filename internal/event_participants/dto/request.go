package dto

type Params struct {
	ID string `uri:"id" binding:"required,uuid"`
}

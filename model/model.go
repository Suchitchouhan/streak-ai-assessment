package model

type RequestPayload struct {
	Numbers []int `json:"numbers" binding:"required"`
	Target  int   `json:"Target" binding:"required"`
}

package models

type Noun struct {
	Article     string `json:"artical" binding:"required"`
	Word        string `json:"word" binding:"required"`
	Translation string `json:"translation" binding:"required"`
}

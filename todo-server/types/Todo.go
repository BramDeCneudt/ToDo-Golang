package types

type Todo struct {
	Id int
	Description string	`json:"description" binding:"required"`
}
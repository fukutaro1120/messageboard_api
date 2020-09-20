package model

// Article ...
type Article struct {
	ID      int    `db:"id", json:"id",form:"id"`
	Name    string `db:"name", json:"name",form:"name"`
	Message string `db:"message", json:"message",form:"message"`
}

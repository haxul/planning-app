package model

import (
	"time"
)

type Card struct {
	Id          string
	CurState    State
	Title       string
	Description string
	Tag         string
	UpdatedOn   time.Time
}

func (c *Card) Push() (string, error) {
	return c.CurState.Move(c)
}

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

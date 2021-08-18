package model

import "time"

type Card struct {
	id          uint64
	curState    State
	title       string
	description string
	tag         string
	updatedOn   time.Time
}

func NewCard(id uint64, state *State, title string, description string, tag string) *Card {
	return &Card{
		id:          id,
		curState:    *state,
		title:       title,
		description: description,
		tag:         tag,
		updatedOn:   time.Now(),
	}
}

func (c *Card) Push() error {
	return c.curState.Move(c)
}

package model

import (
	"errors"
	"time"
)

type State interface {
	Move(c *Card) error
}

var BacklogState = &Backlog{}
var InProgressState = &InProgress{}
var DoneState = &Done{}

type Backlog struct{}

func (state *Backlog) Move(card *Card) error {
	card.UpdatedOn = time.Now()
	card.CurState = &InProgress{}
	return nil
}

type InProgress struct{}

func (state *InProgress) Move(card *Card) error {
	card.UpdatedOn = time.Now()
	card.CurState = &Done{}
	return nil
}

type Done struct{}

func (state *Done) Move(_ *Card) error {
	return errors.New("cannot move card from status done")
}

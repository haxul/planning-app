package model

import (
	"errors"
	"time"
)

type State interface {
	Move(c *Card) error
}

type BacklogState struct {
	card *Card
}

func (state *BacklogState) Move(card *Card) error {
	card.UpdatedOn = time.Now()
	card.CurState = &InProgressState{card: card}
	return nil
}

type InProgressState struct {
	card *Card
}

func (state *InProgressState) Move(card *Card) error {
	card.UpdatedOn = time.Now()
	card.CurState = &DoneState{card: card}
	return nil
}

type DoneState struct {
	card *Card
}

func (state *DoneState) Move(_ *Card) error {
	return errors.New("cannot move card from status done")
}

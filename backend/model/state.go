package model

import (
	"errors"
	"time"
)

type State interface {
	Move(c *Card) error
	Reject(c *Card) error
}

type BacklogState struct{}

func (s *BacklogState) Move(card *Card) error {
	card.UpdatedOn = time.Now()
	card.CurState = &InProgressState{}
	return nil
}

func (s *BacklogState) Reject(c *Card) error {
	c.UpdatedOn = time.Now()
	c.CurState = &RejectState{}
	return nil
}

type InProgressState struct{}

func (s *InProgressState) Move(card *Card) error {
	card.UpdatedOn = time.Now()
	card.CurState = &DoneState{}
	return nil
}

func (s *InProgressState) Reject(c *Card) error {
	c.UpdatedOn = time.Now()
	c.CurState = &RejectState{}
	return nil
}

type DoneState struct{}

func (s *DoneState) Move(_ *Card) error {
	return errors.New("cannot move card from status done")
}

func (s *DoneState) Reject(c *Card) error {
	c.UpdatedOn = time.Now()
	c.CurState = &RejectState{}
	return nil
}

type RejectState struct{}

func (s *RejectState) Move(c *Card) error {
	c.UpdatedOn = time.Now()
	c.CurState = &BacklogState{}
	return nil
}

func (s *RejectState) Reject(_ *Card) error {
	return nil
}

package model

import (
	"errors"
	"github.com/haxul/planning-app/backend/common"
	"time"
)

type State interface {
	Move(c *Card) (string, error)
	Reject(c *Card) (string, error)
	String() string
}

type BacklogState struct{}

func (s *BacklogState) String() string {
	return common.BACKLOG_STATE
}

func (s *BacklogState) Move(card *Card) (string, error) {
	card.UpdatedOn = time.Now()
	card.CurState = &InProgressState{}
	return common.IN_PROGRESS_STATE, nil
}

func (s *BacklogState) Reject(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &RejectState{}
	return common.REJECTED_STATE, nil
}

type InProgressState struct{}

func (s *InProgressState) String() string {
	return common.IN_PROGRESS_STATE
}

func (s *InProgressState) Move(card *Card) (string, error) {
	card.UpdatedOn = time.Now()
	card.CurState = &DoneState{}
	return common.DONE_STATE, nil
}

func (s *InProgressState) Reject(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &RejectState{}
	return common.REJECTED_STATE, nil
}

type DoneState struct{}

func (s *DoneState) String() string {
	return common.DONE_STATE
}

func (s *DoneState) Move(_ *Card) (string, error) {
	return "", errors.New("cannot move card from status done")
}

func (s *DoneState) Reject(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &RejectState{}
	return common.REJECTED_STATE, nil
}

type RejectState struct{}

func (s *RejectState) String() string {
	return common.REJECTED_STATE
}

func (s *RejectState) Move(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &BacklogState{}
	return common.BACKLOG_STATE, nil
}

func (s *RejectState) Reject(_ *Card) (string, error) {
	return common.REJECTED_STATE, nil
}

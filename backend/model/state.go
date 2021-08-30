package model

import (
	"errors"
	"fmt"
	"github.com/haxul/planning-app/backend/common"
	"github.com/haxul/planning-app/backend/custom_errs"
	"time"
)

var stateMap = map[string]State{
	common.IN_PROGRESS_STATE: &InProgressState{},
	common.DONE_STATE:        &DoneState{},
	common.REJECTED_STATE:    &RejectState{},
	common.BOOK_STATE:        &BookState{},
	common.COURSE_STATE:      &CourseState{},
	common.PET_STATE:         &PetState{},
	common.VIDEO_STATE:       &VideoState{},
}

func NewStateFromString(s string) (State, error) {
	state, ok := stateMap[s]
	if ok {
		return state, nil
	}
	return nil, errors.New(fmt.Sprintf("unknown state type: %s", s))
}

type State interface {
	Move(c *Card) (string, error)
	Reject(c *Card) (string, error)
	String() string
}

type VideoState struct {
}

func (v *VideoState) Move(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &InProgressState{}
	return common.IN_PROGRESS_STATE, nil
}

func (v *VideoState) Reject(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &RejectState{}
	return common.REJECTED_STATE, nil
}

func (v *VideoState) String() string {
	return common.VIDEO_STATE
}

type PetState struct {
}

func (p *PetState) Move(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &InProgressState{}
	return common.IN_PROGRESS_STATE, nil
}

func (p *PetState) Reject(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &RejectState{}
	return common.REJECTED_STATE, nil
}

func (p *PetState) String() string {
	return common.COURSE_STATE
}

type CourseState struct {
}

func (cs *CourseState) Move(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &InProgressState{}
	return common.IN_PROGRESS_STATE, nil
}

func (cs *CourseState) Reject(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &RejectState{}
	return common.REJECTED_STATE, nil
}

func (cs *CourseState) String() string {
	return common.COURSE_STATE
}

type BookState struct{}

func (b *BookState) Move(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &InProgressState{}
	return common.IN_PROGRESS_STATE, nil
}

func (b *BookState) Reject(c *Card) (string, error) {
	c.UpdatedOn = time.Now()
	c.CurState = &RejectState{}
	return common.REJECTED_STATE, nil
}

func (b *BookState) String() string {
	return common.BOOK_STATE
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
	return "", &custom_errs.ConflictStateErr{}
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
	c.CurState = &InProgressState{}
	return common.IN_PROGRESS_STATE, nil
}

func (s *RejectState) Reject(_ *Card) (string, error) {
	return "", &custom_errs.ConflictStateErr{}
}

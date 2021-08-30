package dto

import (
	"errors"
	"github.com/haxul/planning-app/backend/common"
	"github.com/haxul/planning-app/backend/model"
	"time"
)

type CardReq struct {
	Title       string `json:"title" validate:"required,NotBlank"`
	Description string `json:"description"`
	Tag         string `json:"tag" validate:"required,IsTag"`
}

type CardResp struct {
	Id          string    `json:"id"`
	CurState    string    `json:"cur_state"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tag         string    `json:"tag"`
	UpdatedOn   time.Time `json:"updated_on"`
}

type ChangeStateCardResp struct {
	NewState string `json:"new_state"`
}

func NewCardResp(card *model.Card) (*CardResp, error) {
	if card == nil {
		return nil, errors.New("card is nil")
	}

	state := ""

	if _, ok := card.CurState.(*model.DoneState); ok {
		state = common.DONE_STATE
	}

	if _, ok := card.CurState.(*model.PetState); ok {
		state = common.PET_STATE
	}

	if _, ok := card.CurState.(*model.VideoState); ok {
		state = common.VIDEO_STATE
	}

	if _, ok := card.CurState.(*model.CourseState); ok {
		state = common.COURSE_STATE
	}

	if _, ok := card.CurState.(*model.BookState); ok {
		state = common.BOOK_STATE
	}

	if _, ok := card.CurState.(*model.InProgressState); ok {
		state = common.IN_PROGRESS_STATE
	}

	if _, ok := card.CurState.(*model.RejectState); ok {
		state = common.REJECTED_STATE
	}

	if state == "" {
		return nil, errors.New("state is not defined in card")
	}

	return &CardResp{
		Id:          card.Id,
		CurState:    state,
		Title:       card.Title,
		Description: card.Description,
		Tag:         card.Tag,
		UpdatedOn:   card.UpdatedOn,
	}, nil
}

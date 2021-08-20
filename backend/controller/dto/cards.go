package dto

import (
	"errors"
	"github.com/haxul/planning-app/backend/model"
	"time"
)

type CardReq struct {
	Title       string `json:"title" validate:"required,NotBlank"`
	Description string `json:"description" validate:"required,NotBlank"`
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

func NewCardResp(card *model.Card) (*CardResp, error) {
	if card == nil {
		return nil, errors.New("card is nil")
	}

	state := ""

	if _, ok := card.CurState.(*model.DoneState); ok {
		state = "DONE"
	}

	if _, ok := card.CurState.(*model.BacklogState); ok {
		state = "BACKLOG"
	}

	if _, ok := card.CurState.(*model.InProgressState); ok {
		state = "IN_PROGRESS"
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

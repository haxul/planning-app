package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/haxul/planning-app/backend/model"
	"log"
	"time"
)

type Cards struct {
	logger log.Logger
}

func (cs *Cards) NewCard(state *model.State, title string, description string, tag string) *model.Card {
	card := &model.Card{
		Id:          uuid.NewString(),
		CurState:    *state,
		Title:       title,
		Description: description,
		Tag:         tag,
		UpdatedOn:   time.Now(),
	}
	msg := fmt.Sprintf("new card with id %s is created", card.Id)
	cs.logger.Println(msg)
	return card
}

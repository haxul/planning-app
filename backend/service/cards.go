package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/haxul/planning-app/backend/common"
	"github.com/haxul/planning-app/backend/model"
	"github.com/haxul/planning-app/backend/persistance/ram"
	"log"
	"time"
)

type Cards struct {
	logger           *log.Logger
	cardsPersistence *ram.Cards
}

var CardsService = &Cards{
	logger:           common.Logger,
	cardsPersistence: ram.CardsPersist,
}

func (cs *Cards) NewCard(title string, description string, tag string) *model.Card {
	card := &model.Card{
		Id:          uuid.NewString(),
		Title:       title,
		Description: description,
		Tag:         tag,
		UpdatedOn:   time.Now(),
	}
	state := model.BacklogState
	card.CurState = state
	msg := fmt.Sprintf("new card with id %s is created", card.Id)
	cs.logger.Println(msg)
	return card
}

func (cs *Cards) SaveCard(card *model.Card) {
	cs.cardsPersistence.AddCard(card)
}

func (cs *Cards) GetAllCards() []*model.Card {
	return cs.cardsPersistence.GetAllCards()
}

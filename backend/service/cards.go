package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/haxul/planning-app/backend/common"
	"github.com/haxul/planning-app/backend/model"
	"github.com/haxul/planning-app/backend/persistance"
	"github.com/haxul/planning-app/backend/persistance/postgres"
	"log"
	"sync"
	"time"
)

var once sync.Once
var instance *CardsSv

type CardsSv struct {
	logger           *log.Logger
	cardsPersistence persistance.CardPersistence
}

func GetCardsSvInstance() *CardsSv {
	once.Do(func() {
		instance = &CardsSv{
			logger:           common.Logger,
			cardsPersistence: postgres.GetCardsPostgresPrs(),
		}
	})
	return instance
}

func (cs *CardsSv) NewCard(title string, description string, tag string) *model.Card {
	card := &model.Card{
		Id:          uuid.NewString(),
		Title:       title,
		Description: description,
		Tag:         tag,
		UpdatedOn:   time.Now(),
		CurState:    &model.BacklogState{},
	}
	msg := fmt.Sprintf("new card with id %s is created", card.Id)
	cs.logger.Println(msg)
	return card
}

func (cs *CardsSv) SaveCard(card *model.Card) error {
	return cs.cardsPersistence.AddCard(card)
}

func (cs *CardsSv) GetAllCards() ([]*model.Card, error) {
	return cs.cardsPersistence.GetAllCards()
}

func (cs *CardsSv) MoveForwardCard(cardId *string) (string, error) {
	card, err := cs.FindCardById(cardId)
	if err != nil {
		return "", err
	}
	return card.CurState.Move(card)
}

func (cs *CardsSv) RejectCard(cardId *string) (string, error) {
	card, err := cs.FindCardById(cardId)
	if err != nil {
		return "", err
	}
	return card.CurState.Reject(card)
}

func (cs *CardsSv) FindCardById(id *string) (*model.Card, error) {
	return cs.cardsPersistence.FindById(id)
}

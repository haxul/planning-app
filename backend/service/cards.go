package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/haxul/planning-app/backend/common"
	"github.com/haxul/planning-app/backend/model"
	"github.com/haxul/planning-app/backend/persistance"
	"github.com/haxul/planning-app/backend/persistance/postgres"
	"log"
	"strings"
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

func (cs *CardsSv) NewCard(title string, description string, tag string) (*model.Card, error) {
	upperTag := strings.ToUpper(tag)
	state, errState := model.NewStateFromString(upperTag)
	if errState != nil {
		return nil, errors.New("unknown state to create card")
	}
	card := &model.Card{
		Id:          uuid.NewString(),
		Title:       title,
		Description: description,
		Tag:         tag,
		UpdatedOn:   time.Now(),
		CurState:    state,
	}
	msg := fmt.Sprintf("new card with id %s is created", card.Id)
	cs.logger.Println(msg)
	return card, nil
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

	newState, stateErr := card.CurState.Move(card)
	if stateErr != nil {
		return "", stateErr
	}

	updErr := cs.cardsPersistence.UpdateCard(card)
	if updErr != nil {
		return "", updErr
	}
	return newState, nil
}

func (cs *CardsSv) RejectCard(cardId *string) (string, error) {
	card, err := cs.FindCardById(cardId)
	if err != nil {
		return "", err
	}
	newState, stateErr := card.CurState.Reject(card)
	if stateErr != nil {
		return "", stateErr
	}

	updErr := cs.cardsPersistence.UpdateCard(card)
	if updErr != nil {
		return "", updErr
	}
	return newState, nil
}

func (cs *CardsSv) FindCardById(id *string) (*model.Card, error) {
	return cs.cardsPersistence.FindById(id)
}

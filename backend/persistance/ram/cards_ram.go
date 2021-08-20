package ram

import (
	"errors"
	"fmt"
	"github.com/haxul/planning-app/backend/model"
	"sync"
	"time"
)

var once sync.Once

type CardsPst struct{}

var instance *CardsPst

var storage = []*model.Card{
	{
		Id:          "test",
		Tag:         "Book",
		Description: "some description",
		CurState:    &model.BacklogState{},
		Title:       "title",
		UpdatedOn:   time.Now(),
	},
}

func GetCardsPrsInstance() *CardsPst {
	once.Do(func() {
		instance = &CardsPst{}
	})

	return instance
}

func (cp *CardsPst) AddCard(c *model.Card) {
	storage = append(storage, c)
}

func (cp *CardsPst) GetAllCards() []*model.Card {
	return storage
}

func (cp *CardsPst) FindById(cardId *string) (*model.Card, error) {
	for _, card := range storage {
		if card.Id == *cardId {
			return card, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("card %s is not found", *cardId))
}

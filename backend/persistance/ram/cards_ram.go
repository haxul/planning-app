package ram

import (
	"errors"
	"fmt"
	"github.com/haxul/planning-app/backend/model"
	"sync"
	"time"
)

var once sync.Once

type CardsRamPst struct{}

var instance *CardsRamPst

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

func GetCardsRamPrsInstance() *CardsRamPst {
	once.Do(func() {
		instance = &CardsRamPst{}
	})

	return instance
}

func (cp *CardsRamPst) UpdateCard(_ *model.Card) error {
	panic("not implemented")
}

func (cp *CardsRamPst) AddCard(c *model.Card) error {
	storage = append(storage, c)
	return nil
}

func (cp *CardsRamPst) GetAllCards() ([]*model.Card, error) {
	return storage, nil
}

func (cp *CardsRamPst) FindById(cardId *string) (*model.Card, error) {
	for _, card := range storage {
		if card.Id == *cardId {
			return card, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("card %s is not found", *cardId))
}

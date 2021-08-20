package ram

import (
	"github.com/haxul/planning-app/backend/model"
	"sync"
)

var once sync.Once

type Cards struct{}

var instance *Cards
var storage = make([]*model.Card, 0)

func GetCardsPrsInstance() *Cards {
	once.Do(func() {
		instance = &Cards{}
	})

	return instance
}

func (cp *Cards) AddCard(c *model.Card) {
	storage = append(storage, c)
}

func (cp *Cards) GetAllCards() []*model.Card {
	return storage
}

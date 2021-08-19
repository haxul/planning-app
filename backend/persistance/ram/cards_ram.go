package ram

import "github.com/haxul/planning-app/backend/model"

type Cards struct{}

var CardsPersist = &Cards{}

var storage = make([]*model.Card, 0)

func (cp *Cards) AddCard(c *model.Card) {
	storage = append(storage, c)
}

func (cp *Cards) GetAllCards() []*model.Card {
	return storage
}

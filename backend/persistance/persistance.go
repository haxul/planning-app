package persistance

import "github.com/haxul/planning-app/backend/model"

type CardPersistence interface {
	AddCard(c *model.Card)
	GetAllCards() []*model.Card
}

package persistance

import "github.com/haxul/planning-app/backend/model"

type CardPersistence interface {
	AddCard(c *model.Card) error
	GetAllCards() ([]*model.Card, error)
	FindById(cardId *string) (*model.Card, error)
	UpdateCard(card *model.Card) error
}

package postgres

import (
	"github.com/haxul/planning-app/backend/model"
	"github.com/jackc/pgx/v4"
	"sync"
)

var PostgreConn *pgx.Conn
var instance *CardsPostgresPersistence
var once *sync.Once

type CardsPostgresPersistence struct{}

func GetCardsPostgresPrs() *CardsPostgresPersistence {
	once.Do(func() {
		instance = &CardsPostgresPersistence{}
	})
	return instance
}

func (cp *CardsPostgresPersistence) AddCard(c *model.Card) {
	panic("implement me")
}

func (cp *CardsPostgresPersistence) GetAllCards() []*model.Card {
	panic("implement me")
}

func (cp *CardsPostgresPersistence) FindById(cardId *string) (*model.Card, error) {
	panic("implement me")
}

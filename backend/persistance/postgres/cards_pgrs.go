package postgres

import (
	"context"
	"github.com/haxul/planning-app/backend/common"
	"github.com/haxul/planning-app/backend/model"
	"github.com/jackc/pgx/v4"
	"sync"
)

var PostgreConn *pgx.Conn
var instance *CardsPostgresPst
var once sync.Once

type CardsPostgresPst struct{}

func GetCardsPostgresPrs() *CardsPostgresPst {
	once.Do(func() {
		instance = &CardsPostgresPst{}
	})
	return instance
}

func (cp *CardsPostgresPst) AddCard(c *model.Card) {
	statement := "INSERT INTO cards (id, state, title, description, tag, update_on) VALUES ($1,$2,$3,$4,$5,$6)"
	_, err := PostgreConn.Exec(context.Background(), statement, c.Id, c.CurState.String(), c.Title, c.Description, c.Tag, c.UpdatedOn)
	if err != nil {
		common.Logger.Fatal(err)
	}
}

func (cp *CardsPostgresPst) GetAllCards() []*model.Card {
	panic("implement me")
}

func (cp *CardsPostgresPst) FindById(cardId *string) (*model.Card, error) {
	panic("implement me")
}

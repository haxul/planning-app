package postgres

import (
	"context"
	"github.com/haxul/planning-app/backend/model"
	"github.com/jackc/pgx/v4"
	"sync"
	"time"
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

func (cp *CardsPostgresPst) AddCard(c *model.Card) error {
	statement := "INSERT INTO cards (id, state, title, description, tag, update_on) VALUES ($1,$2,$3,$4,$5,$6)"
	_, err := PostgreConn.Exec(context.Background(), statement, c.Id, c.CurState.String(), c.Title, c.Description, c.Tag, c.UpdatedOn)
	if err != nil {
		return err
	}
	return nil
}

func (cp *CardsPostgresPst) GetAllCards() ([]*model.Card, error) {
	statement := "SELECT * FROM cards"
	rows, err := PostgreConn.Query(context.Background(), statement)
	if err != nil {
		return nil, err
	}
	var cards []*model.Card
	for rows.Next() {
		var (
			id          string
			title       string
			stateStr    string
			description string
			tag         string
			updateOn    time.Time
		)
		errScan := rows.Scan(&id, &stateStr, &title, &description, &tag, &updateOn)

		if errScan != nil {
			return nil, err
		}

		state, errState := model.NewStateFromString(stateStr)
		if errState != nil {
			return nil, errState
		}

		cards = append(cards, &model.Card{
			Id:          id,
			Title:       title,
			CurState:    state,
			Description: description,
			Tag:         tag,
			UpdatedOn:   updateOn,
		})

	}
	return cards, nil
}

func (cp *CardsPostgresPst) FindById(cardId *string) (*model.Card, error) {
	statement := "SELECT * FROM cards WHERE id = $1"
	var (
		id          string
		title       string
		stateStr    string
		description string
		tag         string
		updateOn    time.Time
	)
	err := PostgreConn.QueryRow(context.Background(), statement, cardId).Scan(&id, &stateStr, &title, &description, &tag, &updateOn)

	if err != nil {
		return nil, err
	}

	state, errState := model.NewStateFromString(stateStr)
	if errState != nil {
		return nil, errState
	}
	return &model.Card{
		Id:          id,
		Title:       title,
		CurState:    state,
		Description: description,
		Tag:         tag,
		UpdatedOn:   updateOn,
	}, nil
}

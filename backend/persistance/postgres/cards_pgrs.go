package postgres

import (
	"context"
	"errors"
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

func (cp *CardsPostgresPst) UpdateCard(newCard *model.Card) error {
	statement := "select exists(select 1 from cards where id =$1)"
	var exists bool
	err := PostgreConn.QueryRow(context.Background(), statement, newCard.Id).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("ID_NOT_FOUND")
	}

	updStatement := "UPDATE cards SET state=$1,title=$2,description=$3,tag=$4,update_on=$5 WHERE id =$6"
	_, updErr := PostgreConn.Exec(context.Background(), updStatement,
		newCard.CurState.String(), newCard.Title, newCard.Description, newCard.Tag, newCard.UpdatedOn, newCard.Id)

	if updErr != nil {
		return err
	}

	return nil
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

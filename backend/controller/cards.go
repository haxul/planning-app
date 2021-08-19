package controller

import (
	json "encoding/json"
	"github.com/haxul/planning-app/backend/common"
	"github.com/haxul/planning-app/backend/controller/dto"
	"github.com/haxul/planning-app/backend/service"
	"log"
	"net/http"
)

type Cards struct {
	logger       *log.Logger
	cardsService *service.Cards
}

var CardsController = &Cards{
	cardsService: service.CardsService,
	logger:       common.Logger}

func (c *Cards) CreateCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cardReq := &dto.CardReq{}

	err := json.NewDecoder(r.Body).Decode(cardReq)
	if err != nil {
		c.logger.Println(err.Error())
		http.Error(w, "cannot decode json cardReq....", http.StatusInternalServerError)
		return
	}

	validateErr := common.JsonValidator.Struct(cardReq)

	if validateErr != nil {
		http.Error(w, validateErr.Error(), http.StatusBadRequest)
		return
	}

	newCard := c.cardsService.NewCard(cardReq.Title, cardReq.Description, cardReq.Tag)
	c.cardsService.SaveCard(newCard)
	cardResp, cardErr := dto.NewCardResp(newCard)

	if cardErr != nil {
		c.logger.Println(cardErr.Error())
		http.Error(w, cardErr.Error(), http.StatusInternalServerError)
		return
	}

	encErr := json.NewEncoder(w).Encode(cardResp)

	if encErr != nil {
		c.logger.Println(encErr.Error())
		http.Error(w, "cannot decode json card", http.StatusInternalServerError)
		return
	}
}

func (c *Cards) GetAllCards(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cards := c.cardsService.GetAllCards()

	list := make([]*dto.CardResp, len(cards))

	for i, card := range cards {
		r, err := dto.NewCardResp(card)
		if err != nil {
			c.logger.Println(err.Error())
			http.Error(w, "cannot encode cards as json", http.StatusInternalServerError)
			return
		}
		list[i] = r
	}

	err := json.NewEncoder(w).Encode(list)
	if err != nil {
		c.logger.Println(err.Error())
		http.Error(w, "cannot encode cards as json", http.StatusInternalServerError)
		return
	}
}

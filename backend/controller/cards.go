package controller

import (
	json "encoding/json"
	"github.com/gorilla/mux"
	"github.com/haxul/planning-app/backend/common"
	"github.com/haxul/planning-app/backend/controller/dto"
	"github.com/haxul/planning-app/backend/service"
	"log"
	"net/http"
	"sync"
)

var once sync.Once
var instance *CardsCtrl

type CardsCtrl struct {
	logger       *log.Logger
	cardsService *service.CardsSv
}

func GetCardsCtrlInstance() *CardsCtrl {
	once.Do(func() {
		instance = &CardsCtrl{
			cardsService: service.GetCardsSvInstance(),
			logger:       common.Logger}
	})
	return instance
}

func (ctrl *CardsCtrl) CreateCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cardReq := &dto.CardReq{}

	err := json.NewDecoder(r.Body).Decode(cardReq)
	if err != nil {
		ctrl.logger.Println(err.Error())
		http.Error(w, "cannot decode json cardReq....", http.StatusInternalServerError)
		return
	}

	validateErr := common.JsonValidator.Struct(cardReq)

	if validateErr != nil {
		http.Error(w, validateErr.Error(), http.StatusBadRequest)
		return
	}

	newCard := ctrl.cardsService.NewCard(cardReq.Title, cardReq.Description, cardReq.Tag)
	newCardErr := ctrl.cardsService.SaveCard(newCard)

	if newCardErr != nil {
		ctrl.logger.Println(newCardErr.Error())
		http.Error(w, "newCardErr", http.StatusInternalServerError)
		return
	}

	cardResp, cardErr := dto.NewCardResp(newCard)

	if cardErr != nil {
		ctrl.logger.Println(cardErr.Error())
		http.Error(w, "Save card error", http.StatusInternalServerError)
		return
	}

	encErr := json.NewEncoder(w).Encode(cardResp)

	if encErr != nil {
		ctrl.logger.Println(encErr.Error())
		http.Error(w, "cannot decode json card", http.StatusInternalServerError)
		return
	}
}

func (ctrl *CardsCtrl) GetAllCards(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cards, err := ctrl.cardsService.GetAllCards()

	if err != nil {
		ctrl.logger.Println(err.Error())
		http.Error(w, "cannot create card response", http.StatusInternalServerError)
		return
	}

	list := make([]*dto.CardResp, len(cards))

	for i, card := range cards {
		r, err := dto.NewCardResp(card)
		if err != nil {
			ctrl.logger.Println(err.Error())
			http.Error(w, "cannot create card response", http.StatusInternalServerError)
			return
		}
		list[i] = r
	}

	errJson := json.NewEncoder(w).Encode(list)
	if errJson != nil {
		ctrl.logger.Println(err.Error())
		http.Error(w, "cannot encode cards as json", http.StatusInternalServerError)
		return
	}
}

func (ctrl *CardsCtrl) MoveCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	cardId := vars["id"]
	newState, err := ctrl.cardsService.MoveForwardCard(&cardId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	resp := &dto.ChangeStateCardResp{NewState: newState}
	errJson := json.NewEncoder(w).Encode(resp)
	if errJson != nil {
		ctrl.logger.Println(err.Error())
		http.Error(w, "changeCardResp encoding error", http.StatusInternalServerError)
		return
	}

}

func (ctrl *CardsCtrl) RejectCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	cardId := vars["id"]
	newState, err := ctrl.cardsService.RejectCard(&cardId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	resp := &dto.ChangeStateCardResp{NewState: newState}
	errJson := json.NewEncoder(w).Encode(resp)
	if errJson != nil {
		ctrl.logger.Println(err.Error())
		http.Error(w, "changeCardResp encoding error", http.StatusInternalServerError)
		return
	}
}

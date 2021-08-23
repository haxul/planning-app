import Constants from "@/commom/constants"

export default {
    actions: {
        async fetchCards(ctx) {
            const resp = await fetch(`${Constants.BASE_URL}/card`)
            const cards = await resp.json()
            ctx.commit("updateCards", cards)
        }
    },
    mutations: {
        updateCards(state, cards) {
            state.cards = cards
        },

        updateCardById(state, payload) {
            const {id, newState} = payload
            console.log(state.cards)
            const card = state.cards.find(el => el.id === id)
            if (card) card.cur_state = newState
        }
    },
    state: {
        cards: []
    },
    getters: {
        getCards(state) {
            return state.cards
        },
        getBacklogList(state) {
            return state.cards.filter(e => e.cur_state === Constants.CARD_STATE.BACKLOG)
        },
        getInProgressList(state) {
            return state.cards.filter(e => e.cur_state === Constants.CARD_STATE.IN_PROGRESS)
        },
        getDoneList(state) {
            return state.cards.filter(e => e.cur_state === Constants.CARD_STATE.DONE)
        },
        getRejectedList(state) {
            return state.cards.filter(e => e.cur_state === Constants.CARD_STATE.REJECTED)
        }
    },
}
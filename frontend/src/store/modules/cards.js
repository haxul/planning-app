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
            return state.cards.filter(e => e.cur_state === "BACKLOG")
        },
        getInProgressList(state) {
            return state.cards.filter(e => e.cur_state === "IN_PROGRESS")
        },
        getDoneList(state) {
            return state.cards.filter(e => e.cur_state === "DONE")
        },
        getRejectedList(state) {
            return state.cards.filter(e => e.cur_state === "REJECTED")
        }
    },
}
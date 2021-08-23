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
        }
    },
}
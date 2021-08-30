import Constants from "@/commom/constants"

export default {
    actions: {
        async fetchCards(ctx) {
            const resp = await fetch(`${Constants.BASE_URL}/card`)
            const cards = await resp.json()
            ctx.commit("updateCards", cards)
        },

        async moveCard(ctx, payload) {
            const {id} = payload
            const resp = await fetch(`${Constants.BASE_URL}/card/${id}/move`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "origin": "localhost:8080",
                },
            })
            if (resp.status === 200) {
                const body = await resp.json()
                const newState = body.new_state
                ctx.commit("updateCardById", {id, newState})
                return
            }
            if (resp.status === 409) {
                const text = await resp.text()
                alert(text)
            }
        },

        async rejectCard(ctx, payload) {
            const {id} = payload
            const resp = await fetch(`${Constants.BASE_URL}/card/${id}/reject`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "origin": "localhost:8080",
                },
            })
            if (resp.status === 200) {
                const body = await resp.json()
                const newState = body.new_state
                ctx.commit("updateCardById", {id, newState})
            }
        }
    },
    mutations: {
        updateCards(state, cards) {
            state.cards = cards
        },

        updateCardById(state, payload) {
            const {id, newState} = payload
            const card = state.cards.find(el => el.id === id)
            if (card) {
                card.cur_state = newState
            }
        }
    },
    state: {
        cards: []
    },
    getters: {
        getCards(state) {
            return state.cards
        },
        getCourseList(state) {
            return state.cards.filter(e => e.cur_state === Constants.CARD_STATE.COURSE)
        },
        getPetList(state) {
            return state.cards.filter(e => e.cur_state === Constants.CARD_STATE.PET)
        },
        getBookList(state) {
            return state.cards.filter(e => e.cur_state === Constants.CARD_STATE.BOOK)
        },
        getVideoList(state) {
            return state.cards.filter(e => e.cur_state === Constants.CARD_STATE.VIDEO)
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
import Vue from "vue";
import Vuex from "vuex"
import cards from "@/store/modules/cards";

Vue.use(Vuex)

export default new Vuex.Store({
    modules: {cards}
})
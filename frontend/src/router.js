import Vue from "vue";
import VueRouter from "vue-router";
import Home from "@/view/Board";
import Account from "@/view/Account";
import CreateCard from "@/view/CreateCard";

Vue.use(VueRouter)

export default new VueRouter({
    mode: "history",
    routes: [
        {
            path: "/",
            component: Home
        },
        {
            path: "/account",
            component: Account
        },
        {
            path: "/create-card",
            component: CreateCard
        }
    ]
})
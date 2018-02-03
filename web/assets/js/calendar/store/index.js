import Vue from "vue"
import Vuex from "vuex"
import * as actions from "./actions"
import * as mutations from "./mutations"
import * as getters from "./getters"

Vue.use(Vuex)

const
    debug = process.env.NODE_ENV !== "production",
    state = Object.assign(
        {
            calendar: {
                hourly: [],
            },
            crons: [],
        },
        typeof window.initialState === "object" ? window.initialState : {},
    )


export default new Vuex.Store({
    state,
    actions,
    mutations,
    getters,
    strict: debug,
})

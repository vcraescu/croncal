import _ from 'lodash'
import Vue from 'vue'

export const updateHourlyCalendar = (state, payload) => {
    state.calendar.hourly = payload
}

export const updateRuntime = (state, payload) => {
    let cron = _.find(state.crons, (cron) => cron.id === payload.id)

    cron.runtime = payload.runtime
}

export const updateCrons = (state, crons) => {
    state.crons = crons
}

export const updateCron = (state, payload) => {
    let i = _.findIndex(state.crons, (cron) => cron.id === payload.id)

    Vue.set(state.crons, i, payload)
}

export const toggleLoading = (state) => {
    state.loading = !state.loading
}

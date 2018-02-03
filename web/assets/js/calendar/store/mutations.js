import _ from "lodash"

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

    state.crons[i].runtime = payload.runtime
    state.crons[i].cmd = payload.cmd
    state.crons[i].name = payload.name
    state.crons[i].ID = payload.ID
}

import calendarApi from './../api/calendar'

export const fetchCrons = async ({commit}) => {
    commit('toggleLoading')

    try {
        let r = await calendarApi.crons()

        commit('updateCrons', r.data)

        return r
    } finally {
        commit('toggleLoading')
    }
}

export const fetchHourlyGrid = async ({commit}) => {
    commit('toggleLoading')

    try {
        let r = await calendarApi.hourly()

        commit('updateHourlyCalendar', r.data.rows)

        return r
    } finally {
        commit('toggleLoading')
    }
}

export const updateCron = async ({commit}, cron) => {
    commit('toggleLoading')

    try {
        let r = await calendarApi.updateCron(cron)

        commit('updateCron', r.data)

        return r
    } finally {
        commit('toggleLoading')
    }
}

export const saveCrontab = async ({commit}) => {
    commit('toggleLoading')

    try {
        return await calendarApi.saveCrontab()
    } finally {
        commit('toggleLoading')
    }
}

export const downloadCrontab = () => {
    calendarApi.downloadCrontab()
}

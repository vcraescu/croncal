import calendarApi from "./../api/calendar"

export const fetchCrons = async ({commit}) => {
    let r = await calendarApi.crons()

    commit("updateCrons", r.data)

    return r.response
}

export const fetchHourlyGrid = async ({commit}) => {
    let r = await calendarApi.hourly()

    commit("updateHourlyCalendar", r.data.rows)

    return r
}

export const updateCron = async ({commit}, cron) => {
    let r = await calendarApi.updateCron(cron)

    commit("updateCron", r.data)

    return r
}

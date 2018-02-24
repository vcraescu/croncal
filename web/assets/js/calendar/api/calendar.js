import axios from 'axios'

export default {
    crons () {
        return axios.get('/api/v1/crons')
    },

    hourly () {
        return axios.get('/api/v1/calendar/hourly')
    },

    updateCron (data) {
        data = {...data}

        let id = data.id

        delete data.id

        return axios.put(`/api/v1/crons/${id}`, data)
    },
}

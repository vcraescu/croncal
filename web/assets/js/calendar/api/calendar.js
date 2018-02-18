import axios from "axios"

export default {
    crons() {
        return axios.get("/api/v1/crons")
    },

    hourly () {
        return axios.get("/api/v1/calendar/hourly")
    },

    updateCron(data) {
        data = {...data}

        let id = data.id

        delete data.id

        if (data.runtime <= 0) {
            data.runtime = 1;
        }

        return axios.put(`/api/v1/crons/${id}`, data)
    }
}

import _ from "lodash"

export const hourlyGrid = ({calendar}, getters) => {
    let grid = []

    calendar.hourly.forEach((row, index) => {
        if (!grid[index]) {
            grid[index] = {label: row.label, crons: []}
        }

        grid[index].label = row.label

        row.crons.forEach((id) => {
            let cron = getters.cron(id),
                i = index + 1,
                runtime = cron.runtime - 1

            grid[index].crons.push(id)

            while (runtime > 0 && calendar.hourly[i]) {
                if (!grid[i]) {
                    grid[i] = {label: null, crons: []}
                }

                grid[i].crons.push(id)
                runtime--
                i++
            }
        })
    })

    return grid
}

export const crons = ({crons}) => {
    return crons
}

export const cron = ({crons}) => (id) => {
    return _.find(crons, (cron) => cron.id === id)
}

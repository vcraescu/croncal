import _ from 'lodash'

export const hourlyGrid = ({calendar}, getters) => {
    let grid = []

    function defaultRow () {
        let row = {
            label: null, count: 0, crons: {},
        }

        getters.crons.forEach((cron) => {
            row.crons[cron.id] = false
        })

        return row
    }

    calendar.hourly.forEach((row, i) => {
        if (!grid[i]) {
            grid[i] = defaultRow()
        }

        grid[i].label = `:${row.label}`

        getters.crons.forEach((cron) => {
            let j = i,
                id = cron.id

            if (row.crons.indexOf(id) === -1) {
                return
            }

            _.forEach(_.range(0, cron.runtime), () => {
                if (!grid[j]) {
                    grid[j] = defaultRow()
                }

                grid[j].crons[id] = true
                grid[j].count++
                j++
            })
        })
    })

    return grid
}

export const crons = ({crons}) => {
    return _.sortBy(crons, ['position'])
}

export const findCron = ({crons}) => (id) => {
    if (!id) {
        return
    }

    return _.find(crons, (cron) => cron.id === id)
}

export const isLoading = ({loading}) => {
    return loading
}

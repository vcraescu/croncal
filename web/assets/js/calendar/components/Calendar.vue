<template>
    <table class="table table-bordered table-hover table-sm">
        <thead class="thead-dark">
        <tr>
            <th class="heading">Cron</th>
            <th class="heading">Count</th>
            <th v-for="cron in crons" :key="cron.id" @dblclick="onStartEditingName(cron.id)">
                {{ cron.name }}

                <form
                        style="min-width: 300px"
                        v-if="editingNames[cron.id]"
                        @submit.prevent="onUpdateName(cron.id)"
                >
                    <div class="row">
                        <div class="col">
                            <input
                                    type="text"
                                    class="form-control form-control-sm"
                                    style="min-width: 100px"
                                    v-model="names[cron.id]"
                            />
                        </div>
                        <div class="col">
                            <button type="submit" class="btn btn-success btn-sm">Save</button>
                            <button type="button" class="btn btn-secondary btn-sm" @click="onCancelEditingName(cron.id)">Cancel</button>
                        </div>
                    </div>
                </form>
            </th>
        </tr>
        </thead>
        <thead class="thead-light">
        <tr>
            <th class="align-middle">Runtime</th>
            <th>&nbsp;</th>
            <th v-for="cron in crons" :key="cron.id">
                <input
                        type="number"
                        v-model.number="runtimes[cron.id]"
                        class="form-control form-control-sm"
                        @input="onRuntimeChanged(cron.id)"
                />
            </th>
        </tr>
        </thead>

        <tbody>
            <row v-for="row in hourlyGrid" :key="row.label" :row="row" :crons="crons" />
        </tbody>
    </table>
</template>

<script>
    import Row from "./Row.vue"
    import { mapGetters, mapActions } from "vuex"

    export default {
        created () {
            this
                .fetchHourlyGrid()
                .then(() => {
                    this.crons.forEach((cron) => {
                        this.$set(this.runtimes, cron.id, cron.runtime)
                        this.$set(this.names, cron.id, cron.name)
                        this.$set(this.editingNames, cron.id, false)
                    })
                })
        },
        data () {
            return {
                runtimes: {},
                names: {},
                grid: [],
                editingNames: {},
            }
        },
        components: {
            Row,
        },
        computed: {
            ...mapGetters([
                "crons",
                "hourlyGrid",
                "cron",
            ]),
        },
        methods: {
            ...mapActions([
                "fetchHourlyGrid",
                "updateCron",
            ]),
            onRuntimeChanged (id) {
                let cron = this.cron(id),
                    data = Object.assign({}, cron, {runtime: this.runtimes[id]})

                this
                    .updateCron(data)
                    .then(() => {
                        this.runtimes[id] = cron.runtime
                    })
                    .catch(() => {
                        this.runtimes[id] = cron.runtime
                    })
            },
            onStartEditingName (id) {
                this.editingNames[id] = true
            },
            onUpdateName (id) {
                let cron = this.cron(id),
                    data = Object.assign({}, cron, {name: this.names[id]})

                this
                    .updateCron(data)
                    .then(() => {
                        this.names[id] = cron.name
                        this.onCancelEditingName(id)
                    })
                    .catch(() => {
                        this.names[id] = cron.name
                    })
            },
            onCancelEditingName (id) {
                this.editingNames[id] = false
            },
        },
    }
</script>

<template>
    <div class="hourly-grid">
        <v-data-table
                :headers="headers"
                :items="items"
                class="grid hourly-grid"
                color="light-blue"
                hide-actions>

            <template slot="headers" slot-scope="props">
                <tr class="blue-grey darken-4 white--text">
                    <th class="white--text sticky-col">:mm</th>
                    <th
                            class="white--text clickable"
                            v-for="header in props.headers"
                            :key="header.id"
                            @click="onEdit(header.id)">
                        <v-tooltip bottom max-width="300" color="blue-grey darken-4">
                            <span slot="activator">{{ header.text }}</span>
                            <span>{{ header.tooltip }}</span>
                        </v-tooltip>
                    </th>
                </tr>
                <tr class="blue-grey darken-3">
                    <th class="white--text sticky-col">Runtime</th>
                    <th
                            class="white--text"
                            v-for="header in props.headers"
                            :key="header.id">
                        {{ header.runtime }}
                    </th>
                </tr>
            </template>

            <template slot="items" slot-scope="props">
                <td class="blue-grey darken-4 white--text sticky-col">
                    <v-badge :color="props.item.count > 1 ? 'red': 'green'">
                        <span slot="badge">{{ props.item.count }}</span>
                        <span>{{ props.item.label }}</span>
                    </v-badge>
                </td>

                <td
                        v-for="(taken, id) in props.item.crons"
                        :key="id"
                        :class="{'blue-grey lighten-3': taken}">
                </td>
            </template>
        </v-data-table>

        <edit-cron-dialog
                v-model="editingCron"
                :visible="editCronDialogVisible"
                @cancel="onCancelEditCronDialog"
                @save="onSaveEditCronDialog">
        </edit-cron-dialog>
    </div>
</template>

<script>
    import EditCronDialog from "./EditCronDialog.vue"
    import { mapGetters, mapActions } from "vuex"

    export default {
        name: "hourly-grid",

        created () {
            this.fetchHourlyGrid()
        },
        data () {
            return {
                editCronDialogVisible: false,
                editingCron: {
                    name: null,
                    cmd: null,
                    id: null,
                    interval: "",
                    runtime: null,
                    position: null,
                },
            }
        },
        components: {
            EditCronDialog,
        },
        computed: {
            ...mapGetters([
                "crons",
                "hourlyGrid",
                "findCron",
            ]),

            headers () {
                let headers = []
                this.crons.forEach((cron) => {
                    headers.push({
                        text: `${cron.name} (${cron.runtime})`,
                        id: cron.id,
                        tooltip: cron.cmd,
                        runtime: cron.runtime,
                    })
                })

                return headers
            },

            items () {
                return this.hourlyGrid
            },
        },
        methods: {
            ...mapActions([
                "fetchHourlyGrid",
                "updateCron",
            ]),
            toggleCronEditDialog () {
                this.editCronDialogVisible = !this.editCronDialogVisible
            },
            onEdit (id) {
                let cron = this.findCron(id)
                if (!cron) {
                    return
                }

                this.toggleCronEditDialog()
                this.editingCron = Object.assign({}, this.editingCron, cron)
            },
            onCancelEditCronDialog () {
                this.toggleCronEditDialog()
            },
            onSaveEditCronDialog () {
                console.info(this.editingCron, this.editingCron.id)
                this
                    .updateCron(this.editingCron)
                    .then(() => {
                        this.$notification.success("Cron data successfully saved")
                        this.toggleCronEditDialog()
                    })
                    .catch((response) => {
                        this.$notification.error("Error occurred")
                    })
            },
        },
    }
</script>

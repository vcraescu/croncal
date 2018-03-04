<template>
  <div class="hourly-grid">
    <v-data-table
      :headers="headers"
      :items="items"
      class="grid hourly-grid"
      color="light-blue"
      hide-actions>

      <template 
        slot="headers" 
        slot-scope="props">
        <tr class="blue-grey darken-4 white--text">
          <th class="white--text sticky-col">:mm</th>
          <th
            class="white--text clickable cron"
            v-for="header in props.headers"
            :key="header.id"
            @click="onEdit(header.id)">
            <v-tooltip 
              bottom 
              max-width="300" 
              color="blue-grey darken-4">
              <span slot="activator">{{ cronName(header.text) }}</span>
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

      <template 
        slot="items" 
        slot-scope="props">
        <td class="blue-grey darken-4 white--text sticky-col">
          <v-badge :color="props.item.count > 1 ? 'red': 'green'">
            <span slot="badge">{{ props.item.count }}</span>
            <span>{{ props.item.label }}</span>
          </v-badge>
        </td>

        <td
          v-for="(taken, id) in props.item.crons"
          :key="id"
          class="cron-slot"
          :class="{'blue-grey lighten-3': taken}"/>
      </template>
    </v-data-table>

    <edit-cron-dialog
      v-model="editingCron"
      :visible="editCronDialogVisible"
      :errors="editCronErrors"
      @cancel="onCancelEditCronDialog"
      @save="onSaveEditCronDialog"/>
  </div>
</template>

<script>
import EditCronDialog from './EditCronDialog.vue'
import {mapGetters, mapActions} from 'vuex'

export default {
    name: 'HourlyGrid',

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
                interval: '',
                runtime: null,
                position: null,
            },
            editCronErrors: {},
        }
    },
    components: {
        EditCronDialog,
    },
    computed: {
        ...mapGetters([
            'crons',
            'hourlyGrid',
            'findCron',
        ]),

        headers () {
            let headers = []
            this.crons.forEach((cron) => {
                headers.push({
                    text: cron.name,
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
            'fetchHourlyGrid',
            'updateCron',
        ]),
        cronName(name) {
            if (!name || name.length < 64) {
                return name
            }

            return `${name.substr(0, 19)} ... ${name.substr(-48)}`
        },
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
            this.editCronErrors = {}
            this.updateCron(this.editingCron).then(() => {
                this.$notification.success('Cron data successfully saved')
                this.toggleCronEditDialog()
            }).catch(({response}) => {
                if (response.status > 499 || !response.data ||
                    !response.data.message
                ) {
                    this.$notification.error('Error occurred')
                    return
                }

                this.$notification.error(response.data.message)
                if (response.data.errors) {
                    this.editCronErrors = response.data.errors
                }
            })
        },
    },
}
</script>

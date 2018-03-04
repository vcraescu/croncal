<template>
  <v-app>
    <progress-bar v-if="isLoading"/>
    <notification delay="5"/>

    <v-toolbar 
      dark 
      color="primary">
      <v-toolbar-title class="white--text">Cron Calendar</v-toolbar-title>
      <v-spacer/>
      <v-tooltip bottom>
        <span slot="activator">
          <v-btn 
            icon 
            @click="onSave">
            <v-icon>save</v-icon>
          </v-btn>
        </span>
        <span>Save Crontab</span>
      </v-tooltip>

      <v-tooltip bottom>
        <span slot="activator">
          <v-btn
            icon
            @click="onDownload">
            <v-icon>file_download</v-icon>
          </v-btn>
        </span>
        <span>Download Crontab file</span>
      </v-tooltip>
    </v-toolbar>

    <v-content>
      <v-container fluid>
        <v-tabs
          color="orange darken-3"
          dark
          slider-color="white">
          <v-tab
            ripple
            key="hourly">
            Hourly
          </v-tab>

          <v-tab-item key="hourly">
            <hourly-grid/>
          </v-tab-item>
        </v-tabs>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import ProgressBar from './ProgressBar.vue'
import HourlyGrid from './HourlyGrid.vue'
import {mapActions, mapGetters} from 'vuex'

export default {
    created () {
        this.fetchCrons()
    },

    components: {
        ProgressBar,
        HourlyGrid,
    },
    computed: {
        ...mapGetters([
            'isLoading',
        ]),
    },
    methods: {
        ...mapActions([
            'fetchCrons',
            'saveCrontab',
            'downloadCrontab',
        ]),

        onSave () {
            this.saveCrontab().then(({data}) => {
                this.$notification.success(data.message)
            }).catch(({data}) => {
                let message = 'Error occurred'
                if (data && data.message) {
                    message = data.message
                }

                this.$notification.error(message)
            })
        },

        onDownload () {
            this.downloadCrontab()
        },
    },
}
</script>

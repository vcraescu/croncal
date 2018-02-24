<template>
  <div class="notification">
    <v-alert
      v-for="(alert, index) in alerts"
      :key="index"
      transition="scale-transition"
      :type="alert.type"
      class="elevation-8"
      :value="true"
      dismissible
      ripple
      @input="dismiss(alert.id)">
      {{ alert.text }}
    </v-alert>
  </div>
</template>

<script>
import _ from 'lodash'

export default {
    name: 'Notification',

    props: {
        delay: {
            type: [Number, String],
            required: false,
            default: 0,
            validator: function (value) {
                return !isNaN(parseInt(value))
            },
        },
    },

    data () {
        return {
            show: true,
            alerts: [],
        }
    },

    methods: {
        generateId () {
            return Math.ceil(Math.random() * 10000000000)
        },

        add (data) {
            let a = Object.assign(data, {
                id: this.generateId(),
            })

            this.alerts.push(a)

            let delay = parseInt(this.delay)
            if (delay > 0) {
                setTimeout(() => {
                    this.dismiss(a.id)
                }, delay * 1000)
            }

            return a
        },

        success (text) {
            this.add({
                type: 'success',
                text,
            })
        },

        error (text) {
            this.add({
                type: 'error',
                text,
            })
        },

        info (text) {
            this.add({
                type: 'info',
                text,
            })
        },

        warning (text) {
            this.add({
                type: 'warning',
                text,
            })
        },

        dismiss (id) {
            let index = _.findIndex(this.alerts, (alert) => {
                return alert.id === id
            })

            this.alerts.splice(index, 1)
        },
    },
}
</script>

<style lang="scss" scoped>
  .notification {
    position: absolute;
    z-index: 9999;
    width: 100%;
    margin: auto;

    .alert {
      max-width: 500px;
      margin-top: 1em;
    }
  }
</style>

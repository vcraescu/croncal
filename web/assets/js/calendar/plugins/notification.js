import Notification from '../components/Notification.vue'

export default {
    install (Vue) {
        let eventBus = new Vue()

        Vue.component(Notification.name, {
            ...Notification,

            created () {
                Notification.created && Notification.created()

                eventBus.$on('success', text => {
                    this.success(text)
                })

                eventBus.$on('error', text => {
                    this.error(text)
                })

                eventBus.$on('info', text => {
                    this.info(text)
                })

                eventBus.$on('warning', text => {
                    this.warning(text)
                })
            },
        })

        Vue.prototype.$notification = {
            success (text) {
                eventBus.$emit('success', text)
            },

            info (text) {
                eventBus.$emit('info', text)
            },

            error (text) {
                eventBus.$emit('error', text)
            },

            warning (text) {
                eventBus.$emit('warning', text)
            },
        }
    },
}

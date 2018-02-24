import Vue from 'vue'
import Vuetify from 'vuetify'
import 'babel-polyfill'
import 'vuetify/dist/vuetify.min.css'
import App from './components/App.vue'
import store from './store'
import Notification from './plugins/notification'

Vue.use(Vuetify)
Vue.use(Notification)

new Vue({
    el: '#calendar',
    store,
    render: h => h(App),
})

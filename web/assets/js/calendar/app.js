import Vue from "vue";
import App from "./components/App.vue";
import store from "./store";

new Vue({
    el: "#calendar",
    store,
    components: {
        App
    }
});

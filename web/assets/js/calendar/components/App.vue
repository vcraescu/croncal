<template>
    <v-app>
        <progress-bar v-if="isLoading"></progress-bar>
        <notification delay="5"></notification>

        <v-content>
            <v-container fluid>
                <v-tabs
                        color="light-blue darken-3"
                        dark
                        slider-color="orange">
                    <v-tab ripple key="hourly">
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
    import ProgressBar from "./ProgressBar.vue"
    import HourlyGrid from "./HourlyGrid.vue"
    import { mapActions, mapGetters } from "vuex"

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
                "isLoading",
            ]),
        },
        methods: {
            ...mapActions([
                "fetchCrons",
            ]),
        },
    }
</script>

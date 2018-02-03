<template>
    <div class="calendar">
        <h1>Cron Calendar</h1>

        <c-calendar/>
    </div>
</template>

<script>
    import Calendar from "./Calendar.vue"
    import { mapActions } from "vuex"

    export default {
        created () {
            this.fetchCrons()
        },
        components: {
            CCalendar: Calendar,
        },
        methods: {
            ...mapActions([
                "fetchCrons",
            ]),
        },
    }
</script>

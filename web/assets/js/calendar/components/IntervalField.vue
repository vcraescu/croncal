<template>
    <div>
        <div
                class="input-group input-group--dirty input-group--text-field "
                :class="{'input-group--required': required, 'input-group--error error--text': hasErrors}">
            <label>Interval</label>
        </div>
        <v-layout>
            <v-flex>
                <v-text-field
                        :value="minute"
                        ref="minute"
                        required
                        :rules="rules.minute"
                        @input="onUpdateMinute"
                ></v-text-field>
            </v-flex>

            <v-flex>
                <v-text-field
                        :value="hour"
                        ref="hour"
                        required
                        :rules="rules.hour"
                        @input="onUpdateHour"
                ></v-text-field>
            </v-flex>

            <v-flex>
                <v-text-field
                        :value="dayOfMonth"
                        ref="dayOfMonth"
                        required
                        :rules="rules.dayOfMonth"
                        @input="onUpdateDayOfMonth"
                ></v-text-field>
            </v-flex>

            <v-flex>
                <v-text-field
                        :value="monthOfYear"
                        ref="monthOfYear"
                        required
                        :rules="rules.monthOfYear"
                        @input="onUpdateMonthOfYear"
                ></v-text-field>
            </v-flex>

            <v-flex>
                <v-text-field
                        :value="dayOfWeek"
                        ref="dayOfWeek"
                        required
                        :rules="rules.dayOfWeek"
                        @input="onUpdateDayOfWeek"
                ></v-text-field>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    export default {
        name: "input-interval",

        props: {
            value: {
                type: String,
                required: false,
                default: "",
            },
            required: {
                type: Boolean,
                required: false,
                default: false,
            },
            rules: {
                type: Object,
                required: false,
                default: {},
            },
        },

        data () {
            return {
                hasErrors: false
            }
        },

        computed: {
            minute () {
                let tokens = this.value.split(" ")

                return tokens[0]
            },
            hour () {
                let tokens = this.value.split(" ")

                return tokens[1]
            },
            dayOfMonth () {
                let tokens = this.value.split(" ")

                return tokens[2]
            },
            monthOfYear () {
                let tokens = this.value.split(" ")

                return tokens[3]
            },
            dayOfWeek () {
                let tokens = this.value.split(" ")

                return tokens[4]
            },
        },

        methods: {
            validate () {
                return !this.$refs.minute.validate()
                    || !this.$refs.hour.validate()
                    || !this.$refs.dayOfMonth.validate()
                    || !this.$refs.monthOfYear.validate()
                    || !this.$refs.dayOfWeek.validate()
            },

            interval () {
                return [
                    this.minute,
                    this.hour,
                    this.dayOfMonth,
                    this.monthOfYear,
                    this.dayOfWeek,
                ]
            },

            onUpdateMinute (value) {
                let interval = this.interval()
                interval[0] = value

                this.hasErrors = this.validate()
                this.$emit("input", interval.join(" "))
            },

            onUpdateHour (value) {
                let interval = this.interval()
                interval[1] = value

                this.hasErrors = this.validate()
                this.$emit("input", interval.join(" "))
            },

            onUpdateDayOfMonth (value) {
                let interval = this.interval()
                interval[2] = value

                this.hasErrors = this.validate()
                this.$emit("input", interval.join(" "))
            },

            onUpdateMonthOfYear (value) {
                let interval = this.interval()
                interval[3] = value

                this.hasErrors = this.validate()
                this.$emit("input", interval.join(" "))
            },

            onUpdateDayOfWeek (value) {
                let interval = this.interval()
                interval[4] = value

                this.hasErrors = this.validate()
                this.$emit("input", interval.join(" "))
            },
        },
    }
</script>

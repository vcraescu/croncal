<template>
    <v-dialog
            :value="visible"
            max-width="600px">
        <v-card>
            <v-card-title>
                <span class="headline">Edit "{{ value.name }}"</span>
            </v-card-title>
            <v-card-text>
                <v-form v-model="valid" lazy-validation ref="form">
                    <v-text-field
                            label="Name"
                            :value="value.name"
                            required
                            @input="onInputName"
                    ></v-text-field>
                    <v-text-field
                            type="number"
                            label="Runtime"
                            :value.number="value.runtime"
                            :rules="validationRules.runtime"
                            required
                            @input="onInputRuntime"
                    ></v-text-field>
                    <interval-field
                            :value="value.interval"
                            :rules="validationRules.interval"
                            required
                    ></interval-field>
                    <v-text-field
                            label="Command"
                            :value="value.cmd"
                            :rules="validationRules.cmd"
                            textarea
                            required
                            rows="3"
                            @input="onInputCommand"
                    ></v-text-field>
                </v-form>
                <small>*indicates required field</small>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" flat @click="onCancel">Cancel</v-btn>
                <v-btn color="success" flat @click="onSave">Save</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
    import IntervalField from "./IntervalField.vue"

    export default {
        name: "edit-cron-dialog",

        data () {
            return {
                valid: true,
                validationRules: {
                    name: [
                        v => !!v || "Name is required",
                    ],
                    runtime: [
                        v => !!v || "Runtime is required",
                        v => v > 0 || "Runtime must be at lest 1 minute",
                        v => v < 59 || "Runtime cannot be greater than 59 minutes",
                    ],
                    cmd: [
                        v => !!v || "Command is required",
                    ],
                    interval: {
                        minute: [
                            v => !!v || "Minute is required",
                        ],
                        hour: [
                            v => !!v || "Hour is required",
                        ],
                        dayOfMonth: [
                            v => !!v || "Day of month is required",
                        ],
                        dayOfWeek: [
                            v => !!v || "Day of week is required",
                        ],
                        monthOfYear: [
                            v => !!v || "Month is required",
                        ]
                    }
                }
            }
        },

        props: {
            visible: {
                type: Boolean,
                required: false,
                default: false,
            },
            value: {
                type: Object,
                required: true,
            },
        },

        components: {
            IntervalField,
        },

        methods: {
            onCancel () {
                this.$emit("cancel")
            },

            onSave () {
                if (!this.$refs.form.validate()) {
                    return
                }

                this.$emit("save", this.form)
            },

            onInputName(value) {
                this.$emit("input", Object.assign({}, this.value, {
                    name: value,
                }))
            },

            onInputCommand(value) {
                this.$emit("input", Object.assign({}, this.value, {
                    cmd: value,
                }))
            },

            onInputRuntime(value) {
                this.$emit("input", Object.assign({}, this.value, {
                    runtime: Number(value),
                }))
            }
        }
    }
</script>

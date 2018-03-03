<template>
  <v-dialog
    :value="visible"
    max-width="600px"
    @input="onInputDialog">
    <v-card>
      <v-card-title>
        <span class="headline">Edit "{{ value.name }}"</span>
      </v-card-title>
      <v-card-text>
        <v-form
          v-model="valid"
          lazy-validation
          ref="form">
          <v-text-field
            label="Name"
            :value="value.name"
            :rules="validationRules.name"
            required
            :error-messages="fieldErrors('name')"
            @input="onInputName"
          />
          <v-text-field
            label="Position"
            type="number"
            :value="value.position"
            :rules="validationRules.position"
            required
            :error-messages="fieldErrors('position')"
            @input="onInputPosition"
          />
          <v-text-field
            type="number"
            label="Runtime"
            :value="value.runtime"
            :rules="validationRules.runtime"
            :error-messages="fieldErrors('runtime')"
            required
            @input="onInputRuntime"
          />
          <interval-field
            :value="value.interval"
            :rules="validationRules.interval"
            :error-messages="fieldErrors('interval')"
            required
            @input="onInputInterval"
          />
          <v-text-field
            label="Command"
            :value="value.cmd"
            :rules="validationRules.cmd"
            :error-messages="fieldErrors('cmd')"
            textarea
            required
            rows="3"
            @input="onInputCommand"
          />
        </v-form>
        <small>*indicates required field</small>
      </v-card-text>
      <v-card-actions>
        <v-spacer/>
        <v-btn
          color="primary"
          flat
          @click="onCancel">Cancel
        </v-btn>
        <v-btn
          color="success"
          flat
          @click="onSave">Save
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import IntervalField from './IntervalField.vue'
import _ from 'lodash'

export default {
    name: 'EditCronDialog',

    data () {
        return {
            valid: true,
            form: {},
            show: false,
            validationRules: {
                name: [
                    v => !!v || 'Name is required',
                    v => v && v.length >= 2 ||
                        'Name must be at least 2 chars long',
                ],
                runtime: [
                    v => v >= 0 || 'Runtime cannot be negative',
                    v => v >= 1 || 'Runtime must be at least 1 minute',
                    v => v <= 59 || 'Runtime cannot be greater than 59 minutes',
                ],
                position: [
                    v => v >= 0 || 'Position cannot be negative',
                ],
                cmd: [
                    v => !!v || 'Command is required',
                    v => v && v.length >= 4 ||
                        'Command must be at least 4 chars long',
                ],
                interval: {
                    minute: [
                        v => !!v || 'Minute is required',
                    ],
                    hour: [
                        v => !!v || 'Hour is required',
                    ],
                    dayOfMonth: [
                        v => !!v || 'Day of month is required',
                    ],
                    dayOfWeek: [
                        v => !!v || 'Day of week is required',
                    ],
                    monthOfYear: [
                        v => !!v || 'Month is required',
                    ],
                },
            },
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
        errors: {
            type: [Object],
            required: false,
            default: () => { return {} },
        }
    },

    components: {
        IntervalField,
    },

    computed: {
        fieldErrors() {
            return (field) => {
                let messages = this.errors[field]
                if (!messages) {
                    return []
                }

                if (typeof messages === 'string') {
                    messages = [messages]
                }

                return messages
            }
        },
    },

    methods: {
        onCancel () {
            this.$emit('cancel')
        },

        onSave () {
            if (!this.$refs.form.validate()) {
                return
            }

            this.$emit('save')
        },

        onInputName (name) {
            let data = {
                ...this.value,
                name,
            }

            this.$emit('input', data)
        },

        onInputCommand (cmd) {
            let data = {
                ...this.value,
                cmd,
            }

            this.$emit('input', data)
        },

        onInputRuntime (runtime) {
            let data = {
                ...this.value,
                runtime: _.toInteger(runtime),
            }

            this.$emit('input', data)
        },

        onInputInterval (interval) {
            let data = {
                ...this.value,
                interval,
            }

            this.$emit('input', data)
        },

        onInputPosition (position) {
            let data = {
                ...this.value,
                position: _.toInteger(position),
            }

            this.$emit('input', data)
        },

        onInputDialog (value) {
            if (value === false) {
                this.$emit('cancel')
            }
        }
    },
}
</script>

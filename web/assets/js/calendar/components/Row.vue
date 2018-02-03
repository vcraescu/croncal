<template>
    <tr>
        <th>:{{ row.label }}</th>
        <td class="text-center">
            <span
                    class="badge"
                    :class="{'badge-danger': row.crons.length > 1, 'badge-success': row.crons.length <= 1}"
            >
                {{ row.crons.length }}
            </span>
        </td>

        <td
                v-for="cron in crons"
                :key="cron.id"
                :class="{'table-primary': row.crons.indexOf(cron.id) >= 0}"
        >
            &nbsp;
        </td>
    </tr>
</template>

<script>
    export default {
        props: {
            row: {
                type: Object,
                required: true
            },
            crons: {
                type: Array,
                required: true
            }
        },
    }
</script>

<style scoped>

</style>

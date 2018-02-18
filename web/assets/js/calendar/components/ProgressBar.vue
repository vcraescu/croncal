<template>
    <div class="progress-bar">
        <v-progress-linear :indeterminate="true" :color="color"></v-progress-linear>
    </div>
</template>

<script>
    export default {
        name: "progress-bar",

        props: {
            color: {
                type: String,
                required: false,
                default: "orange",
            },
        },
    }
</script>

<style lang="scss" scoped>
    .progress-bar {
        position: absolute;
        z-index: 9999;
        width: 100%;

        .progress-linear {
            margin: 0;
        }
    }
</style>

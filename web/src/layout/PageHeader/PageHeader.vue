<script lang="ts" setup>
// components
import Breadcrumb from "./components/Breadcrumb/Breadcrumb.vue"
import HeaderSetup from "./components/HeaderSetup/HeaderSetup.vue"

const emit = defineEmits<{
    (event: "show"): void
}>()

const show = (): void => {
    emit("show")
}

</script>

<template>
    <a-layout-header class="system-header">
        <Breadcrumb />
        <HeaderSetup @show="show" />
    </a-layout-header>
</template>

<style lang="scss">
.system-header {
    z-index: 998;
    justify-content: space-between;
    height: 50px;
    display: flex;
    align-items: center;
    background-color: #fff;
    box-shadow: 0 0 5px 0 #ccc;
}
</style>

<script lang="ts" setup>

</script>

<template>
    <a-layout-content class="system-content">
        <!-- 使用动态过渡名称 -->
        <router-view appear v-slot="{ Component, route }">
            <transition name="slide-fade" mode="out-in">
                <component :is="Component" />
            </transition>
        </router-view>
    </a-layout-content>
</template>

<style lang="less" scoped>
@import "./index.less";
</style>
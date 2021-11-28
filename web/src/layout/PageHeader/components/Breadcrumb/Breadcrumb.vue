<script lang="ts" setup>
import { reactive, watch } from 'vue'
import { useRoute } from 'vue-router'
import type { RouteLocationMatched } from "vue-router"

const route = useRoute()

const breadcrumbConfig = reactive({
    routerArrs: [] as RouteLocationMatched[] | undefined,
})

watch(
    () => route.matched,
    (ov, _) => {
        breadcrumbConfig.routerArrs = ov
    },
    {
        deep: true,
        immediate: true,
    }
)
</script>

<template>
    <a-breadcrumb>
        <a-breadcrumb-item v-for="routeItem in breadcrumbConfig.routerArrs">{{ routeItem.meta.title }}</a-breadcrumb-item>
    </a-breadcrumb>
</template>
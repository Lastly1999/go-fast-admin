<script lang="ts" setup>
import {ref, watch, computed} from "vue"
import {useRoute, useRouter} from "vue-router"
import type {RouteLocationMatched} from "vue-router"

const router = useRouter()
const route = useRoute()

const editableTabs = ref([
    {
        name: "首页",
        path: "/dashboard",
        closable: false
    }
])

const addRoutePane = (route: RouteLocationMatched[]) => {
    const routeMatched = route[route.length - 1]
    const exist = editableTabs.value.find(item => item.path === routeMatched.path)
    if (exist) {
        router.push(routeMatched.path)
    } else {
        editableTabs.value.push({
            path: routeMatched.path,
            name: (routeMatched.meta.title as string),
            closable: true
        })
    }
}

watch(() => route.matched, val => {
    addRoutePane(val)
}, {deep: true, immediate: true})


const activeKey = computed(() => route.path)

const onEdit = (path: string, action: string) => {
    if (action === 'remove') {
        let rmIndex = editableTabs.value.findIndex(item => item.path === path)
        editableTabs.value.splice(rmIndex, 1)
        router.push(editableTabs.value[rmIndex - 1].path)
    }
}

const size = ref("small")

const changeTabs = (path: string) => {
    router.push(path)
}

</script>
<template>
    <a-tabs :activeKey="activeKey" hide-add type="editable-card" :size="size"
            :tabBarStyle="{padding:'3px 5px',background:'#fff'}"
            @edit="onEdit" @change="changeTabs">
        <a-tab-pane class="rouing-item" v-for="pane in editableTabs" :closable="pane.closable" :key="pane.path"
                    :tab="pane.name"/>
    </a-tabs>
</template>
<style lang="less" scoped>
@import "index.less";
</style>
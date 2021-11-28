<script lang="ts" setup>
import {ref, onMounted} from "vue"
import {useStore} from "vuex"
import {Layout} from "ant-design-vue"

import PageHeader from './PageHeader/PageHeader.vue'
import PageContent from './PageContent/PageContent.vue'
import PageSider from './PageSider/PageSider.vue'

// 系统配置抽屉
import SystemSetupDrawer from "@/components/SystemSetupDrawer/SystemSetupDrawer.vue"
// 路由导航
import RoutingRecord from "@/components/RoutingRecord/RoutingRecord.vue"

onMounted(() => {
    loadSystemConf()
})

const store = useStore()

const showSetupDrawer = ref<boolean>(false)

const showSystemSetup = () => {
    showSetupDrawer.value = true
}

/**
 *  理一些系统级的参数请求
 */
const loadSystemConf = () => {
    // 系统图标下拉列表参数
    store.dispatch("systemModule/API_GET_SYS_ICONS")
}
</script>
<template>
    <Layout class="system-container">
        <PageSider/>
        <Layout>
            <PageHeader @show="showSystemSetup"></PageHeader>
            <RoutingRecord/>
            <PageContent></PageContent>
        </Layout>
        <SystemSetupDrawer v-model:visible="showSetupDrawer"/>
    </Layout>
</template>

<style lang="less">
@import "./index.less";
</style>
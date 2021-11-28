<script lang="ts" setup>
import { onMounted, ref } from "vue"
// components
import WorkInfoPanel from './components/WorkInfoPanel/WorkInfoPanel.vue'
import TodoPanel from "./components/TodoPanel/TodoPanel.vue"
import WorkBannerPanel from "./components/WorkBannerPanel/WorkBannerPanel.vue"
import TimeLinePanel from "./components/TimeLinePanel/TimeLinePanel.vue"
import QuickMenusPanel from "./components/QuickMenusPanel/QuickMenusPanel.vue"

const panelLoading = ref(true)

onMounted(() => {
  setTimeout(() => {
    panelLoading.value = false
  }, 500)
  
})

</script>

<template>
  <a-row :gutter="[10, 10]" justify="start" :wrap="true" style="padding: 15px">
    <a-col :xl="18" :md="24" :sm="24">
      <a-skeleton active :loading="panelLoading">
        <WorkInfoPanel />
      </a-skeleton>
      <a-skeleton active :loading="panelLoading">
        <TodoPanel />
      </a-skeleton>
      <a-skeleton active :loading="panelLoading">
        <TimeLinePanel />
      </a-skeleton>
    </a-col>
    <a-col :xl="6" :md="24" :sm="24">
      <a-skeleton active :loading="panelLoading">
        <WorkBannerPanel />
      </a-skeleton>
      <a-skeleton active :loading="panelLoading">
        <QuickMenusPanel />
      </a-skeleton>
    </a-col>
  </a-row>
</template>

<template>
  <div>
    public form
    <a-button type="primary" @click="showDialog">显示</a-button>
    <CommonDialog :title="title" v-model:loading="loading" v-model:visible="visible" @close="closeDialog"></CommonDialog>
  </div>
</template>
<script lang="ts" setup>
import CommonDialog from "@/components/commonDialog/commonDialog.vue"
import useCommonDialog from "@/components/CommonDialog/useCommonDialog"

const { title, visible, loading } = useCommonDialog()

const showDialog = (): void => {
  visible.value = true
}

const closeDialog = (): void => {
  visible.value = false
}
</script>
<script lang="ts" setup>
import { useRouter } from "vue-router"

const router = useRouter()

const back = () => {
  router.go(-1)
  console.log('back')
}
</script>
<template>
  <div>
    <a-result status="403" title="403" sub-title="抱歉，您无权访问此页面。">
      <template #extra>
        <a-button type="primary" @click="back">返回系统</a-button>
      </template>
    </a-result>
  </div>
</template>
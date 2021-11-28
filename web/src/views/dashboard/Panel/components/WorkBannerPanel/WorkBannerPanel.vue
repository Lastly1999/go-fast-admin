<script lang="ts" setup>
import { ref, onMounted } from "vue"
import codingImg from "@/assets/coding.png"

onMounted(() => {
  panelLoading.value = false
})

const deadLine = ref<number>(Date.now() + 1000 * 60 * 60 * 24 * 2 + 1000 * 30)

const panelLoading = ref<boolean>(true)

const onFinish = () => {
  console.log('onFinish')
}

</script>
<template>
  <a-skeleton active :loading="panelLoading">
    <div style="background-color: #fff;box-shadow: 0 1px 3px rgb(18 18 18 / 10%);">
      <img style="width:100%;height:250px" :src="codingImg" alt="#" />
    </div>
  </a-skeleton>
</template>

<style lang="scss">
</style>

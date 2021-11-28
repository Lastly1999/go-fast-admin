<script lang="ts" setup>

defineProps({
  // 标题
  title: {
    type: String,
    default: () => ''
  },
  // 宽度
  width: {
    type: String,
    default: () => "40%"
  },
  // 加载状态
  loading: {
    type: Boolean,
    default: () => false
  },
  // 显示状态
  visible: {
    type: Boolean,
    default: () => false
  },
  // body style
  bodyStyle: {
    type: Object,
    default: () => { }
  },
})

const emits = defineEmits<{
  (event: 'submit'): void,
  (event: 'close'): void
}>()

const onSubmit = (): void => {
  emits('submit')
}

const onClose = (): void => {
  emits('close')
}


</script>
<template>
  <a-drawer :title="title" :width="width" :visible="visible" :body-style="bodyStyle" @close="onClose">
    <slot></slot>
    <div class="f-drawer-body">
      <a-button style="margin-right: 8px" @click="onClose">取消</a-button>
      <a-button type="primary" @click="onSubmit" :loading="loading">保存</a-button>
    </div>
  </a-drawer>
</template>
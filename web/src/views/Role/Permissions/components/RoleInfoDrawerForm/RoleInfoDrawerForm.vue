<script lang="ts" setup>
import { ref } from 'vue';
import type { RoleForm } from '@/services/model/response/role';
import type { PropType } from 'vue';
import { ValidateErrorEntity } from 'ant-design-vue/lib/form/interface';

// components
import FDrawer from '@/components/FDrawer/FDrawer.vue';

const props = defineProps({
  form: {
    type: Object as PropType<RoleForm>,
    default: () => {
      return {};
    },
  },
  value: {
    type: Boolean,
    default: () => false,
  },
  loading: {
    type: Boolean,
    default: () => false,
  },
});

// emits
const emit =
  defineEmits<{
    (event: 'submit', form: RoleForm): RoleForm;
    (event: 'update:value', show: boolean): void;
  }>();

const formRef = ref();

const onSubmit = () => {
  formRef.value
    .validate()
    .then(() => {
      emit('submit', props.form);
    })
    .catch((error: ValidateErrorEntity<RoleForm>) => {
      console.log('error', error);
    });
};

const close = () => {
  formRef.value.resetFields();
  emit('update:value', false);
};

const rules = {
  roleName: [
    {
      required: true,
      message: '请输入',
    },
  ],
  describe: [
    {
      required: true,
      message: '请输入',
    },
  ],
  createDate: [
    {
      required: true,
      message: '请输入',
    },
  ],
  state: [
    {
      required: true,
      message: '请输入',
    },
  ],
};
</script>

<template>
  <FDrawer v-model:value="value" title="角色信息" :width="500" @submit="onSubmit" @close="close" :loading="loading">
    <a-form :model="form" :rules="rules" layout="vertical" ref="formRef">
      <a-form-item label="角色名称" name="roleName">
        <a-input v-model:value="form.roleName" placeholder="请输入" />
      </a-form-item>
      <a-form-item label="角色别名" name="describe">
        <a-input v-model:value="form.describe" placeholder="请输入" />
      </a-form-item>
      <!-- <a-form-item label="创建时间" name="createDate">
        <a-date-picker v-model:value="form.createDate" :mode="form.createDate" show-time placeholder="请选择" />
      </a-form-item> -->
    </a-form>
  </FDrawer>
</template>
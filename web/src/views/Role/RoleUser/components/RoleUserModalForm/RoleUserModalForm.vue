<script lang="ts" setup>
import { computed, ref, watch } from "vue"
import { useStore } from "vuex"
import FModal from "@/components/FModal/FModal.vue"
import { ValidateErrorEntity } from "ant-design-vue/es/form/interface"

// apis
import { getSystemUserRoles } from "@/services/user/user"

import type { PropType } from 'vue'
import type { RoleItem } from "@/store/modules/system"


export type UserForm = {
  id: number | null;
  userName: string;
  userAvatar: string;
  passWord?: string;
  nikeName: string;
  roleIds?: string[];
  roleId: string | null;
  CreatedAt: string;
  UpdatedAt: string;
}

export type SelectValue = {
  value?: string;
  label?: string;
}

const store = useStore()

const props = defineProps({
  loading: {
    type: Boolean,
    default: () => false
  },
  title: {
    type: String,
    default: () => ''
  },
  visible: {
    type: Boolean,
    default: () => false
  },
  form: {
    type: Object as PropType<UserForm>,
    default: () => {
      return {
        id: null
      }
    }
  }
})

const emit = defineEmits<{
  (event: 'submit', form: UserForm): void,
  (event: 'update:visible', show: boolean): void,
  (event: 'update:form', form: object): void,
}>()

watch(() => props.form.id, () => {
  if (props.form.id) {
    getRoles(props.form.id)
  }
}, { deep: true })

const getRoles = async (id: number) => {
  const { code, data } = await getSystemUserRoles(id)
  if (code === 200) {
    roles.value = data.roles
  }
}
const roles = ref<RoleItem[]>([])

const storeRoles = computed<RoleItem[]>(() => store.getters["systemModule/getSysRoles"])

const userFormRef = ref()

const rolesChange = (key: string[], item: SelectValue[]) => {
  // 序列化
  const sync: RoleItem[] = []
  // 参数转换
  item.map(item => {
    sync.push({
      roleId: item.value,
      roleName: item.label
    })
  })
  roles.value = sync
  if (!roles.value.find(item => item.roleId === props.form.roleId)) {
    props.form.roleId = null
  }
}

const onSubmit = (): void => {
  userFormRef.value
    .validate()
    .then(() => {
      emit('submit', props.form)
    })
    .catch((error: ValidateErrorEntity<UserForm>) => {
      console.log('error', error);
    });
}

const modalCancel = (): void => {
  emit('update:visible', false)
  roles.value = []
  emit('update:form', {})
}

const rules = {
  userName: [{ required: true, message: '请输入' }],
  userAvatar: [{ required: true, message: '请输入' }],
  passWord: [{ required: true, message: '请输入' }],
  nikeName: [{ required: true, message: '请输入' }],
  roleIds: [{ required: true, message: '请选择' }],
  roleId: [{ required: true, message: '请选择' }],
  CreatedAt: [{ required: true, message: '请选择' }],
  UpdatedAt: [{ required: true, message: '请选择' }]
}
</script>

<template>
  <FModal :destroyOnClose="true" v-model:value="visible" v-model:title="title" :confirm-loading="loading" @ok="onSubmit" :afterClose="modalCancel">
    <a-form ref="userFormRef" :model="form" :rules="rules" layout="vertical">
      <a-form-item v-if="form.id" label="用户序号" name="pId">
        <a-input disabled v-model:value="form.id" placeholder="请输入" />
      </a-form-item>
      <a-form-item label="账户名称" name="userName">
        <a-input v-model:value="form.userName" placeholder="请输入" />
      </a-form-item>
      <a-form-item label="姓名" name="nikeName">
        <a-input v-model:value="form.nikeName" placeholder="请输入" />
      </a-form-item>
      <a-form-item v-if="!form.id" label="账户密码" name="passWord">
        <a-input v-model:value="form.passWord" placeholder="请输入" />
      </a-form-item>
      <a-form-item label="角色设置" name="roleIds">
        <a-select v-model:value="form.roleIds" style="width: 100%" mode="multiple" placeholder="请先配置我" option-label-prop="label" @change="rolesChange">
          <a-select-option v-for="item in storeRoles" :value="item.roleId" :label="item.roleName">&nbsp;&nbsp;{{ item.roleName }}</a-select-option>
        </a-select>
      </a-form-item>
      <a-form-item label="默认角色" name="roleId">
        <a-select notFoundContent="请先配置角色" v-model:value="form.roleId" style="width: 100%" placeholder="暂无配置角色" option-label-prop="label">
          <a-select-option v-for="item in roles" :value="item.roleId" :label="item.roleName">&nbsp;&nbsp;{{ item.roleName }}</a-select-option>
        </a-select>
      </a-form-item>
      <!-- <a-form-item v-if="form.id" label="创建时间" name="CreatedAt">
        <a-input disabled v-model:value="form.CreatedAt" placeholder="请输入" />
      </a-form-item>
      <a-form-item v-if="form.id" label="更新时间" name="UpdatedAt">
        <a-input disabled v-model:value="form.UpdatedAt" placeholder="请输入" />
      </a-form-item>-->
    </a-form>
  </FModal>
</template>
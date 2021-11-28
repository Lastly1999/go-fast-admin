<script lang="ts" setup>
import {ref, watch} from "vue"

// global components
import FModal from "@/components/FModal/FModal.vue"
import IconSelect from "@/components/IconSelect/IconSelect.vue"

// apis
import {updateBaseMenuInfo, getBaseMenuInfo} from "@/services/auth"
import {ValidateErrorEntity} from "ant-design-vue/es/form/interface";

// typings
import type {MenuFormOptions} from "@/views/Main/Role/RoleMenus/components/MenuDrawerForm/MenuDrawerForm.vue";
import type {UpdateBaseMenuParams} from "@/services/model/request/auth";


type EditMenuForm = {
    id: string;
    icon: string;
    label: string;
    path: string;
}

const props = defineProps({
    visible: {
        type: Boolean,
        default: false
    },
    loading: {
        type: Boolean,
        default: false
    },
    title: {
        type: String,
        default: "编辑菜单"
    },
    id: {
        default: ''
    }
})

const emit = defineEmits<{
    (event: 'success'): void
}>()

watch(() => props.id, val => {
    if (val) {
        getBaseMenuInfo(val).then(res => {
            form.value = {
                ...res.data
            }
        })
    }
})


const form = ref<EditMenuForm>({
    id: '',
    icon: '',
    label: '',
    path: ''
})

const rules = {
    id: [{required: true, message: '无根菜单id'}],
    label: [{required: true, message: '请输入'}],
    icon: [{required: true, message: '请选择'}],
    path: [{required: true, message: '请输入'}],
}

const menuForm = ref()

// 编辑提交
const onSubmit = () => {
    menuForm.value
    .validate()
    .then(() => {
        updateMenu()
    })
    .catch((error: ValidateErrorEntity<MenuFormOptions>) => {
        console.log(error)
    })
}

/**
 * 请求修改菜单详情接口
 */
const updateMenu = async () => {
    const params: UpdateBaseMenuParams = {
        menuName: form.value.label,
        menuId: form.value.id,
        menuIcon: form.value.icon,
        menuPath: form.value.path
    }
    const {data, code} = await updateBaseMenuInfo(params)
    if (code === 200) {
        emit('success')
    }
}

</script>
<template>
    <FModal class="menu-edit-container" :confirm-loading="loading" v-model:title="title" v-model:value="visible"
            @close="$emit('update:visible',false)"
            @ok="onSubmit">
        <a-form ref="menuForm" :model="form" :rules="rules" layout="vertical">
            <div class="dialog-title">
                <h3>菜单编辑</h3>
            </div>
            <a-row :gutter="16">
                <a-col :span="12">
                    <a-form-item label="菜单序号" name="id">
                        <a-input disabled v-model:value="form.id" placeholder="请输入"/>
                    </a-form-item>
                </a-col>
                <a-col :span="12">
                    <a-form-item label="菜单名称" name="label">
                        <a-input v-model:value="form.label" placeholder="请输入"/>
                    </a-form-item>
                </a-col>
            </a-row>
            <a-row :gutter="16">
                <a-col :span="12">
                    <a-form-item label="菜单路径" name="path">
                        <a-input v-model:value="form.path" placeholder="请输入"/>
                    </a-form-item>
                </a-col>
                <a-col :span="12">
                    <a-form-item label="菜单图标" name="icon">
                        <IconSelect v-model:value="form.icon" placeholder="请选择"/>
                    </a-form-item>
                </a-col>
            </a-row>
        </a-form>
    </FModal>
</template>
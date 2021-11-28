<script lang="ts" setup>
import {ref} from "vue"

import type {PropType} from "vue"
import type {ValidateErrorEntity} from "ant-design-vue/es/form/interface"

// components
import FModal from "@/components/FModal/FModal.vue"
import IconSelect from "@/components/IconSelect/IconSelect.vue"

export type MenuFormOptions = {
    icon?: string | null;
    id: number;
    label: string;
    pId?: number;
    pName?: string;
    path?: string;
    pPath?: string;
    pIcon?: string | null;
}

const props = defineProps({
    form: {
        type: Object as PropType<MenuFormOptions> & never
    },
    title: {
        type: String,
        default: () => ""
    },
    visible: {
        type: Boolean,
        default: () => false
    },
    loading: {
        type: Boolean,
        default: () => false
    }
})

const emit = defineEmits<{
    (event: "submit", data: MenuFormOptions): void
}>()

const rules = {
    id: [{required: true, message: '无根菜单id'}],
    label: [{required: true, message: '请输入'}],
    icon: [{required: true, message: '请选择'}],
    path: [{required: true, message: '请输入'}],
    pName: [{required: true, message: ' 请输入'}],
    pId: [{required: true, message: '请输入'}],
    pPath: [{required: true, message: '请输入'}],
    pIcon: [{required: true, message: '清选择'}],
};

const menuForm = ref()

const onSubmit = () => {
    menuForm.value
    .validate()
    .then(() => {
        emit('submit', props.form)
    })
    .catch((error: ValidateErrorEntity<MenuFormOptions>) => {
        console.log(error)
    })
}

</script>

<template>
    <FModal :confirm-loading="loading" v-model:title="title" v-model:value="visible" @ok="onSubmit">
        <a-form ref="menuForm" :model="form" :rules="rules" layout="vertical">
            <div class="dialog-title">
                <h3>根菜单信息</h3>
            </div>
            <a-row :gutter="16">
                <a-col :span="12">
                    <a-form-item label="根菜单序号" name="id">
                        <a-input disabled v-model:value="form.id" placeholder="请输入"/>
                    </a-form-item>
                </a-col>
                <a-col :span="12">
                    <a-form-item label="根序号" name="pId">
                        <a-input disabled v-model:value="form.pId" placeholder="请输入"/>
                    </a-form-item>
                </a-col>
            </a-row>
            <a-row :gutter="16">
                <a-col :span="12">
                    <a-form-item label="根菜单名称" name="label">
                        <a-input :disabled="!!form.id" v-model:value="form.label" placeholder="请输入"/>
                    </a-form-item>
                </a-col>
                <a-col :span="12">
                    <a-form-item label="根菜单路径" name="path">
                        <a-input :disabled="!!form.id" v-model:value="form.path" placeholder="请输入"/>
                    </a-form-item>
                </a-col>
            </a-row>
            <a-row :gutter="16">
                <a-col :span="12">
                    <a-form-item label="根菜单图标" name="icon">
                        <IconSelect v-model:value="form.icon" placeholder="请选择"/>
                    </a-form-item>
                </a-col>
            </a-row>
            <div v-if="form.id">
                <div class="dialog-title">
                    <h3>子菜单设置</h3>
                </div>
                <a-form-item label="子菜单名称" name="pName">
                    <a-input v-model:value="form.pName" placeholder="请输入"/>
                </a-form-item>
                <a-form-item label="子菜单路由" name="pPath">
                    <a-input v-model:value="form.pPath" placeholder="请输入"/>
                </a-form-item>
                <a-form-item label="子菜单图标" name="pIcon">
                    <IconSelect v-model:value="form.pIcon" placeholder="请选择"/>
                </a-form-item>
            </div>
        </a-form>
    </FModal>
</template>

<style lang="scss">
@import "index";
</style>
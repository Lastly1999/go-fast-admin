<script lang="ts" setup>
import {ref, computed, onMounted, createVNode} from "vue"
import {useStore} from 'vuex'
import {useRouter} from "vue-router"

import type {UpdateSystemUserRoleParams} from "@/services/model/response/auth"
import type {RoleItem} from "@/store/modules/system"


import screenfull from 'screenfull'

// components
import {Modal} from 'ant-design-vue'
import {
    SettingOutlined,
    ExpandOutlined,
    SearchOutlined,
    PoweroffOutlined,
    UserSwitchOutlined,
    ExclamationCircleOutlined,
    BgColorsOutlined
} from '@ant-design/icons-vue'
import {alertMsg} from "@/utils/antd/antd"

// apis
import {updateUserRole} from "@/services/auth"

const emit = defineEmits<{
    (event: 'show'): void
}>()


const router = useRouter()
const store = useStore()

const userInfo: any = computed(() => JSON.parse(localStorage.getItem("user-info") as string))
const currentInfo = ref<RoleItem>()

onMounted(() => {
    currentInfo.value = userInfo.value.role.find((item: any) => item.roleId === userInfo.value.roleId)
})

const visible = ref<boolean>(false)
const roleVisible = ref<boolean>(false)

// roleselect 下拉列表
const roleHide = (roleId: string) => {
    Modal.confirm({
        title: '温馨提示',
        icon: createVNode(ExclamationCircleOutlined),
        content: '切换角色后，需要进行重新登录哦，是否继续？',
        okText: '确认',
        cancelText: '算了吧',
        onOk: () => {
            updateRole(roleId)
        }
    });
    roleVisible.value = false;
}
// 请求更新角色默认角色
const updateRole = async (roleId: string) => {
    const param: UpdateSystemUserRoleParams = {
        roleId: String(roleId),
        userId: userInfo.value.id
    }
    const {code} = await updateUserRole(param)
    if (code === 200) {
        alertMsg("success", "切换成功")
    }
    await store.dispatch('authModule/reLogin')
    await router.push('/login')
}
// 当前角色警告
const isCurrentRoleMsg = () => {
    alertMsg('warning', '你现在已经是这个角色啦！')
}

// 显示系统配置抽屉
const showSystemSetup = () => {
    visible.value = false
    emit("show")
}

// 注销系统
const hide = () => {
    store.dispatch('authModule/reLogin')
    router.push("/login")
    visible.value = false
}

const scrrent = () => {
    if (screenfull.isEnabled) {
        screenfull.toggle()
    }
}

</script>
<template>
    <div class="system-setup">
        <!-- 全屏设置 -->
        <div class="setup-item" @click="scrrent">
            <ExpandOutlined/>
        </div>
        <!-- 搜索设置 todo -->
        <div class="setup-item">
            <SearchOutlined/>
        </div>
        <!-- 用户角色切换 -->
        <a-popover overlayClassName="setup-popover" v-model:visible="roleVisible" placement="bottom" trigger="click">
            <template #content>
                <div
                    v-for="item in userInfo.role"
                    :key="item.roleId"
                    :class="['setup-select-item', item.roleId === currentInfo?.roleId ? 'setup-item-active' : '']"
                    @click="item.roleId !== currentInfo?.roleId ? roleHide(item.roleId) : isCurrentRoleMsg()"
                >
                    <PoweroffOutlined/>
                    <span class="setup-select-item-title">{{ item.roleName }}</span>
                </div>
            </template>
            <template #title>
                <div>
                    <UserSwitchOutlined/>
                    <span class="setup-select-title">当前角色：{{ currentInfo?.roleName }}</span>
                </div>
            </template>
            <div class="setup-item">
                <UserSwitchOutlined/>
            </div>
        </a-popover>
        <!-- 系统设置 -->
        <a-popover overlayClassName="setup-popover" v-model:visible="visible" placement="bottom" trigger="click">
            <template #content>
                <div class="setup-select-item" @click="showSystemSetup">
                    <BgColorsOutlined/>
                    <span class="setup-select-item-title">个性设置</span>
                </div>
                <a-popconfirm title="你确认要注销系统吗？" ok-text="是的" cancel-text="算了吧" @confirm="hide">
                    <div class="setup-select-item">
                        <PoweroffOutlined/>
                        <span class="setup-select-item-title">注销系统</span>
                    </div>
                </a-popconfirm>
            </template>
            <template #title>
                <div>
                    <SettingOutlined/>
                    <span class="setup-select-title">系统设置</span>
                </div>
            </template>
            <div class="setup-item">
                <SettingOutlined/>
            </div>
        </a-popover>
    </div>
</template>
<style lang="scss" scoped>
@import "./index.scss";
</style>
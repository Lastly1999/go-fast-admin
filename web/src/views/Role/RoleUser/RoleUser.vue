<script lang="ts" setup>
import {computed, onMounted, reactive, ref, watch} from "vue"
import {useStore} from "vuex"

import type {QueryJsonItem} from "@/components/QueryGroup/QueryGroup.vue"
// components
import QueryGroup from "@/components/QueryGroup/QueryGroup.vue"
import FContainer from "@/components/FContainer/FContainer.vue"
import FTable from "@/components/FTable/FTable.vue"
import type {UserForm} from "./components/RoleUserModalForm/RoleUserModalForm.vue"
import RoleUserModalForm from "./components/RoleUserModalForm/RoleUserModalForm.vue"

// types
import type {RegisterUserForm, UserSystem} from "@/services/model/response/user"
import type {listParams} from "@/services/model/response/public"
import type {RoleItem} from "@/store/modules/system"


// apis
import {createSystemUser, deleteSystemUser, editSystemUser, getUsers} from "@/services/user/user"

// utils
import {alertMsg} from "@/utils/antd/antd"
import {createMd5Pass} from "@/utils/md5/md5"
import {Moment} from "moment";

onMounted(() => {
    getSystemUsers()
})

const store = useStore()

// 修改用户信息
const editUserRow = (data: UserForm): void => {
    console.log(data)
    userForm.value = {
        ...data,
        roleId: data.roleId === "" ? null : data.roleId
    }
    userModalTitle.value = "修改用户信息"
    userModalVisible.value = true
}

// 新增用户信息
const append = (): void => {
    userForm.value = {
        id: null,
        userName: "",
        nikeName: "",
        userAvatar: "",
        roleId: null,
        CreatedAt: "",
        UpdatedAt: ""
    }
    userModalTitle.value = "新增用户"
    userModalVisible.value = true
}

const rangeDateChange = (date: string[]) => {
    queryForm.value.startTime = date[0]
    queryForm.value.endTime = date[1]
}

// 查询组件渲染数据源
const queryJsonData = reactive<QueryJsonItem[]>([
    {
        type: "input",
        placeholder: "请输入姓名、账户名称",
        prop: "keywords"
    },
    {
        type: "rangePicker",
        placeholder: "请选择时间",
        size: "default",
        prop: "time",
        change: rangeDateChange
    },
    {
        type: "button",
        buttonType: "primary",
        text: "新增用户",
        fun: append
    }
])

// 表格头数据
const columns = [
    {
        title: '用户序号',
        dataIndex: 'id',
        align: 'center',
        key: 'id',
        width: 80
    },
    {
        title: '账户名称',
        dataIndex: 'userName',
        key: 'userName',
        align: "center",
        width: 130
    },

    {
        title: '姓名',
        dataIndex: 'nikeName',
        key: 'nikeName',
        width: 130,
    },
    {
        title: "所有角色",
        width: 300,
        slots: {customRender: "tags"},
    },
    {
        title: '创建时间',
        dataIndex: 'CreatedAt',
        key: 'CreatedAt',
        width: 220,
    },
    {
        title: '最近更新',
        dataIndex: 'UpdatedAt',
        key: 'UpdatedAt',
        width: 220,
    },
    {
        title: "操作",
        fixed: "right",
        slots: {customRender: "action"},
    },
]
// 用户列表加载状态
const userTableLoading = ref(false)
// 用户列表数据源
const userData = ref<UserSystem[]>()

// queryGroup 查询表单
const queryForm = ref<listParams>({
    keyWords: "",
    pageSize: 10,
    page: 1,
    startTime: '',
    endTime: '',
    time: []
})

watch(() => queryForm, val => {
    getSystemUsers()
}, {deep: true})

// 用户修改/新增表单按钮加载状态
const formLoading = ref<boolean>(false)

// 用户修改modal标题
const userModalTitle = ref<string>('新增用户')

// 用户修改modal显示状态
const userModalVisible = ref<boolean>(false)

// 用户修改modal表单数据
const userForm = ref<UserForm>({
    id: null,
    userName: "",
    nikeName: "",
    passWord: "",
    userAvatar: "",
    roleId: null,
    CreatedAt: "",
    UpdatedAt: ""
})

// 提交用户操作 添加/修改
const submitForm = async (): Promise<void> => {
    formLoading.value = true
    if (userForm.value.id) {
        await updateFormRequest()
    } else {
        await addFormRequest()
    }
    userModalVisible.value = false
    await getSystemUsers()
    formLoading.value = false
}

// 请求新增用户
const addFormRequest = async (): Promise<void> => {
    const param: RegisterUserForm = {
        userName: userForm.value.userName,
        passWord: userForm.value.passWord,
        nikeName: userForm.value.nikeName,
        roleId: userForm.value.roleId,
        roleIds: userForm.value.roleIds,
    }
    const {code} = await createSystemUser(param)
    if (code === 200) {
        alertMsg("success", "新增成功!")
    }
}

// 请求修改用户信息
const updateFormRequest = async (): Promise<void> => {
    const param: RegisterUserForm = {
        id: userForm.value.id,
        userName: userForm.value.userName,
        nikeName: userForm.value.nikeName,
        roleId: userForm.value.roleId,
        roleIds: userForm.value.roleIds,
    }
    const {code} = await editSystemUser(param)
    if (code === 200) {
        alertMsg("success", "修改成功!")
    }
}

// 请求系统用户列表
const getSystemUsers = async (): Promise<void> => {
    userTableLoading.value = true
    const {code, data} = await getUsers(queryForm.value)
    if (code === 200) {
        userData.value = data.users
        eachRoleIds()
    }
    userTableLoading.value = false
}

// 需要的角色id
const eachRoleIds = (): void => {
    userData.value!.map((item: UserSystem) => {
        item.roleIds = []
        item.role.map(roleItem => {
            item.roleIds.push(roleItem.roleId)
        })
    })
}

// 请求删除系统用户
const delSystemUser = async (id: number): Promise<void> => {
    const {code} = await deleteSystemUser(id)
    if (code === 200) {
        alertMsg("success", "删除成功")
    }
    await getSystemUsers()
}

const roleOptions = computed<RoleItem[]>(() => store.getters["systemModule/getSysRoles"])

</script>
<template>
    <FContainer>
        <template v-slot:header>
            <QueryGroup v-model:jsonData="queryJsonData" v-model:form="queryForm"/>
        </template>
        <template v-slot:main>
            <FTable bordered size="middle" :loading="userTableLoading" :columns="columns" :data-source="userData"
                    rowKey="id">
                <template #tags="{ data }">
                    <a-tag v-for="item in data.role" color="green">{{ item.roleName }}</a-tag>
                </template>
                <template #action="{ data }">
                    <a @click="editUserRow(data)">修改</a>
                    <a-divider type="vertical"/>
                    <a-popconfirm title="你确定要删除该菜单吗?删除后无法恢复!" ok-text="是的" cancel-text="算了吧"
                                  @confirm="delSystemUser(data.id)">
                        <a>删除</a>
                    </a-popconfirm>
                </template>
            </FTable>
            <RoleUserModalForm v-model:loading="formLoading" v-model:visible="userModalVisible"
                               v-model:title="userModalTitle" :form="userForm" @submit="submitForm"/>
        </template>
    </FContainer>
</template>
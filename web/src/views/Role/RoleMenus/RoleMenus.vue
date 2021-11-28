<script lang="ts" setup>
import {reactive, ref, onMounted} from "vue"

// types
import type {QueryJsonItem} from "@/components/QueryGroup/QueryGroup.vue"
// apis
import {getAllSysMenus, putSystemMenu, deleteSystemMenu} from "@/services/auth"
import {alertMsg} from "@/utils/antd/antd"
import {listToTree} from "utils/loadsh/data"
// components
import QueryGroup from "@/components/QueryGroup/QueryGroup.vue"
import FContainer from "@/components/FContainer/FContainer.vue"
import FTable from "@/components/FTable/FTable.vue"
import MenuDrawerForm from "./components/MenuDrawerForm/MenuDrawerForm.vue"
import MenuEdit from "./components/MenuEdit/MenuEdit.vue"
import {Icon} from "@/components/FIcon/FIcon.ts"

import type {MenuInfo} from "@/services/model/response/role"
import type {MenuFormOptions} from "@/views/Main/Role/RoleMenus/components/MenuDrawerForm/MenuDrawerForm.vue"

export type MenuListItem = {
    icon: string;
    id: number;
    label: string;
    pId: number;
    pName: string;
    path: string;
    children?: MenuListItem[]
}

onMounted(() => {
    getRoleMenus()
})

// query 搜索方法
const search = () => {
    console.log(queryForm.value)
    console.log('search')
}
// query 重置方法
const reset = () => {
    queryForm.value = {name: ''}
}
// query 新增方法
const append = () => {
    menuModalForm.value = {
        id: 0,
        icon: null,
        label: '',
        path: '',
        pId: 0,
        pName: '',
        pPath: '',
        pIcon: null
    }
    showMenuTitle.value = "新增根菜单"
    showMenuDrawer.value = true
}
// 查询组件渲染数据源
const queryJsonData = reactive<QueryJsonItem[]>([
    {
        type: "button",
        buttonType: "primary",
        text: "新增根菜单",
        fun: append
    }
])
// 查询的表单
const queryForm = ref({
    name: ""
})
// table loading状态
const menuTableLoading = ref<boolean>(false)
// 表格菜单数据源
const menuData = ref<[]>([])
// 表格头数据
const columns = [
    {
        title: '菜单序号',
        dataIndex: 'id',
        key: 'id',
        width: 150
    },
    {
        title: '菜单名称',
        dataIndex: 'label',
        key: 'label',
        width: 150
    },
    {
        title: "菜单图标",
        width: 200,
        slots: {customRender: "icon"},
    },
    {
        title: '父级菜单序号',
        dataIndex: 'pId',
        key: 'pId',
        width: 150
    },
    {
        title: '父级菜单名称',
        dataIndex: 'pName',
        key: 'pName',
        width: 150
    },
    {
        title: '路由路径',
        dataIndex: 'path',
        key: 'path',
        width: 150
    },
    {
        title: "操作",
        fixed: "right",
        slots: {customRender: "action"},
    },
]
// 系统菜单编辑/新增状态
const showMenuDrawer = ref(false)
// 系统菜单抽屉名称
const showMenuTitle = ref("")
// 编辑权限弹窗按钮加载状态
const menuFormloading = ref(false)
// 系统系统编辑/新增表单
const menuModalForm = ref<MenuFormOptions>({
    id: 0,
    icon: undefined,
    label: '',
    path: '',
    pId: 0,
    pName: '',
    pPath: '',
    pIcon: undefined
})
// 表格行删除菜单操作
const removeMenu = async (menuItem: MenuFormOptions) => {
    const {code} = await deleteSystemMenu(menuItem.id)
    if (code === 200) {
        alertMsg("success", "删除菜单成功！")
        await getRoleMenus()
    }
}
// 请求权限菜单数据源
const getRoleMenus = async () => {
    const {code, data} = await getAllSysMenus()
    if (code === 200) {
        menuData.value = listToTree(data.menus)
    }
}
// 列表添加子菜单按钮方法
const appendChildren = (data: MenuListItem) => {
    menuModalForm.value = {
        id: data.id,
        icon: data.icon,
        label: data.label,
        path: data.path,
        pId: data.pId
    }
    showMenuTitle.value = "添加子菜单"
    showMenuDrawer.value = true
}
// modal 权限菜单修改提交
const formSubmit = async (form: MenuFormOptions) => {
    menuFormloading.value = true
    const params: MenuInfo = {
        menuName: form.id ? form.pName : form.label,
        menuIcon: form.icon,
        menuPath: form.id ? form.pPath : form.path,
        menuParentId: form.id,
        menuParentName: form.pName
    }
    const {code} = await putSystemMenu(params)
    if (code === 200) {
        alertMsg("success", "新增根菜单成功！")
        await getRoleMenus()
    }
    menuFormloading.value = false
    showMenuDrawer.value = false
}

// 编辑菜单的id
const menuEditId = ref<string>("")

// 编辑菜单dialog显示状态
const menuEditVisible = ref(false)

// 权限菜单列表编辑行
const editRow = (data: MenuListItem) => {
    menuEditId.value = String(data.id)
    menuEditVisible.value = true
}

// 权限菜单编辑成功
const editSuccess = () => {
    menuEditVisible.value = false
    alertMsg("success", "编辑菜单成功！")
    getRoleMenus()
}


</script>

<template>
    <FContainer>
        <template v-slot:header>
            <QueryGroup v-model:jsonData="queryJsonData" v-model:form="queryForm"/>
        </template>
        <template v-slot:main>
            <FTable bordered size="middle" :loading="menuTableLoading" :columns="columns" :data-source="menuData"
                    rowKey="id">
                <template #icon="{ data }">
                    <Icon :icon="data.icon"/>
                    {{ data.icon }}
                </template>
                <template #action="{ data }">
                    <a @click="appendChildren(data)">添加子菜单</a>
                    <a-divider type="vertical"/>
                    <a @click="editRow(data)">编辑</a>
                    <a-divider type="vertical"/>
                    <a-popconfirm title="你确定要删除该菜单吗?删除后无法恢复!" ok-text="是的" cancel-text="算了吧"
                                  @confirm="removeMenu(data)">
                        <a>删除</a>
                    </a-popconfirm>
                </template>
            </FTable>
            <MenuDrawerForm v-model:visible="showMenuDrawer" v-model:title="showMenuTitle"
                            v-model:loading="menuFormloading" :form="menuModalForm" @submit="formSubmit"/>
            <!-- 编辑菜单修改 -->
            <MenuEdit v-model:visible="menuEditVisible" :id="menuEditId" @success="editSuccess"/>
        </template>
    </FContainer>
</template>

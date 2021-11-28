import {RouteRecordRaw} from "vue-router"
import Layout from '@/layout/index.vue'


/**
 * role routes config
 * @author lastly
 * @date 2021年8月28日23:31:22
 */
const roleRoutesConf: RouteRecordRaw = {
    path: "/role",
    name: "Role",
    component: Layout,
    meta: {
        title: "权限管理"
    },
    children: [
        {
            path: "permissions",
            name: "role-permissions",
            component: () => import("@/views/Role/Permissions/Permissions.vue"),
            meta: {
                title: "角色权限",
                role: "ROLE:ROLE:PERMISSIONS"
            }
        },
        {
            path: "menus",
            name: "role-menus",
            component: () => import("@/views/Role/RoleMenus/RoleMenus.vue"),
            meta: {
                title: "菜单权限",
                role: "ROLE:ROLE:MENUS"
            }
        },
        {
            path: "user",
            name: "role-user",
            component: () => import("@/views/Role/RoleUser/RoleUser.vue"),
            meta: {
                title: "用户管理",
                role: "ROLE:ROLE:USER"
            }
        },
    ]
}

export default roleRoutesConf
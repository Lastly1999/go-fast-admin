import { RouteRecordRaw } from "vue-router"

// router modules
import dashboardRoutesConf from "@/router/modules/dashboard/dashboard"
import roleRoutesConf from "@/router/modules/role/role"
import systemRoutesConf from "@/router/modules/system/system"
import componentsRoutesConf from "@/router/modules/components/components"


/**
 * 各路由modules在此处做装箱操作 子模块默认格式为以下：
 * 子模块无下级菜单时 请务必将path设置为"" 其他情况下则设置为子路由的自定义路径
 */
/* const testRoutesConf: RouteRecordRaw = {
    path: "/test",
    name: "Test",
    component: Layout,
    meta: {
        title: "组件"
    },
    children: [
        {
            path: "",
            name: "test-index",
            component: () => import("@/views/Main/Test/Test.vue"),
            meta: {
                title: "组件",
                role: "HOME:PANEL:VIEW"
            }
        }
    ]
} */
const basicRoutes: RouteRecordRaw[] = [
    roleRoutesConf,
    dashboardRoutesConf,
    systemRoutesConf,
    componentsRoutesConf
]

export default basicRoutes
import { RouteRecordRaw } from 'vue-router'
import Layout from "@/layout/index.vue"

/**
 * dashboard routes config
 * @author lastly
 * @date 2021年8月28日23:27:44
 */
const dashboardRoutesConf: RouteRecordRaw = {
    path: "/dashboard",
    name: "Dashboard",
    component: Layout,
    meta: {
        title: "首页"
    },
    children: [
        {
            path: "",
            name: "panel-index",
            component: () => import("@/views/dashboard/Panel/Panel.vue"),
            meta: {
                title: "",
                role: "HOME:PANEL:VIEW"
            }
        }
    ]
}

export default dashboardRoutesConf
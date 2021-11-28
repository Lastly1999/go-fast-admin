import {RouteRecordRaw} from "vue-router"
import Layout from '@/layout/index.vue'


/**
 * role routes config
 * @author lastly
 * @date 2021年8月28日23:31:22
 */
const systemRoutesConf: RouteRecordRaw = {
    path: "/system",
    name: "System",
    component: Layout,
    meta: {
        title: "系统管理"
    },
    children: [
        {
            path: "apilogs",
            name: "Apilogs",
            meta: {
                title: "api日志"
            },
            component: () => import("@/views/System/ApiLogs/ApiLogs.vue")
        }
    ]
}

export default systemRoutesConf
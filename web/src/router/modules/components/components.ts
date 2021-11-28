import { RouteRecordRaw } from 'vue-router'
import Layout from "@/layout/index.vue"

/**
 * components routes config
 * @author lastly
 * @date 2021年10月21日10:17:40
 */
const componentsRoutesConf: RouteRecordRaw = {
    path: "/components",
    name: "Components",
    component: Layout,
    meta: {
        title: "组件"
    },
    children: [
        {
            path: "form",
            name: "components-public-form",
            component: () => import("@/views/Component/PublicForm/PublicForm.vue"),
            meta: {
                title: "公用表单",
                role: "HOME:PANEL:VIEW"
            }
        }
    ]
}

export default componentsRoutesConf
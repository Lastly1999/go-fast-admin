import { RouteRecordRaw } from 'vue-router'

const bootPath = "/home"

// project routes
export const routesConf: RouteRecordRaw[] = [
    {
        path: `${bootPath}/panel`,
        name: "panel",
        component: () => import('views/Main/Panel/Panel.vue'),
        meta: {
            role: 'HOME:PANEL:VIEW'
        }
    },
    {
        path: `${bootPath}/component`,
        name: "component",
        component: () => import('views/Main/Component/Component.vue'),
        meta: {
            role: 'HOME:COMPONENTS:VIEW'
        }
    },
    {
        path: `${bootPath}/tools`,
        name: "tools",
        component: () => import('views/Main/Tools/Tools.vue'),
        meta: {
            role: 'HOME:TOOLS:VIEW'
        }
    },
    {
        path: `${bootPath}/role`,
        name: "role",
        component: () => import('views/Main/Role/Role.vue'),
        meta: {
            role: 'HOME:ROLE:VIEW'
        }
    }
]
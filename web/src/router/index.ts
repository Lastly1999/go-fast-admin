import {
    createRouter,
    createWebHistory,
    RouteLocationNormalized,
    NavigationGuardNext,
    NavigationFailure,
    Router,
    RouteRecordRaw
} from "vue-router"
import routesStoreModules from "./modules"
import nprogress from 'nprogress'
import 'nprogress/nprogress.css'

// vuex
import store from '@/store'

// 路由白名单表
const routerAuthList = ["HOME:PANEL:VIEW"]

export const routes: RouteRecordRaw[] = [
    {
        path: "/",
        redirect: "/login",
    },
    {
        path: "/login",
        component: () => import("views/Login/Login.vue"),
    },
    {
        path: "/notAuth",
        name: "NotAuth",
        component: () => import("@/views/NotAuth/NotAuth.vue"),
        meta: {
            title: "暂无权限"
        }
    },
    ...routesStoreModules
]

// 构建路由栈
const router: Router = createRouter({
    history: createWebHistory(),
    routes,
})


const authEachArrays = (routesStore: RouteRecordRaw[]) => {
    const each = () => {
        routesStore.map((item: RouteRecordRaw) => {
            // 增加权限标识
            if (routerAuthList.includes(item.meta?.role as string)) {
                (item as any).isAuth = true
            } else {
                (item.meta as any).isAuth = false
            }
            if (item.children) {
                authEachArrays(item.children)
            }
        })
    }
    each()
}


// todo 权限验证在除开登录页之外pages处理
// 进行用户id请求权限菜单操作
router.beforeEach(async (to: RouteLocationNormalized, form: RouteLocationNormalized, next: NavigationGuardNext) => {
    // 如果是登录页 默认放行 不进行权限验证
    if (to.path === '/login') {
        next()
    } else {
        const menus = store.getters["authModule/getSysMenus"]
        if (!menus || menus.length === 0) {
            await store.dispatch("systemModule/API_GET_SYS_ROLES")
            await store.dispatch("authModule/API_GET_SYS_MENUS") // 全部系统菜单
            await store.dispatch("authModule/API_GET_SYS_USER_INFO") // 系统用户信息
            // 白名单列表
            const whiteList: string[] = store.getters["authModule/getWhiteList"]
            if (whiteList.findIndex(item => item === to.path) === -1) {
                next('/notAuth')
            }
        }
        next()
    }
    // nprogress start...
    nprogress.start()
})


router.afterEach((to: RouteLocationNormalized, form: RouteLocationNormalized, next: void | NavigationFailure | undefined) => {
    // nprogress end...
    nprogress.done()
})

export default router

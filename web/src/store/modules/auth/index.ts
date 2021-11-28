import { Store } from 'vuex'
import router from "@/router"
import { listToTree, generateRouteWhiteList } from "utils/loadsh/data"
import type { LoginForm, UserInfo } from "@/services/model/response/role"

// apis
import { checkAuthUser } from "@/services/auth"
import { getSysMenus, getSystemUserInfo } from '@/services/auth'


export type AuthState = {
    sysMenus: [] | undefined;
    roleWhiteRoles: string[],
    userInfo: UserInfo;
}

export type AuthModule = {
    namespaced: boolean;
    state: () => AuthState;
    getters: {
        getSysMenus: (state: AuthState) => [] | undefined;
        getUserInfo: (state: AuthState) => UserInfo;
        getWhiteList: (state: AuthState) => string[]
    };
    actions: {
        API_POST_SYS_AUTH({ commit }: Store<any>, payload: LoginForm): void;
        API_GET_SYS_MENUS({ commit }: Store<any>, id: string | number): void;
        API_GET_SYS_USER_INFO({ commit }: Store<any>): void;
        reLogin({ commit }: Store<any>): void;
    };
    mutations: {
        SET_SYS_MENUS(state: AuthState, payload: []): void;
        SET_USER_INFO(state: AuthState, payload: UserInfo): void;
        SET_ROLE_WHITE_LIST(state: AuthState, payload: string[]): void;
    };
}

/**
 * auth config module
 * @author lastly
 * @date 2021年8月26日22:57:12
 */
const authModule: AuthModule = {
    namespaced: true,
    state: () => (
        {
            sysMenus: undefined,
            roleWhiteRoles: [],
            userInfo: {
                id: 0,
                nikeName: '',
                role: [],
                roleId: '',
                userAvatar: '',
                userName: '',
            }
        }
    ),
    getters: {
        getSysMenus: (state: AuthState) => state.sysMenus,
        getUserInfo: (state: AuthState) => state.userInfo,
        getWhiteList: (state: AuthState) => state.roleWhiteRoles
    },
    actions: {
        /**
         * 系统登录授权
         * @param commit
         * @param payload
         */
        async API_POST_SYS_AUTH({ commit }: Store<any>, payload: LoginForm) {
            const { code, data } = await checkAuthUser<LoginForm>(payload)
            if (code === 200) {
                localStorage.setItem("auth-token", data.token)
                await router.push('/dashboard')
            } else {
                throw new Error(data)
            }
        },
        /**
         * 注销登录
         * @param param0 
         * @param id 
         */
        async reLogin({ commit }: Store<any>) {
            localStorage.removeItem("user-token")
            localStorage.removeItem("user-info")
            commit('SET_SYS_MENUS', [])
        },
        /**
         * 获取权限系统菜单
         * @param commit
         * @param id
         */
        async API_GET_SYS_MENUS({ commit }: Store<any>, id: string | number) {
            const { code, data } = await getSysMenus(id)
            if (code === 200) {
                commit('SET_ROLE_WHITE_LIST', generateRouteWhiteList(data.menus))
                commit('SET_SYS_MENUS', listToTree(data.menus))
            }
        },
        /**
         * 获取系统用户信息
         * @param param
         */
        async API_GET_SYS_USER_INFO({ commit }: Store<any>) {
            const { code, data } = await getSystemUserInfo()
            if (code === 200) localStorage.setItem('user-info', JSON.stringify(data.userInfo))
        }
    },
    mutations: {
        /**
         * commit 设置系统权限菜单
         * @param state
         * @param payload
         */
        SET_SYS_MENUS(state: AuthState, payload: []) {
            state.sysMenus = payload
        },
        /**
         * 设置系统用户信息
         * @param state
         * @param payload
         */
        SET_USER_INFO(state: AuthState, payload: UserInfo) {
            state.userInfo = payload
        },
        /**
         * 设置用户白名单
         * @param state
         * @param payload
         */
        SET_ROLE_WHITE_LIST(state: AuthState, payload: string[]) {
            state.roleWhiteRoles = payload
        }
    }
}

export default authModule
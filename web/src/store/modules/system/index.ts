import { Store } from "vuex"

// apis
import { getSystemIcons } from "@/services/system/sys"
import { getAllSysMenus } from "@/services/auth"
import { getRoles } from "@/services/role"

export type IconItem = {
    id: number;
    iconName: string;
}

export type MenuItem = {
    CreatedAt: string;
    DeletedAt: null | string;
    Role: null
    UpdatedAt: null | string;
    icon: null | string;
    id: number;
    label: null | string;
    pId: number;
    path: string;
}

export type RoleItem = {
    roleId?: string;
    roleName?: string;
    describe?: string;
    createDate?: string;
    state?: number;
}

export type SystemState = {
    iconSelectDataSource: IconItem[],
    systemMenus: MenuItem[],
    systemRoles: RoleItem[]
}

export type SystemModule = {
    namespaced: boolean;
    state: () => SystemState;
    getters: {
        getSysIcons(state: SystemState): IconItem[],
        getSysMenus(state: SystemState): MenuItem[],
        getSysRoles(state: SystemState): RoleItem[],
    };
    actions: {
        API_GET_SYS_ICONS(action: Store<any>): void,
        API_GET_SYS_MENUS(action: Store<any>): void,
        API_GET_SYS_ROLES(action: Store<any>, userId: number): void
    };
    mutations: {
        SET_SYS_ICONS(state: SystemState, payload: IconItem[]): void,
        SET_SYS_MENUS(state: SystemState, payload: MenuItem[]): void,
        SET_SYS_ROLES(state: SystemState, payload: RoleItem[]): void,
    }
}

export const systemModule: SystemModule = {
    namespaced: true,
    state: () => {
        return {
            iconSelectDataSource: [],
            systemMenus: [],
            systemRoles: []
        }
    },
    getters: {
        getSysIcons(state: SystemState) {
            return state.iconSelectDataSource
        },
        getSysMenus(state: SystemState) {
            return state.systemMenus
        },
        getSysRoles(state: SystemState) {
            return state.systemRoles
        }
    },
    actions: {
        /**
         * 请求系统图标列表
         * @param param0 
         */
        async API_GET_SYS_ICONS({ commit }: Store<any>) {
            const { data, code } = await getSystemIcons()
            if (code === 200) commit('SET_SYS_ICONS', data.icons)
        },
        /**
         * 请求系统全部菜单列表
         * @param param
         */
        async API_GET_SYS_MENUS({ commit }: Store<any>) {
            const { data, code } = await getAllSysMenus()
            if (code === 200) commit('SET_SYS_MENUS', data.menus)
        },
        /**
         * 请求系统角色列表
         * @param param
         */
        async API_GET_SYS_ROLES({ commit }: Store<any>) {
            const { data, code } = await getRoles()
            if (code === 200) commit('SET_SYS_ROLES', data.roles)
        }
    },
    mutations: {
        // set icons
        SET_SYS_ICONS(state: SystemState, payload: IconItem[]) {
            state.iconSelectDataSource.length === 0 && (state.iconSelectDataSource = payload)
        },
        // set menus
        SET_SYS_MENUS(state: SystemState, payload: MenuItem[]) {
            state.systemMenus.length === 0 && (state.systemMenus = payload)
        },
        // set roles
        SET_SYS_ROLES(state: SystemState, payload: RoleItem[]) {
            state.systemRoles = payload
        }
    }
}

export default systemModule
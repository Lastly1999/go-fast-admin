import { Store } from "vuex"

export type SystemTheme = 'dark' | "light"

export type ConfigState = {
  systemTheme: SystemTheme
}

export type configModule = {
  namespaced: boolean;
  state: () => ConfigState;
  getters: {
    getSystemTheme: (state: ConfigState) => SystemTheme
  },
  actions: {

  };
  mutations: {

  }
}
/**
 * 系统全局配置仓库
 * @date 2021年10月13日14:57:58
 */
export const configModule: configModule = {
  namespaced: true,
  state: () => {
    return {
      systemTheme: 'dark'
    }
  },
  getters: {
    getSystemTheme: (state: ConfigState) => {
      return state.systemTheme
    }
  },
  actions: {
    SET_SYSTEM_THEME: async ({ commit }: Store<any>, theme: SystemTheme) => {
      // 本地数据持久化
      localStorage.setItem("sys-theme", theme)
      // store 数据缓存
      await commit("SET_SYSTEM_THEME_CACHE", theme)
    }
  },
  mutations: {
    SET_SYSTEM_THEME_CACHE: (state: ConfigState, payload: SystemTheme) => {
      state.systemTheme = payload
    }
  }
}

export default configModule
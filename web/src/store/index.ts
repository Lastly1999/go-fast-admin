import { createStore, Store } from "vuex"

// 分装仓库 非module式
import State from './apis/state'
import Actions from './apis/actions'
import Getters from './apis/getters'
import Mutations from './apis/mutations'

// modules 仓库分装
import Modules from './modules/install'

/**
 * store instance vuex实例
 * @author lastly
 * @date 2021年8月10日00:36:37
 */
const store: Store<any> = createStore({
    state: { ...State } as any,
    actions: { ...Actions } as any,
    mutations: { ...Mutations } as any,
    getters: { ...Getters } as any,
    modules: { ...Modules } as any
})

export default store
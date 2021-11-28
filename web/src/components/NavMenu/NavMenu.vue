<template>
    <a-menu
        v-model:openKeys="openKeys"
        v-model:selectedKeys="selectedKeys"
        mode="inline"
        :theme="sysTheme"
        :inline-collapsed="collapsed"
        @select="selectMenuItem"
        @openChange="onOpenChange"
    >
        <template v-for="item in $props.list" :key="item.path">
            <template v-if="!item?.children">
                <a-menu-item :key="item.path">
                    <template #icon>
                        <Icon :icon="item.icon"></Icon>
                    </template>
                    {{ item.label }}
                </a-menu-item>
            </template>
            <template v-else>
                <sub-menu :menu-info="item" :key="item.path" />
            </template>
        </template>
    </a-menu>
</template>
<script lang="ts">
import { defineComponent, PropType, watch, reactive, toRefs, onMounted, computed } from 'vue'
import { useStore } from "vuex"
import { Icon } from '@/components/FIcon/FIcon'
import { useRouter, useRoute } from 'vue-router'

export type MenuItem = {
    path: string;
    label: string;
    children: MenuItem[];
    icon: string;
}

export type Theme = "dark" | "light"

// you can rewrite it to a single file component, if not, you should config vue alias to vue/dist/vue.esm-bundler.js
const SubMenu = {
    name: 'SubMenu',
    props: {
        menuInfo: {
            type: Object as PropType<MenuItem>,
            default: () => ({}),
        },
    },
    template: `
        <a-sub-menu :key="menuInfo.path">
        <template #title>
            {{ menuInfo.label }}
        </template>
        <template #icon>
            <Icon :icon="menuInfo.icon"></Icon>
        </template>
        <template v-for="item in menuInfo.children" :key="item.path">
            <template v-if="!item.children">
                <a-menu-item :key="item.path">
                    <template #icon>
                        <Icon :icon="item.icon"></Icon>
                        <i :class="['fa',item.icon]"></i>
                    </template>
                    {{ item.label }}
                </a-menu-item>
            </template>
            <template v-else>
                <sub-menu :menu-info="item" :key="item.path"/>
            </template>
        </template>
        </a-sub-menu>
    `,
    components: {
        Icon
    },
}

export default defineComponent({
    components: {
        'sub-menu': SubMenu,
        Icon
    },
    props: {
        config: {
            type: Object as PropType<MenuItem>,
            default: () => {
                return {}
            },
        },
        list: {
            type: Array as PropType<MenuItem[]>,
            default: () => [],
        },
    },
    setup() {
        onMounted(() => {
            loadInitSystemTheme()
        })
        const router = useRouter()
        const route = useRoute()
        const store = useStore()
        const menuState = reactive({
            collapsed: false as boolean,
            selectedKeys: [] as string[],
            openKeys: [] as string[],
        })
        // watch system 路由栈
        watch(
            () => route.path,
            (nv, _) => {
                menuState.selectedKeys = [route.path]
                // 处理自动展开的需要父级的key值
                const mapKeys: string = '/' + route.path.split('/')[route.path.split('/').length - 2]
                menuState.openKeys = [mapKeys]
            },
            { deep: true, immediate: true }
        )
        // 系统主题
        const sysTheme = computed<Theme>(() => store.getters["configModule/getSystemTheme"])
        // 初始化主题配置
        const loadInitSystemTheme = (): void => {
            console.log(sysTheme.value)
            const cache = localStorage.getItem("sys-theme") as Theme
            if (cache) {
                store.dispatch("configModule/SET_SYSTEM_THEME", cache)
            }
        }
        // 打开菜单item
        const selectMenuItem = (item: any) => {
            router.push(item!.key)
        }
        // 打开菜单
        const onOpenChange = (openKeys: any) => {
        }
        return {
            ...toRefs(menuState),
            sysTheme,
            onOpenChange,
            selectMenuItem,
        }
    },
})
</script>
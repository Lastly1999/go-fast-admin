
<script lang="ts" setup>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { Layout } from 'ant-design-vue'
import config from "@/config/system.config"

// types
import type { SystemTheme } from '@/store/modules/config'

// layout tools
import NavMenu from '@/components/NavMenu/NavMenu.vue'
const Sider = Layout.Sider

const store = useStore()
const roleMenus = store.getters['authModule/getSysMenus']

// 系统cache的主题配置
const systemTheme = computed<SystemTheme>(() => store.getters["configModule/getSystemTheme"])

</script>

<template>
    <Sider :theme="systemTheme" class="system-sider" width="220" collapsible>
        <div class="logo">
            <div class="logo-container">
                <img width="50" height="50" :src="config.logo" />
            </div>
            <div class="logo-title" :style="{ color: systemTheme === 'light' ? '#333' : '#fff' }">Fast Admin</div>
        </div>
        <NavMenu :list="roleMenus" />
    </Sider>
</template>

<style lang="scss" scoped>
.system-sider {
    box-shadow: 0 0 10px 0px #cbc9d2;
    z-index: 999;
    .logo {
        padding: 10px;
        text-align: center;
        display: flex;
        align-items: center;
        .logo-container {
            width: 50px;
            height: 50px;
            overflow: hidden;
            border-radius: 50%;
            display: inline-block;
        }
        .logo-title {
            font-size: 20px;
            color: #fff;
            font-size: 20px;
            margin: 0 0 0 12px;
            font-family: Avenir, Helvetica Neue, Arial, Helvetica, sans-serif;
            font-weight: 600;
            vertical-align: middle;
            display: inline-block;
            color: #fff;
        }
    }
}
</style>
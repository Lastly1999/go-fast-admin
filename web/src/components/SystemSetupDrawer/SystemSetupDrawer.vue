<script lang="ts" setup>
import { ref, onMounted, computed } from "vue"
import { useStore } from "vuex"
import { CheckOutlined, CloseOutlined } from '@ant-design/icons-vue'
import { alertMsg } from "@/utils/antd/antd"
import type { SystemTheme } from "@/store/modules/config"

export type ThemeItem = {
  id: number;
  name: string;
  class: string;
}

export type CacheSystem = {
  treeMenuMode: "dark" | "light";
  theme: "default-theme" | "dark-blue-theme" | "red-theme" | "fag-purple-theme";
  themeId: number;
}

export type GlobalThemeOptions = {
  id: number;
  name: string;
  checked: boolean;
  mode: "default" | SystemTheme
}

defineProps({
  visible: {
    type: Boolean,
    default: () => false
  }
})

onMounted(() => {
  initSystemTheme()
  initGlobalSystemTheme()
})

const emit = defineEmits<{
  (event: 'update:visible', show: boolean): void
}>()

const store = useStore()

// 系统cache的主题配置
const systemTheme = computed(() => store.getters["configModule/getSystemTheme"])

// 抽屉排版方向
const placement = ref<string>('right')

// 默认主题色
const defaultTheme = ref<number>(0)

// 主题色
const themes = ref([
  {
    id: 0,
    name: "default-theme",
    class: 'check-default-theme',
  },
  {
    id: 1,
    name: "dark-blue-theme",
    class: 'check-drak-blue-theme',
  },
  {
    id: 2,
    name: "red-theme",
    class: 'check-red-theme',
  },
  {
    id: 3,
    name: "fag-purple-theme",
    class: 'check-fag-purple-theme',
  }
  // ...可继续添加主题色切换功能，前提是必须在scss文件中配置好相对应的显示颜色值，以及主题包的生成配置 by:lastly 2021年10月16日11:05:24
])

// 全局主题
const globalThemes = ref<GlobalThemeOptions[]>([
  {
    id: 1,
    name: '深灰模式',
    checked: false,
    mode: 'dark'
  },
  {
    id: 2,
    name: '亮色模式',
    checked: false,
    mode: 'light'
  }
])

const onClose = () => {
  emit('update:visible', false)
}

// 初始化全局主题
const initGlobalSystemTheme = (): void => {
  globalThemes.value.map(item => {
    if (item.mode === systemTheme.value) {
      item.checked = true
    }
  })
}

// 初始化主题色
const initSystemTheme = (): void => {
  const cache = localStorage.getItem('sys-config')
  if (!cache) {
    toggleTheme("default-theme", 0)
  } else {
    const sysConfig: CacheSystem = JSON.parse(cache as string)
    defaultTheme.value = sysConfig.themeId
    toggleTheme(sysConfig.theme, sysConfig.themeId)
  }
}

// 主题色切换
const themeChange = (event: ThemeItem): void => {
  defaultTheme.value = event.id
  toggleTheme(event.name, event.id)
  alertMsg("success", "现在你有了新的主题色！快去看看吧")
}

// 全局主题切换
const globalThemeChange = (event: GlobalThemeOptions): void => {
  globalThemes.value.map((item: GlobalThemeOptions) => {
    item.checked = false
  })
  event.checked = true
  store.dispatch("configModule/SET_SYSTEM_THEME", event.mode)
}

// 主题更换
const toggleTheme = (scopeName: string, id: number) => {
  styleSysSwitch(scopeName)
  localStorage.setItem('sys-config', JSON.stringify({ theme: scopeName, themeId: id } as CacheSystem))
}


// 此处 style link 切换的核心函数
// 这只 一个切换的妥协方法 还可继续优化 目标是抽离底层的--theme样式代码
// by lastly 2021年10月15日16:15:14
const styleSysSwitch = (scopeName: string) => {
  let styleLink: HTMLElement = document.getElementById("theme-link-tag") as HTMLElement;
  // 假如已经设置了theme-link-tag的style link 则直接修改为要修改的样式文件css
  if (styleLink) {
    styleLink.href = `/${scopeName}.css`;
    // 如果是removeCssScopeName:true移除了主题文件的权重类名，就可以不用修改className 操作
    document.documentElement.className = scopeName;
  } else {
    // 不存在的话，则新建一个 然后再设置为切换的theme style
    styleLink = document.createElement("link");
    styleLink.type = "text/css";
    styleLink.rel = "stylesheet";
    styleLink.id = "theme-link-tag";
    styleLink.href = `/${scopeName}.css`;
    document.documentElement.className = scopeName;
    document.head.append(styleLink);
  }
}

</script>
<template>
  <a-drawer width="300px" title="个性设置" :placement="placement" :closable="false" :visible="visible" @close="onClose">
    <a-divider>主题切换</a-divider>
    <div class="theme-switch-container">
      <div class="theme-switch-item" v-for="item in globalThemes">
        {{ item.name }}
        <a-switch :checked="item.checked" @change="globalThemeChange(item)">
          <template #checkedChildren>
            <check-outlined />
          </template>
          <template #unCheckedChildren>
            <close-outlined />
          </template>
        </a-switch>
      </div>
    </div>
    <a-divider>主题色</a-divider>
    <div class="theme-container">
      <div v-for="item in themes" :class="['theme-select', item.class]" @click="themeChange(item)">
        <transition appear name="fade" mode="out-in">
          <CheckOutlined v-if="defaultTheme === item.id" />
        </transition>
      </div>
    </div>
  </a-drawer>
</template>
<style lang="scss" scoped>
@import "./index.scss";
</style>
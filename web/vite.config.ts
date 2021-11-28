import {defineConfig, loadEnv} from "vite"
import vue from "@vitejs/plugin-vue"

// jsx plugin
import vueJsx from "@vitejs/plugin-vue-jsx"
import path from "path"
// @ts-ignore
import themePreprocessorPlugin from "@zougt/vite-plugin-theme-preprocessor";

// https://vitejs.dev/config/
export default ({mode}: any) => {
    process.env = {...process.env, ...loadEnv(mode, process.cwd())}

    return defineConfig({
        server: {
            proxy: {
                '/api': {
                    target: `${process.env.VITE_APP_API_URL}`,
                    changeOrigin: true,
                    rewrite: (path) => path.replace(/^\/api/, '')
                },
            }
        },
        css: {
            preprocessorOptions: {
                less: {
                    modifyVars: {
                        // 'primary-color': '#52c41a',
                        // 'link-color': '#1DA57A',
                        // 'border-radius-base': '2px'
                    },
                    javascriptEnabled: true
                }
            }
        },
        resolve: {
            alias: {
                'vue': 'vue/dist/vue.esm-bundler.js', // 定义vue的别名，如果使用其他的插件，可能会用到别名
                // 路径别名配置
                "@": path.resolve(__dirname, "src"),
                // 组件路径别名
                components: path.resolve(__dirname, "src/components"),
                // 视图组件路径别名
                views: path.resolve(__dirname, "src/views"),
                // 布局组件路径别名
                layouts: path.resolve(__dirname, "src/layouts"),
                // 工具类路径别名
                utils: path.resolve(__dirname, "src/utils"),
                // api层路径别名
                services: path.resolve(__dirname, "src/services"),
                // 类型层文件别名
                typings: path.resolve(__dirname, "src/typings"),
            },
        },
        plugins: [
            vueJsx(),
            vue(),
            themePreprocessorPlugin({
                less: {
                    multipleScopeVars: [
                        {
                            scopeName: "default-theme",
                            path: path.resolve('src/theme/default.less')
                        },
                        {
                            scopeName: "fag-purple-theme",
                            path: path.resolve('src/theme/fagPurple.less')
                        },
                        {
                            scopeName: "red-theme",
                            path: path.resolve('src/theme/red.less')
                        },
                        {
                            scopeName: "dark-blue-theme",
                            path: path.resolve('src/theme/darkBlue.less')
                        },
                        {
                            scopeName: "global-dark-theme",
                            path: path.resolve('src/theme/global/dark.less')
                        },
                    ],
                    defaultScopeName: "default-theme",
                    extract: false
                },

            })
        ],
    })

}
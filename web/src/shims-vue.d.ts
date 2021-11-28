declare module '*.vue' {
    import { DefineComponent } from 'vue'
    // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-typings
    const component: DefineComponent<{}, {}, any>
    export default component
}


declare module 'animate.css' // 声明animate.css的类型，否则引入animate.css会报错，提示找不到animate.css模块
declare module "@/*"
declare module "utils/*"
declare module "typings/*"
declare module "services/*"
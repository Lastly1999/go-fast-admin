<!-- 
 * 首页应用主入口
 * @author lastly
 * @date 2021年8月26日21:50:29
 -->
<script lang="ts" setup>
import { ref } from "vue";
import { ConfigProvider } from "ant-design-vue"
import zh_CN from 'ant-design-vue/lib/locale-provider/zh_CN';
import moment from 'moment';
import 'moment/locale/zh-cn';

const lang = ref(zh_CN)
moment.locale('zh-cn');

</script>
<template>
    <ConfigProvider :locale="lang">
        <div id="app">
            <router-view />
        </div>
    </ConfigProvider>
</template>
<style lang="scss">
@import "App.scss";
</style>
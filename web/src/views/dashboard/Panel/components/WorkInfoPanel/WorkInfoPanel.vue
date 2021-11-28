<script lang="ts" setup>
import { ref } from "vue"

const deadLine = ref<number>(Date.now() + 1000 * 60 * 60 * 24 * 2 + 1000 * 30)

const onFinish = () => {
    console.log('onFinish')
}

</script>
<template>
    <div class="work-info-panel">
        <a-row type="flex" justify="center">
            <a-col :span="6">
                <a-statistic class="static-num" title="成品计划" :value="364211" />
            </a-col>
            <a-col :span="6">
                <a-statistic class="static-num" title="正在发布" :value="112893" />
            </a-col>
            <a-col :span="6">
                <a-statistic class="static-num" title="待发布" :value="99">
                    <template #suffix>
                        <span>/ 100</span>
                    </template>
                </a-statistic>
            </a-col>
            <a-col :span="6">
                <a-statistic-countdown
                    title="运行时间"
                    :value="deadLine"
                    style="margin-right: 50px"
                    @finish="onFinish"
                />
            </a-col>
        </a-row>
    </div>
</template>

<style lang="scss">
@import url("./index.scss");
</style>

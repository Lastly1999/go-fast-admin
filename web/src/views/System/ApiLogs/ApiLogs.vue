<script lang="ts" setup>
import {onMounted, reactive, ref, inject} from "vue"
import QueryGroup from "@/components/QueryGroup/QueryGroup.vue"
import FContainer from "@/components/FContainer/FContainer.vue"
import FTable from "@/components/FTable/FTable.vue"

import type {QueryJsonItem} from "@/components/QueryGroup/QueryGroup.vue"
import type {Moment} from 'moment'


// apis
import {getSystemLogs} from "@/services/system/sys"
import type {ApiLogsParams} from "@/services/model/response/role";

export type apiLogsItem = {
    id: number;
    logBody: string;
    logMethod: string;
    logPath: string;
    logTime: string;
    logRequestStatus: number;
    LogCreateTime: string;
}

onMounted(() => {
    getLogs()
})

const getLogs = async () => {
    apiLogsTableLoading.value = true
    const {data, code} = await getSystemLogs(apiLogsTableParams.value)
    if (code === 200) {
        apiLogsData.value = data.list
        apiLogsToal.value = data.total
    }
    apiLogsTableLoading.value = false
}

const apiLogsTableLoading = ref(false)
const apiLogsTableColumns = [
    {
        title: '日志序号',
        dataIndex: 'id',
        key: 'id',
        width: "10%"
    },
    {
        title: '请求信息',
        dataIndex: 'logBody',
        key: 'logBody',
        width: "13%"
    },
    {
        title: "请求方式",
        dataIndex: "logMethod",
        key: "logMethod",
        width: "14%",
        slots: {customRender: "tags"},
    },
    {
        title: 'api路径',
        dataIndex: 'logPath',
        key: 'logPath',
        width: "10%"
    },
    {
        title: '请求耗时',
        dataIndex: 'logTime',
        key: 'logTime',
        width: "13%"
    },
    {
        title: '请求创建时间',
        dataIndex: 'LogCreateTime',
        key: 'LogCreateTime',
        width: "13%"
    },
    {
        title: '请求状态',
        dataIndex: 'logRequestStatus',
        key: 'logRequestStatus',
        width: "13%"
    }
]
const apiLogsTableParams = ref<ApiLogsParams>({
    timer: [],
    page: 1,
    pageSize: 10
})
const apiLogsToal = ref(0)
const apiLogsData = ref<apiLogsItem[]>([])

const paginationChange = () => {
    getLogs()
}

// query 重置方法
const reset = () => {

}

// query 查询方法
const search = () => {
    console.log(queryForm.value)
}
// 筛选时间选择器ok
const timerOK = (time: Moment[]) => {
    apiLogsTableParams.value.timer[0] = time[0].format('YYYY-MM-DD HH:MM:SS')
    apiLogsTableParams.value.timer[1] = time[1].format('YYYY-MM-DD HH:MM:SS')
    getLogs()
}

// 筛选时间change
const timerChange = (value: Moment[], dateString: string[]) => {
    apiLogsTableParams.value.timer = dateString
    getLogs()
}

// 查询组件表单
const queryForm = ref({})
// 查询组件渲染数据源
const queryJsonData = reactive<QueryJsonItem[]>([
    {
        type: "rangePicker",
        placeholder: ["开始时间 ", "结束时间"],
        prop: "timer",
        ok: timerOK,
        change: timerChange
    },
    {
        type: "button",
        buttonType: "primary",
        text: "查询",
        fun: search
    },
    {
        type: "button",
        buttonType: "default",
        text: "重置",
        fun: reset
    }
])

</script>
<template>
    <FContainer>
        <template v-slot:header>
            <QueryGroup v-model:jsonData="queryJsonData" :form="apiLogsTableParams"/>
        </template>
        <template v-slot:main>
            <FTable bordered size="middle" :loading="apiLogsTableLoading" :columns="apiLogsTableColumns"
                    :data-source="apiLogsData"
                    pagination
                    rowKey="id">
                <template #tags="{ data }">
                    <a-tag v-if="data === 'GET'" color="green">GET</a-tag>
                    <a-tag v-if="data === 'POST'" color="warning">POST</a-tag>
                    <a-tag v-if="data === 'PUT'" color="processing">PUT</a-tag>
                    <a-tag v-if="data === 'DELETE'" color="error">DELETE</a-tag>
                </template>
            </FTable>
            <a-pagination
                class="system-pagination"
                v-model:current="apiLogsTableParams.page"
                v-model:pageSize="apiLogsTableParams.pageSize"
                show-size-changer
                :total="apiLogsToal"
                @change="paginationChange"
                @showSizeChange="paginationChange"
            />
        </template>
    </FContainer>
</template>
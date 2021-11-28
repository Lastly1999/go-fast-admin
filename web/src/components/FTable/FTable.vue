<script lang="ts" setup>
import { onMounted, ref, nextTick } from "vue"

defineProps({
    pagination: {
        type: Boolean,
        default: () => false
    }
})

onMounted(() => {
    nextTick(() => {
        tableHeight.value = document.getElementsByClassName('page-content')[0].clientHeight - document.getElementsByClassName('ant-table-thead')[0].clientHeight
    })
    window.onresize = () => {
        tableHeight.value = document.getElementsByClassName('page-content')[0].clientHeight - document.getElementsByClassName('ant-table-thead')[0].clientHeight
    }
})

const tableHeight = ref(0)


const rowClassName = (record: any, index: number) => {
    return index % 2 === 1 ? 'table-striped' : null
}

</script>
<template>
    <a-table
        style="word-break: break-all;"
        class="ant-table-striped"
        {...$props}
        :pagination="false"
        :scroll="{ x: 'max-content', y: tableHeight - (pagination ? 50 : 10) }"
        :rowClassName="rowClassName"
    >
        <template #action="{ text }">
            <slot name="action" :data="text"></slot>
        </template>
        <template #tags="{ text }">
            <slot name="tags" :data="text"></slot>
        </template>
        <template #icon="{ text }">
            <slot name="icon" :data="text"></slot>
        </template>
    </a-table>
</template>

<style scoped>
.ant-table-striped :deep(.table-striped) td {
    background-color: #fafafa;
}
</style>
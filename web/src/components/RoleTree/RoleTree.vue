<script lang="ts" setup>
import type { PropType } from "vue"
import type { TreeDataItem } from 'ant-design-vue/es/tree/Tree'

export type TreeChecked = {
	checked: number[];
	halfChecked: number;
}

const props = defineProps({
	data: {
		type: Array as PropType<TreeDataItem>,
		default: () => []
	},
	checkedKeys: {
		type: Array as PropType<number[]>,
		default: () => []
	}
})

const emit = defineEmits<{
	(event: "check", data: TreeChecked): void,
}>()

const replaceFields = {
	children: 'children',
	title: 'label',
	key: 'id'
};

const treeSelect = (checkedKeys: TreeChecked) => {
	// 已勾选子节点以及半勾选状态的父节点
	emit("check", checkedKeys)
}
</script>

<template>
	<a-tree checkable checkStrictly defaultExpandAll :replaceFields="replaceFields" :tree-data="data" v-model:checkedKeys="checkedKeys" @check="treeSelect"></a-tree>
</template>
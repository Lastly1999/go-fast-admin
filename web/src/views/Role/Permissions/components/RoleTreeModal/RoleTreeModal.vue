<script lang="ts" setup>
import FModal from "@/components/FModal/FModal.vue"
import RoleTree from "@/components/RoleTree/RoleTree.vue"

import type { PropType } from "vue"
import type { TreeChecked } from "@/components/RoleTree/RoleTree.vue"

defineProps({
	value: {
		type: Boolean,
		default: () => false
	},
	data: {
		type: Array,
		default: () => []
	},
	checkedKeys: {
		type: Array as PropType<number[]>,
		default: () => []
	},
	confirmLoading: {
		type: Boolean,
		default: () => false
	}
})


const emit = defineEmits<{
	(event: "update:value", show: boolean): void,
	(event: "check", data: TreeChecked): void,
	(event: "submit"): void
}>()


const close = () => {
	emit("update:value", false)
}

const okSubmit = () => {
	emit("submit")
}

const treeCheck = (checkedKeys: TreeChecked) => {
	emit("check", checkedKeys)
}
</script>																								  																																																																																															

<template>
	<FModal
		v-model:value="value"
		title="权限树"
		:confirmLoading="confirmLoading"
		@close="close"
		@ok="okSubmit"
	>
		<RoleTree v-model:checkedKeys="checkedKeys" :data="data" @check="treeCheck" />
	</FModal>
</template>
<script lang="ts" setup>
import {ref, computed} from "vue"
import {useStore} from "vuex"
import {Icon} from "@/components/FIcon/FIcon.ts"

export type IconsItem = {
    id: number;
    iconName: string;
}

const props = defineProps({
    value: {
        type: [String, Number]
    },
    placeholder: {
        type: String,
        default: () => ""
    }
})

const store = useStore()

// 图标列表
const options = computed<IconsItem[]>(() => store.getters["systemModule/getSysIcons"])

const emit = defineEmits<{
    (event: "change"): void,
    (event: 'update:value', e: string): void
}>()

const handleChange = (event: string) => {
    emit("change")
    emit('update:value', event)
}

const filterOption = (input: string, option: any) => {
    return option.value.toLowerCase().indexOf(input.toLowerCase()) >= 0;
};

</script>
<template>
    <a-select v-model:value="value" show-search style="width: 100%" placeholder="请选择" @change="handleChange"
              :filter-option="filterOption">
        <a-select-option v-for="item in options" :value="item.iconName">
            <span>
                <Icon :icon="item.iconName"></Icon>
                {{ item.iconName }}
            </span>
        </a-select-option>
    </a-select>
</template>
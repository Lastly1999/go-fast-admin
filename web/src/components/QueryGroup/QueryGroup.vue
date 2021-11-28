<script lang="ts" setup>
import type {PropType} from "vue"


export type JsonCompType = "input" | "button" | "rangePicker"

export type QueryJsonItem = {
    type?: JsonCompType;
    label?: string;
    placeholder?: string;
    required?: boolean;
    prop?: any;
} & JsonButton &
    JsonInput &
    JsonRangePicker;

export type JsonInput = {};

export type JsonButton = {
    buttonType?: string;
    text?: string;
    fun?: (data?: any) => void;
    change?: (dateString: string[]) => void;
    ok?: (data: any) => void;
};

export type RangePickerSize = "size" | "default" | "small"

export type JsonRangePicker = {
    size?: RangePickerSize,
    format?: string;
}

defineProps({
    jsonData: {
        type: Array as PropType<QueryJsonItem[]>,
        default: () => [],
    },
    form: {
        type: Object as any,
        default: () => {
        }
    }
});

const formartTime = (val: string, dateStr: string[], jsonItem: QueryJsonItem) => {
    jsonItem.change && jsonItem.change(dateStr)
}

</script>
<template>
    <div class="permissions-fillter">
        <a-form layout="inline" :modal="form">
            <a-form-item v-for="jsonItem in jsonData" :label="jsonItem.label" :required="jsonItem.required">
                <!-- 输入框 -->
                <a-input v-if="jsonItem.type === 'input'" v-model:value="form[jsonItem.prop]"
                         :placeholder="jsonItem.placeholder"/>
                <!-- 按钮 -->
                <a-button v-if="jsonItem.type === 'button'" :type="jsonItem.buttonType"
                          @click.prevent="() => jsonItem.fun && jsonItem.fun()">{{ jsonItem.text }}
                </a-button>
                <!-- 时间选择器 -->
                <a-range-picker
                    v-if="jsonItem.type === 'rangePicker'"
                    v-model:value="form[jsonItem.prop]"
                    :placeholder="jsonItem.placeholder"
                    :size="jsonItem.size"
                    showToday
                    format="YYYY-MM-DD HH:mm:ss"
                    @change="(value, dateString) => formartTime(value,dateString,jsonItem)"
                    @ok="(value) => jsonItem.ok && jsonItem.ok(value)"
                />
            </a-form-item>
        </a-form>
    </div>
</template>

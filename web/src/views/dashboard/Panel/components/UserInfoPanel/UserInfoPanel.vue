<script lang="ts" setup>
import {UserOutlined} from '@ant-design/icons-vue'

</script>
<template>
    <div class="user-info-panel">
        <div>
            <a-avatar :size="64">
                <template #icon>
                    <UserOutlined/>
                </template>
            </a-avatar>
        </div>
    </div>
</template>
<style scoped>
.user-info-panel {
    border-radius: 6px;
    background: #fff;
    display: flex;
    padding: 15px;
    box-shadow: 0 1px 3px rgb(18 18 18 / 10%);
}
</style>
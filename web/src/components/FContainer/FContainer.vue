<script lang="ts" setup></script>
<template>
    <a-layout style="height: 100%; overflow: hidden">
        <!-- headers -->
        <div class="page-content-header">
            <slot name="header"></slot>
        </div>
        <a-layout-content class="page-content system-box-shadow">
            <!-- main -->
            <div class="page-content-container">
                <slot name="main"></slot>
            </div>
        </a-layout-content>
    </a-layout>
</template>
<style>
:deep().ant-table-tbody > tr > td {
    white-space: break-spaces;
}

.page-content {
    padding: 10px 10px 10px 10px;
    position: relative;
    border-radius: 5px;
    box-shadow: 0 1px 3px rgb(18 18 18 / 10%);
}

.page-content-container {
    padding: 20px;
    height: 100%;
    background-color: #fff;
}

.page-content-header{
    background-color: #fff;
    padding: 15px 30px;
}
</style>
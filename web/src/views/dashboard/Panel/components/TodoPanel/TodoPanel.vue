<script lang="ts" setup>
import { ref } from 'vue'

interface DataItem {
  title: string;
}

const data: DataItem[] = [
  {
    title: 'Ant Design Title 1',
  },
  {
    title: 'Ant Design Title 2',
  },
  {
    title: 'Ant Design Title 3',
  },
  {
    title: 'Ant Design Title 4',
  },
];

const activeKey = ref<string>('1')
const tabPosition = ref('top')

</script>

<template>
  <a-tabs class="todo-panel" v-model:activeKey="activeKey" :tab-position="tabPosition">
    <a-tab-pane key="1" tab="待处理消息">
      <a-list item-layout="horizontal" :data-source="data">
        <template #renderItem="{ item }">
          <a-list-item class="todo-container">
            <a-list-item-meta
              description="Ant Design, a design language for background applications, is refined by Ant UED Team"
            >
              <template #title>
                <a href="https://www.antdv.com/">{{ item.title }}</a>
              </template>
              <template #avatar>
                <a-avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
              </template>
            </a-list-item-meta>
          </a-list-item>
        </template>
      </a-list>
    </a-tab-pane>
    <a-tab-pane key="2" tab="我的审批">Content of Tab 2</a-tab-pane>
    <a-tab-pane key="3" tab="我的发布">Content of Tab 3</a-tab-pane>
  </a-tabs>
</template>

<style lang="scss">
.todo-panel {
  margin-top: 10px;
  background-color: #fff;
  box-shadow: 0 1px 3px rgb(18 18 18 / 10%);
  .todo-container {
    padding: 10px;
  }
}
</style>



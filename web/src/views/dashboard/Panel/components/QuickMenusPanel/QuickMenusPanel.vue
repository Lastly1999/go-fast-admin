<script lang="ts" setup>
import { reactive } from "vue"

export type CardDataItem = {
  name: string;
  img: string;
}

const cardData = reactive<CardDataItem[]>([
  {
    name: "我的任务",
    img: "/src/assets/imgs/default.png"
  },
  {
    name: "我的消息",
    img: "/src/assets/imgs/human.png"
  },
  {
    name: "待办事项",
    img: "/src/assets/imgs/positon.png"
  },
  {
    name: "用户管理",
    img: "/src/assets/imgs/work.png"
  }
])

</script>
<template>
  <div style="margin-top:10px">
    <a-card title="快捷入口">
      <a-card-grid v-for="item in cardData" style="width: 25%; text-align: center">
        <div>
          <img width="30" :src="item.img" />
        </div>
        <span>{{ item.name }}</span>
      </a-card-grid>
    </a-card>
  </div>
</template>
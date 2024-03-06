<script setup>
import {onMounted, ref} from 'vue'

const title = ref('商品管理系统')
const time = ref('')

onMounted(() => {
  updateCurrentTime();
  setInterval(updateCurrentTime, 1000);
})

// 更新和格式化当前时间的方法
function updateCurrentTime() {
  const date = new Date();
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hour = String(date.getHours()).padStart(2, '0');
  const minute = String(date.getMinutes()).padStart(2, '0');
  const second = String(date.getSeconds()).padStart(2, '0');

  time.value = `${year}-${month}-${day} ${hour}:${minute}:${second}`;
}

</script>

<template>
    <el-row class="head">
      <el-col :span="16">{{ title }}</el-col>
      <el-col :span="8" class="time">{{ time }}</el-col>
    </el-row>
</template>

<style scoped>

.time {
  text-align: right;
}
</style>
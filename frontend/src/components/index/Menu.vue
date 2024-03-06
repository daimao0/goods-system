<script setup>

import {Box, List, Refrigerator,User} from "@element-plus/icons-vue";
import {onMounted} from "vue";
import {useRouter} from "vue-router";

let router;

onMounted(() => {
  // 在组件挂载后获取路由实例
  router = useRouter();
});

function menuClick(index) {
  if (typeof index === 'string') {
    router.push(index);
  }
}
</script>

<template>

  <el-menu
      default-active="1"
      class="el-menu-vertical-demo"
      @select="menuClick"
  >
    <el-menu-item index="/">
      <el-icon>
        <List/>
      </el-icon>
      <span>商品列表</span>
    </el-menu-item>
    <el-menu-item index="/order">
      <el-icon>
        <Box/>
      </el-icon>
      <span>商品下单</span>
    </el-menu-item>
    <el-menu-item index="/member">
      <el-icon>
        <User/>
      </el-icon>
      <span>会员中心</span>
    </el-menu-item>
    <el-menu-item index="/orderRecord">
      <el-icon>
        <Refrigerator/>
      </el-icon>
      <span>账单记录</span>
    </el-menu-item>
    <el-menu-item index="/system">
      <el-icon>
        <Refrigerator/>
      </el-icon>
      <span>系统管理</span>
    </el-menu-item>
  </el-menu>
</template>

<style scoped>

</style>
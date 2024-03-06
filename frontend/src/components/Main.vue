<script setup>
import Menu from "./index/Menu.vue";
import Head from "./index/Head.vue";

</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <Head/>
      </el-header>
      <el-container>
        <el-aside width="200px">
          <Menu/>
        </el-aside>
        <el-main>
          <router-view/>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<style scoped>
.time {
  text-align: right;
}
</style>
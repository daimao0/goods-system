<script setup>
import {onMounted, ref} from "vue";
import {GoodsTypeSelectAll} from "../../../wailsjs/go/main/App.js";
import {Delete} from "@element-plus/icons-vue";
import GoodsTypeEditDialog from "./GoodsTypeEditDialog.vue";
import GoodsTypeDeleteDialog from "./GoodsTypeDeleteDialog.vue";

//商品类型
const goodsTypes = ref([])
const goodsTypeName = ref()
const goodsTypesEditDialog = ref(false)
const goodsTypesDeleteDialog = ref(false)
//监听参数
onMounted(async () => {
  await listGoodsTypes();

})

//打开编辑对话框
function openEditDialog() {
  goodsTypesEditDialog.value = true
}

//关闭编辑对话框
function closeEditDialog() {
  goodsTypesEditDialog.value = false
}

//刷新商品类型
function listGoodsTypes() {
  GoodsTypeSelectAll().then(resp => [
    goodsTypes.value = resp.data
  ]).catch(err => {
    console.log(err)
  })
}

//打开删除商品类型对话框
function openDeleteGoodsType(row) {
  goodsTypesDeleteDialog.value = true
  goodsTypeName.value = row.name
}

//打开删除商品类型对话框
function closeDeleteGoodsType() {
  goodsTypesDeleteDialog.value = false
}
</script>

<template>
  <el-button type="success" @click="openEditDialog">新增类型</el-button>
  <el-table :data="goodsTypes" style="width: 100%">
    <el-table-column prop="name" label="商品类型"/>
    <el-table-column label="操作">
      <template #default="scope">
        <el-button type="danger" :icon="Delete" circle @click="openDeleteGoodsType(scope.row)"/>
      </template>
    </el-table-column>
  </el-table>

  <!--  商品类型编辑对话框   -->
  <el-dialog
      v-model="goodsTypesEditDialog"
      title="编辑类型"
      :before-close="listGoodsTypes"
  >
    <GoodsTypeEditDialog
        @close="closeEditDialog"
        @refresh="listGoodsTypes"
    />
  </el-dialog>

  <!--  商品类型删除对话框-->
  <el-dialog
      v-model="goodsTypesDeleteDialog"
      title="删除类型"
  >
    <GoodsTypeDeleteDialog
        :name="goodsTypeName"
        @close="closeDeleteGoodsType"
        @refresh="listGoodsTypes"
    />
  </el-dialog>
</template>

<style scoped>

</style>
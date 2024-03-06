<script setup>
import {onMounted, ref} from "vue";
import {GoodsPage} from "../../../wailsjs/go/main/App.js";
import GoodsEdit from "./GoodsEditDialog.vue";
import {ElMessageBox} from "element-plus";
import {Delete, Edit} from "@element-plus/icons-vue";
import GoodsDeleteDialog from "./GoodsDeleteDialog.vue";
import {errElMessage} from "../../utils/el-message-utils.js";
import {getGoodsTypeSelect} from "../../api/select.js";

//对话框可显示
const dialogVisible = ref(false)
const deleteDialogVisible = ref(false)
//表格分页响应结果
const tablePageVOList = ref([])
const pageTotal = ref()
const goodsTypeOpts = ref([])
//表格请求
const goodsPageRequest = ref({
  name: '',
  goodsType: '',
  page: 1,
  size: 10
})

//props传参
const currentGoodsIdToDelete = ref(null)
const currentGoodsVO = ref(null)
//页面加载后执行
onMounted(async () => {
  try {
    await onSubmit()
    goodsTypeOpts.value = await getGoodsTypeSelect()
  } catch (err) {
    console.error("err")
  }
})

// 查询按钮点击事件处理函数
async function onSubmit() {
  try {
    const resp = await GoodsPage(goodsPageRequest.value);
    if (resp.data == null || resp.data.data == null) {
      tablePageVOList.value = []
    } else {
      tablePageVOList.value = resp.data.data;
      pageTotal.value = resp.data.total
    }
  } catch (err) {
    errElMessage(err)
  }
}

// 添加分页切换事件处理函数（假设你已经实现了分页逻辑）
function handleCurrentChange(pageNumber) {
  goodsPageRequest.value.page = pageNumber
  onSubmit()
}

//编辑商品
function addGoods() {
  dialogVisible.value = true
  currentGoodsVO.value = null
}

//编辑商品
function editGoods(row) {
  dialogVisible.value = true
  currentGoodsVO.value = row
}

//删除商品
function deleteGoods(id) {
  deleteDialogVisible.value = true
  currentGoodsIdToDelete.value = id
}

// 关闭删除对话框
function closeDeleteDialog() {
  deleteDialogVisible.value = false
}

//关闭编辑对话框
function closeDialog() {
  dialogVisible.value = false
}

//关闭处理
function handleClose() {
  ElMessageBox.confirm('是否关闭?')
      .then(() => {
        onSubmit()
      })
      .catch(() => {
        // catch error
      }).finally(() => {
    dialogVisible.value = false
    deleteDialogVisible.value = false
  })
}
</script>

<template>
  <el-form :inline="true" :model="goodsPageRequest">
    <el-form-item label="商品名称">
      <el-input v-model="goodsPageRequest.name" placeholder="请输入商品名称" class="form-input" clearable/>
    </el-form-item>
    <el-form-item label="商品类型">
      <el-select
          v-model="goodsPageRequest.goodsType"
          placeholder="请选择类型"
          clearable
          class="form-select"
      >
        <el-option v-for="item in goodsTypeOpts" :key="item.id" :label="item.name" :value="item.name"/>
      </el-select>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">查询</el-button>
    </el-form-item>
    <el-form-item>
      <el-button type="success" @click="addGoods">新增</el-button>
    </el-form-item>
  </el-form>
  <el-table :data="tablePageVOList" stripe style="width: 100%">
    <el-table-column prop="id" label="id"/>
    <el-table-column prop="name" label="商品名称"/>
    <el-table-column prop="goodsType" label="类型"/>
    <el-table-column prop="price" label="价格"/>
    <el-table-column prop="count" label="库存量"/>
    <el-table-column prop="createTime" label="创建时间"/>
    <el-table-column label="操作">
      <template #default="scope">
        <el-button type="primary" :icon="Edit" circle @click="editGoods(scope.row)"/>
        <el-button type="danger" :icon="Delete" circle @click="deleteGoods(scope.row.id)"/>
      </template>
    </el-table-column>
  </el-table>
  <!-- 如果需要根据tableData的实际长度来控制分页组件 -->
  <div class="example-pagination-block" v-if="tablePageVOList.length > 0">
    <el-pagination
        layout="prev, pager, next"
        :total="pageTotal"
        @current-change="handleCurrentChange"
    />
  </div>
  <!--  商品编辑对话框  -->
  <el-dialog
      v-model="dialogVisible"
      title="编辑商品"
      :before-close="handleClose"
  >
    <GoodsEdit
        :goods-v-o="currentGoodsVO"
        @refresh="onSubmit"
        @close="closeDialog"
    />
  </el-dialog>
  <!--  商品删除对话框  -->
  <el-dialog
      v-model="deleteDialogVisible"
      title="删除商品"
      :before-close="handleClose"
  >
    <GoodsDeleteDialog
        :goods-id="currentGoodsIdToDelete"
        @refresh="onSubmit"
        @close="closeDeleteDialog"
    />
  </el-dialog>
</template>

<style scoped>
.form-select {
  --el-select-width: 180px;
}

.form-input {
  --el-select-width: 120px;
}
</style>
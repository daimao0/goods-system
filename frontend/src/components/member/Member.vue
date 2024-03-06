<script setup>
import {onMounted, ref} from "vue";
import {MemberPage} from "../../../wailsjs/go/main/App.js";
import {ElMessageBox} from "element-plus";
import {Delete, Edit} from "@element-plus/icons-vue";
import MemberEditDialog from "./MemberEditDialog.vue";
import MemberDeleteDialog from "./MemberDeleteDialog.vue";
import {getMemberLevelSelect} from "../../api/select.js";
//表格请求
const memberPageRequest = ref({
  name: '',
  phone: '',
  levelId: '',
  account: '',
  page: 1,
  size: 10
})

const memberPageVO = ref([])
const pageTotal = ref(0)
const memberVO = ref(null)
const memberLevelSelect = ref([])
//会员编辑对话框
const memberEditDialogVisible = ref(false)
const memberDeleteDialogVisible = ref(false)

//页面加载后执行
onMounted(async () => {
  searchPageMember()
  memberLevelSelect.value = await getMemberLevelSelect()
})

//搜索会员
function searchPageMember() {
  MemberPage(memberPageRequest.value).then(resp => {
    memberPageVO.value = resp.data.data
  })
}

//打开编辑对话框
function openEditDialog(row) {
  memberEditDialogVisible.value = true
  if (row) {
    memberVO.value = row
  } else {
    memberVO.value = null
  }
}

// 添加分页切换事件处理函数（假设你已经实现了分页逻辑）
function handleCurrentChange(pageNumber) {
  memberPageRequest.value.page = pageNumber
  searchPageMember()
}

//删除编辑对话框
function closeEditDialog() {
  memberEditDialogVisible.value = false
  searchPageMember()
}

//打开删除对话框
function openDeleteDialog(row) {
  memberDeleteDialogVisible.value = true
  memberVO.value = row
}

//删除对话框
function closeDeleteDialog() {
  memberDeleteDialogVisible.value = false
  searchPageMember()
}

//关闭会员编辑对话框
function memberEditDialogClose() {
  ElMessageBox.confirm('是否关闭?')
      .then(() => {
        //重新搜索会员
        searchPageMember()
      })
      .catch(() => {
        // catch error
      }).finally(() => {
    memberEditDialogVisible.value = false
  })
}
</script>

<template>
  <!--  查询表单  -->
  <el-form :inline="true" :model="memberPageRequest">
    <el-form-item label="会员名称">
      <el-input v-model="memberPageRequest.name" placeholder="请输入会员名称" class="form-input" clearable/>
    </el-form-item>
    <el-form-item label="会员等级">
      <el-select
          v-model="memberPageRequest.levelId"
          placeholder="请选择等级"
          clearable
          class="form-select"
          style="width: 120px"
      >
        <el-option v-for="item in memberLevelSelect" :key="item.id" :label="item.name" :value="item.id"/>
      </el-select>
    </el-form-item>
    <el-form-item label="手机号">
      <el-input v-model="memberPageRequest.phone" placeholder="请输入手机号" clearable/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="searchPageMember">查询</el-button>
    </el-form-item>
    <el-form-item>
      <el-button type="success" @click="openEditDialog(null)">新增</el-button>
    </el-form-item>
  </el-form>

  <!--查询表格-->
  <el-table :data="memberPageVO" stripe style="width: 100%">
    <el-table-column prop="id" label="id"/>
    <el-table-column prop="name" label="会员名称"/>
    <el-table-column prop="level.name" label="会员等级"/>
    <el-table-column prop="level.discount" label="折扣"/>
    <el-table-column prop="phone" label="手机号"/>
    <el-table-column prop="account" label="余额"/>
    <el-table-column prop="createTime" label="创建时间"/>
    <el-table-column label="操作">
      <template #default="scope">
        <el-button type="primary" :icon="Edit" circle @click="openEditDialog(scope.row)"/>
        <el-button type="danger" :icon="Delete" circle @click="openDeleteDialog(scope.row)"/>
      </template>
    </el-table-column>
  </el-table>
  <!-- 如果需要根据的实际长度来控制分页组件 -->
  <el-pagination
      layout="prev, pager, next"
      :total="pageTotal"
      @current-change="handleCurrentChange"
  />
  <!-- 会员编辑对话框 -->
  <el-dialog
      v-model="memberEditDialogVisible"
      title="会员编辑"
      :before-close="memberEditDialogClose"
  >
    <member-edit-dialog
        :member-v-o="memberVO"
        @close="closeEditDialog"
    />
  </el-dialog>
  <!-- 会员删除对话框 -->
  <el-dialog
      v-model="memberDeleteDialogVisible"
      title="删除会员"
  >
    <member-delete-dialog
        :member-v-o="memberVO"
        @close="closeDeleteDialog"
    />
  </el-dialog>
</template>

<style scoped>
</style>
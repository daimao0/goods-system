<script setup>

import {Delete, Edit} from "@element-plus/icons-vue";
import {onMounted, ref} from "vue";
import {MemberLevelSelectAll} from "../../../wailsjs/go/main/App.js";
import MemberLevelEditDialog from "./MemberLevelEditDialog.vue";
import MemberLevelDeleteDialog from "./MemberLevelDeleteDialog.vue";

const memberLevels = ref([])

//页面加载后执行
onMounted(() => {
  searchAllMemberLevel()
})
const memberLevelVO = ref(null)
//对话框
const memberLevelEditDialogVisible = ref(false)
const memberLevelDeleteDialogVisible = ref(false)

//搜索会员等级
function searchAllMemberLevel() {
  MemberLevelSelectAll().then(resp => {
    memberLevels.value = resp.data
  })
}

//打开编辑对话框
function openEditDialog(row) {
  if (row != null) {
    memberLevelVO.value = row
  } else {
    memberLevelVO.value = null
  }
  memberLevelEditDialogVisible.value = true
}

//打开删除对话框
function openDeleteDialog(row) {
  memberLevelVO.value = row
  memberLevelDeleteDialogVisible.value = true
}

//关闭编辑对话框
function closeEditDialog() {
  memberLevelEditDialogVisible.value = false
}

//关闭删除对话框
function closeDeleteDialog() {
  memberLevelDeleteDialogVisible.value = false
  searchAllMemberLevel()
}
</script>

<template>
  <el-button type="success" @click="openEditDialog(null)">新增类型</el-button>
  <el-table :data="memberLevels" style="width: 100%">
    <el-table-column prop="name" label="会员等级"/>
    <el-table-column prop="discount" label="会员折扣"/>
    <el-table-column label="操作">
      <template #default="scope">
        <el-button type="primary" :icon="Edit" circle @click="openEditDialog(scope.row)"/>
        <el-button type="danger" :icon="Delete" circle @click="openDeleteDialog(scope.row)"/>
      </template>
    </el-table-column>
  </el-table>

  <!--  会员等级编辑对话框   -->
  <el-dialog
      v-model="memberLevelEditDialogVisible"
      title="编辑等级"
  >
    <member-level-edit-dialog
        :member-level-v-o="memberLevelVO"
        @close="closeEditDialog"
        @refresh="searchAllMemberLevel"
    />
  </el-dialog>
  <!--  会员等级删除对话框   -->
  <el-dialog
      v-model="memberLevelDeleteDialogVisible"
      title="删除等级"
  >
    <member-level-delete-dialog
        :member-level-v-o="memberLevelVO"
        @close="closeDeleteDialog"
        @refresh="searchAllMemberLevel"
    />
  </el-dialog>
</template>

<style scoped>

</style>
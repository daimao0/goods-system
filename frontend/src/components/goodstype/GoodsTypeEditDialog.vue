<script setup>
import {defineEmits, ref} from "vue";
import {GoodsTypeSave} from "../../../wailsjs/go/main/App.js";
import {commonElMessage, errElMessage} from "../../utils/el-message-utils.js";

const emit = defineEmits(['refresh', 'close']);
//商品类型名称参数
const goodsTypeName = ref()

//提交表单
function onSubmit() {
  let goodsTypeSaveRequest = {
    name: goodsTypeName.value
  }
  GoodsTypeSave(goodsTypeSaveRequest).then(resp => {
    commonElMessage(resp)
  }).catch(err => {
    errElMessage(err)
  }).finally(() => {
    closeGoodsTypeDialog()
  })
}

//关闭商品类型对话框
function closeGoodsTypeDialog() {
  emit('close')
  emit('refresh')
}
</script>

<template>
  <el-form>
    <el-form-item label="类型名称">
      <el-input v-model="goodsTypeName"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">确认</el-button>
      <el-button @click="closeGoodsTypeDialog">取消</el-button>
    </el-form-item>
  </el-form>
</template>

<style scoped>

</style>
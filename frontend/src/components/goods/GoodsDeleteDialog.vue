<script setup>
import {defineEmits} from 'vue';
import {GoodsDelete} from "../../../wailsjs/go/main/App.js";
import {ElMessage} from "element-plus";
// 定义你希望触发的自定义事件
const emit = defineEmits(['refresh', 'close']);

//删除商品
function deleteGoods() {
  GoodsDelete(props.goodsId).then(() => {
    ElMessage.success('删除成功')
  }).catch(err => {
    ElMessage.error('出现异常' + err)
  })
  //关闭对话框
  emit('close')
  //刷新分页
  emit('refresh')
}

//关闭对话框
function close() {
  //关闭对话框
  emit('close')

}

//接收外部传入的商品id
const props = defineProps({
  goodsId: {
    type: String,
    required: true
  }
})
</script>

<template>

  <span>提示，删除后不可恢复</span>
  <div class="dialog-footer">
    <el-button @click="close">取消</el-button>
    <el-button type="primary" @click="deleteGoods ">
      确认
    </el-button>
  </div>
</template>

<style scoped>

</style>
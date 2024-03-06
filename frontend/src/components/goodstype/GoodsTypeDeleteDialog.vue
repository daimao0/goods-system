<script setup>
import {defineEmits} from 'vue';
import {GoodsTypeDeleteByName} from "../../../wailsjs/go/main/App.js";
import {ElMessage} from "element-plus";

const emit = defineEmits(['close','refresh'])
//接收外部传入的商品id
const props = defineProps({
  name: {
    type: String,
    required: true
  }
})

function deleteGoodsType() {
  let request={
    name:props.name
  }
  GoodsTypeDeleteByName(request).then(resp => {
    if (resp.code === 200) {
      ElMessage.success('删除成功')
      emit("refresh")
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(err => {
    ElMessage.error(err)
  }).finally(() => {
    closeDialog()
  })
}

function closeDialog() {
  emit('close')
}

</script>
<template>
  <span>提示，删除后不可恢复</span>
  <div class="dialog-footer">
    <el-button @click="closeDialog">取消</el-button>
    <el-button type="primary" @click="deleteGoodsType">
      确认
    </el-button>
  </div>
</template>

<style scoped>

</style>
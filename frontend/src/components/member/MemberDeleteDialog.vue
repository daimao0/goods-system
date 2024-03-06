<script setup>
import {defineEmits} from "vue";
import {MemberDeleteById} from "../../../wailsjs/go/main/App.js";
import {commonElMessage, errElMessage} from "../../utils/el-message-utils.js";

const emit = defineEmits(['close']);
//props接收参数
const props = defineProps({
  memberVO: {
    type: [Object, null],
    required: true
  }
})

//删除会员
function deleteMember() {
  MemberDeleteById(props.memberVO.id).then(resp => {
    commonElMessage(resp)
  }).catch(err => {
    errElMessage(err)
  })
  close()
}

//关闭对话框
function close() {
  emit('close')
}
</script>

<template>

  <span>提示，删除后不可恢复</span>
  <div class="dialog-footer">
    <el-button @click="close">取消</el-button>
    <el-button type="primary" @click="deleteMember">
      确认
    </el-button>
  </div>
</template>

<style scoped>

</style>
<script setup>
import {defineEmits, reactive, ref, watchEffect} from "vue";
import {MemberLevelSave, MemberLevelUpdate} from "../../../wailsjs/go/main/App.js";
import {commonElMessage, errElMessage} from "../../utils/el-message-utils.js";

const emits = defineEmits(['close', 'refresh'])
const props = defineProps({
  memberLevelVO:{
    type: [Object, null],
    required: true
  }
})


const isEdit = ref(false)

const form = ref()

watchEffect(()=>{
  if (props.memberLevelVO){
    form.value = reactive(props.memberLevelVO)
    isEdit.value = true
  }else{
    form.value = {
      name:'',
      discount:''
    }
    isEdit.value = false
  }
})
//提交会员等级
function onSubmit() {
  if (isEdit.value){
    MemberLevelUpdate(form.value).then(resp=>{
      commonElMessage(resp)
    }).catch(err=>{
      errElMessage(err)
    })
  }else{
    MemberLevelSave(form.value).then(resp => {
      commonElMessage(resp)
    }).catch(err=>{
      errElMessage(err)
    })
    closeMemberLevelDialog()
    emits('refresh')
  }

}

//关闭会员等级编辑对话框
function closeMemberLevelDialog() {
  emits('close')
}
</script>

<template>
  <el-form :model="form">
    <el-form-item label="会员等级">
      <el-input v-model="form.name"/>
    </el-form-item>
    <el-form-item label="折扣">
      <el-input v-model="form.discount"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">确认</el-button>
      <el-button @click="closeMemberLevelDialog">取消</el-button>
    </el-form-item>
  </el-form>
</template>

<style scoped>

</style>
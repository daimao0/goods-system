<script setup>
import {defineEmits, onMounted, reactive, ref, watchEffect} from "vue";
import {MemberSave, MemberUpdate} from "../../../wailsjs/go/main/App.js";
import {commonElMessage, errElMessage} from "../../utils/el-message-utils.js";
import MemberLevelDialog from "../memberlevel/MemberLevelDialog.vue";
import {getMemberLevelSelect} from "../../api/select.js";

const emit = defineEmits(['close']);
//props接收参数
const props = defineProps({
  memberVO: {
    type: [Object, null],
    required: true
  }
})
const isEdit = ref(false)
const form = ref({})
const memberLevelSelect = ref([])
//对话框
const memberLevelDialogVisible = ref(false)

onMounted(async () => {
  memberLevelSelect.value = await getMemberLevelSelect()
})

//监听属性变化
watchEffect(() => {
  if (props.memberVO) {
    form.value = reactive(props.memberVO)
    isEdit.value = true
  } else {
    form.value = {
      name: '',
      phone: '',
      account: '',
      level: {}
    }
    isEdit.value = false
  }
})

//提交表单
function onSubmit() {
  let request = {
    name: form.value.name,
    phone: form.value.phone,
    levelId: form.value.level.id,
    account: form.value.account,
  }
  if (isEdit.value) {
    request['id'] = props.memberVO.id
    MemberUpdate(request).then(resp => {
      commonElMessage(resp)
    }).catch(err => {
      errElMessage(err)
    })
  } else {
    MemberSave(request).then(resp => {
      commonElMessage(resp)
    }).catch(err => {
      errElMessage(err)
    })
  }
  closeDialog()
}

//点击选择框
async function clickSelect() {
  memberLevelSelect.value = await getMemberLevelSelect()
}

//关闭对话框
function closeDialog() {
  emit('close')
}

//打开会员等级对话框
function openMemberLevelDialog() {
  memberLevelDialogVisible.value = true
}

//关闭会员等级对话框
function closeMemberLevelDialog() {
  memberLevelDialogVisible.value = false
}
</script>

<template>

  <el-form :model="form" label-width="120px">
    <el-form-item label="会员名称">
      <el-input v-model="form.name"/>
    </el-form-item>
    <el-form-item label="手机号">
      <el-input v-model="form.phone"/>
    </el-form-item>
    <el-form-item label="会员等级" class="member-level">
      <el-select v-model="form.level.id" @click="clickSelect" placeholder="请选择会员等级" class="member-select">
        <el-option v-for="item in memberLevelSelect" :key="item.id" :label="item.name" :value="item.id"/>
      </el-select>
      <el-button class="member-button" @click="openMemberLevelDialog">新增类型</el-button>
    </el-form-item>
    <el-form-item label="账户余额">
      <el-input v-model="form.account"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">确认</el-button>
      <el-button @click="closeDialog">取消</el-button>
    </el-form-item>
  </el-form>
  <!--对话框-->
  <el-dialog
      v-model="memberLevelDialogVisible"
      title="会员等级"
  >
    <member-level-dialog/>
  </el-dialog>

</template>


<style scoped>

.member-level {
  display: flex;
}

.member-select {
  --el-select-width: 220px;
}

.member-button {
  margin-left: 1px;
}
</style>
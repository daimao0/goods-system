<script setup>
import {defineEmits, onMounted, reactive, ref, watchEffect} from 'vue'
import {GoodsEdit, GoodsSave, GoodsTypeSelectAll} from '../../../wailsjs/go/main/App.js'
import {ElMessage} from "element-plus";
import GoodsTypeDialog from "../goodstype/GoodsTypeDialog.vue";
import {commonElMessage, errElMessage} from '../../utils/el-message-utils.js'

const emit = defineEmits(['refresh', 'close']);
//props接受参数
const props = defineProps({
  goodsVO: {
    type: [Object, null],
    required: true
  }
})
//商品类型对话框
const goodsTypeDialogVisible = ref(false)
//判断是否为编辑
const isEdit = ref(false)
// 表单
let form = ref() // 初始化为空ref对象
const goodsTypeOpts = ref([])

//页面加载后
onMounted(async () => {
  await goodsTypes()
})

// 表单参数，这里使用watchEffect来监听props的变化
watchEffect(() => {
  if (props.goodsVO) {
    form.value = reactive(props.goodsVO)
    isEdit.value = true
  } else {
    form.value = {
      name: '',
      goodsNumber: '',
      goodsType: '',
      price: '',
      count: '',
      desc: '',
    }
    isEdit.value = false
  }
})
// 查询列表
const onSubmit = () => {
  let goodsSaveRequest = reactive({...form.value})
  if (isEdit.value) {
    let goodsEditRequest = {
      id: form.value.id,
      ...form.value
    }
    GoodsEdit(goodsEditRequest).then((resp) => {
      commonElMessage(resp)
      emit('close')
      //刷新分页
      emit('refresh')
    }).catch(err => {
      errElMessage(err)
    })
  } else {
    GoodsSave(goodsSaveRequest).then((resp) => {
      commonElMessage(resp)
    }).catch(err => {
      ElMessage.error('出现异常' + err)
    })
  }
}

//查询商品类型
async function goodsTypes() {
  GoodsTypeSelectAll().then(resp => {
    goodsTypeOpts.value = resp.data
  }).catch(err => {
    console.log(err)
  })
}

//关闭商品对话框
function closeGoodsDialog() {
  emit('close')
}

// 打开商品对话框
function openGoodsTypeDialog() {
  goodsTypeDialogVisible.value = true
}

//关闭商品类型对话框
function closeGoodsTypeDialog() {
  goodsTypeDialogVisible.value = false
}
</script>

<template>
  <el-form :model="form" label-width="120px">
    <el-form-item label="商品名称">
      <el-input v-model="form.name"/>
    </el-form-item>
    <el-form-item label="商品编号">
      <el-input v-model="form.goodsNumber"/>
    </el-form-item>
    <el-form-item label="商品类型" class="goodsType">
      <el-select v-model="form.goodsType" @click="goodsTypes" placeholder="请选择商品类型" class="goodsType-select">
        <el-option v-for="item in goodsTypeOpts" :key="item.id" :label="item.name" :value="item.name"/>
      </el-select>
      <el-button class="goodsType-button" @click="openGoodsTypeDialog">新增类型</el-button>
    </el-form-item>
    <el-form-item label="商品价格">
      <el-input v-model="form.price"/>
    </el-form-item>
    <el-form-item label="商品数量">
      <el-input v-model="form.count"/>
    </el-form-item>
    <el-form-item label="商品描述">
      <el-input v-model="form.desc" type="textarea"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">确认</el-button>
      <el-button @click="closeGoodsDialog">取消</el-button>
    </el-form-item>
  </el-form>

  <!--  商品类型对话框   -->
  <el-dialog
      v-model="goodsTypeDialogVisible"
      title="商品类型"
      :before-close="closeGoodsTypeDialog"
      @close="closeGoodsTypeDialog"
  >
    <GoodsTypeDialog/>
  </el-dialog>
</template>


<style scoped>
.goodsType {
  display: flex;
}

.goodsType-select {
  --el-select-width: 220px;
}

.goodsType-button {
  margin-left: 1px;
}
</style>
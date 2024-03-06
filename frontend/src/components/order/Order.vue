<script setup>

import {onMounted, ref} from "vue";
import {allGoodsSelect, allMemberSelect} from "../../api/select.js";
import {GoodsGetById, MemberGetById, Order} from "../../../wailsjs/go/main/App.js";
import {Delete} from "@element-plus/icons-vue";
import {ElMessage} from "element-plus";
import {errElMessage, msgElMessage} from "../../utils/el-message-utils.js";
//对话框
const dialogVisible = ref(false)
//是否抵扣
const isDeduction = ref(false)
const form = ref({
  memberId: '',
  realPrice: '',
  num: 1,
})
const goodsForm = ref({
  id: '',
  name: '',
  goodsNumber: '',
  goodsType: '',
  price: '',
  count: ''
})
const memberForm = ref({
  id: '',
  name: '',
  phone: '',
  level: '',
  account: '',
  discount: ''
})
const bills = ref([])
const totalPrice = ref('0')
const pay = ref('0')
const goodsSelect = ref([])
const memberSelect = ref([])


onMounted(async () => {
  goodsSelect.value = await allGoodsSelect()
  memberSelect.value = await allMemberSelect()
})

// 下单
function onSubmit() {
  let request = {
    pay: totalPrice.value,
    memberId: memberForm.value.id,
    orderItem: []
  }
  for (let i = 0; i < bills.value.length; i++) {
    let bill = bills.value[i]
    let order = {
      goodsId: bill.goodsId,
      realPrice: bill.realPrice,
      num: bill.num
    }
    request.orderItem.push(order)
  }
  //如果购物车价格超过会员账户，提示
  if (!dialogVisible.value && memberForm.value.id !== "" && Number(request.pay) > Number(memberForm.value.account)) {
    dialogVisible.value = true
  } else {
    //下订单
    Order(request).then(resp => {
      msgElMessage(resp, "下单成功")
      bills.value = []
      totalPrice.value = 0
      isDeduction.value = false
    }).catch(err => {
      errElMessage(err)
    })
  }
}

//加入购物车
function addCart() {
  let bill = {
    goodsId: goodsForm.value.id,
    goodsName: goodsForm.value.name,
    goodsNumber: goodsForm.value.goodsNumber,
    goodsType: goodsForm.value.goodsType,
    price: goodsForm.value.price,
    realPrice: form.value.realPrice,
    num: form.value.num,
  }
  if (goodsForm.value.count < bill.num) {
    ElMessage.warning(goodsForm.value.name + "库存不足，库存数量为" + goodsForm.value.count)
    return
  }
  for (let i = 0; i < bills.value.length; i++) {
    if (bills.value[i].goodsId === bill.goodsId) {
      ElMessage.warning(bills.value[i].goodsName + "商品已加入购物车")
      return
    }
  }
  bills.value.push(bill)
  calcTotalPrice()
}

// 切换商品
function changeGoods(goodsId) {
  GoodsGetById(goodsId).then(resp => {
    goodsForm.value.id = resp.data.id
    goodsForm.value.name = resp.data.name
    goodsForm.value.goodsType = resp.data.goodsType
    goodsForm.value.goodsNumber = resp.data.goodsNumber
    goodsForm.value.price = resp.data.price
    goodsForm.value.count = resp.data.count
    calcRealPrice()
  })
}

//deleteGoods 删除商品
function deleteGoods(row) {
  for (let i = 0; i < bills.value.length; i++) {
    if (bills.value[i].goodsId === row.goodsId) {
      bills.value.splice(i, 1)
      break
    }
  }
  calcTotalPrice()
}

// 根据手机号过滤
async function filterMember(phone) {
  memberSelect.value = await allMemberSelect()
  if (phone) {
    let select = []
    for (let i = 0; i < memberSelect.value.length; i++) {
      let member = memberSelect.value[i]
      if (member.phone.includes(phone)) {
        select.push(member)
      }
    }
    memberSelect.value = select
  }
}

// 切换会员
function changeMember(memberId) {
  MemberGetById(memberId).then(resp => {
    memberForm.value.id = resp.data.id
    memberForm.value.name = resp.data.name
    memberForm.value.level = resp.data.level.name
    memberForm.value.discount = resp.data.level.discount
    memberForm.value.account = resp.data.account
    calcRealPrice()
  })
}

//修改商品数量
function changeGoodsNum() {
  calcTotalPrice()
}

//计算实际支付
function calcRealPrice() {
  if (goodsForm.value.id !== "") {
    if (memberForm.value.id !== "" && memberForm.value.id !== '0') {
      form.value.realPrice = (Number(goodsForm.value.price) * Number(memberForm.value.discount)).toFixed(2)
    } else {
      form.value.realPrice = Number(goodsForm.value.price).toFixed(2)
    }
  }
}

//计算总价
function calcTotalPrice() {
  totalPrice.value = '0'
  for (let i = 0; i < bills.value.length; i++) {
    let bill = bills.value[i]
    totalPrice.value = (Number(totalPrice.value) + Number(bill.realPrice) * bill.num).toFixed(2)
  }
}

//会员抵扣对话框确认
function memberDeduction() {
  isDeduction.value = true
  onSubmit()
  dialogVisible.value = false
}
</script>

<template>
  <h3>商品下单</h3>
  <el-container>
    <el-aside width="360px">
      <el-form :model="form" class="form">
        <el-form-item label="商品名称">
          <el-select
              v-model="form.goodsId"
              filterable
              placeholder="输入商品名称"
              @change="changeGoods"
          >
            <el-option
                v-for="item in goodsSelect"
                :key="item.id"
                :label="item.name"
                :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="会员手机号">
          <el-select
              v-model="form.memberId"
              filterable
              :filter-method="filterMember"
              placeholder="输入手机号"
              @change="changeMember"
              :disabled="bills.length>0"
              clearable
          >
            <el-option
                v-for="item in memberSelect"
                :key="item.id"
                :label="item.name"
                :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="购买单价">
          <el-input v-model="form.realPrice"/>
        </el-form-item>
        <el-form-item label="数量">
          <el-input-number v-model="form.num" :min="1" :max="999"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="addCart">加入购物车</el-button>
        </el-form-item>
      </el-form>
    </el-aside>

    <el-main>
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>购物清单</span><br/><br/>
            <span v-if="memberForm.id!=='' && memberForm.id!=='0'">{{ memberForm.name }} {{
                memberForm.level
              }} 折扣:{{ memberForm.discount }} 余额:{{
                memberForm.account
              }} 元</span>
          </div>
        </template>
        <div>
          <el-table :data="bills" stripe style="width: 100%">
            <el-table-column prop="goodsName" label="商品名称"/>
            <el-table-column prop="goodsNumber" label="商品编号"/>
            <el-table-column prop="goodsType" label="商品类型"/>
            <el-table-column prop="price" label="商品单价"/>
            <el-table-column prop="realPrice" label="实际单价"/>
            <el-table-column label="购买数量">
              <template #default="scope">
                <el-input-number v-model="scope.row.num" :min="1" :max="999" size="small" @change="changeGoodsNum"
                                 style="width: 80px"/>
              </template>
            </el-table-column>
            <el-table-column label="操作">
              <template #default="scope">
                <el-button type="danger" :icon="Delete" circle size="small" @click="deleteGoods(scope.row)"/>
              </template>
            </el-table-column>
          </el-table>
          <el-form-item label="总价">
            <el-input v-model="totalPrice"/>
          </el-form-item>
        </div>
        <template #footer>
          <el-button type="primary" @click="onSubmit">下单</el-button>
        </template>
      </el-card>
    </el-main>
  </el-container>

  <!--  用户余额不足对话框  -->
  <el-dialog
      v-model="dialogVisible"
      title="余额不足"
      width="500"
  >
    <span>用户余额不足,是否直接抵扣(完全扣除余额)，用户余额{{memberForm.account}}元</span>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="memberDeduction">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<style scoped>
</style>
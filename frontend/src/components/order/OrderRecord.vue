<script setup>

import {onMounted, ref, watchEffect} from "vue";
import {OrderPage, OrderStatistics} from "../../../wailsjs/go/main/App.js";
import {errElMessage} from "../../utils/el-message-utils.js";
import {datetimeCuts, dateTimeToDateStr} from "../../utils/datetime.js"
import {allMemberSelect} from "../../api/select.js";
//订单统计
const orderStatistics = ref({
  orderCount: 0,
  goodsPrice: 0,
  memberPay: 0,
  preferentialPrice: 0,
})
//订单分页参数
const orderPageRequest = ref({
  memberId: '',
  startTime: '',
  endTime: '',
  page: 1,
  size: 10
})
//订单分页
const orderPageVOList = ref([])

// 日期参数
const datetime = ref([])
// 日期选择器
const shortcuts = datetimeCuts
const memberSelect = ref([])
const pageTotal = ref(0)

watchEffect(() => {
  //订单分页
  searchOrderPage()
  //订单统计
  searchOrderStatistics()
})
onMounted(async () => {
  memberSelect.value = await allMemberSelect()
})

//分页查询订单
function searchOrderPage() {
  if (datetime.value == null) {
    orderPageRequest.value.startTime = null
    orderPageRequest.value.endTime = null
  }
  if (datetime.value != null && datetime.value.length === 2) {
    orderPageRequest.value.startTime = dateTimeToDateStr(datetime.value[0])
    orderPageRequest.value.endTime = dateTimeToDateStr(datetime.value[1])
  }
  OrderPage(orderPageRequest.value).then(resp => {
    orderPageVOList.value = resp.data.data
    pageTotal.value = resp.data.total
  }).catch(err => {
    errElMessage(err)
  })
}

//订单统计搜索
function searchOrderStatistics() {
  let request = {
    memberId: orderPageRequest.value.memberId,
    startTime: orderPageRequest.value.startTime,
    endTime: orderPageRequest.value.endTime
  }
  OrderStatistics(request).then(resp => {
    orderStatistics.value.orderCount = Number(resp.data.orderCount)
    orderStatistics.value.goodsPrice = Number(resp.data.goodsPrice)
    orderStatistics.value.memberPay = Number(resp.data.memberPay)
    orderStatistics.value.preferentialPrice = Number(resp.data.preferentialPrice)
  })
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

// 添加分页切换事件处理函数（假设你已经实现了分页逻辑）
function handleCurrentChange(pageNumber) {
  orderPageRequest.value.page = pageNumber
  searchOrderPage()
}
</script>

<template>
  <el-form :inline="true" :model="orderPageRequest">
    <el-form-item label="选择会员">
      <el-select
          v-model="orderPageRequest.memberId"
          filterable
          :filter-method="filterMember"
          placeholder="请输入手机号"
          clearable
          style="width: 180px"
      >
        <el-option v-for="item in memberSelect" :key="item.id" :label="item.name" :value="item.id"/>
      </el-select>
    </el-form-item>
    <el-form-item label="创建时间">
      <el-date-picker
          v-model="datetime"
          type="daterange"
          unlink-panels
          range-separator="至"
          start-placeholder="开始时间"
          end-placeholder="结束时间"
          :shortcuts="shortcuts"
      />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="searchOrderPage">查询</el-button>
    </el-form-item>
  </el-form>
  <!--  统计  -->
  <el-row>
    <el-col :span="6">
      <el-statistic title="订单数" :value="orderStatistics.orderCount"/>
    </el-col>
    <el-col :span="6">
      <el-statistic title="商品总价" :value="orderStatistics.goodsPrice"/>
    </el-col>
    <el-col :span="6">
      <el-statistic title="支付总价" :value="orderStatistics.memberPay"/>
    </el-col>
    <el-col :span="6">
      <el-statistic title="优惠总价" :value="orderStatistics.preferentialPrice"/>
    </el-col>
  </el-row>
  <!--  明细  -->
  <el-table :data="orderPageVOList" stripe style="width: 100%">
    <el-table-column prop="id" label="id"/>
    <el-table-column prop="member.name" label="会员名称"/>
    <el-table-column prop="name" label="购买商品">
      <template #default="scope">
        <div v-for="(item, index) in scope.row.orderItems" :key="index">
          {{ item.goods.name }}
          <el-tag size="small" type="info">
            {{ item.realPrice }} x{{ item.num }}
          </el-tag>
        </div>
      </template>
    </el-table-column>
    <el-table-column prop="pay" label="总价"/>
    <el-table-column prop="createTime" label="订单时间"/>
  </el-table>

  <!-- 分页组件 -->
  <!--  <div class="example-pagination-block" v-if="tablePageVOList.length > 0">-->
  <div class="example-pagination-block" v-if="true">
    <el-pagination
        layout="prev, pager, next"
        :total="pageTotal"
        @current-change="handleCurrentChange"
    />
  </div>
</template>

<style scoped>

</style>
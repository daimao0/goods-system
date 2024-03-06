import {createRouter, createWebHistory} from "vue-router";
import Goods from "../components/goods/Goods.vue";
import Member from "../components/member/Member.vue";
import Bill from "../components/order/Order.vue";
import OrderRecord from "../components/order/OrderRecord.vue";
import System from "../components/system/System.vue";

const routes = [
    {
        path: '/',
        name: 'Goods',
        component: Goods,
    }, {
        path: '/member',
        name: 'Member',
        component: Member,
    },{
        path: '/order',
        name: 'Bill',
        component: Bill,
    },{
        path: '/orderRecord',
        name: 'OrderRecord',
        component: OrderRecord,
    },{
        path: '/system',
        name: 'System',
        component: System,
    }
]
const router = createRouter({
    history: createWebHistory(), // 路由模式：history模式
    routes: routes
})
export default router
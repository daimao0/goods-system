import {GoodsAll, GoodsTypeSelectAll, MemberAll, MemberLevelSelectAll} from "../../wailsjs/go/main/App.js";

//会员等级
export async function getMemberLevelSelect() {
    let select
    await MemberLevelSelectAll().then(resp => {
        select = resp.data
    })
    return select
}

//商品类型
export async function getGoodsTypeSelect(){
    let select
    await GoodsTypeSelectAll().then(resp=>{
        select = resp.data
    })
    return select
}

//获得全部商品
export async function allGoodsSelect(){
    let select
    await GoodsAll().then(resp=>{
        select = resp.data
    })
    return select
}

//获得所有会员
export async function allMemberSelect(){
   let select
    await MemberAll().then(resp=>{
        select = resp.data
    })
    return select
}

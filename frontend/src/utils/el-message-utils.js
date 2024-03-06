//默认饿了么消息
import {ElMessage} from "element-plus";

export function commonElMessage(resp){
    if (resp.code===200){
        ElMessage.success('编辑成功')
    }else{
        ElMessage.error(resp.message)
    }
}
export function msgElMessage(resp,msg){
    if (resp.code===200){
        ElMessage.success(msg)
    }else{
        ElMessage.error(resp.message)
    }
}

export function errElMessage(err){
    ElMessage.error('出现异常,请联系管理员',err)
}

export function errMsgElMessage(msg){
    ElMessage.error(msg)
}
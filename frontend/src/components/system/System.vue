<script setup>
import {ref} from 'vue'
import {ImportBackup, SystemBackup, SystemDesktopPath} from "../../../wailsjs/go/main/App.js";
import {commonElMessage, errElMessage, errMsgElMessage, msgElMessage} from "../../utils/el-message-utils.js";

const backupPath = ref(null)

// 备份文件
function backup() {
  if (!backupPath.value) {
    errMsgElMessage("要先填写路径才能备份文件")
  }
  SystemBackup(backupPath.value).then(resp => {
    commonElMessage(resp)
  })
}

// 获得桌面路径
function getDesktopPath() {
  SystemDesktopPath().then(resp => {
    backupPath.value = resp.data.path
  })
}

// 同步备份
function syncBackup(file) {
  let reader = new FileReader();

  reader.onload = function (e) {
    let fileContentArrayBuffer = e.target.result; // 文件内容的 ArrayBuffer 形式
    let fileContentBase64 = btoa(new Uint8Array(fileContentArrayBuffer).reduce((data, byte) => data + String.fromCharCode(byte), '')); // 转换为 Base64 字符串

    // 假设 Wails 后端接受 Base64 字符串
    ImportBackup(fileContentBase64).then(resp => {
      msgElMessage(resp,"同步成功")
    }).catch(err => {
      errElMessage(err)
    });
  };

  reader.onerror = function () {
    console.error('Error reading the file');
  };

  // 开始读取文件为 ArrayBuffer
  reader.readAsArrayBuffer(file.raw);
}
</script>

<template>
  <!--  备份  -->
  <el-form>
    <el-form-item>
      <!--  备份提示  -->
      <el-text class="mx-1" size="large">该系统当您更换主机、重做系统、删除数据库时需要先导出备份，再同步备份</el-text>
    </el-form-item>
    <el-form-item>
      <!--  备份文件  -->
      <el-button type="info" @click="backup">导出备份</el-button>
      <el-input v-model="backupPath" style="width: 360px" placeholder="请输入导出备份文件的路径"/>
    </el-form-item>
    <el-form-item>
      <el-button type="info" plain @click="getDesktopPath">获得桌面路径</el-button>
    </el-form-item>
    <el-form-item>
      <el-upload
          action=""
          class="upload-demo"
          :auto-upload="false"
          @change="syncBackup"
      >
        <!--  上传文件  -->
        <el-button type="primary">同步备份</el-button>
        <template #tip>
          <div class="el-upload__tip">
            只能提交备份文件
          </div>
        </template>
      </el-upload>
    </el-form-item>
  </el-form>
</template>

<style scoped>

</style>
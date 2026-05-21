<template>
  <div class="home-container">
    <el-container>
      <el-header>
        <div class="header-content">
          <span class="system-name">IMS 制造运营管理系统</span>
          <div class="user-info">
            <span>{{ user.nickname || user.username }} ({{ user.role }})</span>
            <el-button type="text" @click="handleLogout">退出登录</el-button>
          </div>
        </div>
      </el-header>
      <el-main>
        <h1>欢迎使用 MES1 系统</h1>
        <p>您已成功登录。</p>
      </el-main>
    </el-container>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()
const user = JSON.parse(localStorage.getItem('user') || '{}')

const handleLogout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/login')
}
</script>

<style scoped>
.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 60px;
  border-bottom: 1px solid #eee;
}

.system-name {
  font-size: 18px;
  font-weight: bold;
}

.user-info {
  display: flex;
  gap: 20px;
  align-items: center;
}
</style>

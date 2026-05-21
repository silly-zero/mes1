<template>
  <div class="home-container">
    <el-container class="layout-container">
      <el-aside width="200px">
        <el-menu
          default-active="1"
          class="el-menu-vertical"
          background-color="#304156"
          text-color="#bfcbd9"
          active-text-color="#409EFF"
          router
        >
          <el-menu-item index="/">
            <el-icon><Monitor /></el-icon>
            <span>首页看板</span>
          </el-menu-item>
          <el-menu-item index="/inventory">
            <el-icon><Box /></el-icon>
            <span>库存查询</span>
          </el-menu-item>
          <el-menu-item index="/inbound">
            <el-icon><Download /></el-icon>
            <span>入库管理</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      
      <el-container>
        <el-header>
          <div class="header-content">
            <span class="system-name">IMS 制造运营管理系统</span>
            <div class="user-info">
              <el-tag :type="user.role === 'core_admin' ? 'danger' : 'success'" size="small">
                {{ user.role === 'core_admin' ? '核心管理员' : '管理员' }}
              </el-tag>
              <span class="username">{{ user.nickname || user.username }}</span>
              <el-button type="info" link @click="handleLogout">退出登录</el-button>
            </div>
          </div>
        </el-header>
        
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { Monitor, Box, Download } from '@element-plus/icons-vue'

const router = useRouter()
const user = JSON.parse(localStorage.getItem('user') || '{}')

const handleLogout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/login')
}
</script>

<style scoped>
.layout-container {
  height: 100vh;
}

.el-aside {
  background-color: #304156;
}

.el-menu-vertical {
  border-right: none;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 60px;
  border-bottom: 1px solid #e6e6e6;
  padding: 0 20px;
}

.system-name {
  font-size: 18px;
  font-weight: bold;
  color: #303133;
}

.user-info {
  display: flex;
  gap: 15px;
  align-items: center;
}

.username {
  font-size: 14px;
  color: #606266;
}
</style>

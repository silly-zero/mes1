<template>
  <div class="home-container">
    <el-container class="layout-container">
      <el-aside width="220px">
        <div class="logo-area">
          <span class="logo-text">MES SYSTEM</span>
        </div>
        <el-menu
          :default-active="$route.path"
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
          
          <el-sub-menu index="inventory-mgt">
            <template #title>
              <el-icon><Box /></el-icon>
              <span>库存管理</span>
            </template>
            <el-menu-item index="/inventory">库存查询</el-menu-item>
            <el-menu-item index="/receipt">收料管理</el-menu-item>
            <el-menu-item index="/iqc">来料检验</el-menu-item>
            <el-menu-item index="/inbound">入库管理</el-menu-item>
            <el-menu-item index="/outbound">出库管理</el-menu-item>
          </el-sub-menu>
        </el-menu>
      </el-aside>
      
      <el-container>
        <el-header>
          <div class="header-content">
            <span class="breadcrumb">{{ $route.name }}</span>
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
import { useRouter, useRoute } from 'vue-router'
import { Monitor, Box, Download, Upload, List } from '@element-plus/icons-vue'

const router = useRouter()
const route = useRoute()
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

.logo-area {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #2b3648;
}

.logo-text {
  color: #fff;
  font-size: 18px;
  font-weight: bold;
  letter-spacing: 1px;
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

.breadcrumb {
  font-size: 16px;
  color: #606266;
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

<template>
  <div class="login-container">
    <div class="login-box">
      <div class="logo">
        <h1 class="brand-text">Pangus</h1>
      </div>
      <h2 class="title">IMS 制造运营管理系统</h2>
      
      <el-form :model="loginForm" :rules="rules" ref="loginFormRef" class="login-form">
        <el-form-item prop="username">
          <el-input 
            v-model="loginForm.username" 
            placeholder="用户"
            :prefix-icon="User"
          />
        </el-form-item>
        
        <el-form-item prop="password">
          <el-input 
            v-model="loginForm.password" 
            type="password" 
            placeholder="密码"
            :prefix-icon="Lock"
            show-password
          />
        </el-form-item>

        <div class="options">
          <el-link type="info" :underline="false">忘记密码</el-link>
          <div class="checkboxes">
            <el-checkbox v-model="loginForm.rememberUser">记住用户</el-checkbox>
            <el-checkbox v-model="loginForm.rememberPassword">记住密码</el-checkbox>
          </div>
        </div>

        <el-form-item>
          <el-button 
            type="primary" 
            class="login-button" 
            :loading="loading"
            @click="handleLogin"
          >
            登录
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { User, Lock } from '@element-plus/icons-vue'
import { login } from '../api'
import { ElMessage } from 'element-plus'

const router = useRouter()
const loginFormRef = ref(null)
const loading = ref(false)

const loginForm = reactive({
  username: '',
  password: '',
  rememberUser: false,
  rememberPassword: false
})

const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

const handleLogin = async () => {
  if (!loginFormRef.value) return
  
  await loginFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        const res = await login({
          username: loginForm.username,
          password: loginForm.password
        })
        localStorage.setItem('token', res.token)
        localStorage.setItem('user', JSON.stringify(res.user))
        ElMessage.success('登录成功')
        router.push('/')
      } catch (error) {
        console.error(error)
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style scoped>
.login-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f5f7fa;
}

.login-box {
  width: 400px;
  padding: 40px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.08);
  text-align: center;
}

.brand-text {
  font-size: 48px;
  color: #e67e22; /* 橙色系，类似图片中的颜色 */
  margin-bottom: 10px;
  font-weight: bold;
}

.title {
  font-size: 20px;
  color: #333;
  margin-bottom: 30px;
  font-weight: 500;
}

.login-form {
  margin-top: 20px;
}

:deep(.el-input__wrapper) {
  padding: 8px 12px;
}

.options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  font-size: 14px;
}

.checkboxes {
  display: flex;
  gap: 15px;
}

.login-button {
  width: 100%;
  height: 44px;
  font-size: 16px;
  background-color: #e67e22 !important;
  border-color: #e67e22 !important;
}

.login-button:hover {
  background-color: #d35400 !important;
  border-color: #d35400 !important;
}
</style>

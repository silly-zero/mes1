<template>
  <div class="dashboard">
    <el-row :gutter="20">
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>今日入库单</template>
          <div class="card-value">{{ stats.today_inbound }}</div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>待检验批次</template>
          <div class="card-value">{{ stats.pending_qc }}</div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>库存预警</template>
          <div class="card-value color-danger">{{ stats.low_stock }}</div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from '../api'

const stats = ref({
  today_inbound: 0,
  pending_qc: 0,
  low_stock: 0
})

const fetchStats = async () => {
  try {
    const res = await axios.get('/core/dashboard')
    stats.value = res
  } catch (error) {
    console.error('获取统计数据失败', error)
  }
}

onMounted(fetchStats)
</script>

<style scoped>
.card-value {
  font-size: 32px;
  font-weight: bold;
  text-align: center;
}
.color-danger {
  color: #F56C6C;
}
</style>

<template>
  <div class="inventory">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>实时库存查询</span>
          <el-button type="primary" @click="fetchData">刷新</el-button>
        </div>
      </template>
      
      <el-table :data="tableData" border style="width: 100%" v-loading="loading">
        <el-table-column prop="material_code" label="物料编码" />
        <el-table-column prop="quantity" label="库存数量" />
        <el-table-column prop="warehouse" label="存放仓库" />
        <el-table-column prop="UpdatedAt" label="最后更新时间">
          <template #default="scope">
            {{ new Date(scope.row.UpdatedAt).toLocaleString() }}
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from '../api'

const tableData = ref([])
const loading = ref(false)

const fetchData = async () => {
  loading.value = true
  try {
    const res = await axios.get('/inventory/list')
    tableData.value = res
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

onMounted(fetchData)
</script>

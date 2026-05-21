<template>
  <div class="inbound">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>入库管理</span>
          <el-button type="primary" @click="fetchData">刷新</el-button>
        </div>
      </template>
      
      <el-table :data="tableData" border style="width: 100%" v-loading="loading">
        <el-table-column prop="inbound_no" label="入库单号" width="180" />
        <el-table-column prop="material_code" label="物料编码" />
        <el-table-column prop="qualified_qty" label="合格数量" />
        <el-table-column prop="is_qc_confirmed" label="品质确认">
          <template #default="scope">
            <el-tag :type="scope.row.is_qc_confirmed ? 'success' : 'info'">
              {{ scope.row.is_qc_confirmed ? '已确认' : '未确认' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="is_inbound" label="入库状态">
          <template #default="scope">
            <el-tag :type="scope.row.is_inbound ? 'success' : 'warning'">
              {{ scope.row.is_inbound ? '已入库' : '待入库' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-button 
              v-if="!scope.row.is_inbound" 
              type="success" 
              size="small" 
              @click="handleConfirm(scope.row.ID)"
            >
              确认入库
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from '../api'
import { ElMessage } from 'element-plus'

const tableData = ref([])
const loading = ref(false)

const fetchData = async () => {
  loading.value = true
  try {
    const res = await axios.get('/inventory/inbound')
    tableData.value = res
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const handleConfirm = async (id) => {
  try {
    await axios.post(`/inventory/inbound/${id}/confirm`)
    ElMessage.success('入库成功')
    fetchData()
  } catch (error) {
    console.error(error)
  }
}

onMounted(fetchData)
</script>

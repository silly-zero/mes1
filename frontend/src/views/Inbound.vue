<template>
  <div class="inbound">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>入库管理</span>
          <el-button type="primary" @click="fetchData">刷新</el-button>
        </div>
      </template>
      
      <el-table :data="tableData" border style="width: 100%" v-loading="loading" height="calc(100vh - 250px)">
        <el-table-column type="index" label="行号" width="60" fixed="left" />
        <el-table-column prop="inbound_no" label="入库确认单号" width="220" fixed="left">
          <template #default="scope">
            <span style="color: #E6A23C;">{{ scope.row.inbound_no }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="receive_no" label="收料单编码" width="220" />
        <el-table-column prop="material_code" label="物料编码" width="150" />
        <el-table-column prop="total_qty" label="收料数量" width="100" />
        <el-table-column prop="qualified_qty" label="合格数量" width="100" />
        <el-table-column prop="unqualified_qty" label="不合格数量" width="100" />
        <el-table-column prop="is_qc_confirmed" label="品质确认" width="100" align="center">
          <template #default="scope">
            <el-tag :type="scope.row.is_qc_confirmed ? 'success' : 'info'" size="small">
              {{ scope.row.is_qc_confirmed ? '已确认' : '未确认' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="qc_user" label="品质确认人" width="120" />
        <el-table-column prop="qc_time" label="品质确认时间" width="180" />
        <el-table-column prop="is_inbound" label="入库状态" width="100" align="center">
          <template #default="scope">
            <el-tag :type="scope.row.is_inbound ? 'success' : 'warning'" size="small">
              {{ scope.row.is_inbound ? '已入库' : '待入库' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="inbound_user" label="入库确认人" width="120" />
        <el-table-column prop="inbound_time" label="入库确认时间" width="180" />
        <el-table-column prop="org_name" label="组织名称" width="150" />
        <el-table-column label="操作" width="120" fixed="right" align="center">
          <template #default="scope">
            <el-button 
              v-if="!scope.row.is_inbound && scope.row.is_qc_confirmed" 
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
    await axios.post(`/inventory/inbound/${id}`)
    ElMessage.success('入库成功')
    fetchData()
  } catch (error) {
    console.error(error)
  }
}

onMounted(fetchData)
</script>

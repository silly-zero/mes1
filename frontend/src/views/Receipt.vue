<template>
  <div class="receipt">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>收料管理</span>
          <el-button type="primary" @click="dialogVisible = true">新增收料单</el-button>
        </div>
      </template>

      <el-table :data="tableData" border style="width: 100%" v-loading="loading">
        <el-table-column prop="inbound_no" label="入库确认单号" width="180" />
        <el-table-column prop="receive_no" label="收料单号" width="180" />
        <el-table-column prop="material_code" label="物料编码" />
        <el-table-column prop="total_qty" label="收料总数" />
        <el-table-column prop="CreatedAt" label="收料时间">
          <template #default="scope">
            {{ new Date(scope.row.CreatedAt).toLocaleString() }}
          </template>
        </el-table-column>
        <el-table-column label="状态" width="120">
          <template #default="scope">
            <el-tag v-if="!scope.row.is_qc_confirmed" type="warning">待检验</el-tag>
            <el-tag v-else-if="!scope.row.is_inbound" type="primary">待入库</el-tag>
            <el-tag v-else type="success">已完成</el-tag>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" title="新增收料记录" width="500px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="入库单号">
          <el-input v-model="form.inbound_no" placeholder="自动生成或输入" />
        </el-form-item>
        <el-form-item label="收料单号">
          <el-input v-model="form.receive_no" />
        </el-form-item>
        <el-form-item label="物料编码">
          <el-input v-model="form.material_code" />
        </el-form-item>
        <el-form-item label="收料数量">
          <el-input-number v-model="form.total_qty" :min="0" />
        </el-form-item>
        <el-form-item label="加急">
          <el-switch v-model="form.is_urgent" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import axios from '../api'
import { ElMessage } from 'element-plus'

const tableData = ref([])
const loading = ref(false)
const dialogVisible = ref(false)

const form = reactive({
  inbound_no: 'MES' + Date.now(),
  receive_no: 'REC' + Date.now(),
  material_code: '',
  total_qty: 0,
  is_urgent: false,
  confirm_type: '来料入库确认'
})

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

const handleSubmit = async () => {
  try {
    await axios.post('/inventory/receipt', form)
    ElMessage.success('收料成功')
    dialogVisible.ref = false
    fetchData()
  } catch (error) {
    console.error(error)
  }
}

onMounted(fetchData)
</script>

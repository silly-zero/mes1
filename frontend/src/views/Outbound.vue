<template>
  <div class="outbound">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>出库管理</span>
          <el-button type="primary" @click="dialogVisible = true">新增出库单</el-button>
        </div>
      </template>

      <el-table :data="tableData" border style="width: 100%" v-loading="loading">
        <el-table-column prop="outbound_no" label="出库单号" width="180" />
        <el-table-column prop="material_code" label="物料编码" />
        <el-table-column prop="quantity" label="出库数量" />
        <el-table-column prop="receiver" label="领料人" />
        <el-table-column prop="outbound_time" label="出库时间" />
        <el-table-column prop="status" label="状态">
          <template #default="scope">
            <el-tag type="success">{{ scope.row.status === 'completed' ? '已完成' : '进行中' }}</el-tag>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" title="新增出库单" width="500px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="出库单号">
          <el-input v-model="form.outbound_no" />
        </el-form-item>
        <el-form-item label="物料编码">
          <el-select v-model="form.material_code" placeholder="选择物料" filterable @focus="fetchStock">
            <el-option 
              v-for="item in stockOptions" 
              :key="item.material_code" 
              :label="`${item.material_code} (库存: ${item.quantity})`" 
              :value="item.material_code" 
            />
          </el-select>
        </el-form-item>
        <el-form-item label="出库数量">
          <el-input-number v-model="form.quantity" :min="0" />
        </el-form-item>
        <el-form-item label="领料人">
          <el-input v-model="form.receiver" />
        </el-form-item>
        <el-form-item label="仓库">
          <el-input v-model="form.warehouse" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确认出库</el-button>
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
const stockOptions = ref([])
const loading = ref(false)
const dialogVisible = ref(false)

const form = reactive({
  outbound_no: 'OUT' + Date.now(),
  material_code: '',
  quantity: 0,
  receiver: '',
  warehouse: '默认仓库'
})

const fetchData = async () => {
  loading.value = true
  try {
    const res = await axios.get('/inventory/outbound')
    tableData.value = res
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const fetchStock = async () => {
  try {
    const res = await axios.get('/inventory/stock')
    stockOptions.value = res
  } catch (error) {
    console.error(error)
  }
}

const handleSubmit = async () => {
  try {
    await axios.post('/inventory/outbound', form)
    ElMessage.success('出库成功')
    dialogVisible.value = false
    fetchData()
  } catch (error) {
    console.error(error)
  }
}

onMounted(fetchData)
</script>

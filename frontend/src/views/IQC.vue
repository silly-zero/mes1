<template>
  <div class="iqc">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>来料检验 (IQC)</span>
          <el-button type="primary" @click="fetchData">刷新</el-button>
        </div>
      </template>

      <el-table :data="pendingOrders" border style="width: 100%" v-loading="loading">
        <el-table-column prop="inbound_no" label="入库确认单号" width="180" />
        <el-table-column prop="material_code" label="物料编码" />
        <el-table-column prop="total_qty" label="收料数量" />
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-button type="warning" size="small" @click="handleInspect(scope.row)">进行检验</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" title="录入检验结果" width="500px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="检验单号">
          <el-input v-model="form.qc_no" />
        </el-form-item>
        <el-form-item label="合格数量">
          <el-input-number v-model="form.qualified_qty" :min="0" :max="currentRow?.total_qty" />
        </el-form-item>
        <el-form-item label="不合格数">
          <el-input-number v-model="form.unqualified_qty" :min="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">提交检验</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import axios from '../api'
import { ElMessage } from 'element-plus'

const tableData = ref([])
const loading = ref(false)
const dialogVisible = ref(false)
const currentRow = ref(null)

const form = reactive({
  qc_no: 'QC' + Date.now(),
  qualified_qty: 0,
  unqualified_qty: 0
})

const pendingOrders = computed(() => {
  return tableData.value.filter(item => !item.is_qc_confirmed)
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

const handleInspect = (row) => {
  currentRow.value = row
  form.qualified_qty = row.total_qty
  form.unqualified_qty = 0
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    await axios.post(`/inventory/iqc/${currentRow.value.ID}`, form)
    ElMessage.success('检验结果已录入')
    dialogVisible.value = false
    fetchData()
  } catch (error) {
    console.error(error)
  }
}

onMounted(fetchData)
</script>

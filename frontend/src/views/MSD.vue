<template>
  <div class="msd-management">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>湿敏物料管理 (MSD)</span>
          <el-button type="primary" @click="dialogVisible = true">新增 MSD 记录</el-button>
        </div>
      </template>

      <el-table :data="tableData" border style="width: 100%" v-loading="loading">
        <el-table-column prop="lot_no" label="批次/卷盘号" width="180" />
        <el-table-column prop="material_code" label="物料编码" width="150" />
        <el-table-column prop="msl_level" label="MSL等级" width="100" />
        <el-table-column label="状态" width="120">
          <template #default="scope">
            <el-tag :type="statusType(scope.row.status)">{{ statusText(scope.row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="剩余寿命 (h)" width="120">
          <template #default="scope">
            <span :class="{'text-danger': scope.row.remaining_life < 24}">
              {{ scope.row.remaining_life === -1 ? '不限制' : scope.row.remaining_life.toFixed(2) }}
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="total_exposed_time" label="累计暴露 (h)" width="120">
          <template #default="scope">
            {{ scope.row.total_exposed_time.toFixed(2) }}
          </template>
        </el-table-column>
        <el-table-column prop="operator" label="操作人" width="120" />
        <el-table-column label="操作" fixed="right" width="250">
          <template #default="scope">
            <el-button v-if="scope.row.status === 'InBag' || scope.row.status === 'Baking'" 
                       type="primary" size="small" @click="handleAction(scope.row.ID, 'open')">开封</el-button>
            <el-button v-if="scope.row.status === 'Exposed'" 
                       type="warning" size="small" @click="handleAction(scope.row.ID, 'close')">封袋</el-button>
            <el-button v-if="scope.row.status === 'InBag' || scope.row.status === 'Exposed' || scope.row.status === 'Expired'" 
                       type="danger" size="small" @click="handleAction(scope.row.ID, 'bake')">烘烤</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" title="新增 MSD 记录" width="500px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="物料编码">
          <el-input v-model="form.material_code" />
        </el-form-item>
        <el-form-item label="批次/卷盘号">
          <el-input v-model="form.lot_no" />
        </el-form-item>
        <el-form-item label="MSL等级">
          <el-select v-model="form.msl_level" placeholder="请选择">
            <el-option label="MSL 1" value="1" />
            <el-option label="MSL 2" value="2" />
            <el-option label="MSL 2a" value="2a" />
            <el-option label="MSL 3" value="3" />
            <el-option label="MSL 4" value="4" />
            <el-option label="MSL 5" value="5" />
            <el-option label="MSL 5a" value="5a" />
            <el-option label="MSL 6" value="6" />
          </el-select>
        </el-form-item>
        <el-form-item label="所在仓库">
          <el-input v-model="form.warehouse" />
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
  material_code: '',
  lot_no: 'LOT' + Date.now(),
  msl_level: '3',
  warehouse: '电子料仓'
})

const statusText = (status) => {
  const maps = {
    'InBag': '袋装中',
    'Exposed': '暴露中',
    'Baking': '烘烤中',
    'Expired': '已过期'
  }
  return maps[status] || status
}

const statusType = (status) => {
  const maps = {
    'InBag': 'success',
    'Exposed': 'warning',
    'Baking': 'primary',
    'Expired': 'danger'
  }
  return maps[status] || 'info'
}

const fetchData = async () => {
  loading.value = true
  try {
    const res = await axios.get('/inventory/msd/')
    tableData.value = res
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const handleSubmit = async () => {
  try {
    await axios.post('/inventory/msd/', form)
    ElMessage.success('记录创建成功')
    dialogVisible.value = false
    fetchData()
  } catch (error) {
    console.error(error)
  }
}

const handleAction = async (id, action) => {
  try {
    await axios.post(`/inventory/msd/${action}/${id}`)
    ElMessage.success('操作成功')
    fetchData()
  } catch (error) {
    console.error(error)
  }
}

onMounted(fetchData)
</script>

<style scoped>
.text-danger {
  color: #F56C6C;
  font-weight: bold;
}
</style>

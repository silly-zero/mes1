<template>
  <div class="receipt">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>收料管理</span>
          <el-button type="primary" @click="dialogVisible = true">新增收料单</el-button>
        </div>
      </template>

      <el-table :data="tableData" border style="width: 100%" v-loading="loading" height="calc(100vh - 250px)">
        <el-table-column type="index" label="行号" width="60" fixed="left" />
        <el-table-column prop="receive_no" label="收料单编码" width="220" fixed="left">
          <template #default="scope">
            <span style="color: #E6A23C; cursor: pointer;">{{ scope.row.receive_no }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="receipt_type_code" label="收料类型编码" width="150" />
        <el-table-column prop="receipt_type_name" label="收料类型名称" width="150" />
        <el-table-column prop="supplier_code" label="供应商编码" width="150" />
        <el-table-column prop="supplier_name" label="供应商名称" width="200" />
        <el-table-column prop="customer_code" label="客户编码" width="150" />
        <el-table-column prop="customer_name" label="客户名称" width="200" />
        <el-table-column prop="asn" label="ASN" width="150" />
        <el-table-column prop="erp_order_type_code" label="ERP单据类型编码" width="180" />
        <el-table-column prop="erp_order_type_name" label="ERP单据类型名称" width="180" />
        <el-table-column prop="dept_code" label="部门编码" width="120" />
        <el-table-column prop="dept_name" label="部门名称" width="150" />
        <el-table-column prop="project_code" label="项目编码" width="150" />
        <el-table-column prop="project_name" label="项目名称" width="200" />
        <el-table-column prop="project_abbr" label="项目简称" width="120" />
        <el-table-column prop="is_urgent" label="加急" width="80" align="center">
          <template #default="scope">
            <el-tag :type="scope.row.is_urgent ? 'danger' : 'info'" size="small">
              {{ scope.row.is_urgent ? '是' : '否' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="收料单状态" width="120" align="center">
          <template #default="scope">
            <el-tag :type="scope.row.status === '已完成' ? 'success' : 'warning'" size="small">
              {{ scope.row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="source" label="来源" width="150" />
        <el-table-column prop="org_code" label="组织编码" width="150" />
        <el-table-column prop="org_name" label="组织名称" width="150" />
        <el-table-column prop="remark" label="备注" width="200" show-overflow-tooltip />
        <el-table-column prop="modifier_id" label="修改人" width="120" />
        <el-table-column prop="modifier_name" label="修改人名称" width="120" />
        <el-table-column prop="modify_time" label="修改时间" width="180" />
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" title="新增收料记录" width="600px">
      <el-form :model="form" label-width="120px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="收料单编码">
              <el-input v-model="form.receive_no" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="入库确认单号">
              <el-input v-model="form.inbound_no" placeholder="自动生成" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="收料类型">
              <el-select v-model="form.receipt_type_name" placeholder="请选择">
                <el-option label="采购单收料" value="采购单收料" />
                <el-option label="客供收料" value="客供收料" />
                <el-option label="客退收料" value="客退收料" />
                <el-option label="初始化收料" value="初始化收料" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="物料编码">
              <el-input v-model="form.material_code" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="收料数量">
              <el-input-number v-model="form.total_qty" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="供应商名称">
              <el-input v-model="form.supplier_name" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="组织名称">
              <el-input v-model="form.org_name" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="加急">
              <el-switch v-model="form.is_urgent" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="备注">
          <el-input v-model="form.remark" type="textarea" />
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
  receipt_type_name: '采购单收料',
  material_code: '',
  total_qty: 0,
  is_urgent: false,
  confirm_type: '来料入库确认',
  supplier_name: '',
  org_name: '默认组织',
  remark: '',
  status: '未完成',
  source: '手动创建'
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
    dialogVisible.value = false
    fetchData()
  } catch (error) {
    console.error(error)
  }
}

onMounted(fetchData)
</script>

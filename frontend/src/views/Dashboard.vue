<template>
  <div class="dashboard">
    <el-row :gutter="20">
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>今日入库单</template>
          <div class="card-value">12</div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>待检验批次</template>
          <div class="card-value">5</div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>库存预警</template>
          <div class="card-value color-danger">2</div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

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

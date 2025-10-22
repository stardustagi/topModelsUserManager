<template>
  <div class="dashboard-container">
    <el-card class="section-card">
      <div class="section-title">预算预警</div>

      <div v-for="item in budgets" :key="item.id" class="mb-6">
        <div class="flex justify-between items-center mb-2">
          <div>{{ item.name }}</div>
          <div>{{ item.percentage }}%</div>
        </div>
        <el-progress :percentage="item.percentage" :status="item.status" />
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
type ProgressStatus = '' | 'success' | 'warning' | 'exception'

interface BudgetItem {
  id: number
  name: string
  percentage: number
  status: ProgressStatus
}

const budgets: BudgetItem[] = [
  // { id: 1, name: '总预算使用情况', percentage: 90, status: 'warning' },
  // { id: 2, name: '分组A预算使用', percentage: 70, status: 'success' },
  // { id: 3, name: '分组B预算使用', percentage: 100, status: 'exception' }
]
</script>

<style scoped>
.dashboard-container {
  padding: 24px;
}
.section-card {
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  background: white;
  padding: 24px;
}
.section-title {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 16px;
}
</style>

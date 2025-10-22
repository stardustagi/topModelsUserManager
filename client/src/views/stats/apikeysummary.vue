<template>
  <div class="dashboard-container">
    <el-card class="section-card">
      <div class="section-title">多 APIKey 使用情况</div>
      <el-table :data="apiKeys" style="width: 100%" border>
        <el-table-column prop="key" label="APIKey" width="250" />
        <el-table-column prop="usage" label="已用 tokens" />
        <el-table-column prop="limit" label="总配额" />
        <el-table-column prop="calls" label="调用次数" />
        <el-table-column prop="status" label="状态" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
const apiKeys = [
  // { key: 'key-abc123', usage: '1.2M', limit: '2M', calls: 320, status: '正常' },
  // { key: 'key-def456', usage: '1.9M', limit: '2M', calls: 490, status: '预警' },
  // { key: 'key-ghi789', usage: '2.0M', limit: '2M', calls: 530, status: '超额' }
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

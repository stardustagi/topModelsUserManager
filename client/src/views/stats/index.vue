<template>
  <div class="dashboard-container">
    <!-- 实时仪表盘 -->
    <el-card class="section-card mb-4">
      <div class="section-title">实时仪表盘</div>
      <div class="stats-grid">
        <div v-for="item in realTimeStats" :key="item.label" class="stat-item">
          <div class="stat-label">{{ item.label }}</div>
          <div class="stat-value">{{ item.value }}</div>
        </div>
      </div>
    </el-card>

    <!-- 历史趋势 -->
    <el-card class="section-card mb-4">
      <template #header>
        <div class="chart-header">
          <h3>历史趋势</h3>
          <div class="chart-controls">
            <el-date-picker
              v-model="dateRange"
              type="daterange"
              size="small"
              range-separator="至"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
              @change="generateMockData"
            />
            <el-select
              v-model="granularity"
              size="small"
              class="chart-select"
              @change="generateMockData"
            >
              <el-option label="按天" value="day" />
              <el-option label="按周" value="week" />
              <el-option label="按月" value="month" />
            </el-select>
          </div>
        </div>
      </template>

      <div v-if="trendData.length > 0" class="chart-container">
        <v-chart class="chart" :option="chartOption" :key="chartKey" autoresize />
      </div>
      <div v-else class="chart-placeholder">
        <el-icon><TrendCharts /></el-icon>
        <p>暂无数据</p>
      </div>
    </el-card>

    <!-- 模型占比 -->
    <el-card class="section-card">
      <template #header>
        <div class="chart-header">
          <h3>模型使用占比</h3>
          <el-select
            v-model="pieType"
            size="small"
            class="chart-select"
            @change="generatePieData"
          >
            <el-option label="按 Token 占比" value="tokens" />
            <el-option label="按 费用 占比" value="cost" />
          </el-select>
        </div>
      </template>

      <div v-if="pieData.length > 0" class="chart-container">
        <v-chart class="chart" :option="pieOption" autoresize />
      </div>
      <div v-else class="chart-placeholder">
        <el-icon><TrendCharts /></el-icon>
        <p>暂无数据</p>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { TrendCharts } from '@element-plus/icons-vue'

// ================== 实时仪表盘 ==================
const realTimeStats = ref([
  // { label: '今日用量', value: '132,400 tokens' },
  // { label: '本月用量', value: '3,234,500 tokens' },
  // { label: '剩余配额', value: '1,765,500 tokens' },
  // { label: '累计费用', value: '$12.80' },
  // { label: '月度预算', value: '$20.00' },
  // { label: '调用次数', value: '456 次' }
   { label: '今日用量', value: '0 tokens' },
  { label: '本月用量', value: '0 tokens' },
  { label: '剩余配额', value: '0 tokens' },
  { label: '累计费用', value: '$0' },
  { label: '月度预算', value: '$0' },
  { label: '调用次数', value: '0 次' }
])

// ================== 历史趋势折线图 ==================
type DataPoint = {
  date: string
  tokens: number
  cost: number
}

const trendData = ref<DataPoint[]>([])
const dateRange = ref<[Date, Date]>([
  new Date('2025-09-01'),
  new Date('2025-10-01')
])
const granularity = ref<'day' | 'week' | 'month'>('day')

const chartKey = computed(
  () =>
    `${granularity.value}-${dateRange.value[0].toDateString()}-${dateRange.value[1].toDateString()}`
)

const generateMockData = () => {
  trendData.value = []
  // const start = dateRange.value[0]
  // const end = dateRange.value[1]
  // const result: DataPoint[] = []
  // const current = new Date(start)
  // const step = granularity.value === 'day' ? 1 : granularity.value === 'week' ? 7 : 30

  // while (current <= end) {
  //   const tokens = Math.floor(Math.random() * 50000) + 10000
  //   const cost = parseFloat((tokens / 10000 * 0.3).toFixed(2))
  //   result.push({
  //     date: current.toISOString().split('T')[0],
  //     tokens,
  //     cost
  //   })
  //   current.setDate(current.getDate() + step)
  // }
  // trendData.value = result
}
onMounted(() => {
  generateMockData()
})

const chartOption = computed(() => ({
  tooltip: {
    trigger: 'axis',
    backgroundColor: 'rgba(50,50,50,0.8)',
    textStyle: { color: '#fff' }
  },
  legend: {
    data: ['消耗 Tokens', '费用 ($)'],
    bottom: 0
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '12%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    data: trendData.value.map(i => i.date),
    axisLabel: { rotate: 40 }
  },
  yAxis: [
    { type: 'value', name: 'Tokens' },
    { type: 'value', name: '费用 ($)' }
  ],
  series: [
    {
      name: '消耗 Tokens',
      type: 'line',
      smooth: true,
      areaStyle: { color: 'rgba(64,158,255,0.2)' },
      data: trendData.value.map(i => i.tokens)
    },
    {
      name: '费用 ($)',
      type: 'line',
      smooth: true,
      yAxisIndex: 1,
      areaStyle: { color: 'rgba(103,194,58,0.2)' },
      data: trendData.value.map(i => i.cost)
    }
  ]
}))

// ================== 模型占比饼状图 ==================
type ModelData = {
  name: string
  tokens: number
  cost: number
}

const models = ref<ModelData[]>([])
const pieType = ref<'tokens' | 'cost'>('tokens')
const pieData = ref<{ name: string; value: number }[]>([])

// 生成模拟模型数据
const generatePieData = () => {
  // const modelNames = ['gpt-4-turbo', 'claude-3-sonnet', 'mistral-large', 'gemini-1.5-pro', 'gpt-4o-mini']
  const modelNames = []
  models.value = modelNames.map(name => ({
    name,
    tokens: Math.floor(Math.random() * 300000) + 100000,
    cost: parseFloat((Math.random() * 50 + 20).toFixed(2))
  }))

  pieData.value = models.value.map(m => ({
    name: m.name,
    value: pieType.value === 'tokens' ? m.tokens : m.cost
  }))
}

// 初次加载生成
onMounted(() => {
  generatePieData()
})

// 饼图配置
const pieOption = computed(() => ({
  tooltip: {
    trigger: 'item',
    formatter: '{b}: {c} ({d}%)'
  },
  legend: {
    orient: 'horizontal',
    bottom: 0
  },
  series: [
    {
      name: pieType.value === 'tokens' ? 'Token 占比' : '费用占比',
      type: 'pie',
      radius: ['30%', '70%'],
      avoidLabelOverlap: false,
      itemStyle: { borderRadius: 10, borderColor: '#fff', borderWidth: 2 },
      label: { show: true, formatter: '{b}\n{d}%' },
      emphasis: { label: { show: true, fontSize: 16, fontWeight: 'bold' } },
      data: pieData.value
    }
  ]
}))
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
  margin-bottom: 24px;
}
.section-title {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 16px;
}
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 16px;
}
.stat-item {
  background: #f8f9fa;
  padding: 12px;
  border-radius: 8px;
  text-align: center;
}
.stat-label {
  font-size: 13px;
  color: #666;
}
.stat-value {
  font-size: 16px;
  font-weight: bold;
  color: #333;
}
.chart-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.chart-controls {
  display: flex;
  gap: 12px;
  align-items: center;
}
.chart-container {
  height: 400px;
}
.chart {
  width: 100%;
  height: 100%;
}
.chart-placeholder {
  background: #f2f2f2;
  height: 280px;
  border-radius: 8px;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #999;
  flex-direction: column;
}
@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: 1fr;
  }
}
</style>

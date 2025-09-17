<template>
  <div class="dashboard-container">
    <!-- 顶部欢迎区域 -->
    <div class="welcome-section">
      <h1 class="welcome-title">早上好, {{ useStore.username }}</h1>
      <p class="welcome-subtitle">
        今天是{{ currentDate }}，祝您有愉快的一天！
      </p>
    </div>

    <!-- 数据卡片区域 -->
    <div class="stats-grid">
      <el-card class="stat-card" shadow="hover">
        <div class="card-content">
          <h3>账户数据</h3>
          <div class="card-content-one">
            <div class="left-part">
              <div class="card-icon blue-icon">
                <el-icon><Wallet /></el-icon>
              </div>
              <div class="stat-item">
                <span class="label">当前余额</span>
                <span class="value">$0.50</span>
              </div>
            </div>

            <el-button type="primary" tag="a">充值</el-button>
          </div>
          <div class="card-content-one">
            <div class="left-part">
              <div class="card-icon blue-icon">
                <el-icon><TrendCharts /></el-icon>
              </div>
              <div class="stat-item">
                <span class="label">历史消耗</span>
                <span class="value">$0.00</span>
              </div>
            </div>
          </div>
        </div>
      </el-card>

      <el-card class="stat-card" shadow="hover">
        <div class="card-content">
          <h3>使用统计</h3>
          <div class="card-content-one">
            <div class="left-part">
              <div class="card-icon green-icon">
                <el-icon><Connection /></el-icon>
              </div>
              <div class="stat-item">
                <span class="label">请求次数</span>
                <span class="value">0</span>
              </div>
            </div>
          </div>
          <div class="card-content-one">
            <div class="left-part">
              <div class="card-icon green-icon">
                <el-icon><DataAnalysis /></el-icon>
              </div>
              <div class="stat-item">
                <span class="label">统计次数</span>
                <span class="value">0</span>
              </div>
            </div>
          </div>
        </div>
      </el-card>

      <el-card class="stat-card" shadow="hover">
        <div class="card-content">
          <h3>资源消耗</h3>
          <div class="card-content-one">
            <div class="left-part">
              <div class="card-icon purple-icon">
                <el-icon><Ticket /></el-icon>
              </div>
              <div class="stat-item">
                <span class="label">统计额度</span>
                <span class="value">$0.00</span>
              </div>
            </div>
          </div>
          <div class="card-content-one">
            <div class="left-part">
              <div class="card-icon purple-icon">
                <el-icon><Cpu /></el-icon>
              </div>
              <div class="stat-item">
                <span class="label">统计Tokens</span>
                <span class="value">0</span>
              </div>
            </div>
          </div>
        </div>
      </el-card>

      <el-card class="stat-card" shadow="hover">
        <div class="card-content">
          <h3>性能指标</h3>
          <div class="card-content-one">
            <div class="left-part">
              <div class="card-icon orange-icon">
                <el-icon><Timer /></el-icon>
              </div>
              <div class="stat-item">
                <span class="label">平均RPM</span>
                <span class="value">0.000</span>
              </div>
            </div>
          </div>
          <div class="card-content-one">
            <div class="left-part">
              <div class="card-icon orange-icon">
                <el-icon><Odometer /></el-icon>
              </div>
              <div class="stat-item">
                <span class="label">平均TPM</span>
                <span class="value">0</span>
              </div>
            </div>
          </div>
        </div>
      </el-card>
    </div>

    <!-- 模型数据分析区域 -->
    <el-card class="chart-section" shadow="never">
      <template #header>
        <div class="chart-header">
          <h2>模型数据分析</h2>
          <div class="chart-tabs">
            <el-button
              :type="activeTab === 'consumption' ? 'primary' : ''"
              @click="activeTab = 'consumption'"
              >消耗分布</el-button
            >
            <el-button
              :type="activeTab === 'trend' ? 'primary' : ''"
              @click="activeTab = 'trend'"
              >消耗趋势</el-button
            >
            <el-button
              :type="activeTab === 'distribution' ? 'primary' : ''"
              @click="activeTab = 'distribution'"
              >调用次数分布</el-button
            >
            <el-button
              :type="activeTab === 'ranking' ? 'primary' : ''"
              @click="activeTab = 'ranking'"
              >调用次数排行</el-button
            >
          </div>
        </div>
      </template>

      <div class="chart-content">
        <h3>模型消耗分布</h3>
        <p class="total-text">总计：$0.00</p>

        <div class="chart-placeholder">
          <i class="fas fa-chart-pie"></i>
          <p>暂无数据</p>
        </div>
      </div>
    </el-card>

    <!-- API信息区域 -->
    <el-card class="api-section" shadow="never">
      <template #header>
        <h2>API信息</h2>
      </template>

      <div class="api-grid">
        <div v-for="(api, index) in apiInfo" :key="index" class="api-card">
          <div class="api-header">
            <span class="api-name">{{ api.name }}</span>
            <span
              class="api-status"
              :class="{
                'status-online': api.status === 'online',
                'status-offline': api.status !== 'online'
              }"
            >
              {{ api.status === "online" ? "在线" : "离线" }}
            </span>
          </div>
          <p class="api-url">{{ api.url }}</p>
          <p class="api-desc">{{ api.description }}</p>
          <div class="api-actions">
            <el-button size="small" @click="testSpeed(api)">测速</el-button>
            <el-button size="small" type="primary" @click="copyUrl(api.url)"
              >复制</el-button
            >
            <el-button size="small" @click="openUrl(api.url)">跳转</el-button>
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from "vue";
import { ElCard, ElButton, ElMessage } from "element-plus";
import { useUserStoreHook } from "@/store/modules/user";
import {
  Wallet,
  TrendCharts,
  Connection,
  DataAnalysis,
  Ticket,
  Cpu,
  Timer,
  Odometer
} from "@element-plus/icons-vue";

const useStore = useUserStoreHook();
const activeTab = ref("consumption");

const currentDate = computed(() => {
  const date = new Date();
  return date.toLocaleDateString("zh-CN", {
    year: "numeric",
    month: "long",
    day: "numeric",
    weekday: "long"
  });
});

const apiInfo = [
  {
    name: "海外主站 (推荐)",
    url: "https://api.wenwen-ai.com",
    description: "CF 集群，高防，国外快",
    status: "online"
  },
  {
    name: "国内主站",
    url: "https://api.wenwen-ai.com",
    description: "无界面，带查询明细，方便代理进行二次销售",
    status: "online"
  },
  {
    name: "海外线路(备用)",
    url: "https://proxy.wenwen-ai.com",
    description: "备用线路，稳定可靠",
    status: "online"
  },
  {
    name: "海外主站 (推荐)",
    url: "https://apiro.ai",
    description: "高性能API节点",
    status: "online"
  },
  {
    name: "国内线路",
    url: "https://proxy.wenwen-ai.com",
    description: "国内优化线路",
    status: "online"
  },
  {
    name: "海外线路(备用)",
    url: "https://aprox.ai.com",
    description: "全球覆盖，多区域备份",
    status: "offline"
  },
  {
    name: "海外主站 (推荐)",
    url: "https://api.wenwen-ai.com",
    description: "CF 集群，高防，国外快",
    status: "online"
  },
  {
    name: "国内主站",
    url: "https://api.wenwen-ai.com",
    description: "无界面，带查询明细，方便代理进行二次销售",
    status: "online"
  },
  {
    name: "海外线路(备用)",
    url: "https://proxy.wenwen-ai.com",
    description: "备用线路，稳定可靠",
    status: "online"
  },
  {
    name: "海外主站 (推荐)",
    url: "https://apiro.ai",
    description: "高性能API节点",
    status: "online"
  },
  {
    name: "国内线路",
    url: "https://proxy.wenwen-ai.com",
    description: "国内优化线路",
    status: "online"
  },
  {
    name: "海外线路(备用)",
    url: "https://aprox.ai.com",
    description: "全球覆盖，多区域备份",
    status: "offline"
  },
  {
    name: "海外主站 (推荐)",
    url: "https://api.wenwen-ai.com",
    description: "CF 集群，高防，国外快",
    status: "online"
  },
  {
    name: "国内主站",
    url: "https://api.wenwen-ai.com",
    description: "无界面，带查询明细，方便代理进行二次销售",
    status: "online"
  },
  {
    name: "海外线路(备用)",
    url: "https://proxy.wenwen-ai.com",
    description: "备用线路，稳定可靠",
    status: "online"
  },
  {
    name: "海外主站 (推荐)",
    url: "https://apiro.ai",
    description: "高性能API节点",
    status: "online"
  },
  {
    name: "国内线路",
    url: "https://proxy.wenwen-ai.com",
    description: "国内优化线路",
    status: "online"
  },
  {
    name: "海外线路(备用)",
    url: "https://aprox.ai.com",
    description: "全球覆盖，多区域备份",
    status: "offline"
  }
];

const testSpeed = (api: any) => {
  ElMessage.info(`正在测试 ${api.name} 的速度...`);
};

const copyUrl = (url: string) => {
  navigator.clipboard.writeText(url);
  ElMessage.success("已复制到剪贴板");
};

const openUrl = (url: string) => {
  window.open(url, "_blank");
};
</script>

<style scoped>
.dashboard-container {
  padding: 24px;
  background-color: #f5f7fa;
  min-height: 100vh;
}

.welcome-section {
  margin-bottom: 24px;
}

.welcome-title {
  font-size: 28px;
  font-weight: 600;
  color: #1f2f3d;
  margin-bottom: 8px;
}

.welcome-subtitle {
  font-size: 14px;
  color: #87909c;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 20px;
  margin-bottom: 24px;
}

.stat-card {
  border-radius: 12px;
  border: none;
  overflow: hidden;
  transition: transform 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-5px);
}

.stat-card :deep(.el-card__body) {
  display: flex;
  padding: 0;
  align-items: stretch;
  justify-content: flex-start;
  /* justify-content: center; */
}

.card-icon {
  width: 44px;
  height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 24px;
  /* height: 100%; */
}

/* 为不同卡片设置不同颜色 */
.blue-icon .el-icon {
  color: #3498db;
}

.green-icon .el-icon {
  color: #2ecc71;
}

.purple-icon .el-icon {
  color: #9b59b6;
}

.orange-icon .el-icon {
  color: #f39c12;
}

.card-content {
  flex: 1;
  width: 100%;
  padding: 20px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.card-content-one {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  gap: 12px;
}

.card-content-one > .el-button {
  white-space: nowrap;
}

.left-part {
  display: flex;
  align-items: center;
  gap: 20px;
  min-width: 0;
}

.card-content h3 {
  margin: 0 0 16px 0;
  font-size: 16px;
  font-weight: 600;
  color: #1f2f3d;
}

.stat-item {
  display: flex;
  flex-direction: column;
  /* justify-content: space-between; */
  justify-content: center;
  margin-bottom: 12px;
}

.stat-item .label {
  color: #87909c;
}

.stat-item .value {
  font-weight: 600;
  color: #1f2f3d;
}

.action-btn {
  width: 100%;
  margin-top: 12px;
}

.chart-section,
.api-section {
  border-radius: 12px;
  border: none;
  margin-bottom: 24px;
  background: white;
}

.chart-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.chart-header h2 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
}

.chart-content {
  padding: 20px 0;
}

.chart-content h3 {
  margin: 0 0 8px 0;
  font-size: 16px;
  font-weight: 600;
}

.total-text {
  margin: 0 0 20px 0;
  color: #87909c;
}

.chart-placeholder {
  height: 300px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: #f9fafc;
  border-radius: 8px;
  color: #c0c4cc;
}

.chart-placeholder i {
  font-size: 48px;
  margin-bottom: 12px;
}

.api-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
}

.api-card {
  border: 1px solid #e6e8eb;
  border-radius: 8px;
  padding: 16px;
  transition: all 0.3s ease;
}

.api-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transform: translateY(-2px);
}

.api-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.api-name {
  font-weight: 600;
  color: #1f2f3d;
}

.api-status {
  font-size: 12px;
  padding: 4px 8px;
  border-radius: 12px;
}

.status-online {
  background-color: #f0f9eb;
  color: #67c23a;
}

.status-offline {
  background-color: #fef0f0;
  color: #f56c6c;
}

.api-url {
  margin: 0 0 8px 0;
  color: #3498db;
  font-family: monospace;
  word-break: break-all;
}

.api-desc {
  margin: 0 0 12px 0;
  font-size: 13px;
  color: #87909c;
}

.api-actions {
  display: flex;
  gap: 8px;
}

@media (max-width: 768px) {
  .dashboard-container {
    padding: 16px;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }

  .chart-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .chart-tabs {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
  }

  .api-grid {
    grid-template-columns: 1fr;
  }
}
</style>

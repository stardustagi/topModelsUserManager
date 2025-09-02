<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-4">下午好, yuanmu</h1>
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
      <el-card class="w-full">
        <div class="flex items-center mb-2">
          <i class="el-icon-folder mr-2 text-xl"></i>
          <span>账户数据</span>
        </div>
        <div class="mb-2">
          <span>当前余额:</span>
          <span class="font-bold">$0.50</span>
        </div>
        <div>
          <span>历史消耗:</span>
          <span class="font-bold">$0.00</span>
        </div>
      </el-card>
      <el-card class="w-full">
        <div class="flex items-center mb-2">
          <i class="el-icon-chart-bar mr-2 text-xl"></i>
          <span>使用统计</span>
        </div>
        <div class="mb-2">
          <span>请求次数:</span>
          <span class="font-bold">0</span>
        </div>
        <div>
          <span>统计次数:</span>
          <span class="font-bold">0</span>
        </div>
      </el-card>
      <el-card class="w-full">
        <div class="flex items-center mb-2">
          <i class="el-icon-dollar-sign mr-2 text-xl"></i>
          <span>资源消耗</span>
        </div>
        <div class="mb-2">
          <span>统计额度:</span>
          <span class="font-bold">$0.00</span>
        </div>
        <div>
          <span>统计Tokens:</span>
          <span class="font-bold">0</span>
        </div>
      </el-card>
      <el-card class="w-full">
        <div class="flex items-center mb-2">
          <i class="el-icon-speedometer mr-2 text-xl"></i>
          <span>性能指标</span>
        </div>
        <div class="mb-2">
          <span>平均RPM:</span>
          <span class="font-bold">0.000</span>
        </div>
        <div>
          <span>平均TPM:</span>
          <span class="font-bold">0</span>
        </div>
      </el-card>
    </div>
    <div class="mt-4">
      <el-card class="w-full">
        <div class="flex justify-between items-center mb-2">
          <h2 class="text-lg font-bold">模型数据分析</h2>
          <div class="flex space-x-2">
            <el-button type="primary" plain>消耗分布</el-button>
            <el-button plain>消耗趋势</el-button>
            <el-button plain>调用次数分布</el-button>
            <el-button plain>调用次数排行</el-button>
          </div>
        </div>
        <div class="mb-2">
          <h3 class="font-bold">模型消耗分布</h3>
          <p>总计：$0.00</p>
        </div>
        <!-- 在这里添加你的图表代码 -->
      </el-card>
    </div>
    <div class="mt-4">
      <el-card class="w-full">
        <h2 class="text-lg font-bold mb-2">API信息</h2>
        <ul>
          <li v-for="api in apiInfo" :key="api.name" class="mb-2">
            <span>{{ api.name }}:</span>
            <a :href="api.url" target="_blank">{{ api.url }}</a>
          </li>
        </ul>
      </el-card>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { ElCard, ElButton } from 'element-plus';

export default defineComponent({
  components: {
    ElCard,
    ElButton,
  },
  data() {
    return {
      apiInfo: [
        { name: '国内线路', url: 'https://api.wenwen-ai.com' },
        { name: '海外线路', url: 'https://key.wenwen-ai.com' },
        { name: '国内备用', url: 'https://a.wenwen-ai.com' },
        { name: '新节点', url: 'https://apipro.ai' },
      ],
    };
  },
});
</script>

<style scoped>
/* 添加一些自定义样式 */
</style>
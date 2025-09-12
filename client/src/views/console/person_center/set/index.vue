<template>
  <div class="p-4">
    <!-- 用户信息 -->
    <el-card class="mb-4 bg-gray-100">
      <div class="flex items-center mb-2">
        <!-- <img src="@/assets/user-avatar.png" alt="User Avatar" class="w-12 h-12 rounded-full mr-2"> -->
        <div>
          <h2 class="text-xl font-bold">YU yuanmu</h2>
          <span class="text-sm text-gray-500">普通用户 ID: 2450</span>
        </div>
      </div>
      <div class="mt-4">
        <h3 class="text-lg font-bold">$0.50</h3>
        <div class="flex mt-2">
          <div class="mr-4">
            <span>历史消耗:</span>
            <span class="font-bold">$0.00</span>
          </div>
          <div class="mr-4">
            <span>请求次数:</span>
            <span class="font-bold">0</span>
          </div>
          <div>
            <span>用户分组:</span>
            <span class="font-bold">default</span>
          </div>
        </div>
      </div>
    </el-card>

    <!-- 导航标签 -->
    <el-tabs v-model="activeTab" class="mb-4">
      <el-tab-pane label="可用模型" name="models"></el-tab-pane>
      <!-- <el-tab-pane label="账户绑定" name="account"></el-tab-pane>
      <el-tab-pane label="安全设置" name="security"></el-tab-pane>
      <el-tab-pane label="其他设置" name="other"></el-tab-pane> -->
    </el-tabs>

    <!-- 模型列表 -->
    <div v-if="activeTab === 'models'">
      <h2 class="text-lg font-bold mb-2">模型列表</h2>
      <p class="text-sm text-gray-500">点击模型名称可复制</p>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-2 mt-4">
        <el-tag v-for="model in models" :key="model.name" type="info" class="cursor-pointer" @click="copyModelName(model.name)">
          {{ model.name }}
        </el-tag>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { ElCard, ElTabs, ElTabPane, ElTag, ElMessage } from 'element-plus';

export default defineComponent({
  components: {
    ElCard,
    ElTabs,
    ElTabPane,
    ElTag,
  },
  setup() {
    const activeTab = ref('models');
    const models = [
      // 添加你的模型列表数据
      { name: 'claude-3-5-haiku-20241022' },
      { name: 'claude-3-5-sonnet-20240620' },
      // 更多模型...
    ];

    const copyModelName = (name: string) => {
      navigator.clipboard.writeText(name);
      ElMessage.success(`已复制：${name}`)
    };

    return {
      activeTab,
      models,
      copyModelName,
    };
  },
});
</script>

<style scoped>
/* 添加一些自定义样式 */
</style>
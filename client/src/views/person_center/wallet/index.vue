<template>
  <div class="p-4">
    <!-- 页面标题 -->
    <el-card class="mb-4 bg-gray-100 flex items-center justify-between">
      <div class="flex items-center">
        <!-- <img src="@/assets/icon.png" alt="Icon" class="w-8 h-8 mr-2"> -->
        <div>
          <h2 class="text-xl font-bold">在线充值</h2>
          <span class="text-sm text-gray-500">快速方便的充值方式</span>
        </div>
      </div>
      <div>
        <span class="mr-2"><i class="el-icon-user"></i> yuanmu (普通用户)</span>
      </div>
    </el-card>

    <!-- 当前余额和历史消耗 -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
      <el-card class="bg-white p-4">
        <div class="text-center">
          <span>当前余额</span>
          <h2 class="text-2xl font-bold mt-2">$0.50</h2>
        </div>
      </el-card>
      <el-card class="bg-white p-4">
        <div class="text-center">
          <span>历史消耗</span>
          <h2 class="text-2xl font-bold mt-2">$0.00</h2>
        </div>
      </el-card>
    </div>

    <!-- 选择充值额度 -->
    <div class="mb-4">
      <h3 class="text-lg font-bold mb-2">选择充值额度</h3>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
        <el-card v-for="option in rechargeOptions" :key="option.value" class="bg-white p-4 cursor-pointer" @click="selectOption(option)">
          <div class="text-center">
            <span>{{ option.label }}</span>
            <p class="mt-2">{{ option.description }}</p>
          </div>
        </el-card>
      </div>
    </div>

    <!-- 自定义金额输入 -->
    <div class="mb-4">
      <h3 class="text-lg font-bold mb-2">或输入自定义金额</h3>
      <div class="flex justify-between items-center">
        <span>充值数量</span>
        <span>实付金额：{{ customAmount }} 元</span>
      </div>
      <el-input v-model="customAmount" type="number" placeholder="请输入充值金额" class="w-full mt-2"></el-input>
    </div>

    <!-- 确定按钮 -->
    <div class="text-center">
      <el-button type="primary" @click="confirmRecharge">确定</el-button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { ElCard, ElInput, ElButton } from 'element-plus';

export default defineComponent({
  components: {
    ElCard,
    ElInput,
    ElButton,
  },
  setup() {
    const customAmount = ref('60');
    const selectedOption = ref(null);

    const rechargeOptions = [
      { label: '5', description: '实付 ¥25.00', value: 5 },
      { label: '10', description: '实付 ¥50.00', value: 10 },
      { label: '30', description: '实付 ¥150.00', value: 30 },
      { label: '50', description: '实付 ¥250.00', value: 50 },
      { label: '100', description: '实付 ¥500.00', value: 100 },
      { label: '300', description: '实付 ¥1500.00', value: 300 },
      { label: '500', description: '实付 ¥2500.00', value: 500 },
      { label: '1000', description: '实付 ¥5000.00', value: 1000 },
    ];

    const selectOption = (option) => {
      selectedOption.value = option;
      customAmount.value = option.value.toString();
    };

    const confirmRecharge = () => {
      if (selectedOption.value) {
        alert(`您选择了充值 ${selectedOption.value.value}，实付金额为 ${selectedOption.value.description}`);
      } else {
        alert(`您选择了自定义充值金额 ${customAmount.value} 元`);
      }
    };

    return {
      customAmount,
      selectedOption,
      rechargeOptions,
      selectOption,
      confirmRecharge,
    };
  },
});
</script>

<style scoped>
/* 添加一些自定义样式 */
</style>
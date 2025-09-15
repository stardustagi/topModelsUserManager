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
        <el-card
          v-for="option in rechargeOptions"
          :key="option.value"
          class="bg-white p-4 cursor-pointer"
          @click="selectOption(option)"
        >
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
      <el-input
        v-model="customAmount"
        type="number"
        placeholder="请输入充值金额"
        class="w-full mt-2"
      ></el-input>
    </div>

    <div class="mb-4">
      <!-- <h3 class="text-lg font-bold mb-2"></h3> -->
      <div class="flex justify-between items-center">
        <span>用户ID</span>
      </div>
      <el-input
        v-model="rechargeUserId"
        type="number"
        placeholder="用户ID"
        class="w-full mt-2"
      ></el-input>
    </div>

    <!-- 确定按钮 -->
    <div class="text-center">
      <el-button type="primary" @click="confirmRecharge">确定</el-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { ElCard, ElInput, ElButton } from "element-plus";
import { adminPaymentApi } from "@/api/managerApi";

const customAmount = ref("60");
const rechargeUserId = ref("");
const selectedOption = ref(null);

const rechargeOptions = [
  { label: "5", description: "实付 ¥25.00", value: 5, money: 2500 },
  { label: "10", description: "实付 ¥50.00", value: 10, money: 5000 },
  { label: "30", description: "实付 ¥150.00", value: 30, money: 15000 },
  { label: "50", description: "实付 ¥250.00", value: 50, money: 25000 },
  { label: "100", description: "实付 ¥500.00", value: 100, money: 50000 },
  { label: "300", description: "实付 ¥1500.00", value: 300, money: 150000 },
  { label: "500", description: "实付 ¥2500.00", value: 500, money: 250000 },
  { label: "1000", description: "实付 ¥5000.00", value: 1000, money: 500000 }
];

const selectOption = option => {
  selectedOption.value = option;
  customAmount.value = option.value.toString();
};

const confirmRecharge = async () => {
  let money = 0;
  if (selectedOption.value) {
    money = selectedOption.value.money;
    // alert(
    //   `您选择了充值 ${selectedOption.value.value}，实付金额为 ${selectedOption.value.description}`
    // );
  } else {
    money = Number(customAmount.value) * 100;
    // alert(`您选择了自定义充值金额 ${customAmount.value} 元`);
  }
  let data = {
    user_id: Number(rechargeUserId.value),
    amount: money,
    reason: "充值"
  };
  const res = await adminPaymentApi(data);
  if (res.errcode === 0) {
    console.log("充值成功");
  }
};
</script>

<style scoped>
/* 添加一些自定义样式 */
</style>

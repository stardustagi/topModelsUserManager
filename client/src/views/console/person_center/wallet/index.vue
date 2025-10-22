<template>
  <div class="recharge-container">
    <!-- 页面标题区域 -->
    <div class="header-section">
      <div class="header-content">
        <div class="title-group">
          <h1 class="main-title">在线充值</h1>
          <p class="subtitle">快速方便的充值方式</p>
        </div>
        <div class="user-info">
          <span class="user-tag"
            ><i class="el-icon-user"></i> yuanmu (普通用户)</span
          >
        </div>
      </div>
    </div>

    <!-- 余额信息卡片 -->
    <div class="balance-cards">
      <div class="balance-card blue-card">
        <div class="card-icon">
          <i class="el-icon-wallet"></i>
        </div>
        <div class="card-content">
          <span class="card-label">当前余额</span>
          <h2 class="card-value">$0.50</h2>
        </div>
      </div>

      <div class="balance-card green-card">
        <div class="card-icon">
          <i class="el-icon-data-analysis"></i>
        </div>
        <div class="card-content">
          <span class="card-label">历史消耗</span>
          <h2 class="card-value">$0.00</h2>
        </div>
      </div>
    </div>

    <!-- 充值选项区域 -->
    <div class="recharge-section">
      <h3 class="section-title">选择充值额度</h3>
      <div class="recharge-options">
        <div
          v-for="option in rechargeOptions"
          :key="option.value"
          class="option-card"
          :class="{ 'option-selected': selectedOption?.value === option.value }"
          @click="selectOption(option)"
        >
          <div class="option-main">
            <span class="option-amount">{{ option.label }}</span>
            <span class="option-currency">额度</span>
          </div>
          <p class="option-price">{{ option.description }}</p>
        </div>
      </div>
    </div>

    <!-- 自定义金额区域 -->
    <div class="custom-amount-section">
      <h3 class="section-title">或输入自定义金额</h3>
      <div class="amount-input-group">
        <div class="input-header">
          <span class="input-label">充值数量</span>
          <span class="actual-amount">实付金额：{{ customAmount }} 元</span>
        </div>
        <el-input
          v-model="customAmount"
          type="number"
          placeholder="请输入充值金额"
          class="custom-input"
        >
          <template #prefix>
            <span class="input-prefix">¥</span>
          </template>
        </el-input>
      </div>
    </div>

    <!-- 用户ID输入区域 -->
    <div class="user-id-section">
      <div class="input-header">
        <span class="input-label">用户ID</span>
      </div>
      <el-input
        v-model="rechargeUserId"
        type="number"
        placeholder="请输入用户ID"
        class="user-id-input"
      />
    </div>

    <!-- 操作按钮 -->
    <div class="action-buttons">
      <el-button type="primary" class="confirm-btn" @click="confirmRecharge">
        确认充值
      </el-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { ElInput, ElButton } from "element-plus";
// import { adminPaymentApi } from "@/api/managerApi";

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

const selectOption = (option: any) => {
  selectedOption.value = option;
  customAmount.value = option.value.toString();
};

const confirmRecharge = async () => {
  // let money = 0;
  // if (selectedOption.value) {
  //   money = selectedOption.value.money;
  // } else {
  //   money = Number(customAmount.value) * 100;
  // }

  // const data = {
  //   user_id: Number(rechargeUserId.value),
  //   amount: money,
  //   reason: "充值"
  // };

  // const res = await adminPaymentApi(data);
  // if (res.errcode === 0) {
  //   console.log("充值成功");
  // }
};
</script>

<style scoped>
.recharge-container {
  padding: 24px;
  background-color: #f5f7fa;
  min-height: 100vh;
}

/* 头部样式 */
.header-section {
  background: white;
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.title-group .main-title {
  font-size: 24px;
  font-weight: 600;
  color: #1f2f3d;
  margin: 0 0 8px 0;
}

.title-group .subtitle {
  font-size: 14px;
  color: #87909c;
  margin: 0;
}

.user-info .user-tag {
  background: #f0f9eb;
  color: #67c23a;
  padding: 8px 16px;
  border-radius: 20px;
  font-size: 14px;
}

/* 余额卡片样式 */
.balance-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-bottom: 32px;
}

.balance-card {
  background: white;
  border-radius: 12px;
  padding: 24px;
  display: flex;
  align-items: center;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.blue-card {
  border-left: 4px solid #3498db;
}

.green-card {
  border-left: 4px solid #2ecc71;
}

.card-icon {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 16px;
  font-size: 20px;
}

.blue-card .card-icon {
  background: #e3f2fd;
  color: #3498db;
}

.green-card .card-icon {
  background: #e8f5e8;
  color: #2ecc71;
}

.card-label {
  font-size: 14px;
  color: #87909c;
  display: block;
  margin-bottom: 4px;
}

.card-value {
  font-size: 24px;
  font-weight: 600;
  color: #1f2f3d;
  margin: 0;
}

/* 充值选项样式 */
.recharge-section {
  background: white;
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 24px;
}

.section-title {
  font-size: 18px;
  font-weight: 600;
  color: #1f2f3d;
  margin: 0 0 20px 0;
}

.recharge-options {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  gap: 16px;
}

.option-card {
  border: 2px solid #e6e8eb;
  border-radius: 8px;
  padding: 16px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
}

.option-card:hover {
  border-color: #3498db;
  transform: translateY(-2px);
}

.option-selected {
  border-color: #3498db;
  background: #f0f9ff;
}

.option-main {
  margin-bottom: 8px;
}

.option-amount {
  font-size: 20px;
  font-weight: 600;
  color: #1f2f3d;
}

.option-currency {
  font-size: 12px;
  color: #87909c;
  margin-left: 4px;
}

.option-price {
  font-size: 12px;
  color: #3498db;
  margin: 0;
}

/* 自定义金额样式 */
.custom-amount-section {
  background: white;
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 24px;
}

.amount-input-group {
  margin-top: 16px;
}

.input-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.input-label {
  font-size: 14px;
  color: #1f2f3d;
  font-weight: 500;
}

.actual-amount {
  font-size: 14px;
  color: #3498db;
  font-weight: 500;
}

.custom-input :deep(.el-input__inner) {
  height: 48px;
  font-size: 16px;
}

.input-prefix {
  color: #87909c;
  font-weight: 500;
}

/* 用户ID输入样式 */
.user-id-section {
  background: white;
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 32px;
}

.user-id-input :deep(.el-input__inner) {
  height: 48px;
}

/* 按钮样式 */
.action-buttons {
  text-align: center;
}

.confirm-btn {
  width: 200px;
  height: 48px;
  font-size: 16px;
  font-weight: 500;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .recharge-container {
    padding: 16px;
  }

  .header-content {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .balance-cards {
    grid-template-columns: 1fr;
  }

  .recharge-options {
    grid-template-columns: repeat(2, 1fr);
  }

  .confirm-btn {
    width: 100%;
  }
}
</style>

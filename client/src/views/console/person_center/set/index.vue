<template>
  <div class="dashboard-container">
    <!-- 用户信息 -->
    <el-card class="user-info-card">
      <div class="user-header">
        <div class="user-main-info">
          <h2 class="user-name">YU yuanmu</h2>
          <span class="user-id">普通用户 ID: 2450</span>
        </div>
        <el-button
          type="primary"
          size="small"
          class="bind-company-btn"
          @click="showCompanyDialog = true"
        >
          绑定公司
        </el-button>
      </div>

      <div class="balance-section">
        <h3 class="balance-amount">$0.50</h3>
        <div class="stats-grid">
          <div class="stat-item">
            <span class="stat-label">历史消耗:</span>
            <span class="stat-value">$0.00</span>
          </div>
          <div class="stat-item">
            <span class="stat-label">请求次数:</span>
            <span class="stat-value">0</span>
          </div>
          <div class="stat-item">
            <span class="stat-label">用户分组:</span>
            <span class="stat-value">default</span>
          </div>
        </div>
      </div>
    </el-card>

    <!-- 公司信息弹窗 -->
    <el-dialog
      v-model="showCompanyDialog"
      title="绑定公司信息"
      width="500px"
      class="company-dialog"
    >
      <el-form :model="companyForm" label-width="100px">
        <el-form-item label="公司名字" required>
          <el-input
            v-model="companyForm.company_name"
            placeholder="请输入公司全称"
          />
        </el-form-item>
        <el-form-item label="公司地址" required>
          <el-input
            v-model="companyForm.address"
            placeholder="请输入公司详细地址"
          />
        </el-form-item>
        <el-form-item label="营业执照号" required>
          <el-input
            v-model="companyForm.license_id"
            placeholder="请输入营业执照编号"
          />
        </el-form-item>
        <el-form-item label="联系电话" required>
          <el-input
            v-model="companyForm.phone"
            placeholder="请输入公司联系电话"
          />
        </el-form-item>
        <el-form-item label="电子邮箱" required>
          <el-input
            v-model="companyForm.email"
            placeholder="请输入公司联系邮箱"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showCompanyDialog = false">取消</el-button>
        <el-button type="primary" @click="submitCompanyInfo"
          >确认绑定</el-button
        >
      </template>
    </el-dialog>

    <!-- 导航标签 -->
    <el-tabs v-model="activeTab" class="main-tabs">
      <el-tab-pane label="可用模型" name="models"></el-tab-pane>
    </el-tabs>

    <!-- 模型列表 -->
    <div v-if="activeTab === 'models'" class="models-section">
      <div class="section-header">
        <h2 class="section-title">模型列表</h2>
        <p class="section-subtitle">点击模型名称可复制</p>
      </div>

      <div class="models-grid">
        <div
          v-for="model in models"
          :key="model.name"
          class="model-card"
          @click="copyModelName(model.name)"
        >
          <div class="model-icon">
            <i class="el-icon-cpu"></i>
          </div>
          <div class="model-info">
            <h3 class="model-name">{{ model.name }}</h3>
            <span class="model-status">可用</span>
          </div>
          <div class="copy-hint">
            <i class="el-icon-document-copy"></i>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import {
  ElCard,
  ElTabs,
  ElTabPane,
  ElTag,
  ElMessage,
  ElDialog,
  ElForm,
  ElFormItem,
  ElInput,
  ElButton
} from "element-plus";
import { bindCompanyInfoApi, getCompanyInfoApi } from "@/api/managerApi";
import { useUserStoreHook } from "@/store/modules/user";

const activeTab = ref("models");
const showCompanyDialog = ref(false);
const companyForm = ref({
  company_name: "",
  address: "",
  license_id: "",
  phone: "",
  email: ""
});

const models = [
  { name: "claude-3-5-haiku-20241022" },
  { name: "claude-3-5-sonnet-20240620" }
  // 更多模型...
];

const copyModelName = (name: string) => {
  navigator.clipboard.writeText(name);
  ElMessage.success(`已复制：${name}`);
};

onMounted(() => {
  getCompanyInfo();
});

const getCompanyInfo = async () => {
  const res = await getCompanyInfoApi({});
  console.log("公司信息=", res);
};

const submitCompanyInfo = async () => {
  if (
    !companyForm.value.company_name ||
    !companyForm.value.address ||
    !companyForm.value.license_id ||
    !companyForm.value.phone ||
    !companyForm.value.email
  ) {
    ElMessage.error("请填写所有必填字段");
    return;
  }

  let data = {
    user_id: Number(useUserStoreHook().userId),
    company_name: companyForm.value.company_name,
    license_id: companyForm.value.license_id,
    address: companyForm.value.address,
    phone: companyForm.value.phone,
    email: companyForm.value.email
  };
  const res = await bindCompanyInfoApi(data);
  if (res.errcode === 0) {
    console.log("绑定公司成功");
  }

  showCompanyDialog.value = false;
};
</script>

<style scoped>
.dashboard-container {
  padding: 24px;
  background: linear-gradient(135deg, #f5f7fa 0%, #e4e7ed 100%);
  min-height: 100vh;
}

/* 用户信息卡片 */
.user-info-card {
  border-radius: 12px;
  border: none;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  background: white;
  margin-bottom: 24px;
  padding: 24px;
}

.user-header {
  display: flex;
  justify-content: between;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #e8e8e8;
}

.user-main-info {
  flex: 1;
}

.user-name {
  font-size: 24px;
  font-weight: 600;
  color: #1f2f3d;
  margin: 0 0 8px 0;
}

.user-id {
  font-size: 14px;
  color: #87909c;
}

.bind-company-btn {
  background: linear-gradient(45deg, #409eff, #66b1ff);
  border: none;
  border-radius: 8px;
  padding: 10px 20px;
  font-weight: 500;
}

.bind-company-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(64, 158, 255, 0.3);
}

/* 余额部分 */
.balance-section {
  text-align: center;
}

.balance-amount {
  font-size: 32px;
  font-weight: 700;
  color: #67c23a;
  margin: 0 0 20px 0;
  text-shadow: 0 2px 4px rgba(103, 194, 58, 0.2);
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  gap: 16px;
  margin-top: 16px;
}

.stat-item {
  background: #f8f9fa;
  padding: 12px;
  border-radius: 8px;
  text-align: center;
}

.stat-label {
  display: block;
  font-size: 12px;
  color: #87909c;
  margin-bottom: 4px;
}

.stat-value {
  display: block;
  font-size: 14px;
  font-weight: 600;
  color: #1f2f3d;
}

/* 标签页 */
.main-tabs {
  margin-bottom: 24px;
}

.main-tabs :deep(.el-tabs__header) {
  margin: 0;
}

.main-tabs :deep(.el-tabs__nav-wrap::after) {
  height: 1px;
}

.main-tabs :deep(.el-tabs__item) {
  font-size: 16px;
  font-weight: 500;
  padding: 0 24px;
}

/* 模型部分 */
.models-section {
  background: white;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.section-header {
  margin-bottom: 24px;
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  color: #1f2f3d;
  margin: 0 0 8px 0;
}

.section-subtitle {
  font-size: 14px;
  color: #87909c;
  margin: 0;
}

.models-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 16px;
}

.model-card {
  display: flex;
  align-items: center;
  padding: 16px;
  border: 1px solid #e8e8e8;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  background: #fafbfc;
}

.model-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-color: #409eff;
  background: #f0f7ff;
}

.model-icon {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  background: linear-gradient(45deg, #409eff, #66b1ff);
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 16px;
  color: white;
  font-size: 18px;
}

.model-info {
  flex: 1;
}

.model-name {
  font-size: 14px;
  font-weight: 600;
  color: #1f2f3d;
  margin: 0 0 4px 0;
}

.model-status {
  font-size: 12px;
  color: #67c23a;
  background: #f0f9eb;
  padding: 2px 8px;
  border-radius: 4px;
}

.copy-hint {
  color: #409eff;
  font-size: 16px;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.model-card:hover .copy-hint {
  opacity: 1;
}

/* 弹窗样式 */
.company-dialog :deep(.el-dialog) {
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
}

.company-dialog :deep(.el-dialog__header) {
  border-bottom: 1px solid #e8e8e8;
  margin: 0;
  padding: 20px;
}

.company-dialog :deep(.el-dialog__title) {
  font-size: 18px;
  font-weight: 600;
  color: #1f2f3d;
}

.company-dialog :deep(.el-dialog__body) {
  padding: 20px;
}

.company-dialog :deep(.el-dialog__footer) {
  border-top: 1px solid #e8e8e8;
  padding: 16px 20px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .dashboard-container {
    padding: 16px;
  }

  .user-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }

  .models-grid {
    grid-template-columns: 1fr;
  }

  .model-card {
    flex-direction: column;
    text-align: center;
    gap: 12px;
  }

  .model-icon {
    margin-right: 0;
  }
}
</style>

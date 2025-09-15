<template>
  <div class="p-4">
    <!-- 用户信息 -->
    <el-card class="mb-4 bg-gray-100">
      <div class="flex items-center justify-between mb-2">
        <div class="flex items-center">
          <!--  -->
          <div>
            <h2 class="text-xl font-bold">YU yuanmu</h2>
            <span class="text-sm text-gray-500">普通用户 ID: 2450</span>
          </div>
        </div>
        <el-button
          type="primary"
          size="small"
          @click="showCompanyDialog = true"
        >
          绑定公司
        </el-button>
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

    <!-- 公司信息弹窗 -->
    <el-dialog v-model="showCompanyDialog" title="绑定公司信息" width="500px">
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
        <el-tag
          v-for="model in models"
          :key="model.name"
          type="info"
          class="cursor-pointer"
          @click="copyModelName(model.name)"
        >
          {{ model.name }}
        </el-tag>
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
  // 这里可以添加表单验证
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

  // 这里可以添加API调用
  console.log("提交公司信息:", companyForm.value);
  // ElMessage.success("公司信息绑定成功");

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
/* 添加一些自定义样式 */
</style>

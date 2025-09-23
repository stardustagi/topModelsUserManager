<template>
  <div class="dashboard-container">
    <div class="p-4">
      <!-- 顶部操作区 -->
      <div class="flex items-center gap-2 mb-4">
        <el-button type="primary" @click="dialogVisible = true"
          >添加令牌</el-button
        >
        <!-- <el-button>复制所选令牌</el-button> -->
        <!-- <el-button type="danger">删除所选令牌</el-button> -->

        <!-- <el-input
        v-model="search.keyword"
        placeholder="搜索关键字"
        clearable
        class="w-60 ml-auto"
      /> -->
        <!-- <el-input
        v-model="search.secret"
        placeholder="密钥"
        clearable
        class="w-60"
      /> -->
        <!-- <el-button type="primary" @click="handleSearch">查询</el-button> -->
        <!-- <el-button @click="resetSearch">重置</el-button> -->
      </div>

      <!-- 表格 -->
      <el-table
        :data="tableData"
        border
        style="width: 100%"
        v-loading="loading"
      >
        <!-- <el-table-column type="selection" width="55" /> -->
        <el-table-column prop="name" label="名称" min-width="120" />
        <el-table-column prop="status" label="状态" min-width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '启用' ? 'success' : 'info'">
              {{ row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <!-- <el-table-column prop="quota" label="剩余额度/总额度" min-width="150" /> -->
        <!-- <el-table-column prop="group" label="分组" min-width="120">
          <template #default="{ row }">
            <el-tag type="warning">{{ row.group }}</el-tag>
          </template>
        </el-table-column> -->
        <el-table-column prop="accessKey" label="密钥" min-width="200" />
        <el-table-column prop="model" label="可用模型" min-width="120" />
        <el-table-column label="操作" fixed="right" width="200">
          <template #default="{ row }">
            <el-button type="danger" size="small" @click="delAksk(row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 添加令牌对话框（无输入框） -->
    <el-dialog v-model="dialogVisible" title="添加令牌" width="400px">
      <div class="text-center py-6">确定要添加一个新的令牌吗？</div>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleAdd">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import {
  nodeUserAddAccessKeyAndSecurityKeyApi,
  nodeUserDelAkSkApi,
  nodeUserGetAkSkApi
} from "@/api/managerApi";
import { useUserStoreHook } from "@/store/modules/user";
import { onMounted, ref } from "vue";

const dialogVisible = ref(false);
const loading = ref(false);

// const search = ref({
//   keyword: "",
//   secret: ""
// });

const tableData = ref([
  // {
  //   name: "yuanmu的初始令牌",
  //   status: "启用",
  //   quota: "无限额度",
  //   group: "智能测试",
  //   secret: "sk-iktN**********gA7h",
  //   model: "无限制"
  // }
]);

onMounted(async () => {
  getAksk();
});

// 添加令牌（调用接口时由服务器返回实际数据）
const handleAdd = async () => {
  const res = await nodeUserAddAccessKeyAndSecurityKeyApi();
  if (res.errcode === 0) {
    console.log(res.data);
    getAksk();
  }
  dialogVisible.value = false;
};

// 获取aksk
const getAksk = async () => {
  const uid = useUserStoreHook().userId;
  const res = await nodeUserGetAkSkApi({ user_id: Number(uid) });
  if (res.errcode === 0) {
    tableData.value = [];
    for (let i = 0; i < res.data.keys.length; i++) {
      const item = res.data.keys[i];
      tableData.value.push({
        name: `节点 ${item.node_id}`,
        status: "启用",
        accessKey: `${item.access_key}`,
        model: "无限制",
        secretKey: `${item.secret_key}`,
        nodeId: `${item.node_id}`
      });
    }
  }
};

// 删除aksk
const delAksk = async row => {
  const uid = useUserStoreHook().userId;
  const res = await nodeUserDelAkSkApi({
    user_id: Number(uid),
    node_id: row.nodeId,
    ak: row.accessKey,
    sk: row.secretKey
  });
  if (res.errcode === 0) {
    getAksk();
    console.log(res.data);
  }
};
</script>

<style scoped>
.dashboard-container {
  padding: 24px;
  background: linear-gradient(135deg, #f5f7fa 0%, #e4e7ed 100%);
  min-height: 100vh;
}
</style>

<template>
  <el-main class="key-management-container">
    <!-- 标题和按钮区域 -->
    <div class="header">
      <h2>我的KEY</h2>
      <el-button
        type="primary"
        size="large"
        class="add-key-btn"
        @click="addNewKey"
      >
        <el-icon><Plus /></el-icon> 新增Key
      </el-button>
    </div>

    <!-- 提示文本 -->
    <div class="tips-text">
      将本地 Key 复制到本机其它应用中使用，Key 只可用于本机
    </div>

    <!-- Endpoint URL 区域 -->
    <div class="endpoint-section">
      <div class="endpoint-label">
        <span class="endpoint-label">Endpoint URL：</span>
        <span class="endpoint-url">http://localhost:8833</span>
      </div>
      <el-tooltip content="复制Endpoint" placement="top">
        <el-button
          type="primary"
          class="copy-endpoint-btn"
          @click="copyEndpoint"
        >
          <el-icon><DocumentCopy /></el-icon>
          复制
        </el-button>
      </el-tooltip>
    </div>

    <!-- 表格区域 -->
    <div class="table-container">
      <el-table
        :data="apiKeysStore.apiKeys"
        style="width: 100%"
        class="keys-table"
      >
        <el-table-column
          prop="name"
          label="Key名称"
          width="180"
        ></el-table-column>

        <el-table-column prop="status" label="状态" width="120">
          <template #default="{ row }">
            <el-tag
              :type="row.status === '1' ? 'success' : 'danger'"
              size="small"
              effect="dark"
            >
              {{ row.status === "1" ? "启用" : "禁用" }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column prop="apiKey" label="API Key">
          <template #default="{ row }">
            <div class="api-key-container">
              <span class="api-key-text">{{ maskApiKey(row.apiKey) }}</span>
              <el-tooltip content="复制Key" placement="top">
                <el-button
                  type="text"
                  class="copy-key-btn"
                  @click="copyText(row.apiKey)"
                >
                  <el-icon><DocumentCopy /></el-icon>
                </el-button>
              </el-tooltip>
            </div>
          </template>
        </el-table-column>

        <el-table-column
          prop="createTime"
          label="创建时间"
          width="180"
        ></el-table-column>

        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <div class="action-buttons">
              <el-tooltip
                :content="row.status === '1' ? '禁用此Key' : '启用此Key'"
                placement="top"
              >
                <el-switch
                  :model-value="row.status === '1'"
                  inline-prompt
                  active-text="启用"
                  inactive-text="禁用"
                  @change="handleStatusChange(row, $event)"
                  class="status-switch"
                ></el-switch>
              </el-tooltip>

              <el-tooltip content="删除Key" placement="top">
                <el-button
                  type="danger"
                  size="small"
                  class="delete-btn"
                  @click="deleteKey(row)"
                >
                  <el-icon><Delete /></el-icon>
                </el-button>
              </el-tooltip>

              <el-tooltip content="复制Key" placement="top">
                <el-button
                  type="text"
                  class="action-btn"
                  @click="copyAllText(row)"
                >
                  <el-icon><DocumentCopy /></el-icon>
                </el-button>
              </el-tooltip>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- <el-table :data="apiKeysStore.apiKeys" style="width: 100%">
      <el-table-column
        prop="name"
        label="Key名称"
        width="180"
      ></el-table-column>

      <el-table-column prop="status" label="状态" width="120">
        <template #default="{ row }">
          <el-tag
            :type="row.status === '1' ? 'success' : 'danger'"
            size="small"
          >
            {{ row.status === "1" ? "启用" : "禁用" }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column prop="apiKey" label="API Key">
        <template #default="{ row }">
          <span class="api-key-text">{{ maskApiKey(row.apiKey) }}</span>
        </template>
      </el-table-column>

      <el-table-column
        prop="createTime"
        label="创建时间"
        width="180"
      ></el-table-column>

      <el-table-column label="操作" width="220">
        <template #default="{ row }">
          <el-tooltip
            :content="row.status === '1' ? '禁用此Key' : '启用此Key'"
            placement="top"
          >
            <el-switch
              :model-value="row.status === '1'"
              inline-prompt
              active-text="启用"
              inactive-text="禁用"
              @change="handleStatusChange(row, $event)"
              style="margin-right: 12px"
            ></el-switch>
          </el-tooltip>

          <el-tooltip content="复制Key" placement="top">
            <el-button
              type="text"
              class="action-btn"
              @click="copyText(row.apiKey)"
            >
              <el-icon><DocumentCopy /></el-icon>
            </el-button>
          </el-tooltip>

          <el-tooltip content="删除Key" placement="top">
            <el-button type="text" class="action-btn" @click="deleteKey(row)">
              <el-icon><Delete /></el-icon>
            </el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table> -->

    <!-- 分页 -->
    <div class="pagination">
      <div class="page-info">共{{ total }}条</div>
    </div>
  </el-main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { DocumentCopy, Delete, Plus } from "@element-plus/icons-vue";
import { getModelKeysApi, setModelKeysApi } from "@/api/managerApi";
import { useMyKeysStore, type MyKeyItem } from "@/store/myKeyStore";

const apiKeysStore = useMyKeysStore();

function isChineseEnglishOrNumber(text: string): boolean {
  return /^[\u4e00-\u9fa5a-zA-Z0-9]+$/.test(text);
}

onMounted(() => {
  loadApikeys();
});

const loadApikeys = async () => {
  if (apiKeysStore.keysCount === 0) {
    const res = await getModelKeysApi();
    console.log("服务拉取apikey=",res);
    let keys: MyKeyItem[] = apiKeysStore.convertToMyKeyItems(res.data.api_keys);
    apiKeysStore.setApiKeys(keys);
  }
};

const total = computed(() => apiKeysStore.keysCount);

// 复制文本函数
const copyText = async (text: string) => {
  console.log("text ===== ", text);
  try {
    if (navigator.clipboard && navigator.clipboard.writeText) {
      await navigator.clipboard.writeText(text);
      ElMessage.success("复制成功");
    } else {
      // fallback 方法
      const textarea = document.createElement("textarea");
      textarea.value = text;
      document.body.appendChild(textarea);
      textarea.select();
      document.execCommand("copy");
      document.body.removeChild(textarea);
      ElMessage.success("复制成功");
    }
  } catch (err) {
    console.error("复制失败:", err);
    ElMessage.error("复制失败");
  }
};

// 复制 Endpoint
const copyEndpoint = () => {
  copyText("http://localhost:8833");
};

// 复制整个key
const copyAllText = async (row: any)=>{
  const str = row.name+"|"+row.apiKey+"|"+row.status+"|"+row.createTime
  console.log("all text = ",str);
  try {
    if (navigator.clipboard && navigator.clipboard.writeText) {
      await navigator.clipboard.writeText(str);
      ElMessage.success("复制成功");
    } else {
      // fallback 方法
      const textarea = document.createElement("textarea");
      textarea.value = str;
      document.body.appendChild(textarea);
      textarea.select();
      document.execCommand("copy");
      document.body.removeChild(textarea);
      ElMessage.success("复制成功");
    }
    
  } catch (err) {
    console.error("复制失败:", err);
    ElMessage.error("复制失败");
  }
}

// 处理状态变化
const handleStatusChange = async (row: MyKeyItem, newVal: boolean|number|string) => {
  const boolVal = Boolean(newVal);
  row.status = boolVal ? "1" : "0";
   const parasm = {api_keys:apiKeysStore.toServerData}
  const setResp = await setModelKeysApi(parasm);
  if (setResp.errcode === 0) {
    if (boolVal) {
      // 启用当前Key时，先禁用其他所有已启用的Key
      // apiKeysStore.apiKeys.forEach((item: MyKeyItem) => {
      //   if (item.status === "1" && item !== row) {
      //     item.status = "0";
      //     DelKey(item.apiKey);
      //     ElMessage.info(`已禁用 Key: ${item.name}`);
      //   }
      // });
      ElMessage.success(`已启用 Key: ${row.name}`);
    } else {
      // 直接禁用当前Key，不自动启用其他Key
      ElMessage.info(`已禁用 Key: ${row.name}`);
    }
  }
};

// 新增Key函数
const addNewKey = async () => {
  try {
    const { value: keyName } = await ElMessageBox.prompt(
      "请输入新 Key 的名称",
      "新增 Key",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        inputPattern: /^.{1,20}$/,
        inputErrorMessage: "名称长度应在 1-20 个字符之间",
      }
    );

    if (apiKeysStore.hasKeyByName(keyName)) {
      ElMessage.success("名字已存在");
      return;
    }

    if (!isChineseEnglishOrNumber(keyName)) {
      ElMessage.warning("名字非法");
      return;
    }

    // name|key|enable|time
    const curTime = String(Math.floor(Date.now() / 1000));
    const newKeyItem: MyKeyItem = {
      name: keyName,
      apiKey: generateApiKey(),
      status: "0",
      createTime: curTime,
    };
    // const newApi = keyName + "|" + generateApiKey() + "|" + "0|" + curTime;
    let keys = Array.from(apiKeysStore.apiKeys);
    keys.push(newKeyItem);

    const toServer = keys.map(
      (item) => `${item.name}|${item.apiKey}|${item.status}|${item.createTime}`
    );
    const parasm = {api_keys:toServer}
    const resp = await setModelKeysApi(parasm);
    console.log("resp ================== ", resp);
    if (resp.errcode === 0) {
      apiKeysStore.addApiKey(newKeyItem);
    }

    // ElMessage.success("新增 Key 成功");
  } catch (error) {
    if (error !== "cancel") {
      ElMessage.error("新增 Key 失败");
    }
  }
};

// 删除Key函数
const deleteKey = async (row: MyKeyItem) => {
  if (row.status === "1") {
    ElMessage.warning("无法删除已启用的 Key，请先禁用后再删除");
    return;
  }

  try {
    await ElMessageBox.confirm(
      `确定要删除 Key "${row.name}" 吗？此操作不可恢复。`,
      "确认删除",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }
    );

    const index = apiKeysStore.apiKeys.findIndex((item) => item === row);
    if (index !== -1) {
      let delData = apiKeysStore.apiKeys.splice(index, 1);
      const setResp = await setModelKeysApi(apiKeysStore.toServerData);
      if (setResp.errcode === 0) {
        ElMessage.success("删除成功");
      } else {
        if (delData && delData.length > 0) apiKeysStore.addApiKey(delData[0]);
      }
    }
  } catch (error) {
    // 用户取消删除
  }
};

// 生成随机的 API Key
const generateApiKey = (): string => {
  const chars = "abcdefghijklmnopqrstuvwxyz0123456789";
  let result = "sk-";
  for (let i = 0; i < 32; i++) {
    result += chars[Math.floor(Math.random() * chars.length)];
  }
  return result;
};

// 隐藏部分 API Key（安全显示）
const maskApiKey = (key: string): string => {
  if (key.length <= 8) return key;
  return key.substring(0, 8) + "*".repeat(16) + key.substring(key.length - 8);
};
</script>

<style scoped>
.key-management-container {
  padding: 24px;
  background: linear-gradient(135deg, #f5f7fa 0%, #e3e8f0 100%);
  min-height: 100vh;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

h2 {
  margin: 0;
  font-size: 24px;
  color: #673ab7;
  font-weight: 600;
}

.add-key-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  border-radius: 8px;
  padding: 12px 24px;
  font-weight: 600;
}

.add-key-btn:hover {
  background: linear-gradient(135deg, #5a67d8 0%, #6b46c1 100%);
}

.tips-text {
  font-size: 14px;
  color: #666;
  margin-bottom: 20px;
  padding: 0 8px;
}

.endpoint-section {
  margin-bottom: 24px;
  padding: 20px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.endpoint-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.endpoint-label {
  font-weight: 600;
  color: #333;
}

.endpoint-url {
  color: #409eff;
  font-family: monospace;
  font-size: 14px;
}

.copy-endpoint-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  border-radius: 6px;
}

.copy-endpoint-btn:hover {
  background: linear-gradient(135deg, #5a67d8 0%, #6b46c1 100%);
}

.table-container {
  background: white;
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  margin-bottom: 20px;
}

.api-key-container {
  display: flex;
  align-items: center;
  gap: 8px;
}

.api-key-text {
  font-family: monospace;
  font-size: 13px;
  color: #555;
}

.copy-key-btn {
  color: #409eff;
  padding: 4px;
}

.copy-key-btn:hover {
  color: #66b1ff;
}

.action-buttons {
  display: flex;
  align-items: center;
  gap: 12px;
}

.status-switch {
  margin-right: 8px;
}

.pagination {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  margin-top: 20px;
}

.page-info {
  font-size: 14px;
  color: #666;
  background: white;
  padding: 12px 20px;
  border-radius: 8px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.05);
}

:deep(.el-table th) {
  background-color: #f8fafc;
  font-weight: 600;
  color: #2d3748;
}

:deep(.el-table td) {
  padding: 16px 0;
}

:deep(.el-table .cell) {
  padding: 0 16px;
}
</style>
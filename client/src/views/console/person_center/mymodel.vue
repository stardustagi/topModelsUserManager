<template>
  <el-main class="my-models-container">
    <div class="header-section">
      <h2 class="page-title">我的模型</h2>
      <el-button type="primary" size="large" class="refresh-btn" @click="refreshAll">
        <!-- <el-icon>
          <Refresh />
        </el-icon> -->
        刷新列表
      </el-button>
    </div>

    <el-scrollbar :style="{ height: scrollbarHeight }" class="models-scrollbar">
      <div v-for="r in myModelInfos" :key="r.map_node_id + '-' + r.map_model_id" class="model-card">
        <div class="card-header">
          <h3 class="model-name">{{ r.name }}</h3>
        </div>

        <div class="model-details">
          <div class="detail-row">
            <span class="detail-label">输入价格:</span>
            <span class="detail-value">¥{{ r.input_price / 100 }}/百万token</span>
          </div>
          <div class="detail-row">
            <span class="detail-label">输出价格:</span>
            <span class="detail-value">¥{{ r.output_price / 100 }}/百万token</span>
          </div>
          <div class="detail-row">
            <span class="detail-label">缓存价格:</span>
            <span class="detail-value">¥{{ r.cache_price / 100 }}/百万token</span>
          </div>
          <div class="detail-row">
            <span class="detail-label">延迟:</span>
            <span class="detail-value" :class="getLatencyClass(r.latency)">
              {{ r.latency }}ms
            </span>
          </div>
        </div>

        <div class="card-footer">
          <span class="update-time">更新时间: {{ formatTime(r.last_updated) }}</span>
          <el-button type="danger" size="small" class="unsubscribe-btn" @click="onClickUnSubscribe(r)">
            取消订阅
          </el-button>
        </div>
      </div>
    </el-scrollbar>
  </el-main>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { unsubscribeModelApi, userGetSelectLLMInfo } from "@/api/managerApi";
import { useRouter } from "vue-router";
import { ModelEntity } from "@/api/apiParamsResp";
import { getToken } from "@/utils/auth";
import { useMyModelStore } from "@/store/myModelsStore";
import { useLatencyStore } from "@/store/latencyStore";


const myModelInfos = ref<ModelEntity[]>([]);
const router = useRouter();
const scrollbarHeight = ref("500px"); // 默认高度
const latencyStore = useLatencyStore();


onMounted(async () => {
  setScrollbarHeight();
  await getModelList();
  await refreshLatency();
});

// 刷新所有
const refreshAll = async () => {
  await getModelList();

  // 或强制刷新延迟
  const nodeIds = Array.from(
    new Set(myModelInfos.value.map(r => r.map_node_id).filter(id => !!id))
  );

  await latencyStore.refreshLatency(nodeIds);

  myModelInfos.value = myModelInfos.value.map(r => ({
    ...r,
    latency: latencyStore.getLatency(r.map_node_id)
  }));
};


// 获取模型列表
const getModelList = async () => {
  if (useMyModelStore().models.length <= 0) {
    let params = {
      skip: 0,
      limit: 10000,
      sort: ""
    }
    const resp = await userGetSelectLLMInfo(params);
    console.log("我的模型=== ", resp);
    if (resp.errcode === 0) {
      if (resp.data && resp.data.models_config && resp.data.models_config.length > 0) {
        myModelInfos.value = [];
        for (let i = 0; i < resp.data.models_config.length; i++) {
          const rule = resp.data.models_config[i];
          const reqRule: ModelEntity = {
            // id: rule.map_id,
            name: rule.model_name,
            provider: "",
            address: rule.domain,
            input_price: rule.model_input_price,
            output_price: rule.model_output_price,
            cache_price: rule.model_cache_price,
            latency: 0,
            health_score: 0,
            last_updated: rule.last_update,
            map_node_id: rule.map_node_id,
            map_model_id: rule.map_model_id,
          };
          myModelInfos.value.push(reqRule);
          useMyModelStore().addMyModel(reqRule);
        }
      }
    }
  } else {
    myModelInfos.value = useMyModelStore().models;
  }
};

// 取消订阅
const onClickUnSubscribe = async (rule: ModelEntity) => {
  const userId = getToken().userId;
  if (!userId) {
    router.push("/login");
    return;
  }
  console.log("userid ====== ", userId);

  const nodeIdStr = String(rule.map_node_id);
  const modelIds = myModelInfos.value
    .filter(
      m => m.map_node_id === rule.map_node_id && m.map_model_id !== rule.map_model_id
    )
    .map(m => String(m.map_model_id));

  // 如果已经是最后一个模型，取消后就这个节点下没有模型了，可以考虑直接删除该节点

  const usreq: UserSaveSelectLLMInfoReq = {
    user_id: Number(userId),
    select_models: [
      {
        node_id: nodeIdStr,
        model_ids: modelIds,
      },
    ],
  };

  const res1 = await unsubscribeModelApi(usreq);
  if (res1.errcode === 0) {
    myModelInfos.value = myModelInfos.value.filter(
      m => !(m.map_node_id === rule.map_node_id && m.map_model_id === rule.map_model_id)
    );
    useMyModelStore().removeModel(rule.map_node_id, rule.map_model_id);
  }
  getModelList();
};

// 获取视口高度并设置滚动条高度
const setScrollbarHeight = () => {
  const vh = window.innerHeight * 0.9; // 视口高度的90%
  scrollbarHeight.value = `${vh}px`;
};

// 根据延迟获取样式类
const getLatencyClass = (latency: number) => {
  if (latency < 500) return "latency-low";
  if (latency < 1000) return "latency-medium";
  return "latency-high";
};

// 格式化时间
const formatTime = (timestamp: number) => {
  if (!timestamp) return "未知";
  const date = new Date(timestamp * 1000);
  return date.toLocaleString();
};

// 刷新延迟
const refreshLatency = async () => {
  console.log("my model info = ", myModelInfos.value);
  const nodeIds = Array.from(
    new Set(myModelInfos.value.map(r => r.map_node_id).filter(id => !!id))
  );
  console.log("my node ids === ", nodeIds);
  await latencyStore.loadLatency(nodeIds);

  myModelInfos.value = myModelInfos.value.map(r => ({
    ...r,
    latency: latencyStore.getLatency(r.map_node_id)
  }));
};

</script>



<style scoped>
.my-models-container {
  padding: 24px;
  background: linear-gradient(135deg, #f5f7fa 0%, #e3e8f0 100%);
  min-height: 100vh;
}

.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.page-title {
  color: #673ab7;
  font-size: 24px;
  font-weight: 600;
  margin: 0;
}

.refresh-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  border-radius: 8px;
  padding: 12px 24px;
  font-weight: 600;
}

.refresh-btn:hover {
  background: linear-gradient(135deg, #5a67d8 0%, #6b46c1 100%);
}

.models-scrollbar {
  margin: 0 -24px;
  padding: 0 24px;
}

.model-card {
  background: white;
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 16px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
  border: 1px solid #e2e8f0;
}

.model-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.12);
}

.card-header {
  display: flex;
  align-items: center;
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid #f0f2f5;
}

.drag-handle {
  color: #a0aec0;
  margin-right: 12px;
  cursor: move;
}

.model-name {
  font-size: 18px;
  font-weight: 600;
  color: #2d3748;
  margin: 0;
  flex: 1;
}

.provider-tag {
  background: #667eea;
  color: white;
  border: none;
  margin-left: 12px;
}

.model-details {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
  margin-bottom: 16px;
}

.detail-row {
  display: flex;
  align-items: center;
}

.detail-label {
  color: #718096;
  font-size: 14px;
  font-weight: 500;
  margin-right: 8px;
  min-width: 80px;
}

.detail-value {
  color: #2d3748;
  font-size: 14px;
  font-weight: 600;
}

.latency-low {
  color: #38a169;
}

.latency-medium {
  color: #d69e2e;
}

.latency-high {
  color: #e53e3e;
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 16px;
  padding-top: 16px;
  border-top: 1px solid #f0f2f5;
}

.update-time {
  color: #a0aec0;
  font-size: 12px;
}

.unsubscribe-btn {
  background: #f56c6c;
  border: none;
  border-radius: 6px;
  padding: 8px 16px;
  font-weight: 500;
}

.unsubscribe-btn:hover {
  background: #f78989;
}
</style>
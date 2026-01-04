<template>
  <el-main class="model-market-container">
    <div class="market-header">
      <h2 class="market-title">模型市场</h2>
      <div class="search-filter">
        <el-input v-model="searchQuery" placeholder="搜索模型名称..." class="search-input" size="large">
          <template #prefix>
            <!-- <el-icon>
              <Search />
            </el-icon> -->
          </template>
        </el-input>
        <el-select v-model="sortOrder" placeholder="排序方式" size="large" class="sort-select">
          <el-option label="价格从低到高" value="price-asc" />
          <el-option label="价格从高到低" value="price-desc" />
          <el-option label="延迟从低到高" value="latency-asc" />
          <!--  <el-option label="评分从高到低" value="rating-desc" /> -->
        </el-select>

        <!-- 新增模型类型切换 -->
        <!-- <el-radio-group v-model="modelType" size="large">
            <el-radio-button label="0" @click="handleModelTypeChange(0)">公有模型</el-radio-button>
            <el-radio-button label="1" @click="handleModelTypeChange(1)">私有模型</el-radio-button>
          </el-radio-group> -->
      </div>
    </div>

    <el-scrollbar :style="{ height: scrollbarHeight }" class="market-scrollbar">
      <div class="models-grid">
        <!-- <div v-for="r in filteredRules" :key="r.id" class="model-card"> -->
        <div v-for="r in filteredRules" :key="r.map_node_id + '-' + r.map_model_id" class="model-card">
          <div class="card-header">
            <h3 class="model-name">{{ r.name }}</h3>
            <!-- <el-tag type="success" size="small" class="provider-tag">
              {{ r.provider }}
            </el-tag> -->
          </div>

          <div class="model-stats">
            <div class="stat-item">
              <span class="stat-label">输入价格:</span>
              <span class="stat-value">¥{{ r.input_price / 100 }}/百万token</span>
            </div>
            <div class="stat-item">
              <span class="stat-label">输出价格:</span>
              <span class="stat-value">¥{{ r.output_price / 100 }}/百万token</span>
            </div>
            <div class="stat-item">
              <span class="stat-label">缓存价格:</span>
              <span class="stat-value">¥{{ r.cache_price / 100 }}/百万token</span>
            </div>
            <div class="stat-item">
              <span class="stat-label">延迟:</span>
              <span class="stat-value" :class="getLatencyClass(r.latency)">
                {{ r.latency }}ms
              </span>
            </div>
            <!-- <div class="stat-item">
              <span class="stat-label">性能:</span>
              <span class="stat-value">{{ r.health_score }} token/s</span>
            </div> -->
          </div>

          <div class="card-footer">
            <div class="model-meta">
              <span class="update-time">更新于: {{ formatTime(r.last_updated) }}</span>
            </div>
            <!-- <el-button v-if="r.subscribed" class="subscribe-btn" type="primary" size="large">
              已订阅
            </el-button>
            <el-button v-else class="subscribe-btn" type="primary" size="large" @click="onClickSubscribe(r)">
              订阅模型
            </el-button> -->
          </div>
        </div>
      </div>
    </el-scrollbar>

    <!-- :page-sizes="[2]" 每页固定数量，多个则可以选择每页显示数量-->
    <!-- layout里面有个sizes,加上就显示可以选择每页显示的数量-->
    <div class="pagination-container">
      <el-pagination v-model:current-page="pagination.currentPage" v-model:page-size="pagination.pageSize"
        :total="pagination.total" layout="total, prev, pager, next" @size-change="getModelList"
        @current-change="getModelList" />
    </div>
  </el-main>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from "vue";
import {
  subscribeModelApi,
  getModelMarketApi,
  userGetSelectLLMInfo,
} from "@/api/managerApi";
import type { modelMarketEntity } from "@/store/modelMarektEntity";
import { useRouter } from "vue-router";
import { useMyModelStore } from "@/store/myModelsStore";
import type { ModelEntity, ModelMarketDTO, ModelMarketResp } from "@/api/apiParamsResp";
import { type MyKeyItem, useMyKeysStore } from "@/store/myKeyStore";
import { ElMessage } from "element-plus";
import { getToken } from "@/utils/auth";
import { useLatencyStore } from "@/store/latencyStore";

const myModelStore = useMyModelStore();
const router = useRouter();
const scrollbarHeight = ref("500px"); // 默认高度

// 搜索和排序
const searchQuery = ref("");
const sortOrder = ref("");

// 模型类型  0公有
const modelType = ref("0");

// 已订阅模型
const modelInfos = ref<ModelEntity[]>([]);

// 模型市场数据
const rules = ref<modelMarketEntity[]>([]);

// 延迟数据
const delayMap = ref<Record<string, number>>({});

// 分页数据
const pagination = reactive({
  currentPage: 1,
  pageSize: 50,
  total: 0,
  sort: "", // im.id
});

const latencyStore = useLatencyStore();

// 计算属性：过滤 + 排序
const filteredRules = computed(() => {
  console.log("-------------------------------------1111111111")
  // let data = [...rules.value];

  let data = rules.value.map(r => ({
    ...r,
    latency: latencyStore.getLatency(r.map_node_id)
  }));

  for (let i = 0; i < rules.value.length; i++) {
    console.log(rules.value[i].latency, "        111")
  }


  // 搜索过滤
  if (searchQuery.value.trim()) {
    const keyword = searchQuery.value.trim().toLowerCase();
    data = data.filter((item) =>
      item.name.toLowerCase().includes(keyword)
    );
  }

  // 排序
  switch (sortOrder.value) {
    case "price-asc":
      data.sort((a, b) => a.input_price - b.input_price);
      break;
    case "price-desc":
      data.sort((a, b) => b.input_price - a.input_price);
      break;
    case "latency-asc":
      data.sort((a, b) => a.latency - b.latency);
      break;
    case "rating-desc":
      data.sort((a, b) => b.health_score - a.health_score);
      break;
  }

  return data;
});

onMounted(() => {
  // modelType.value = getCurUserModel();
  modelType.value = "0";
  setScrollbarHeight();
  getMyModelList(false);
});

// 拉取我的模型
// 获取模型列表
const getMyModelList = async (del: boolean) => {
  const res = useMyModelStore().models;

  console.log("0000000000000000   拉取我的模型=", res);

  // let ask: boolean = false;
  // if (del) {
  //   ask = true;
  //   modelInfos.value = [];
  // } else {
  //   if (res && res.length > 0) {
  //     modelInfos.value = res;
  //   } else {
  //     ask = true;
  //   }
  // }
  // if (ask) {
  //   let params = {
  //     skip: 0,
  //     limit: 10000,
  //     sort: ""
  //   }
  //   const resp = await userGetSelectLLMInfo(params);
  //   if (resp.errcode === 0) {
  //     console.log("当前模型是: ", resp.data);
  //     if (resp.data && resp.data.models_config && resp.data.models_config.length > 0) {
  //       for (let i = 0; i < resp.data.models_config.length; i++) {
  //         const rule = resp.data.models_config[i];
  //         const reqRule: ModelEntity = {
  //           // id: rule.map_id,
  //           name: rule.model_name,
  //           provider: "",
  //           address: rule.domain,
  //           input_price: rule.model_input_price,
  //           output_price: rule.model_output_price,
  //           cache_price: rule.model_cache_price,
  //           latency: 0,
  //           health_score: 0,
  //           last_updated: rule.last_update,
  //           map_node_id: rule.map_node_id,
  //           map_model_id: rule.map_model_id,
  //         };
  //         modelInfos.value.push(reqRule);
  //         useMyModelStore().addMyModel(reqRule);
  //       }
  //     }
  //   }
  // }
  await getModelList();
  await loadLatency();
};

// 切换模型类型,模型市场不变，私有模型后，清理我的模型，返回得到私有的模型
const handleModelTypeChange = async (type: number) => {
  //   if (type === Number(modelType.value)) {
  //     return;
  //   }
  //   const userId = currentUserId.value;
  //   console.log("type === ", type)
  //   const res = await userSetPrivateModelApi(Number(userId), type);
  //   if (res.errcode === 0) {
  //     pagination.currentPage = 1;
  //     setCurUserModel(type);
  //     modelType.value = String(type);

  //     getMyModelList(true);
  //   }
};

// 拉取模型信息
const getModelList = async () => {
  // const res = await ListModels();
  // let marktes = marketStore.getModelMarket();
  const params = {
    skip: (pagination.currentPage - 1) * pagination.pageSize,
    limit: pagination.pageSize,
    sort: pagination.sort,
  };

  const respMarktes = await getModelMarketApi(params);
  console.log("markets === ", respMarktes);
  if (respMarktes.errcode === 0) {
    // 不存储在store
    rules.value = [];
    if (respMarktes.data && respMarktes.data.length >= 2) {
      for (let i = 0; i < respMarktes.data[0].length; i++) {
        let dto: ModelMarketDTO = respMarktes.data[0][i];

        let hasSub: boolean = isModleExistSubscribe(dto.map_node_id, dto.map_model_id,);
        // console.log("has sub === ", hasSub, "   ", dto);
        // 取自己模型数据
        let mme: modelMarketEntity = {
          // id: dto.info_id,
          provider: dto.provider_provider_id,
          name: dto.info_name,
          address: dto.info_address,
          input_price: dto.info_input_price,
          output_price: dto.info_output_price,
          cache_price: dto.info_cache_price,
          latency: 0,
          health_score: 0,
          last_updated: dto.info_last_update,
          subscribed: hasSub,
          map_node_id: dto.map_node_id,
          map_model_id: dto.map_model_id,
        };
        rules.value.push(mme);
        // }
      }

      if (respMarktes.data[1].length > 0) {
        pagination.total = respMarktes.data[1][0].total_count;
      }
    }
  } else {
    ElMessage.error("拉取模型失败，请检查网络111");
  }
};


// 是否已经订阅了
const isModleExistSubscribe = (mapNodeId: number, mapModelId: number) => {
  return modelInfos.value.some(m =>
    m.map_node_id === mapNodeId && m.map_model_id === mapModelId
  );
};

// 排序时候，回到第一页
const handleSortChange = (val: string) => {
  pagination.sort = val;
  pagination.currentPage = 1;
  getModelList();
};

// 订阅模型
// const onClickSubscribe = async (rule: modelMarketEntity) => {
//   if (modelType.value !== "0") {
//     ElMessage.warning('请切换到公有模型在订阅');
//     return;
//   }
//   console.log("ruleId === ", rule);
//   // 验证模型是否存在 暂时不做

//   console.log(myModelStore.models);

//   // 发送给中心服务器
//   const userId = getToken().userId;
//   if (!userId) {
//     router.push("/login");
//     return;
//   }

//   // 找出当前 node_id 下的所有已订阅模型 ID（排除重复）
//   const nodeIdStr = String(rule.map_node_id);

//   const sameNodeModelIds = modelInfos.value
//     .filter(m => m.map_node_id === rule.map_node_id)
//     .map(m => String(m.map_model_id));

//   // 包含当前点击的 model_id
//   const fullModelIdsSet = new Set(sameNodeModelIds);
//   fullModelIdsSet.add(String(rule.map_model_id));

//   const model_ids = Array.from(fullModelIdsSet);

//   const usreq: UserSaveSelectLLMInfoReq = {
//     user_id: Number(userId),
//     select_models: [
//       {
//         node_id: nodeIdStr,
//         model_ids: model_ids,
//       },
//     ],
//   };
//   const res1 = await subscribeModelApi(usreq);
//   if (res1.errcode === 0) {
//     // 发送订阅
//     const reqRule: ModelEntity = {
//       // id: rule.id,
//       name: rule.name,
//       provider: rule.provider,
//       address: rule.address,
//       input_price: rule.input_price,
//       output_price: rule.output_price,
//       cache_price: rule.cache_price,
//       latency: rule.latency,
//       health_score: rule.health_score,
//       last_updated: rule.last_updated,
//       map_node_id: rule.map_node_id,
//       map_model_id: rule.map_model_id,
//     };

//     console.log(reqRule, "111");

//     const idx = rules.value.findIndex(
//       item =>
//         item.map_node_id === rule.map_node_id &&
//         item.map_model_id === rule.map_model_id
//     );
//     if (idx !== -1) {
//       rules.value[idx] = {
//         ...rules.value[idx],
//         subscribed: true,
//       };
//     }

//     myModelStore.addMyModel(reqRule);
//   } else {
//     ElMessage.warning("订阅失败");
//   }
// };

// 获取视口高度并设置滚动条高度
const setScrollbarHeight = () => {
  const vh = window.innerHeight * 0.9; // 视口高度的90%
  scrollbarHeight.value = `${vh}px`;
};

const getLatencyClass = (latency: number) => {
  if (latency < 500) return "latency-low";
  if (latency < 1000) return "latency-medium";
  return "latency-high";
};

const formatTime = (timestamp: number) => {
  if (!timestamp) return "刚刚";
  const date = new Date(timestamp * 1000);
  return date.toLocaleString();
};

// 获取延迟数据
const loadLatency = async () => {
  const nodeIds = Array.from(
    new Set(rules.value.map(r => r.map_node_id).filter(id => !!id))
  );

  console.log("node ids === ", nodeIds);

  await latencyStore.loadLatency(nodeIds);


  // for (const r of rules.value) {
  //   const v = latencyStore.getLatency(r.map_node_id);
  //   console.log("vvvvvvvvvvvvvvvvvvvvvvv === ", v)
  //   if (v !== undefined && v !== null) {
  //     console.log("33333")
  //     r.latency = v;   // 🔥 这里是关键
  //   }
  // }
  rules.value = [...rules.value];
};

</script>



<style lang="scss" scoped>
.model-market-container {
  padding: 20px;
  background: linear-gradient(135deg, #f5f7fa 0%, #e3e8f0 100%);
  min-height: 100vh;
}

.market-header {
  background: white;
  padding: 20px;
  border-radius: 12px;
  margin-bottom: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.market-title {
  color: #673ab7;
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 20px;
  text-align: center;
}

.search-filter {
  display: flex;
  gap: 16px;
  align-items: center;
}

.search-input {
  flex: 1;
}

.sort-select {
  width: 200px;
}

.market-scrollbar {
  // margin: 0 -20px;
  height: calc(100vh - 280px) !important;
}

.models-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 20px;
  padding: 0 20px;
}

.model-card {
  background: white;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
  border: 1px solid #e2e8f0;
}

.model-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.model-name {
  color: #2d3748;
  font-size: 18px;
  font-weight: 600;
  margin: 0;
}

.provider-tag {
  background: #667eea;
  color: white;
  border: none;
}

.model-stats {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 20px;
}

.stat-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.stat-label {
  color: #718096;
  font-size: 14px;
  font-weight: 500;
}

.stat-value {
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
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid #e2e8f0;
}

.update-time {
  color: #a0aec0;
  font-size: 12px;
}

.subscribe-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  border-radius: 8px;
  padding: 12px 24px;
  font-weight: 600;
}

.subscribe-btn:hover {
  background: linear-gradient(135deg, #5a67d8 0%, #6b46c1 100%);
  transform: translateY(-1px);
}

// 下面的分页
.pagination-container {
  display: flex;
  align-items: center;
  justify-content: center;
  /* 关键：水平居中 */
  margin-top: 20px;
  padding: 16px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  // 适配移动端
  // @media (max-width: 768px) {
  //   :deep(.el-pagination) {
  //     --el-pagination-button-width: 28px;
  //     --el-pagination-button-height: 28px;
  //   }
  // }
}

.total-text {
  margin-right: 15px;
  /* 与分页器间距 */
  color: #606266;
  font-size: 14px;
  white-space: nowrap;
  /* 防止文字换行 */
}
</style>
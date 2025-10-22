import { defineStore } from "pinia";
import { computed, ref } from "vue";

export interface MyKeyItem {
  name: string;
  apiKey: string;
  status: "1" | "0";
  createTime: string;
}

export const useMyKeysStore = defineStore("mykeys", () => {
  // state
  const apiKeys = ref<MyKeyItem[]>([]);

  // actions
  const setApiKeys = (keys: MyKeyItem[]) => {
    apiKeys.value = keys;
  };

  const addApiKey = (key: MyKeyItem) => {
    if (key.name.trim() && !apiKeys.value.some(k => k.apiKey === key.apiKey)) {
      apiKeys.value.push(key);
    }
  };

  const removeApiKey = (target: string) => {
    apiKeys.value = apiKeys.value.filter(k => k.apiKey !== target);
  };

  const updateKeyStatus = (apiKey: string, status: "1" | "0") => {
    const key = apiKeys.value.find(k => k.apiKey === apiKey);
    if (key) {
      key.status = status;
    }
  };

  // 服务器数据转换成apiKeys
  const convertToMyKeyItems = (apiKeysString: string[]): MyKeyItem[] => {
    const curTime = String(Math.floor(Date.now() / 1000));
    return apiKeysString
      .map(item => {
        // 使用 "|" 分割字符串
        const parts = item.split("|");

        // 确保有足够的组成部分
        if (parts.length < 4) {
          console.warn("Invalid API key format:", item);
          return null;
        }

        // 构建 MyKeyItem 对象
        return {
          name: parts[0] || "",
          apiKey: parts[1] || "",
          status: (parts[2] === "1" ? "1" : "0") as "1" | "0", // 确保类型安全
          createTime: parts[3] || curTime
        };
      })
      .filter(item => item !== null) as MyKeyItem[]; // 过滤掉无效项
  };

  // getter
  const keysCount = computed(() => apiKeys.value.length);
  const toServerData = computed(() => {
    return apiKeys.value.map(
      item => `${item.name}|${item.apiKey}|${item.status}|${item.createTime}`
    );
  });

  const enabledKeys = computed(() =>
    apiKeys.value.filter(k => k.status === "1")
  );

  const disabledKeys = computed(() =>
    apiKeys.value.filter(k => k.status === "0")
  );

  const hasKeyByName = (name: string): boolean => {
    return apiKeys.value.some(k => k.name === name);
  };

  return {
    apiKeys,
    setApiKeys,
    addApiKey,
    removeApiKey,
    keysCount,
    updateKeyStatus,
    toServerData,
    enabledKeys,
    disabledKeys,
    hasKeyByName,
    convertToMyKeyItems
  };
});

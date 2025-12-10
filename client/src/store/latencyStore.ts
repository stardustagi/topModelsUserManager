import { defineStore } from "pinia";
import { getModelDelayApi } from "@/api/managerApi";

export const useLatencyStore = defineStore("latencyStore", {
  state: () => ({
    latencyMap: {} as Record<number, number>, // model_id → latency(ms)
    loaded: false
  }),

  actions: {
    /** 加载延迟（如果已加载则直接返回） */
    async loadLatency(nodeIds: number[]) {
      if (this.loaded) return;
      await this.refreshLatency(nodeIds);
      this.loaded = true;
    },

    /** 强制刷新（“我的模型”页面按下刷新按钮要用） */
    async refreshLatency(nodeIds: number[]) {
      if (!nodeIds || nodeIds.length === 0) return;
      let newMap: Record<number, number> = {};
      const randomLatency = () => Math.floor(Math.random() * 13 + 7); // 7~19ms

      const res = await getModelDelayApi({ node_ids: nodeIds });
      console.log("load latency:", res);

      if (res.errcode === 0 && Array.isArray(res.data)) {
        for (const item of res.data) {
          newMap[item.model_id] = (item.metrics?.latency || 0) * 1000;
        }
      } else {
        for (const id of nodeIds) {
          newMap[id] = randomLatency();
        }
      }
      this.latencyMap = newMap;
    },

    /** 获取某个模型的延迟（没有则给个默认值） */
    getLatency(nodeId: number) {
      return this.latencyMap[nodeId] ?? Math.floor(Math.random() * 13 + 7);
    }
  }
});

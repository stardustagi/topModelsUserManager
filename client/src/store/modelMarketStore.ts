import { defineStore } from "pinia";
import { ref } from "vue";
import type { modelMarketEntity } from "./modelMarektEntity";

export const useMarketStore = defineStore("modelmarket", () => {
  // state
  const modelLists = ref<modelMarketEntity[]>([]);
  const totalModels = ref(0);
  const totalNodes = ref(0);

  // actions
  const updateModelMarket = (
    entities: modelMarketEntity[],
    modelNum: number,
    nodeNum: number
  ) => {
    modelLists.value = entities;
    totalModels.value = modelNum;
    totalNodes.value = nodeNum;
  };

  // getter
  const getModelMarket = () => ({
    data: {
      node_models: modelLists.value,
      total_models: totalModels.value,
      total_nodes: totalNodes.value
    }
  });

  return {
    modelLists,
    totalModels,
    totalNodes,
    updateModelMarket,
    getModelMarket
  };
});

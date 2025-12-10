import { ModelEntity } from "@/api/apiParamsResp";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useMyModelStore = defineStore("mymodel", () => {
  // state
  const models = ref<ModelEntity[]>([]);

  // actions
  // const updateModels = (newModels: any[]) => {
  //     models.value = newModels.map(m => model.Model.createFrom(m));
  // };

  const isSameModel = (a: ModelEntity, b: ModelEntity) => {
    return a.map_node_id === b.map_node_id && a.map_model_id === b.map_model_id;
  };

  const addMyModel = (newModel: ModelEntity) => {
    const exists = models.value.some(m => isSameModel(m, newModel));
    if (!exists) {
      models.value.push(newModel);
    }
  };

  const removeModel = (node_id: number, model_id: number) => {
    models.value = models.value.filter(
      m => !(m.map_node_id === node_id && m.map_model_id === model_id)
    );
  };

  // getter

  return {
    models,
    // updateModels,
    addMyModel,
    removeModel
  };
});

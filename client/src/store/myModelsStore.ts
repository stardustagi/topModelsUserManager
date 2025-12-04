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

  const addMyModel = (newModel: ModelEntity) => {
    const exists = models.value.some(m => newModel.id === m.id);
    if (!exists) {
      models.value.push(newModel);
    }
  };

  const removeModel = (id: number) => {
    models.value = models.value.filter(m => m.id != id);
  };

  // getter

  return {
    models,
    // updateModels,
    addMyModel,
    removeModel
  };
});

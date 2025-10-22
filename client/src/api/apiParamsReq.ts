interface UserSelectModelIdsDTO {
  node_id: string;
  model_ids: string[];
}

interface UserSaveSelectLLMInfoReq {
  user_id: number;
  select_models: UserSelectModelIdsDTO[];
}
